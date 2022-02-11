package convert

import (
	"encoding/binary"
)

type Uint32 struct {
	Value uint32
}

// ToByte uint32 to byte
func (u Uint32) ToByte() []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, u.Value)
	return bs
}
