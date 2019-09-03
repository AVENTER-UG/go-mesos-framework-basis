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
	CommandChan       chan string
}
