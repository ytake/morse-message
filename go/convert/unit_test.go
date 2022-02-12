package convert

import (
	"encoding/binary"
	"testing"
)

func TestUint32_ToByte(t *testing.T) {
	u := Uint32{Value: uint32(12345)}
	i := binary.LittleEndian.Uint32(u.ToByte())
	if i != uint32(12345) {
		t.Error("mismatched value")
	}
}
