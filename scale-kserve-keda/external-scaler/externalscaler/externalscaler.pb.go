// Code generated by protoc-gen-go. DO NOT EDIT.
// source: externalscaler.proto

package externalscaler

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ScaledObjectRef struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string            `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ScalerMetadata       map[string]string `protobuf:"bytes,3,rep,name=scalerMetadata,proto3" json:"scalerMetadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ScaledObjectRef) Reset()         { *m = ScaledObjectRef{} }
func (m *ScaledObjectRef) String() string { return proto.CompactTextString(m) }
func (*ScaledObjectRef) ProtoMessage()    {}
func (*ScaledObjectRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d382708546499d1, []int{0}
}

func (m *ScaledObjectRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScaledObjectRef.Unmarshal(m, b)
}
func (m *ScaledObjectRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScaledObjectRef.Marshal(b, m, deterministic)
}
func (m *ScaledObjectRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScaledObjectRef.Merge(m, src)
}
func (m *ScaledObjectRef) XXX_Size() int {
	return xxx_messageInfo_ScaledObjectRef.Size(m)
}
func (m *ScaledObjectRef) XXX_DiscardUnknown() {
	xxx_messageInfo_ScaledObjectRef.DiscardUnknown(m)
}

var xxx_messageInfo_ScaledObjectRef proto.InternalMessageInfo

func (m *ScaledObjectRef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ScaledObjectRef) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ScaledObjectRef) GetScalerMetadata() map[string]string {
	if m != nil {
		return m.ScalerMetadata
	}
	return nil
}

type IsActiveResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsActiveResponse) Reset()         { *m = IsActiveResponse{} }
func (m *IsActiveResponse) String() string { return proto.CompactTextString(m) }
func (*IsActiveResponse) ProtoMessage()    {}
func (*IsActiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d382708546499d1, []int{1}
}

