// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/zk/enc/enc.proto

package zkenc

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	paillier "github.com/taurusgroup/cmp-ecdsa/pkg/paillier"
	github_com_taurusgroup_cmp_ecdsa_proto "github.com/taurusgroup/cmp-ecdsa/proto"
	io "io"
	math "math"
	math_big "math/big"
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

type Commitment struct {
	// S = sᵏtᵘ
	S *math_big.Int `protobuf:"bytes,1,opt,name=S,proto3,casttypewith=math/big.Int;github.com/taurusgroup/cmp-ecdsa/proto.IntCaster" json:"S,omitempty"`
	// A = Enc₀ (α, r)
	A *paillier.Ciphertext `protobuf:"bytes,2,opt,name=A,proto3" json:"A,omitempty"`
	// C = sᵃtᵍ
	C *math_big.Int `protobuf:"bytes,4,opt,name=C,proto3,casttypewith=math/big.Int;github.com/taurusgroup/cmp-ecdsa/proto.IntCaster" json:"C,omitempty"`
}

func (m *Commitment) Reset()         { *m = Commitment{} }
func (m *Commitment) String() string { return proto.CompactTextString(m) }
func (*Commitment) ProtoMessage()    {}
func (*Commitment) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1eddf36c8dd773b, []int{0}
}
func (m *Commitment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Commitment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Commitment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commitment.Merge(m, src)
}
func (m *Commitment) XXX_Size() int {
	return m.Size()
}
func (m *Commitment) XXX_DiscardUnknown() {
	xxx_messageInfo_Commitment.DiscardUnknown(m)
}

var xxx_messageInfo_Commitment proto.InternalMessageInfo

type Proof struct {
	*Commitment `protobuf:"bytes,1,opt,name=C,proto3,embedded=C" json:"C,omitempty"`
	// Z₁ = α + e⋅k
	Z1 *math_big.Int `protobuf:"bytes,2,opt,name=Z1,proto3,casttypewith=math/big.Int;github.com/taurusgroup/cmp-ecdsa/proto.IntCaster" json:"Z1,omitempty"`
	// Z₂ = r ⋅ ρᵉ mod N₀
	Z2 *math_big.Int `protobuf:"bytes,3,opt,name=Z2,proto3,casttypewith=math/big.Int;github.com/taurusgroup/cmp-ecdsa/proto.IntCaster" json:"Z2,omitempty"`
	// Z₃ = γ + e⋅μ
	Z3 *math_big.Int `protobuf:"bytes,4,opt,name=Z3,proto3,casttypewith=math/big.Int;github.com/taurusgroup/cmp-ecdsa/proto.IntCaster" json:"Z3,omitempty"`
}

func (m *Proof) Reset()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) ProtoMessage()    {}
func (*Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1eddf36c8dd773b, []int{1}
}
func (m *Proof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proof.Merge(m, src)
}
func (m *Proof) XXX_Size() int {
	return m.Size()
}
func (m *Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_Proof.DiscardUnknown(m)
}

var xxx_messageInfo_Proof proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Commitment)(nil), "zkenc.Commitment")
	proto.RegisterType((*Proof)(nil), "zkenc.Proof")
}

func init() { proto.RegisterFile("pkg/zk/enc/enc.proto", fileDescriptor_b1eddf36c8dd773b) }

