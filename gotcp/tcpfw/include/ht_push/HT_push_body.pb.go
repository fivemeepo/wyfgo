// Code generated by protoc-gen-go.
// source: HT_push_body.proto
// DO NOT EDIT!

package ht_push

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 呼叫
type NotifyRemarkNameReqBody struct {
	Remarkname       []byte `protobuf:"bytes,1,opt,name=remarkname" json:"remarkname,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *NotifyRemarkNameReqBody) Reset()                    { *m = NotifyRemarkNameReqBody{} }
func (m *NotifyRemarkNameReqBody) String() string            { return proto.CompactTextString(m) }
func (*NotifyRemarkNameReqBody) ProtoMessage()               {}
func (*NotifyRemarkNameReqBody) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *NotifyRemarkNameReqBody) GetRemarkname() []byte {
	if m != nil {
		return m.Remarkname
	}
	return nil
}

// 不需要回包
type NotifyRemarkNameRspBody struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *NotifyRemarkNameRspBody) Reset()                    { *m = NotifyRemarkNameRspBody{} }
func (m *NotifyRemarkNameRspBody) String() string            { return proto.CompactTextString(m) }
func (*NotifyRemarkNameRspBody) ProtoMessage()               {}
func (*NotifyRemarkNameRspBody) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type PushMsgReqBody struct {
	// optional uint32 term_type			= 1;//终端类型 0-Ios, 1-Android
	ChatType         *uint32 `protobuf:"varint,2,opt,name=chat_type" json:"chat_type,omitempty"`
	RoomId           *uint32 `protobuf:"varint,3,opt,name=room_id" json:"room_id,omitempty"`
	PushType         *uint32 `protobuf:"varint,4,opt,name=push_type" json:"push_type,omitempty"`
	Nickname         []byte  `protobuf:"bytes,5,opt,name=nickname" json:"nickname,omitempty"`
	Content          []byte  `protobuf:"bytes,6,opt,name=content" json:"content,omitempty"`
	Sound            *uint32 `protobuf:"varint,7,opt,name=sound" json:"sound,omitempty"`
	Light            *uint32 `protobuf:"varint,8,opt,name=light" json:"light,omitempty"`
	MessageId        []byte  `protobuf:"bytes,9,opt,name=message_id" json:"message_id,omitempty"`
	ActionId         *uint32 `protobuf:"varint,10,opt,name=action_id" json:"action_id,omitempty"`
	NeedRsp          *uint32 `protobuf:"varint,50,opt,name=need_rsp" json:"need_rsp,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PushMsgReqBody) Reset()                    { *m = PushMsgReqBody{} }
func (m *PushMsgReqBody) String() string            { return proto.CompactTextString(m) }
func (*PushMsgReqBody) ProtoMessage()               {}
func (*PushMsgReqBody) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *PushMsgReqBody) GetChatType() uint32 {
	if m != nil && m.ChatType != nil {
		return *m.ChatType
	}
	return 0
}

func (m *PushMsgReqBody) GetRoomId() uint32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *PushMsgReqBody) GetPushType() uint32 {
	if m != nil && m.PushType != nil {
		return *m.PushType
	}
	return 0
}

func (m *PushMsgReqBody) GetNickname() []byte {
	if m != nil {
		return m.Nickname
	}
	return nil
}

func (m *PushMsgReqBody) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *PushMsgReqBody) GetSound() uint32 {
	if m != nil && m.Sound != nil {
		return *m.Sound
	}
	return 0
}

func (m *PushMsgReqBody) GetLight() uint32 {
	if m != nil && m.Light != nil {
		return *m.Light
	}
	return 0
}

func (m *PushMsgReqBody) GetMessageId() []byte {
	if m != nil {
		return m.MessageId
	}
	return nil
}

func (m *PushMsgReqBody) GetActionId() uint32 {
	if m != nil && m.ActionId != nil {
		return *m.ActionId
	}
	return 0
}

