// Code generated by protoc-gen-go.
// source: messages.proto
// DO NOT EDIT!

/*
Package dht is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	PBDHTMessage
*/
package dht

import proto "github.com/jbenet/go-ipfs/Godeps/_workspace/src/code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type PBDHTMessage_MessageType int32

const (
	PBDHTMessage_PUT_VALUE     PBDHTMessage_MessageType = 0
	PBDHTMessage_GET_VALUE     PBDHTMessage_MessageType = 1
	PBDHTMessage_ADD_PROVIDER  PBDHTMessage_MessageType = 2
	PBDHTMessage_GET_PROVIDERS PBDHTMessage_MessageType = 3
	PBDHTMessage_FIND_NODE     PBDHTMessage_MessageType = 4
	PBDHTMessage_PING          PBDHTMessage_MessageType = 5
	PBDHTMessage_DIAGNOSTIC    PBDHTMessage_MessageType = 6
)

var PBDHTMessage_MessageType_name = map[int32]string{
	0: "PUT_VALUE",
	1: "GET_VALUE",
	2: "ADD_PROVIDER",
	3: "GET_PROVIDERS",
	4: "FIND_NODE",
	5: "PING",
	6: "DIAGNOSTIC",
}
var PBDHTMessage_MessageType_value = map[string]int32{
	"PUT_VALUE":     0,
	"GET_VALUE":     1,
	"ADD_PROVIDER":  2,
	"GET_PROVIDERS": 3,
	"FIND_NODE":     4,
	"PING":          5,
	"DIAGNOSTIC":    6,
}

func (x PBDHTMessage_MessageType) Enum() *PBDHTMessage_MessageType {
	p := new(PBDHTMessage_MessageType)
	*p = x
	return p
}
func (x PBDHTMessage_MessageType) String() string {
	return proto.EnumName(PBDHTMessage_MessageType_name, int32(x))
}
func (x *PBDHTMessage_MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PBDHTMessage_MessageType_value, data, "PBDHTMessage_MessageType")
	if err != nil {
		return err
	}
	*x = PBDHTMessage_MessageType(value)
	return nil
}

type PBDHTMessage struct {
	Type             *PBDHTMessage_MessageType `protobuf:"varint,1,req,name=type,enum=dht.PBDHTMessage_MessageType" json:"type,omitempty"`
	Key              *string                   `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value            []byte                    `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	Id               *string                   `protobuf:"bytes,4,req,name=id" json:"id,omitempty"`
	Response         *bool                     `protobuf:"varint,5,opt,name=response" json:"response,omitempty"`
	Success          *bool                     `protobuf:"varint,6,opt,name=success" json:"success,omitempty"`
	Peers            []*PBDHTMessage_PBPeer    `protobuf:"bytes,7,rep,name=peers" json:"peers,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *PBDHTMessage) Reset()         { *m = PBDHTMessage{} }
func (m *PBDHTMessage) String() string { return proto.CompactTextString(m) }
func (*PBDHTMessage) ProtoMessage()    {}

func (m *PBDHTMessage) GetType() PBDHTMessage_MessageType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return PBDHTMessage_PUT_VALUE
}

func (m *PBDHTMessage) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *PBDHTMessage) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *PBDHTMessage) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *PBDHTMessage) GetResponse() bool {
	if m != nil && m.Response != nil {
		return *m.Response
	}
	return false
}

func (m *PBDHTMessage) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *PBDHTMessage) GetPeers() []*PBDHTMessage_PBPeer {
	if m != nil {
		return m.Peers
	}
	return nil
}

type PBDHTMessage_PBPeer struct {
	Id               *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Addr             *string `protobuf:"bytes,2,req,name=addr" json:"addr,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PBDHTMessage_PBPeer) Reset()         { *m = PBDHTMessage_PBPeer{} }
func (m *PBDHTMessage_PBPeer) String() string { return proto.CompactTextString(m) }
func (*PBDHTMessage_PBPeer) ProtoMessage()    {}

func (m *PBDHTMessage_PBPeer) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *PBDHTMessage_PBPeer) GetAddr() string {
	if m != nil && m.Addr != nil {
		return *m.Addr
	}
	return ""
}

func init() {
	proto.RegisterEnum("dht.PBDHTMessage_MessageType", PBDHTMessage_MessageType_name, PBDHTMessage_MessageType_value)
}
