package wire

import (
	"bytes"
	"fmt"
	"io"
)

// MsgProtoConf defines a bitcoin protoconf message which is used for...
type MsgProtoConf struct {
	NumberOfFields         uint8
	MaxRecvPayloadLengthIn uint32
}

func (msg *MsgProtoConf) BtcDecode(r io.Reader, pver uint32) error {
	buf, ok := r.(*bytes.Buffer)
	if !ok {
		return fmt.Errorf("MsgProtoConf.BtcDecode reader is not a " +
			"*bytes.Buffer")
	}
	err := readElement(buf, &msg.NumberOfFields)
	if err != nil {
		return err
	}
	if msg.NumberOfFields >= 1 {
		err = readElement(buf, &msg.MaxRecvPayloadLengthIn)
		if err != nil {
			return err
		}
	}
	return nil
}

func (msg *MsgProtoConf) BtcEncode(w io.Writer, pver uint32) error {
	return nil
}

func (msg *MsgProtoConf) Command() string {
	return CmdProtoConf
}

func (msg *MsgProtoConf) MaxPayloadLength(pver uint32) uint32 {
	return 27
}

func NewMsgProtoConf() *MsgProtoConf {
	return &MsgProtoConf{}
}
