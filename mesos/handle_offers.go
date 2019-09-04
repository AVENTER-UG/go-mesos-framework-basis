package mesos

import (
	"fmt"
	"sync/atomic"

	"github.com/sirupsen/logrus"

	"../proto"
)

func defaultResources() []*mesosproto.Resource {
	CPU := "cpus"
	MEM := "mem"
	cpu := float64(0.1)

	return []*mesosproto.Resource{
		{
			Name:   &CPU,
			Type:   mesosproto.Value_SCALAR.Enum(),
			Scalar: &mesosproto.Value_Scalar{Value: &cpu},
		},
		{
			Name:   &MEM,
			Type:   mesosproto.Value_SCALAR.Enum(),
			Scalar: &mesosproto.Value_Scalar{Value: &cpu},
		},
	}
}

// HandleOffers will handle the offers event of mesos
func HandleOffers(offers *mesosproto.Event_Offers) error {
	offerIds := []*mesosproto.OfferID{}
	for _, offer := range offers.Offers {
		offerIds = append(offerIds, offer.Id)
	}

	select {
	case cmd := <-config.CommandChan:
		firstOffer := offers.Offers[0]

		newTaskID := fmt.Sprint(atomic.AddUint64(&config.TaskID, 1))
		taskInfo := []*mesosproto.TaskInfo{{
			Name: &cmd.Command,
			TaskId: &mesosproto.TaskID{
				Value: &newTaskID,
			},
			AgentId:   firstOffer.AgentId,
			Resources: defaultResources(),
			Command: &mesosproto.CommandInfo{
				Shell:       &cmd.Shell,
				Value:       &cmd.Command,
				Uris:        cmd.Uris,
				Environment: &cmd.Environment,
			}}}

		logrus.Debug("HandleOffers cmd: ", taskInfo)

		accept := &mesosproto.Call{
			Type: mesosproto.Call_ACCEPT.Enum(),
			Accept: &mesosproto.Call_Accept{
				OfferIds: offerIds,
				Operations: []*mesosproto.Offer_Operation{{
					Type: mesosproto.Offer_Operation_LAUNCH.Enum(),
					Launch: &mesosproto.Offer_Operation_Launch{
						TaskInfos: taskInfo,
					}}}}}
		return Call(accept)
	default:
		decline := &mesosproto.Call{
			Type:    mesosproto.Call_DECLINE.Enum(),
			Decline: &mesosproto.Call_Decline{OfferIds: offerIds},
		}
		return Call(decline)
	}
}
