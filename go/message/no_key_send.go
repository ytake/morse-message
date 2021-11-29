package message

import (
	pbd "github.com/ytake/morse-message/publisher/pbdef"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/rand"
	"time"
)

func makeTimestampForProto() *timestamppb.Timestamp {
	t, _ := time.Parse(time.RFC3339, "2020-09-01T21:46:43Z")
	return timestamppb.New(t)
}

// NoKeyMessages key指定なしサンプルイベントスライス
func NoKeyMessages() ([]*pbd.UserAction, error) {
	var sua []*pbd.UserAction
	sua = append(sua, &pbd.UserAction{
		UserId:  uint32(1),
		Event:   pbd.UserAction_CREATED,
		Name:    "aaa1",
		Created: makeTimestampForProto(),
	}, &pbd.UserAction{
		UserId:  uint32(2),
		Event:   pbd.UserAction_CREATED,
		Name:    "aaa2",
		Created: makeTimestampForProto(),
	}, &pbd.UserAction{
		UserId:  uint32(1),
		Event:   pbd.UserAction_DELETED,
		Name:    "aaa1",
		Created: makeTimestampForProto(),
	}, &pbd.UserAction{
		UserId:  uint32(3),
		Event:   pbd.UserAction_CREATED,
		Name:    "aaa3",
		Created: makeTimestampForProto(),
	}, &pbd.UserAction{
		UserId:  uint32(6),
		Event:   pbd.UserAction_CREATED,
		Name:    "aaa6",
		Created: makeTimestampForProto(),
	}, &pbd.UserAction{
		UserId:  uint32(18),
		Event:   pbd.UserAction_CREATED,
		Name:    "aaa18",
		Created: makeTimestampForProto(),
	}, &pbd.UserAction{
		UserId:  uint32(90),
		Event:   pbd.UserAction_CREATED,
		Name:    "aaa90",
		Created: makeTimestampForProto(),
	}, &pbd.UserAction{
		UserId:  uint32(93),
		Event:   pbd.UserAction_CREATED,
		Name:    "aaa93",
		Created: makeTimestampForProto(),
	}, &pbd.UserAction{
		UserId:  uint32(90),
		Event:   pbd.UserAction_DELETED,
		Name:    "aaa90",
		Created: makeTimestampForProto(),
	},&pbd.UserAction{
		UserId:  uint32(100),
		Event:   pbd.UserAction_CREATED,
		Name:    "aaa100",
		Created: makeTimestampForProto(),
	},
	)
	return sua, nil
}

func randEventType() pbd.UserAction_EventType {
	es := []pbd.UserAction_EventType{pbd.UserAction_CREATED, pbd.UserAction_DELETED}
	rand.Seed(time.Now().UnixNano())
	return es[rand.Intn(len(es))]
}
