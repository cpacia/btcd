package wire

import (
	"bytes"
	"fmt"
	"io"
)

// MsgCreateStream defines a bitcoin protoconf message which is used for...
type MsgCreateStream struct {
}

func (msg *MsgCreateStream) BtcDecode(r io.Reader, pver uint32) error {
	_, ok := r.(*bytes.Buffer)
	if !ok {
		return fmt.Errorf("MsgCreateStream.BtcDecode reader is not a " +
			"*bytes.Buffer")
	}
	return nil
}

func (msg *MsgCreateStream) BtcEncode(w io.Writer, pver uint32) error {
	return nil
}

func (msg *MsgCreateStream) Command() string {
	return CmdCreateStream
}

func (msg *MsgCreateStream) MaxPayloadLength(pver uint32) uint32 {
	return 0
}

func NewMsgCreateStream() *MsgCreateStream {
	return &MsgCreateStream{}
}
