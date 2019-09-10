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

	networkIsolator := "weave"

	contype := mesosproto.ContainerInfo_DOCKER.Enum()

	if cmd.ContainerType == "MESOS" {
		contype = mesosproto.ContainerInfo_MESOS.Enum()
	}

	// Save state of the task
	state := config.State[newTaskID]
	state.Command = cmd
	config.State[newTaskID] = state

	return []*mesosproto.TaskInfo{{
		Name: &cmd.TaskName,
		TaskId: &mesosproto.TaskID{
			Value: &newTaskID,
		},
		AgentId:   agent,
		Resources: defaultResources(),
		Command: &mesosproto.CommandInfo{
			Shell: &cmd.Shell,
			Value: &cmd.Command,
			Uris:  cmd.Uris,
		},
		Container: &mesosproto.ContainerInfo{
			Type: contype,
			Docker: &mesosproto.ContainerInfo_DockerInfo{
				Image: &cmd.ContainerImage,
			},
			NetworkInfos: []*mesosproto.NetworkInfo{{
				Name: &networkIsolator,
			}},
		},
	}}, nil
}
