package mesos

import (
	"../proto"
	cfg "../types"
	"git.aventer.biz/AVENTER/util"
)

func prepareTaskInfoExecuteCommand(agent *mesosproto.AgentID, cmd cfg.Command) ([]*mesosproto.TaskInfo, error) {

	newTaskID, _ := util.GenUUID()

	return []*mesosproto.TaskInfo{{
		Name: &cmd.TaskName,
		TaskId: &mesosproto.TaskID{
			Value: &newTaskID,
		},
		AgentId:   agent,
		Resources: defaultResources(),
		Command: &mesosproto.CommandInfo{
			Shell:       &cmd.Shell,
			Value:       &cmd.Command,
			Uris:        cmd.Uris,
			Environment: &cmd.Environment,
		},
	}}, nil
}

func prepareTaskInfoExecuteContainer(agent *mesosproto.AgentID, cmd cfg.Command) ([]*mesosproto.TaskInfo, error) {
	newTaskID, _ := util.GenUUID()

	return []*mesosproto.TaskInfo{{
		Name: &cmd.TaskName,
		TaskId: &mesosproto.TaskID{
			Value: &newTaskID,
		},
		AgentId:   agent,
		Resources: defaultResources(),
		Executor:  prepareExecuteInfoDockerContainer(cmd),
	}}, nil
}

func prepareExecuteInfoDockerContainer(cmd cfg.Command) *mesosproto.ExecutorInfo {

	networkIsolator := "weave"
	//networkHostname := "testhostname"

	newExecutorID, _ := util.GenUUID()

	return &mesosproto.ExecutorInfo{
		Type: mesosproto.ExecutorInfo_CUSTOM.Enum(),
		ExecutorId: &mesosproto.ExecutorID{
			Value: &newExecutorID,
		},
		Name: &cmd.TaskName,
		Command: &mesosproto.CommandInfo{
			Shell: &cmd.Shell,
			Value: &cmd.Command,
			Uris:  cmd.Uris,
		},
		Container: &mesosproto.ContainerInfo{
			Type: mesosproto.ContainerInfo_DOCKER.Enum(),
			Docker: &mesosproto.ContainerInfo_DockerInfo{
				Image: &cmd.ContainerImage,
			},
			NetworkInfos: []*mesosproto.NetworkInfo{{
				Name: &networkIsolator,
			}},
		},
		Resources: defaultResources(),
	}
}

func prepareExecuteInfoMesosContainer(cmd cfg.Command) *mesosproto.ExecutorInfo {

	networkIsolator := "weave"
	//networkHostname := "testhostname"

	newExecutorId := "default"

	return &mesosproto.ExecutorInfo{
		Type: mesosproto.ExecutorInfo_CUSTOM.Enum(),
		ExecutorId: &mesosproto.ExecutorID{
			Value: &newExecutorId,
		},
		Name: &cmd.TaskName,
		Command: &mesosproto.CommandInfo{
			Shell: &cmd.Shell,
			Value: &cmd.Command,
			Uris:  cmd.Uris,
		},
		Container: &mesosproto.ContainerInfo{
			Type:  mesosproto.ContainerInfo_MESOS.Enum(),
			Mesos: &mesosproto.ContainerInfo_MesosInfo{},
			NetworkInfos: []*mesosproto.NetworkInfo{{
				Name: &networkIsolator,
			}},
		},
		Resources: defaultResources(),
	}
}