func (m *PushMsgReqBody) GetNeedRsp() uint32 {
	if m != nil && m.NeedRsp != nil {
		return *m.NeedRsp
	}
	return 0
}

type PushMsgRspBody struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *PushMsgRspBody) Reset()                    { *m = PushMsgRspBody{} }
func (m *PushMsgRspBody) String() string            { return proto.CompactTextString(m) }
func (*PushMsgRspBody) ProtoMessage()               {}
func (*PushMsgRspBody) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func init() {
	proto.RegisterType((*NotifyRemarkNameReqBody)(nil), "ht.push.NotifyRemarkNameReqBody")
	proto.RegisterType((*NotifyRemarkNameRspBody)(nil), "ht.push.NotifyRemarkNameRspBody")
	proto.RegisterType((*PushMsgReqBody)(nil), "ht.push.PushMsgReqBody")
	proto.RegisterType((*PushMsgRspBody)(nil), "ht.push.PushMsgRspBody")
}

func init() { proto.RegisterFile("HT_push_body.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x15, 0xa0, 0xa4, 0x3d, 0x51, 0x28, 0x5e, 0x30, 0x1b, 0xca, 0xc4, 0x42, 0x07, 0x1e,
	0x81, 0x89, 0x85, 0x0a, 0x55, 0xec, 0x96, 0x49, 0x8e, 0xc4, 0x82, 0xf8, 0x8c, 0x7d, 0x1d, 0xf2,
	0x78, 0xbc, 0x19, 0xf6, 0x45, 0xc0, 0xc2, 0x78, 0x9f, 0xfe, 0xfb, 0xbf, 0x3b, 0x50, 0x8f, 0x2f,
	0x26, 0x1c, 0xd2, 0x60, 0x5e, 0xa9, 0x9b, 0xb6, 0x21, 0x12, 0x93, 0xaa, 0x07, 0xde, 0x16, 0xd6,
	0xdc, 0xc1, 0xd5, 0x8e, 0xd8, 0xbd, 0x4d, 0x7b, 0x1c, 0x6d, 0x7c, 0xdf, 0xd9, 0x11, 0xf7, 0xf8,
	0xf9, 0x90, 0x93, 0x4a, 0x01, 0x44, 0x81, 0x3e, 0x43, 0x5d, 0xdd, 0x54, 0xb7, 0x67, 0xcd, 0xf5,
	0x3f, 0xf1, 0x14, 0x4a, 0xbc, 0xf9, 0xaa, 0xe0, 0xfc, 0x39, 0x57, 0x3e, 0xa5, 0xfe, 0xa7, 0xe1,
	0x12, 0x56, 0xed, 0x60, 0xd9, 0xf0, 0x14, 0x50, 0x1f, 0xe5, 0x82, 0xb5, 0xba, 0x80, 0x3a, 0x12,
	0x8d, 0xc6, 0x75, 0xfa, 0x58, 0x40, 0xce, 0xc8, 0x71, 0x92, 0x39, 0x11, 0xb4, 0x81, 0xa5, 0x77,
	0xed, 0xac, 0x5d, 0x14, 0x6d, 0xd9, 0x6a, 0xc9, 0x33, 0x7a, 0xd6, 0xa7, 0x02, 0xd6, 0xb0, 0x48,
	0x74, 0xf0, 0x9d, 0xae, 0x65, 0x23, 0x8f, 0x1f, 0xae, 0x1f, 0x58, 0x2f, 0x65, 0xcc, 0x97, 0x8f,
	0x98, 0x92, 0xed, 0xb1, 0x78, 0x56, 0xb2, 0x91, 0x3d, 0xb6, 0x65, 0x47, 0xbe, 0x20, 0xf8, 0xf5,
	0x20, 0x76, 0x26, 0xa6, 0xa0, 0xef, 0x0b, 0x69, 0x36, 0x7f, 0x2f, 0xcc, 0x5f, 0x7d, 0x07, 0x00,
	0x00, 0xff, 0xff, 0xb9, 0xe3, 0x31, 0xa1, 0x3d, 0x01, 0x00, 0x00,
}
