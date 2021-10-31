package message

import (
	"fmt"
	pbd "github.com/ytake/morse-message/publisher/pbdef"
	"google.golang.org/protobuf/proto"
)

type NoKeyMessagesReceive struct {
}

func (nk *NoKeyMessagesReceive) Proceed(message []byte) error {
	u := &pbd.UserAction{}
	if err := proto.Unmarshal(message, u); err != nil {
		return err
	}
	fmt.Println(u)
	return nil
}
