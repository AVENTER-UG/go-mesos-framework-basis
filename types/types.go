package types

import "../proto"

type Config struct {
	FrameworkPort     string
	FrameworkBind     string
	FrameworkUser     string
	FrameworkName     string
	FrameworkInfo     mesosproto.FrameworkInfo
	Principal         string
	Username          string
	Password          string
	MesosMasterServer string
	MesosStreamID     string
	TaskID            uint64
	SSL               bool
	LogLevel          string
	MinVersion        string
	AppName           string
	EnableSyslog      bool
	Hostname          string
	Listen            string
	CommandChan       chan Command
}

type Command struct {
	ContainerImage string                        `json:"container_image,omitempty"`
	ContainerType  string                        `json:"container_type,omitempty"`
	TaskName       string                        `json:"task_name,omitempty"`
	Command        string                        `json:"command,omitempty"`
	Shell          bool                          `protobuf:"varint,6,opt,name=shell,def=1" json:"shell,omitempty"`
	Uris           []*mesosproto.CommandInfo_URI `protobuf:"bytes,1,rep,name=uris" json:"uris,omitempty"`
	Environment    mesosproto.Environment        `protobuf:"bytes,2,opt,name=environment" json:"environment,omitempty"`
	Arguments      []string                      `protobuf:"bytes,7,rep,name=arguments" json:"arguments,omitempty"`
	Executor       mesosproto.ExecutorInfo
}
