package k8s

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"regexp"
	"strings"

	. "pilot_agent/apps/funcs"
)

type K8SPods struct {
	Name        string
	Namespace   string
	NodeName    string
	Manager     string
	Runtime     string
	ContainerID string
	SandboxID   string
}

func getSandboxID(pod *K8SPods) {
	var fpath string

	if strings.EqualFold("k3s", pod.Manager) {
		fpath = fmt.Sprintf(
			"/rootfs/run/k3s/containerd/io.containerd.runtime.v2.task/k8s.io/%s/config.json", pod.ContainerID)
	} else {
		fpath = fmt.Sprintf(
			"/rootfs/run/containerd/io.containerd.runtime.v2.task/k8s.io/%s/config.json", pod.ContainerID)
	}

	data, err := os.Open(fpath)
	if err != nil {
		log.Printf("# error: getSandboxID faild(%v)\n", err)
		return
	}

	bytes, _ := ioutil.ReadAll(data)

	var json_data JSON_DATA
	json.Unmarshal(bytes, &json_data)

	/*
		{
			"annotations": {
				"io.kubernetes.cri.sandbox-id": "{sendbox_id}"
			}
		}
	*/
	value, stype := JsonValue(json_data, "annotations", "io.kubernetes.cri.sandbox-id")
	if "string" == stype {
		pod.SandboxID = value.(string)
	}
}

func GetK8SPods() (bool, []K8SPods) {
	var results []K8SPods = nil

	server, token := getK8SAccount()

	headers := url.Values{}
	headers.Add("Authorization", "Bearer "+token)

	path := fmt.Sprintf("%s/api/v1/pods", server)

	status, output := HttpRequest(path, "GET", headers)
	if 200 != status {
		log.Printf("# error: HttpRequest failed(status=%d).\n", status)
		log.Println(string(output))
		return false, results
	}

	var jdata JSON_DATA
	json.Unmarshal(output, &jdata)

	reg, _ := regexp.Compile(`(.*):\/\/+([0-9a-zA-Z]+)`)

	var data JSON_DATA
	var datas []JSON_DATA
	var buff string

	for _, item := range jdata["items"].([]interface{}) {
		var k8s_pod = K8SPods{}

		json.Unmarshal([]byte(JsonDumps(item)), &data)
		value, stype := JsonValue(data, "metadata", "namespace")
		if "string" == stype {
			k8s_pod.Namespace = value.(string)
		}
		if strings.EqualFold(k8s_pod.Namespace, "kube-system") {
			continue
		}

		value, stype = JsonValue(data, "spec", "nodeName")
		if "string" == stype {
			k8s_pod.NodeName = value.(string)
		}
		value, stype = JsonValue(data, "metadata", "name")
		if "string" == stype {
			k8s_pod.Name = value.(string)
		}

		fields := data["metadata"].(map[string]interface{})["managedFields"].([]interface{})
		if nil != fields {
			buff = JsonDumps(fields)
			json.Unmarshal([]byte(buff), &datas)
			k8s_pod.Manager = datas[0]["manager"].(string)
		}

		json.Unmarshal([]byte(JsonDumps(data["status"])), &data)
		buff = JsonDumps(GetNC(data["containerStatuses"]))

		json.Unmarshal([]byte(buff), &datas)

		match := reg.FindStringSubmatch(datas[0]["containerID"].(string))
		if len(match) < 3 {
			continue
		}
		k8s_pod.Runtime = match[1]
		k8s_pod.ContainerID = match[2]

		if strings.EqualFold("containerd", k8s_pod.Runtime) {
			getSandboxID(&k8s_pod)
		}

		results = append(results, k8s_pod)
	} // end of for

	return IfThenElse(len(results) > 0, true, false).(bool), results
}

func SearchK8SPods(container_id string, pods []K8SPods) (bool, K8SPods) {
	var result = false
	var result_pod = K8SPods{}

	for _, pod := range pods {
		if strings.EqualFold("containerd", pod.Runtime) {
			if strings.EqualFold(container_id, pod.SandboxID) {
				result = true
				result_pod = pod
				break
			}
		} else {
			if strings.EqualFold(container_id, pod.ContainerID) {
				result = true
				result_pod = pod
				break
			}
		} // end of if
	} // end of for

	return result, result_pod
}
