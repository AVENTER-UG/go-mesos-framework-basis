package mesos

import (
	"github.com/sirupsen/logrus"

	"../proto"
)

// HandleUpdate will handle the offers event of mesos
func HandleUpdate(event *mesosproto.Event_Update) error {

	logrus.Debug("HandleUpate cmd: ", event)

	return nil

}
