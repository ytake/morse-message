package message

import (
	pbd "github.com/ytake/morse-message/publisher/pbdef"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/rand"
	"time"
)

func NoKeyMessages() ([]*pbd.UserAction, error) {
	var sua []*pbd.UserAction
	for i := 0; i < 100; i++ {
		t, err := time.Parse(time.RFC3339, "2020-09-01T21:46:43Z")
		if err != nil {
			return sua, err
		}
		sua = append(sua, &pbd.UserAction{
			UserId:        rand.Uint32(),
			CorrelationId: rand.Uint64(),
			Event:         randEventType(),
			Name:          "aaa",
			Created:       timestamppb.New(t),
		})
	}
	return sua, nil
}

func randEventType() pbd.UserAction_EventType {
	es := []pbd.UserAction_EventType{pbd.UserAction_CREATED, pbd.UserAction_DELETED}
	rand.Seed(time.Now().UnixNano())
	return es[rand.Intn(len(es))]
}
