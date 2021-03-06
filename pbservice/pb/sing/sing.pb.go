// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sing.proto

package sing

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SingInput struct {
	Query         string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageNumber    int64  `protobuf:"zigzag64,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	ResultPerPage int64  `protobuf:"zigzag64,3,opt,name=result_per_page,json=resultPerPage,proto3" json:"result_per_page,omitempty"`
}

func (m *SingInput) Reset()         { *m = SingInput{} }
func (m *SingInput) String() string { return proto.CompactTextString(m) }
func (*SingInput) ProtoMessage()    {}
func (*SingInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_sing_de4ff7d2969bf8fe, []int{0}
}
func (m *SingInput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SingInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SingInput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SingInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingInput.Merge(dst, src)
}
func (m *SingInput) XXX_Size() int {
	return m.Size()
}
func (m *SingInput) XXX_DiscardUnknown() {
	xxx_messageInfo_SingInput.DiscardUnknown(m)
}

var xxx_messageInfo_SingInput proto.InternalMessageInfo

func (m *SingInput) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SingInput) GetPageNumber() int64 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *SingInput) GetResultPerPage() int64 {
	if m != nil {
		return m.ResultPerPage
	}
	return 0
}

type SingOutput struct {
	Url      string        `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Title    string        `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Snippets []string      `protobuf:"bytes,3,rep,name=snippets" json:"snippets,omitempty"`
	Results  []*SingResult `protobuf:"bytes,4,rep,name=results" json:"results,omitempty"`
}

func (m *SingOutput) Reset()         { *m = SingOutput{} }
func (m *SingOutput) String() string { return proto.CompactTextString(m) }
func (*SingOutput) ProtoMessage()    {}
func (*SingOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_sing_de4ff7d2969bf8fe, []int{1}
}
func (m *SingOutput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SingOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SingOutput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SingOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingOutput.Merge(dst, src)
}
func (m *SingOutput) XXX_Size() int {
	return m.Size()
}
func (m *SingOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_SingOutput.DiscardUnknown(m)
}

var xxx_messageInfo_SingOutput proto.InternalMessageInfo

func (m *SingOutput) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SingOutput) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SingOutput) GetSnippets() []string {
	if m != nil {
		return m.Snippets
	}
	return nil
}

func (m *SingOutput) GetResults() []*SingResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type SingResult struct {
	Url      string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Title    string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Snippets []string `protobuf:"bytes,3,rep,name=snippets" json:"snippets,omitempty"`
}

func (m *SingResult) Reset()         { *m = SingResult{} }
func (m *SingResult) String() string { return proto.CompactTextString(m) }
func (*SingResult) ProtoMessage()    {}
func (*SingResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_sing_de4ff7d2969bf8fe, []int{2}
}
func (m *SingResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SingResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SingResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SingResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingResult.Merge(dst, src)
}
func (m *SingResult) XXX_Size() int {
	return m.Size()
}
func (m *SingResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SingResult.DiscardUnknown(m)
}

var xxx_messageInfo_SingResult proto.InternalMessageInfo

func (m *SingResult) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SingResult) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SingResult) GetSnippets() []string {
	if m != nil {
		return m.Snippets
	}
	return nil
}

func init() {
	proto.RegisterType((*SingInput)(nil), "sing.SingInput")
	proto.RegisterType((*SingOutput)(nil), "sing.SingOutput")
	proto.RegisterType((*SingResult)(nil), "sing.SingResult")
}
func (m *SingInput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SingInput) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Query) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSing(dAtA, i, uint64(len(m.Query)))
		i += copy(dAtA[i:], m.Query)
	}
	if m.PageNumber != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSing(dAtA, i, uint64((uint64(m.PageNumber)<<1)^uint64((m.PageNumber>>63))))
	}
	if m.ResultPerPage != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintSing(dAtA, i, uint64((uint64(m.ResultPerPage)<<1)^uint64((m.ResultPerPage>>63))))
	}
	return i, nil
}

func (m *SingOutput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SingOutput) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Url) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSing(dAtA, i, uint64(len(m.Url)))
		i += copy(dAtA[i:], m.Url)
	}
	if len(m.Title) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSing(dAtA, i, uint64(len(m.Title)))
		i += copy(dAtA[i:], m.Title)
	}
	if len(m.Snippets) > 0 {
		for _, s := range m.Snippets {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Results) > 0 {
		for _, msg := range m.Results {
			dAtA[i] = 0x22
			i++
			i = encodeVarintSing(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *SingResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SingResult) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Url) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSing(dAtA, i, uint64(len(m.Url)))
		i += copy(dAtA[i:], m.Url)
	}
	if len(m.Title) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSing(dAtA, i, uint64(len(m.Title)))
		i += copy(dAtA[i:], m.Title)
	}
	if len(m.Snippets) > 0 {
		for _, s := range m.Snippets {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeVarintSing(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SingInput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Query)
	if l > 0 {
		n += 1 + l + sovSing(uint64(l))
	}
	if m.PageNumber != 0 {
		n += 1 + sozSing(uint64(m.PageNumber))
	}
	if m.ResultPerPage != 0 {
		n += 1 + sozSing(uint64(m.ResultPerPage))
	}
	return n
}

func (m *SingOutput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovSing(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovSing(uint64(l))
	}
	if len(m.Snippets) > 0 {
		for _, s := range m.Snippets {
			l = len(s)
			n += 1 + l + sovSing(uint64(l))
		}
	}
	if len(m.Results) > 0 {
		for _, e := range m.Results {
			l = e.Size()
			n += 1 + l + sovSing(uint64(l))
		}
	}
	return n
}

func (m *SingResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovSing(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovSing(uint64(l))
	}
	if len(m.Snippets) > 0 {
		for _, s := range m.Snippets {
			l = len(s)
			n += 1 + l + sovSing(uint64(l))
		}
	}
	return n
}

func sovSing(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSing(x uint64) (n int) {
	return sovSing(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SingInput) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSing
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SingInput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SingInput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSing
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Query = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PageNumber", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.PageNumber = int64(v)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResultPerPage", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.ResultPerPage = int64(v)
		default:
			iNdEx = preIndex
			skippy, err := skipSing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSing
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
func (m *SingOutput) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSing
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SingOutput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SingOutput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSing
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSing
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Snippets", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSing
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Snippets = append(m.Snippets, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSing
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, &SingResult{})
			if err := m.Results[len(m.Results)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSing
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
func (m *SingResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSing
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SingResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SingResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSing
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSing
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Snippets", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSing
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Snippets = append(m.Snippets, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSing
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
func skipSing(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSing
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
					return 0, ErrIntOverflowSing
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSing
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSing
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSing
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSing(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSing = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSing   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("sing.proto", fileDescriptor_sing_de4ff7d2969bf8fe) }

var fileDescriptor_sing_de4ff7d2969bf8fe = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x4b, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0xd3, 0x66, 0x7c, 0xa4, 0x82, 0xcc, 0xd0, 0xb8, 0x08, 0x59, 0xb4, 0x21, 0x0b, 0x09,
	0x2e, 0x06, 0x19, 0x0f, 0x20, 0xb8, 0x1a, 0x37, 0x1a, 0x7a, 0x0e, 0x10, 0x66, 0xa4, 0x08, 0x91,
	0x98, 0xb4, 0xfd, 0x10, 0x04, 0x0f, 0xe1, 0xb1, 0x5c, 0xce, 0xd2, 0xa5, 0x24, 0x17, 0x91, 0xee,
	0xf6, 0x75, 0x80, 0xd9, 0xd5, 0xff, 0x57, 0xd1, 0x5f, 0xf5, 0x5f, 0x00, 0xaa, 0xe9, 0xea, 0xb9,
	0x90, 0xbd, 0xee, 0xe9, 0xc4, 0xd6, 0xf9, 0x03, 0x44, 0xab, 0xa6, 0xab, 0x6f, 0x3a, 0x61, 0x34,
	0x3d, 0x81, 0xfd, 0x27, 0x83, 0xf2, 0x25, 0x21, 0x19, 0x29, 0x22, 0xee, 0x05, 0x3d, 0x85, 0x58,
	0xac, 0x6b, 0xac, 0x3a, 0xf3, 0xb8, 0x41, 0x99, 0xec, 0x65, 0xa4, 0xa0, 0x1c, 0xac, 0x75, 0xeb,
	0x1c, 0x7a, 0x06, 0x53, 0x89, 0xca, 0xb4, 0xba, 0x12, 0x28, 0x2b, 0xdb, 0x48, 0x42, 0x37, 0x74,
	0xec, 0xed, 0x12, 0x65, 0xb9, 0xae, 0x31, 0x7f, 0x05, 0xb0, 0xac, 0x3b, 0xa3, 0x2d, 0x6c, 0x06,
	0xa1, 0x91, 0xed, 0x37, 0xca, 0x96, 0x16, 0xaf, 0x1b, 0xdd, 0xa2, 0x43, 0x44, 0xdc, 0x0b, 0x9a,
	0xc2, 0x91, 0xea, 0x1a, 0x21, 0x50, 0xab, 0x24, 0xcc, 0xc2, 0x22, 0xe2, 0xbf, 0x9a, 0x9e, 0xc3,
	0xa1, 0x47, 0xa8, 0x64, 0x92, 0x85, 0x45, 0xbc, 0x98, 0xcd, 0xdd, 0x0f, 0x2d, 0x86, 0xbb, 0x06,
	0xff, 0x19, 0xc8, 0x4b, 0x4f, 0xf7, 0xf6, 0x2e, 0xe8, 0x8b, 0x2b, 0x88, 0xed, 0x8b, 0x2b, 0x94,
	0xcf, 0xcd, 0x3d, 0xd2, 0x0b, 0x1f, 0xe5, 0x12, 0xdb, 0xb6, 0xa7, 0xd3, 0xbf, 0x45, 0x5c, 0xb6,
	0xe9, 0xbf, 0xcd, 0x7c, 0x00, 0x79, 0x70, 0x9d, 0xbe, 0x0f, 0x8c, 0x6c, 0x07, 0x46, 0x3e, 0x07,
	0x46, 0xde, 0x46, 0x16, 0x6c, 0x47, 0x16, 0x7c, 0x8c, 0x2c, 0x58, 0x92, 0xcd, 0x81, 0xbb, 0xd2,
	0xe5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xee, 0x6d, 0x5c, 0xb3, 0x01, 0x00, 0x00,
}
