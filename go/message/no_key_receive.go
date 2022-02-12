package message

import (
	"fmt"
	pbd "github.com/ytake/morse-message/publisher/pbdef"
	"google.golang.org/protobuf/proto"
)

type MessagesReceive struct {
}

func (nk *MessagesReceive) Proceed(message []byte) error {
	u := &pbd.UserAction{}
	if err := proto.Unmarshal(message, u); err != nil {
		return err
	}
	fmt.Printf(
		"Event: %v\nName: %v\nUserID: %v\nCorrelationID %v",
		u.GetEvent(),
		u.GetName(),
		u.GetUserId(),
		u.GetCorrelationId())
	return nil
}
