package k8s

import (
	"regexp"
	"strings"

	. "pilot_agent/apps/funcs"
	. "pilot_agent/apps/pstree"
)

type PodProcess struct {
	Pod K8SPods

	Name    string
	CmdLine string
	Stat    ProcessStat
}

func addPods(procs []PodProcess, pod K8SPods, pid int, tree *Tree) []PodProcess {
	for _, cid := range tree.Procs[pid].Children {
		var pod_proc PodProcess

		pod_proc.Pod = pod
		pod_proc.Name = tree.Procs[cid].Name
		pod_proc.CmdLine = tree.Procs[cid].CmdLine
		pod_proc.Stat = tree.Procs[cid].Stat

		procs = append(procs, pod_proc)
		procs = addPods(procs, pod, cid, tree)
	} // end of for

	return procs
}

var _regs_proc []*regexp.Regexp = []*regexp.Regexp{
	regexp.MustCompile(`.*containerd.*-namespace[\s](.*)[\s].*-id[\s]([0-9a-zA-Z]+)[\s]+`),
	regexp.MustCompile(`.*conmon.*\/overlay-containers\/(.*).*-c[\s]([0-9a-zA-Z]+)[\s]+`)}

func GetProcContainerID(cmdline string) (bool, string) {
	var result bool = false
	var container_id string

	for _, reg := range _regs_proc {
		matchs := reg.FindStringSubmatch(cmdline)
		if len(matchs) >= 3 {
			result = true
			container_id = matchs[2]
			break
		}
	} // end of for

	return result, container_id
}

func GetPods(pods []K8SPods, tree *Tree) (bool, []PodProcess) {
	var pod_procs []PodProcess = nil

	for _, proc := range tree.Procs {

		cmds := strings.Split(proc.CmdLine, "\x00")
		cmdline := strings.Join(cmds, " ")

		rtn, container_id := GetProcContainerID(cmdline)
		if false == rtn {
			continue
		}

		rtn, pod := SearchK8SPods(container_id, pods)
		if false == rtn {
			continue
		}

		pod_procs = addPods(pod_procs, pod, proc.Stat.Pid, tree)

	} // end of for

	return IfThenElse(len(pod_procs) > 0, true, false).(bool), pod_procs
}
