package command

import (
	pbd "github.com/ytake/morse-message/publisher/pbdef"
	"github.com/ytake/morse-message/publisher/pub"
	"google.golang.org/protobuf/proto"
)

type UserAction struct {
	Client *pub.Messenger
}

func (u *UserAction) UserRegistrationForNoKey(messages []*pbd.UserAction) error {
	for _, v := range messages {
		o, err := proto.Marshal(v)
		if err != nil {
			return err
		}
		if err = u.Client.Publish(pub.RequestParameter{Byte: o}); err != nil {
			return err
		}
	}
	defer u.Client.Close()
	return nil
}
