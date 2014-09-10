// Code generated by protoc-gen-go.
// source: entry.proto
// DO NOT EDIT!

/*
Package namesys is a generated protocol buffer package.

It is generated from these files:
	entry.proto

It has these top-level messages:
	InpsEntry
*/
package namesys

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type InpsEntry struct {
	Value            []byte `protobuf:"bytes,1,req,name=value" json:"value,omitempty"`
	Signature        []byte `protobuf:"bytes,2,req,name=signature" json:"signature,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *InpsEntry) Reset()         { *m = InpsEntry{} }
func (m *InpsEntry) String() string { return proto.CompactTextString(m) }
func (*InpsEntry) ProtoMessage()    {}

func (m *InpsEntry) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *InpsEntry) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
}