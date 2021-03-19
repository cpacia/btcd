package wire

import (
	"bytes"
	"fmt"
	"io"
)

// MsgStreamAck defines a bitcoin protoconf message which is used for...
type MsgStreamAck struct {
}

func (msg *MsgStreamAck) BtcDecode(r io.Reader, pver uint32) error {
	_, ok := r.(*bytes.Buffer)
	if !ok {
		return fmt.Errorf("MsgStreamAck.BtcDecode reader is not a " +
			"*bytes.Buffer")
	}
	return nil
}

func (msg *MsgStreamAck) BtcEncode(w io.Writer, pver uint32) error {
	return nil
}

func (msg *MsgStreamAck) Command() string {
	return CmdStreamAck
}

func (msg *MsgStreamAck) MaxPayloadLength(pver uint32) uint32 {
	return 0
}

func NewMsgStreamAck() *MsgStreamAck {
	return &MsgStreamAck{}
}
