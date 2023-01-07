package cmds

import (
	"net/http"
	"sort"
	"strings"
	"time"

	"pilot_agent/apps"
	"pilot_agent/apps/funcs"
	"pilot_agent/apps/k8s"
	"pilot_agent/apps/pstree"
)

func getPodsProclist() (int, string) {
	rtn, pods := k8s.GetK8SPods()
	if false == rtn {
		return apps.GetHttpError(http.StatusBadRequest, "get K8S pods failed")
	}

	tree, err := pstree.New("/rootfs/proc")
	if err != nil {
		return apps.GetHttpError(http.StatusBadRequest, "get proc failed")
	}

	rtn, procs := k8s.GetPods(pods, tree)
	if false == rtn {
		return apps.GetHttpError(http.StatusBadRequest, "get pods failed")
	}

	sort.Slice(procs, func(i, j int) bool {
		stri := procs[i].Pod.Namespace + "/" + procs[i].Pod.Name
		strj := procs[j].Pod.Namespace + "/" + procs[j].Pod.Name

		cmp := strings.Compare(stri, strj)
		if 0 == cmp {
			return procs[i].Stat.Pid < procs[j].Stat.Pid
		}
		return cmp < 0
	})

	type ProcListData struct {
		Name string `json:"proc"`
		PID  int    `json:"pid"`
		Cmds string `json:"cmds"`
	}
	type PodListData struct {
		Namespace   string `json:"namespace"`
		PodName     string `json:"podName"`
		ContainerID string `json:"containerID"`

		ProcList []ProcListData `json:"procList"`
	}
	type JsonData struct {
		Mode      string `json:"mode"`
		NodeName  string `json:"nodeName"`
		TimeStamp string `json:"timestamp"`

		PodList []PodListData `json:"podList"`
	}

	var json_data JsonData
	json_data.Mode = "GET_PODS_PROCLIST"
	json_data.TimeStamp = time.Now().Format(time.RFC3339)
	for _, proc := range procs {
		json_data.NodeName = proc.Pod.NodeName
		break
	}

	var save_podlist PodListData
	var comp_str1, comp_str2 string

	for _, proc := range procs {
		comp_str2 = proc.Pod.Namespace + "/" + proc.Pod.Name
		if len(comp_str1) > 0 && !strings.EqualFold(comp_str1, comp_str2) {
			json_data.PodList = append(json_data.PodList, save_podlist)
			save_podlist.ProcList = nil
		}
		comp_str1 = proc.Pod.Namespace + "/" + proc.Pod.Name

		save_podlist.Namespace = proc.Pod.Namespace
		save_podlist.PodName = proc.Pod.Name
		save_podlist.ContainerID = proc.Pod.ContainerID

		var proclist ProcListData
		proclist.Name = proc.Name
		proclist.PID = proc.Stat.Pid

		cmds := strings.Join(strings.Split(proc.CmdLine, "\x00"), " ")
		proclist.Cmds = strings.TrimRight(cmds, " ")

		save_podlist.ProcList = append(save_podlist.ProcList, proclist)
	} // end of for

	if len(save_podlist.ProcList) > 0 {
		json_data.PodList = append(json_data.PodList, save_podlist)
	}

	return http.StatusOK, funcs.JsonDumps(json_data)
}