func (m *IsActiveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsActiveResponse.Unmarshal(m, b)
}
func (m *IsActiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsActiveResponse.Marshal(b, m, deterministic)
}
func (m *IsActiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsActiveResponse.Merge(m, src)
}
func (m *IsActiveResponse) XXX_Size() int {
	return xxx_messageInfo_IsActiveResponse.Size(m)
}
func (m *IsActiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IsActiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IsActiveResponse proto.InternalMessageInfo

func (m *IsActiveResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type GetMetricSpecResponse struct {
	MetricSpecs          []*MetricSpec `protobuf:"bytes,1,rep,name=metricSpecs,proto3" json:"metricSpecs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetMetricSpecResponse) Reset()         { *m = GetMetricSpecResponse{} }
func (m *GetMetricSpecResponse) String() string { return proto.CompactTextString(m) }
func (*GetMetricSpecResponse) ProtoMessage()    {}
func (*GetMetricSpecResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d382708546499d1, []int{2}
}

func (m *GetMetricSpecResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMetricSpecResponse.Unmarshal(m, b)
}
func (m *GetMetricSpecResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMetricSpecResponse.Marshal(b, m, deterministic)
}
func (m *GetMetricSpecResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMetricSpecResponse.Merge(m, src)
}
func (m *GetMetricSpecResponse) XXX_Size() int {
	return xxx_messageInfo_GetMetricSpecResponse.Size(m)
}
func (m *GetMetricSpecResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMetricSpecResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMetricSpecResponse proto.InternalMessageInfo

func (m *GetMetricSpecResponse) GetMetricSpecs() []*MetricSpec {
	if m != nil {
		return m.MetricSpecs
	}
	return nil
}

type MetricSpec struct {
	MetricName           string   `protobuf:"bytes,1,opt,name=metricName,proto3" json:"metricName,omitempty"`
	TargetSize           int64    `protobuf:"varint,2,opt,name=targetSize,proto3" json:"targetSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricSpec) Reset()         { *m = MetricSpec{} }
func (m *MetricSpec) String() string { return proto.CompactTextString(m) }
func (*MetricSpec) ProtoMessage()    {}
func (*MetricSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d382708546499d1, []int{3}
}

func (m *MetricSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricSpec.Unmarshal(m, b)
}
func (m *MetricSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricSpec.Marshal(b, m, deterministic)
}
func (m *MetricSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricSpec.Merge(m, src)
}
func (m *MetricSpec) XXX_Size() int {
	return xxx_messageInfo_MetricSpec.Size(m)
}
func (m *MetricSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricSpec.DiscardUnknown(m)
}

var xxx_messageInfo_MetricSpec proto.InternalMessageInfo

func (m *MetricSpec) GetMetricName() string {
	if m != nil {
		return m.MetricName
	}
	return ""
}

func (m *MetricSpec) GetTargetSize() int64 {
	if m != nil {
		return m.TargetSize
	}
	return 0
}

type GetMetricsRequest struct {
	ScaledObjectRef      *ScaledObjectRef `protobuf:"bytes,1,opt,name=scaledObjectRef,proto3" json:"scaledObjectRef,omitempty"`
	MetricName           string           `protobuf:"bytes,2,opt,name=metricName,proto3" json:"metricName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetMetricsRequest) Reset()         { *m = GetMetricsRequest{} }
func (m *GetMetricsRequest) String() string { return proto.CompactTextString(m) }
func (*GetMetricsRequest) ProtoMessage()    {}
func (*GetMetricsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d382708546499d1, []int{4}
}

func (m *GetMetricsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMetricsRequest.Unmarshal(m, b)
}
func (m *GetMetricsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMetricsRequest.Marshal(b, m, deterministic)
}
func (m *GetMetricsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMetricsRequest.Merge(m, src)
}
func (m *GetMetricsRequest) XXX_Size() int {
	return xxx_messageInfo_GetMetricsRequest.Size(m)
}
func (m *GetMetricsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMetricsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMetricsRequest proto.InternalMessageInfo

func (m *GetMetricsRequest) GetScaledObjectRef() *ScaledObjectRef {
	if m != nil {
		return m.ScaledObjectRef
	}
	return nil
}

func (m *GetMetricsRequest) GetMetricName() string {
	if m != nil {
		return m.MetricName
	}
	return ""
}

type GetMetricsResponse struct {
	MetricValues         []*MetricValue `protobuf:"bytes,1,rep,name=metricValues,proto3" json:"metricValues,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetMetricsResponse) Reset()         { *m = GetMetricsResponse{} }
func (m *GetMetricsResponse) String() string { return proto.CompactTextString(m) }
func (*GetMetricsResponse) ProtoMessage()    {}
func (*GetMetricsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d382708546499d1, []int{5}
}

func (m *GetMetricsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMetricsResponse.Unmarshal(m, b)
}
func (m *GetMetricsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMetricsResponse.Marshal(b, m, deterministic)
}
func (m *GetMetricsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMetricsResponse.Merge(m, src)
}
func (m *GetMetricsResponse) XXX_Size() int {
	return xxx_messageInfo_GetMetricsResponse.Size(m)
}
func (m *GetMetricsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMetricsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMetricsResponse proto.InternalMessageInfo

func (m *GetMetricsResponse) GetMetricValues() []*MetricValue {
	if m != nil {
		return m.MetricValues
	}
	return nil
}

type MetricValue struct {
	MetricName           string   `protobuf:"bytes,1,opt,name=metricName,proto3" json:"metricName,omitempty"`
	MetricValue          int64    `protobuf:"varint,2,opt,name=metricValue,proto3" json:"metricValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricValue) Reset()         { *m = MetricValue{} }
func (m *MetricValue) String() string { return proto.CompactTextString(m) }
func (*MetricValue) ProtoMessage()    {}
func (*MetricValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d382708546499d1, []int{6}
}

func (m *MetricValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricValue.Unmarshal(m, b)
}
func (m *MetricValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricValue.Marshal(b, m, deterministic)
}
func (m *MetricValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricValue.Merge(m, src)
}
func (m *MetricValue) XXX_Size() int {
	return xxx_messageInfo_MetricValue.Size(m)
}
func (m *MetricValue) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricValue.DiscardUnknown(m)
}

var xxx_messageInfo_MetricValue proto.InternalMessageInfo

func (m *MetricValue) GetMetricName() string {
	if m != nil {
		return m.MetricName
	}
	return ""
}

func (m *MetricValue) GetMetricValue() int64 {
	if m != nil {
		return m.MetricValue
	}
	return 0
}

func init() {
	proto.RegisterType((*ScaledObjectRef)(nil), "externalscaler.ScaledObjectRef")
	proto.RegisterMapType((map[string]string)(nil), "externalscaler.ScaledObjectRef.ScalerMetadataEntry")
	proto.RegisterType((*IsActiveResponse)(nil), "externalscaler.IsActiveResponse")
	proto.RegisterType((*GetMetricSpecResponse)(nil), "externalscaler.GetMetricSpecResponse")
	proto.RegisterType((*MetricSpec)(nil), "externalscaler.MetricSpec")
	proto.RegisterType((*GetMetricsRequest)(nil), "externalscaler.GetMetricsRequest")
	proto.RegisterType((*GetMetricsResponse)(nil), "externalscaler.GetMetricsResponse")
	proto.RegisterType((*MetricValue)(nil), "externalscaler.MetricValue")
}

func init() { proto.RegisterFile("externalscaler.proto", fileDescriptor_3d382708546499d1) }

var fileDescriptor_3d382708546499d1 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdf, 0x6b, 0xd4, 0x40,
	0x10, 0x6e, 0x12, 0x2d, 0xed, 0x44, 0xd3, 0x73, 0xac, 0x12, 0xa2, 0x68, 0x5c, 0x10, 0x8a, 0x0f,
	0x87, 0x5c, 0x5f, 0x44, 0x05, 0xa9, 0x50, 0xa4, 0x60, 0x3d, 0xd8, 0x70, 0x15, 0xf5, 0x69, 0x9b,
	0x8e, 0x72, 0x9a, 0xcb, 0xc5, 0xdd, 0xbd, 0xe2, 0xf9, 0xe0, 0x3f, 0xeb, 0xab, 0x7f, 0x84, 0xe4,
	0x77, 0xb2, 0x9c, 0xe6, 0xa5, 0x4f, 0xd9, 0x9d, 0xf9, 0xe6, 0xdb, 0x99, 0x6f, 0x3e, 0x02, 0xfb,
	0xf4, 0x43, 0x93, 0x4c, 0x45, 0xa2, 0x62, 0x91, 0x90, 0x1c, 0x67, 0x72, 0xa9, 0x97, 0xe8, 0xf5,
	0xa3, 0xec, 0xb7, 0x05, 0x7b, 0x51, 0x7e, 0xbc, 0x98, 0x9e, 0x7f, 0xa5, 0x58, 0x73, 0xfa, 0x8c,
	0x08, 0xd7, 0x52, 0xb1, 0x20, 0xdf, 0x0a, 0xad, 0x83, 0x5d, 0x5e, 0x9c, 0xf1, 0x3e, 0xec, 0xe6,
	0x5f, 0x95, 0x89, 0x98, 0x7c, 0xbb, 0x48, 0xb4, 0x01, 0xfc, 0x04, 0x5e, 0xc9, 0x77, 0x4a, 0x5a,
	0x5c, 0x08, 0x2d, 0x7c, 0x27, 0x74, 0x0e, 0xdc, 0xc9, 0xe1, 0xd8, 0x68, 0xc2, 0x78, 0xaa, 0xbc,
	0x37, 0x55, 0xc7, 0xa9, 0x96, 0x6b, 0x6e, 0x50, 0x05, 0x47, 0x70, 0x7b, 0x03, 0x0c, 0x47, 0xe0,
	0x7c, 0xa3, 0x75, 0xd5, 0x64, 0x7e, 0xc4, 0x7d, 0xb8, 0x7e, 0x29, 0x92, 0x55, 0xdd, 0x5f, 0x79,
	0x79, 0x6e, 0x3f, 0xb3, 0xd8, 0x13, 0x18, 0x9d, 0xa8, 0xa3, 0x58, 0xcf, 0x2f, 0x89, 0x93, 0xca,
	0x96, 0xa9, 0x22, 0xbc, 0x0b, 0xdb, 0x92, 0xd4, 0x2a, 0xd1, 0x05, 0xc5, 0x0e, 0xaf, 0x6e, 0x6c,
	0x06, 0x77, 0xde, 0x90, 0x3e, 0x25, 0x2d, 0xe7, 0x71, 0x94, 0x51, 0xdc, 0x14, 0xbc, 0x04, 0x77,
	0xd1, 0x44, 0x95, 0x6f, 0x15, 0x13, 0x06, 0xe6, 0x84, 0x9d, 0xc2, 0x2e, 0x9c, 0xbd, 0x05, 0x68,
	0x53, 0xf8, 0x00, 0xa0, 0x4c, 0xbe, 0x6b, 0x85, 0xee, 0x44, 0xf2, 0xbc, 0x16, 0xf2, 0x0b, 0xe9,
	0x68, 0xfe, 0xb3, 0x9c, 0xc7, 0xe1, 0x9d, 0x08, 0xfb, 0x05, 0xb7, 0x9a, 0x26, 0x15, 0xa7, 0xef,
	0x2b, 0x52, 0x1a, 0x4f, 0x60, 0x4f, 0xf5, 0xf5, 0x2d, 0x98, 0xdd, 0xc9, 0xc3, 0x81, 0x35, 0x70,
	0xb3, 0xce, 0xe8, 0xcf, 0x36, 0xfb, 0x63, 0x33, 0xc0, 0xee, 0xfb, 0x95, 0x42, 0xaf, 0xe0, 0x46,
	0x89, 0x39, 0xcb, 0x95, 0xaf, 0x25, 0xba, 0xb7, 0x59, 0xa2, 0x02, 0xc3, 0x7b, 0x05, 0x6c, 0x0a,
	0x6e, 0x27, 0x39, 0xa8, 0x52, 0x58, 0x6f, 0xe4, 0xac, 0x59, 0xbb, 0xc3, 0xbb, 0xa1, 0xc9, 0x1f,
	0x1b, 0xbc, 0xe3, 0xea, 0xf5, 0xd2, 0x44, 0x38, 0x85, 0x9d, 0xda, 0x0b, 0x38, 0x24, 0x4c, 0x10,
	0x9a, 0x00, 0xd3, 0x46, 0x6c, 0x0b, 0xdf, 0x83, 0x17, 0x69, 0x49, 0x62, 0x71, 0xa5, 0xb4, 0x4f,
	0x2d, 0xfc, 0x00, 0x37, 0x7b, 0x4e, 0x1c, 0xe6, 0x7d, 0x6c, 0x02, 0x36, 0x3a, 0x99, 0x6d, 0xe1,
	0x0c, 0xa0, 0xdd, 0x1f, 0x3e, 0xfa, 0x67, 0x59, 0xed, 0xad, 0x80, 0xfd, 0x0f, 0x52, 0xd3, 0xbe,
	0xc6, 0x8f, 0xa3, 0xf1, 0x8b, 0x3e, 0xf0, 0x7c, 0xbb, 0xf8, 0xf1, 0x1c, 0xfe, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0xdc, 0xf1, 0x73, 0xce, 0x90, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExternalScalerClient is the client API for ExternalScaler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExternalScalerClient interface {
	IsActive(ctx context.Context, in *ScaledObjectRef, opts ...grpc.CallOption) (*IsActiveResponse, error)
	StreamIsActive(ctx context.Context, in *ScaledObjectRef, opts ...grpc.CallOption) (ExternalScaler_StreamIsActiveClient, error)
	GetMetricSpec(ctx context.Context, in *ScaledObjectRef, opts ...grpc.CallOption) (*GetMetricSpecResponse, error)
	GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error)
}

type externalScalerClient struct {
	cc *grpc.ClientConn
}

func NewExternalScalerClient(cc *grpc.ClientConn) ExternalScalerClient {
	return &externalScalerClient{cc}
}

func (c *externalScalerClient) IsActive(ctx context.Context, in *ScaledObjectRef, opts ...grpc.CallOption) (*IsActiveResponse, error) {
	out := new(IsActiveResponse)
	err := c.cc.Invoke(ctx, "/externalscaler.ExternalScaler/IsActive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalScalerClient) StreamIsActive(ctx context.Context, in *ScaledObjectRef, opts ...grpc.CallOption) (ExternalScaler_StreamIsActiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ExternalScaler_serviceDesc.Streams[0], "/externalscaler.ExternalScaler/StreamIsActive", opts...)
	if err != nil {
		return nil, err
	}
	x := &externalScalerStreamIsActiveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExternalScaler_StreamIsActiveClient interface {
	Recv() (*IsActiveResponse, error)
	grpc.ClientStream
}

type externalScalerStreamIsActiveClient struct {
	grpc.ClientStream
}

func (x *externalScalerStreamIsActiveClient) Recv() (*IsActiveResponse, error) {
	m := new(IsActiveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *externalScalerClient) GetMetricSpec(ctx context.Context, in *ScaledObjectRef, opts ...grpc.CallOption) (*GetMetricSpecResponse, error) {
	out := new(GetMetricSpecResponse)
	err := c.cc.Invoke(ctx, "/externalscaler.ExternalScaler/GetMetricSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalScalerClient) GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error) {
	out := new(GetMetricsResponse)
	err := c.cc.Invoke(ctx, "/externalscaler.ExternalScaler/GetMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExternalScalerServer is the server API for ExternalScaler service.
type ExternalScalerServer interface {
	IsActive(context.Context, *ScaledObjectRef) (*IsActiveResponse, error)
	StreamIsActive(*ScaledObjectRef, ExternalScaler_StreamIsActiveServer) error
	GetMetricSpec(context.Context, *ScaledObjectRef) (*GetMetricSpecResponse, error)
	GetMetrics(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error)
}

// UnimplementedExternalScalerServer can be embedded to have forward compatible implementations.
type UnimplementedExternalScalerServer struct {
}

func (*UnimplementedExternalScalerServer) IsActive(ctx context.Context, req *ScaledObjectRef) (*IsActiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsActive not implemented")
}
func (*UnimplementedExternalScalerServer) StreamIsActive(req *ScaledObjectRef, srv ExternalScaler_StreamIsActiveServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamIsActive not implemented")
}
func (*UnimplementedExternalScalerServer) GetMetricSpec(ctx context.Context, req *ScaledObjectRef) (*GetMetricSpecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricSpec not implemented")
}
func (*UnimplementedExternalScalerServer) GetMetrics(ctx context.Context, req *GetMetricsRequest) (*GetMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetrics not implemented")
}

func RegisterExternalScalerServer(s *grpc.Server, srv ExternalScalerServer) {
	s.RegisterService(&_ExternalScaler_serviceDesc, srv)
}

func _ExternalScaler_IsActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScaledObjectRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalScalerServer).IsActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/externalscaler.ExternalScaler/IsActive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalScalerServer).IsActive(ctx, req.(*ScaledObjectRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalScaler_StreamIsActive_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ScaledObjectRef)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExternalScalerServer).StreamIsActive(m, &externalScalerStreamIsActiveServer{stream})
}

type ExternalScaler_StreamIsActiveServer interface {
	Send(*IsActiveResponse) error
	grpc.ServerStream
}

type externalScalerStreamIsActiveServer struct {
	grpc.ServerStream
}

func (x *externalScalerStreamIsActiveServer) Send(m *IsActiveResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ExternalScaler_GetMetricSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScaledObjectRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalScalerServer).GetMetricSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/externalscaler.ExternalScaler/GetMetricSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalScalerServer).GetMetricSpec(ctx, req.(*ScaledObjectRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalScaler_GetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalScalerServer).GetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/externalscaler.ExternalScaler/GetMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalScalerServer).GetMetrics(ctx, req.(*GetMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExternalScaler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "externalscaler.ExternalScaler",
	HandlerType: (*ExternalScalerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsActive",
			Handler:    _ExternalScaler_IsActive_Handler,
		},
		{
			MethodName: "GetMetricSpec",
			Handler:    _ExternalScaler_GetMetricSpec_Handler,
		},
		{
			MethodName: "GetMetrics",
			Handler:    _ExternalScaler_GetMetrics_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamIsActive",
			Handler:       _ExternalScaler_StreamIsActive_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "externalscaler.proto",
}
