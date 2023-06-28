// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: saw/citizen/citizen.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Citizen struct {
	SaId         string `protobuf:"bytes,1,opt,name=saId,proto3" json:"saId,omitempty"`
	CitizenCode  string `protobuf:"bytes,2,opt,name=citizenCode,proto3" json:"citizenCode,omitempty"`
	AddressOwner string `protobuf:"bytes,3,opt,name=addressOwner,proto3" json:"addressOwner,omitempty"`
	Name         string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	CreateAt     uint64 `protobuf:"varint,5,opt,name=createAt,proto3" json:"createAt,omitempty"`
	UpdateAt     uint64 `protobuf:"varint,6,opt,name=updateAt,proto3" json:"updateAt,omitempty"`
	AvatarUrl    string `protobuf:"bytes,7,opt,name=avatarUrl,proto3" json:"avatarUrl,omitempty"`
}

func (m *Citizen) Reset()         { *m = Citizen{} }
func (m *Citizen) String() string { return proto.CompactTextString(m) }
func (*Citizen) ProtoMessage()    {}
func (*Citizen) Descriptor() ([]byte, []int) {
	return fileDescriptor_37a5a67f642dc391, []int{0}
}
func (m *Citizen) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Citizen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Citizen.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Citizen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Citizen.Merge(m, src)
}
func (m *Citizen) XXX_Size() int {
	return m.Size()
}
func (m *Citizen) XXX_DiscardUnknown() {
	xxx_messageInfo_Citizen.DiscardUnknown(m)
}

var xxx_messageInfo_Citizen proto.InternalMessageInfo

func (m *Citizen) GetSaId() string {
	if m != nil {
		return m.SaId
	}
	return ""
}

func (m *Citizen) GetCitizenCode() string {
	if m != nil {
		return m.CitizenCode
	}
	return ""
}

func (m *Citizen) GetAddressOwner() string {
	if m != nil {
		return m.AddressOwner
	}
	return ""
}

func (m *Citizen) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Citizen) GetCreateAt() uint64 {
	if m != nil {
		return m.CreateAt
	}
	return 0
}

func (m *Citizen) GetUpdateAt() uint64 {
	if m != nil {
		return m.UpdateAt
	}
	return 0
}

func (m *Citizen) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*Citizen)(nil), "saw.citizen.Citizen")
}

func init() { proto.RegisterFile("saw/citizen/citizen.proto", fileDescriptor_37a5a67f642dc391) }

var fileDescriptor_37a5a67f642dc391 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x4e, 0x2c, 0xd7,
	0x4f, 0xce, 0x2c, 0xc9, 0xac, 0x4a, 0xcd, 0x83, 0xd1, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0xdc, 0xc5, 0x89, 0xe5, 0x7a, 0x50, 0x21, 0xa5, 0xb3, 0x8c, 0x5c, 0xec, 0xce, 0x10, 0xb6, 0x90,
	0x10, 0x17, 0x4b, 0x71, 0xa2, 0x67, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x2d,
	0xa4, 0xc0, 0xc5, 0x0d, 0x55, 0xea, 0x9c, 0x9f, 0x92, 0x2a, 0xc1, 0x04, 0x96, 0x42, 0x16, 0x12,
	0x52, 0xe2, 0xe2, 0x49, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0xf6, 0x2f, 0xcf, 0x4b, 0x2d, 0x92,
	0x60, 0x06, 0x2b, 0x41, 0x11, 0x03, 0x99, 0x9c, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0x02, 0x31, 0x19,
	0xc4, 0x16, 0x92, 0xe2, 0xe2, 0x48, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x75, 0x2c, 0x91, 0x60, 0x55,
	0x60, 0xd4, 0x60, 0x09, 0x82, 0xf3, 0x41, 0x72, 0xa5, 0x05, 0x29, 0x10, 0x39, 0x36, 0x88, 0x1c,
	0x8c, 0x2f, 0x24, 0xc3, 0xc5, 0x99, 0x58, 0x96, 0x58, 0x92, 0x58, 0x14, 0x5a, 0x94, 0x23, 0xc1,
	0x0e, 0x36, 0x10, 0x21, 0xe0, 0xe4, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f,
	0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c,
	0x51, 0xea, 0xe9, 0x99, 0x25, 0x39, 0x89, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xf9, 0x79, 0xa9,
	0xc9, 0x19, 0x89, 0x99, 0x79, 0xfa, 0xa0, 0x50, 0xaa, 0x80, 0x87, 0x53, 0x49, 0x65, 0x41, 0x6a,
	0x71, 0x12, 0x1b, 0x38, 0x98, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xcd, 0x2f, 0x73,
	0x43, 0x01, 0x00, 0x00,
}

func (m *Citizen) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Citizen) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Citizen) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AvatarUrl) > 0 {
		i -= len(m.AvatarUrl)
		copy(dAtA[i:], m.AvatarUrl)
		i = encodeVarintCitizen(dAtA, i, uint64(len(m.AvatarUrl)))
		i--
		dAtA[i] = 0x3a
	}
	if m.UpdateAt != 0 {
		i = encodeVarintCitizen(dAtA, i, uint64(m.UpdateAt))
		i--
		dAtA[i] = 0x30
	}
	if m.CreateAt != 0 {
		i = encodeVarintCitizen(dAtA, i, uint64(m.CreateAt))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintCitizen(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AddressOwner) > 0 {
		i -= len(m.AddressOwner)
		copy(dAtA[i:], m.AddressOwner)
		i = encodeVarintCitizen(dAtA, i, uint64(len(m.AddressOwner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.CitizenCode) > 0 {
		i -= len(m.CitizenCode)
		copy(dAtA[i:], m.CitizenCode)
		i = encodeVarintCitizen(dAtA, i, uint64(len(m.CitizenCode)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SaId) > 0 {
		i -= len(m.SaId)
		copy(dAtA[i:], m.SaId)
		i = encodeVarintCitizen(dAtA, i, uint64(len(m.SaId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCitizen(dAtA []byte, offset int, v uint64) int {
	offset -= sovCitizen(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Citizen) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SaId)
	if l > 0 {
		n += 1 + l + sovCitizen(uint64(l))
	}
	l = len(m.CitizenCode)
	if l > 0 {
		n += 1 + l + sovCitizen(uint64(l))
	}
	l = len(m.AddressOwner)
	if l > 0 {
		n += 1 + l + sovCitizen(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCitizen(uint64(l))
	}
	if m.CreateAt != 0 {
		n += 1 + sovCitizen(uint64(m.CreateAt))
	}
	if m.UpdateAt != 0 {
		n += 1 + sovCitizen(uint64(m.UpdateAt))
	}
	l = len(m.AvatarUrl)
	if l > 0 {
		n += 1 + l + sovCitizen(uint64(l))
	}
	return n
}

func sovCitizen(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCitizen(x uint64) (n int) {
	return sovCitizen(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Citizen) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCitizen
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Citizen: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Citizen: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SaId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCitizen
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCitizen
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCitizen
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SaId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CitizenCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCitizen
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCitizen
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCitizen
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CitizenCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCitizen
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCitizen
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCitizen
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AddressOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCitizen
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCitizen
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCitizen
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateAt", wireType)
			}
			m.CreateAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCitizen
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreateAt |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateAt", wireType)
			}
			m.UpdateAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCitizen
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdateAt |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AvatarUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCitizen
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCitizen
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCitizen
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AvatarUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCitizen(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCitizen
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCitizen(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCitizen
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCitizen
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCitizen
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCitizen
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCitizen
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCitizen
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCitizen        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCitizen          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCitizen = fmt.Errorf("proto: unexpected end of group")
)