var fileDescriptor_b1eddf36c8dd773b = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xbf, 0x6e, 0xea, 0x30,
	0x14, 0x87, 0x73, 0x72, 0xe1, 0x0e, 0x6e, 0x97, 0x46, 0x0c, 0x11, 0x52, 0x0d, 0x42, 0xaa, 0xc4,
	0x82, 0x2d, 0x60, 0x44, 0x1d, 0x20, 0x53, 0xa7, 0xb6, 0xb0, 0x65, 0x4b, 0x52, 0x13, 0x2c, 0x70,
	0x6c, 0x19, 0x47, 0xaa, 0x78, 0x8a, 0xbe, 0x45, 0xa5, 0x3e, 0x09, 0x23, 0xea, 0xc4, 0xd4, 0x3f,
	0xe1, 0x45, 0xaa, 0x18, 0x95, 0x76, 0xeb, 0x92, 0xc1, 0x92, 0x8f, 0x8e, 0xfd, 0xfd, 0x3e, 0x1f,
	0x19, 0x35, 0xd4, 0x32, 0xa5, 0x9b, 0x25, 0x65, 0x59, 0x52, 0x2e, 0xa2, 0xb4, 0x34, 0xd2, 0xab,
	0x6f, 0x96, 0x2c, 0x4b, 0x9a, 0xbd, 0x94, 0x9b, 0x45, 0x1e, 0x93, 0x44, 0x0a, 0x9a, 0xca, 0x54,
	0x52, 0xdb, 0x8d, 0xf3, 0xb9, 0xad, 0x6c, 0x61, 0x77, 0xc7, 0x5b, 0xcd, 0xcb, 0x92, 0xa5, 0x22,
	0xbe, 0x5a, 0x71, 0xa6, 0x69, 0xc2, 0xd5, 0x82, 0x69, 0xc3, 0x1e, 0xcd, 0xb1, 0xdd, 0x79, 0x05,
	0x84, 0x02, 0x29, 0x04, 0x37, 0x82, 0x65, 0xc6, 0xbb, 0x45, 0x30, 0xf3, 0xa1, 0x0d, 0xdd, 0xf3,
	0xc9, 0xf8, 0xe5, 0xbd, 0x75, 0x2d, 0x22, 0xb3, 0xa0, 0x31, 0x4f, 0xc9, 0x4d, 0x66, 0x46, 0xbf,
	0x82, 0x4d, 0x94, 0xeb, 0x7c, 0x9d, 0x6a, 0x99, 0x2b, 0x9a, 0x08, 0xd5, 0x63, 0xc9, 0xc3, 0x3a,
	0x3a, 0x9a, 0x94, 0x47, 0x83, 0x68, 0x6d, 0x98, 0x9e, 0xc2, 0xcc, 0xeb, 0x20, 0x18, 0xfb, 0x6e,
	0x1b, 0xba, 0x67, 0x83, 0x06, 0xf9, 0xd6, 0x20, 0xc1, 0x49, 0x63, 0x0a, 0xe3, 0x32, 0x34, 0xf0,
	0x6b, 0x95, 0x85, 0x06, 0x9d, 0x67, 0x17, 0xd5, 0xef, 0xb4, 0x94, 0x73, 0xef, 0xaa, 0x44, 0x83,
	0x8d, 0xbf, 0x20, 0x76, 0x7e, 0xe4, 0xe7, 0xb5, 0x93, 0xda, 0xee, 0xad, 0x05, 0x53, 0x08, 0xbc,
	0x7b, 0xe4, 0x86, 0x7d, 0xab, 0x59, 0x89, 0x82, 0x1b, 0xf6, 0x2d, 0x72, 0xe0, 0xff, 0xab, 0x0e,
	0x39, 0xb0, 0xc8, 0x61, 0x75, 0x83, 0x72, 0xc3, 0xe1, 0x64, 0xb6, 0xfd, 0xc4, 0xce, 0xb6, 0xc0,
	0xb0, 0x2b, 0x30, 0xec, 0x0b, 0x0c, 0x1f, 0x05, 0x86, 0xa7, 0x03, 0x76, 0x76, 0x07, 0xec, 0xec,
	0x0f, 0xd8, 0x09, 0xfb, 0x7f, 0x73, 0x4f, 0x9f, 0x75, 0x64, 0x27, 0x1c, 0xff, 0xb7, 0x49, 0xc3,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x62, 0xc5, 0x35, 0xc7, 0x02, 0x00, 0x00,
}

func (m *Commitment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Commitment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Commitment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
		size := __caster.Size(m.C)
		i -= size
		if _, err := __caster.MarshalTo(m.C, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEnc(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.A != nil {
		{
			size, err := m.A.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEnc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	{
		__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
		size := __caster.Size(m.S)
		i -= size
		if _, err := __caster.MarshalTo(m.S, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEnc(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Proof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
		size := __caster.Size(m.Z3)
		i -= size
		if _, err := __caster.MarshalTo(m.Z3, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEnc(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
		size := __caster.Size(m.Z2)
		i -= size
		if _, err := __caster.MarshalTo(m.Z2, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEnc(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
		size := __caster.Size(m.Z1)
		i -= size
		if _, err := __caster.MarshalTo(m.Z1, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEnc(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Commitment != nil {
		{
			size, err := m.Commitment.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEnc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEnc(dAtA []byte, offset int, v uint64) int {
	offset -= sovEnc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Commitment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	{
		__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
		l = __caster.Size(m.S)
		n += 1 + l + sovEnc(uint64(l))
	}
	if m.A != nil {
		l = m.A.Size()
		n += 1 + l + sovEnc(uint64(l))
	}
	{
		__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
		l = __caster.Size(m.C)
		n += 1 + l + sovEnc(uint64(l))
	}
	return n
}

func (m *Proof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Commitment != nil {
		l = m.Commitment.Size()
		n += 1 + l + sovEnc(uint64(l))
	}
	{
		__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
		l = __caster.Size(m.Z1)
		n += 1 + l + sovEnc(uint64(l))
	}
	{
		__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
		l = __caster.Size(m.Z2)
		n += 1 + l + sovEnc(uint64(l))
	}
	{
		__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
		l = __caster.Size(m.Z3)
		n += 1 + l + sovEnc(uint64(l))
	}
	return n
}

func sovEnc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEnc(x uint64) (n int) {
	return sovEnc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Commitment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnc
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
			return fmt.Errorf("proto: Commitment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Commitment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field S", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEnc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEnc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.S = tmp
				}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEnc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.A == nil {
				m.A = &paillier.Ciphertext{}
			}
			if err := m.A.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEnc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEnc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.C = tmp
				}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEnc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEnc
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
func (m *Proof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnc
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
			return fmt.Errorf("proto: Proof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commitment", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEnc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Commitment == nil {
				m.Commitment = &Commitment{}
			}
			if err := m.Commitment.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Z1", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEnc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEnc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Z1 = tmp
				}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Z2", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEnc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEnc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Z2 = tmp
				}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Z3", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEnc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEnc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Z3 = tmp
				}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEnc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEnc
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
func skipEnc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEnc
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
					return 0, ErrIntOverflowEnc
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
					return 0, ErrIntOverflowEnc
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
				return 0, ErrInvalidLengthEnc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEnc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEnc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEnc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEnc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEnc = fmt.Errorf("proto: unexpected end of group")
)
