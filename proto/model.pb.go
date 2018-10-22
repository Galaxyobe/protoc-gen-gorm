// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: model.proto

package gorm

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Model struct {
	ID        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primary_key"`
	CreatedAt int64  `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt int64  `protobuf:"varint,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt int64  `protobuf:"varint,4,opt,name=deletedAt,proto3" json:"deletedAt,omitempty" sql:"index"`
}

func (m *Model) Reset()         { *m = Model{} }
func (m *Model) String() string { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()    {}
func (*Model) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_8ec8256ddd8b28f9, []int{0}
}
func (m *Model) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Model.Unmarshal(m, b)
}
func (m *Model) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Model.Marshal(b, m, deterministic)
}
func (dst *Model) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Model.Merge(dst, src)
}
func (m *Model) XXX_Size() int {
	return xxx_messageInfo_Model.Size(m)
}
func (m *Model) XXX_DiscardUnknown() {
	xxx_messageInfo_Model.DiscardUnknown(m)
}

var xxx_messageInfo_Model proto.InternalMessageInfo

func (m *Model) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Model) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Model) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *Model) GetDeletedAt() int64 {
	if m != nil {
		return m.DeletedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*Model)(nil), "gorm.Model")
}
func (this *Model) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*Model)
	if !ok {
		that2, ok := that.(Model)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if this.ID != that1.ID {
		if this.ID < that1.ID {
			return -1
		}
		return 1
	}
	if this.CreatedAt != that1.CreatedAt {
		if this.CreatedAt < that1.CreatedAt {
			return -1
		}
		return 1
	}
	if this.UpdatedAt != that1.UpdatedAt {
		if this.UpdatedAt < that1.UpdatedAt {
			return -1
		}
		return 1
	}
	if this.DeletedAt != that1.DeletedAt {
		if this.DeletedAt < that1.DeletedAt {
			return -1
		}
		return 1
	}
	return 0
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_model_8ec8256ddd8b28f9) }

var fileDescriptor_model_8ec8256ddd8b28f9 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xcd, 0x4f, 0x49,
	0xcd, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0xcf, 0x2f, 0xca, 0x95, 0xd2, 0x4d,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07,
	0x4b, 0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa4, 0xb4, 0x89, 0x91, 0x8b,
	0xd5, 0x17, 0x64, 0x88, 0x90, 0x0e, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x8b,
	0x93, 0xcc, 0xa3, 0x7b, 0xf2, 0x4c, 0x9e, 0x2e, 0x9f, 0xee, 0xc9, 0x0b, 0x81, 0x0c, 0xb5, 0x52,
	0x2a, 0x28, 0xca, 0xcc, 0x4d, 0x2c, 0xaa, 0x8c, 0xcf, 0x4e, 0xad, 0x54, 0x0a, 0x62, 0xca, 0x4c,
	0x11, 0x92, 0xe1, 0xe2, 0x4c, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x4d, 0x71, 0x2c, 0x91, 0x60, 0x52,
	0x60, 0xd4, 0x60, 0x0e, 0x42, 0x08, 0x80, 0x64, 0x4b, 0x0b, 0x52, 0xa0, 0xb2, 0xcc, 0x10, 0x59,
	0xb8, 0x80, 0x90, 0x2e, 0x17, 0x67, 0x4a, 0x6a, 0x4e, 0x2a, 0x44, 0x96, 0x05, 0x24, 0xeb, 0xc4,
	0xff, 0xe9, 0x9e, 0x3c, 0x77, 0x71, 0x61, 0x8e, 0x95, 0x52, 0x66, 0x5e, 0x4a, 0x6a, 0x85, 0x52,
	0x10, 0x42, 0x85, 0x95, 0xc0, 0x85, 0x85, 0xf2, 0x0c, 0x2f, 0x16, 0xca, 0x33, 0x4e, 0x58, 0x24,
	0xcf, 0x30, 0x63, 0x91, 0x3c, 0x43, 0x12, 0x1b, 0xd8, 0xed, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x98, 0x91, 0xe7, 0x02, 0xff, 0x00, 0x00, 0x00,
}