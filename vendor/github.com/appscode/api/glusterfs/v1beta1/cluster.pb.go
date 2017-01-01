// Code generated by protoc-gen-go.
// source: cluster.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	cluster.proto
	volume.proto

It has these top-level messages:
	Cluster
	ClusterListRequest
	ClusterListResponse
	ClusterDescribeRequest
	ClusterDescribeResponse
	ClusterCreateRequest
	ClusterDeleteRequest
	VolumeListRequest
	VolumeListResponse
	Volume
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Cluster struct {
	Phid          string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	KubeCluster   string `protobuf:"bytes,3,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	KubeNamespace string `protobuf:"bytes,4,opt,name=kube_namespace,json=kubeNamespace" json:"kube_namespace,omitempty"`
	Replica       int32  `protobuf:"varint,5,opt,name=replica" json:"replica,omitempty"`
	Endpoint      string `protobuf:"bytes,6,opt,name=endpoint" json:"endpoint,omitempty"`
	CreatedAt     int64  `protobuf:"varint,7,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	Status        string `protobuf:"bytes,8,opt,name=status" json:"status,omitempty"`
}

func (m *Cluster) Reset()                    { *m = Cluster{} }
func (m *Cluster) String() string            { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()               {}
func (*Cluster) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Cluster) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

func (m *Cluster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cluster) GetKubeCluster() string {
	if m != nil {
		return m.KubeCluster
	}
	return ""
}

func (m *Cluster) GetKubeNamespace() string {
	if m != nil {
		return m.KubeNamespace
	}
	return ""
}

func (m *Cluster) GetReplica() int32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

func (m *Cluster) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *Cluster) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Cluster) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type ClusterListRequest struct {
	KubeCluster string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	// List of status to get the agent filterd on the status
	// values in
	//   PENDING
	//   FAILED
	//   READY
	//   DELETED
	Status []string `protobuf:"bytes,2,rep,name=status" json:"status,omitempty"`
}

func (m *ClusterListRequest) Reset()                    { *m = ClusterListRequest{} }
func (m *ClusterListRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterListRequest) ProtoMessage()               {}
func (*ClusterListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ClusterListRequest) GetKubeCluster() string {
	if m != nil {
		return m.KubeCluster
	}
	return ""
}

func (m *ClusterListRequest) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

type ClusterListResponse struct {
	Status   *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Clusters []*Cluster              `protobuf:"bytes,2,rep,name=clusters" json:"clusters,omitempty"`
}

func (m *ClusterListResponse) Reset()                    { *m = ClusterListResponse{} }
func (m *ClusterListResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterListResponse) ProtoMessage()               {}
func (*ClusterListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClusterListResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ClusterListResponse) GetClusters() []*Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type ClusterDescribeRequest struct {
	KubeCluster   string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	KubeNamespace string `protobuf:"bytes,2,opt,name=kube_namespace,json=kubeNamespace" json:"kube_namespace,omitempty"`
	Name          string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *ClusterDescribeRequest) Reset()                    { *m = ClusterDescribeRequest{} }
func (m *ClusterDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterDescribeRequest) ProtoMessage()               {}
func (*ClusterDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ClusterDescribeRequest) GetKubeCluster() string {
	if m != nil {
		return m.KubeCluster
	}
	return ""
}

func (m *ClusterDescribeRequest) GetKubeNamespace() string {
	if m != nil {
		return m.KubeNamespace
	}
	return ""
}

func (m *ClusterDescribeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ClusterDescribeResponse struct {
	Status    *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Glusterfs *Cluster                `protobuf:"bytes,2,opt,name=glusterfs" json:"glusterfs,omitempty"`
}

func (m *ClusterDescribeResponse) Reset()                    { *m = ClusterDescribeResponse{} }
func (m *ClusterDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterDescribeResponse) ProtoMessage()               {}
func (*ClusterDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ClusterDescribeResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ClusterDescribeResponse) GetGlusterfs() *Cluster {
	if m != nil {
		return m.Glusterfs
	}
	return nil
}

type ClusterCreateRequest struct {
	Name          string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Node          int32  `protobuf:"varint,2,opt,name=node" json:"node,omitempty"`
	DiskSizeGb    int32  `protobuf:"varint,6,opt,name=disk_size_gb,json=diskSizeGb" json:"disk_size_gb,omitempty"`
	KubeCluster   string `protobuf:"bytes,4,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	KubeNamespace string `protobuf:"bytes,5,opt,name=kube_namespace,json=kubeNamespace" json:"kube_namespace,omitempty"`
	StorageClass  string `protobuf:"bytes,7,opt,name=storage_class,json=storageClass" json:"storage_class,omitempty"`
}

func (m *ClusterCreateRequest) Reset()                    { *m = ClusterCreateRequest{} }
func (m *ClusterCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterCreateRequest) ProtoMessage()               {}
func (*ClusterCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ClusterCreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClusterCreateRequest) GetNode() int32 {
	if m != nil {
		return m.Node
	}
	return 0
}

func (m *ClusterCreateRequest) GetDiskSizeGb() int32 {
	if m != nil {
		return m.DiskSizeGb
	}
	return 0
}

func (m *ClusterCreateRequest) GetKubeCluster() string {
	if m != nil {
		return m.KubeCluster
	}
	return ""
}

func (m *ClusterCreateRequest) GetKubeNamespace() string {
	if m != nil {
		return m.KubeNamespace
	}
	return ""
}

func (m *ClusterCreateRequest) GetStorageClass() string {
	if m != nil {
		return m.StorageClass
	}
	return ""
}

type ClusterDeleteRequest struct {
	KubeCluster   string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	KubeNamespace string `protobuf:"bytes,2,opt,name=kube_namespace,json=kubeNamespace" json:"kube_namespace,omitempty"`
	Name          string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *ClusterDeleteRequest) Reset()                    { *m = ClusterDeleteRequest{} }
func (m *ClusterDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterDeleteRequest) ProtoMessage()               {}
func (*ClusterDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ClusterDeleteRequest) GetKubeCluster() string {
	if m != nil {
		return m.KubeCluster
	}
	return ""
}

func (m *ClusterDeleteRequest) GetKubeNamespace() string {
	if m != nil {
		return m.KubeNamespace
	}
	return ""
}

func (m *ClusterDeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Cluster)(nil), "appscode.glusterfs.v1beta1.Cluster")
	proto.RegisterType((*ClusterListRequest)(nil), "appscode.glusterfs.v1beta1.ClusterListRequest")
	proto.RegisterType((*ClusterListResponse)(nil), "appscode.glusterfs.v1beta1.ClusterListResponse")
	proto.RegisterType((*ClusterDescribeRequest)(nil), "appscode.glusterfs.v1beta1.ClusterDescribeRequest")
	proto.RegisterType((*ClusterDescribeResponse)(nil), "appscode.glusterfs.v1beta1.ClusterDescribeResponse")
	proto.RegisterType((*ClusterCreateRequest)(nil), "appscode.glusterfs.v1beta1.ClusterCreateRequest")
	proto.RegisterType((*ClusterDeleteRequest)(nil), "appscode.glusterfs.v1beta1.ClusterDeleteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Clusters service

type ClustersClient interface {
	// Glusterfs cluster list. Needs to work with two modes.
	// First is to list all the glusterfs cluster through the
	// space with out considering the kubernetes cluster. if the
	// cluster_name is provided then list all the glusterfs cluster
	// with respect to the provided kube cluster space.
	List(ctx context.Context, in *ClusterListRequest, opts ...grpc.CallOption) (*ClusterListResponse, error)
	Describe(ctx context.Context, in *ClusterDescribeRequest, opts ...grpc.CallOption) (*ClusterDescribeResponse, error)
	Create(ctx context.Context, in *ClusterCreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *ClusterDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type clustersClient struct {
	cc *grpc.ClientConn
}

func NewClustersClient(cc *grpc.ClientConn) ClustersClient {
	return &clustersClient{cc}
}

func (c *clustersClient) List(ctx context.Context, in *ClusterListRequest, opts ...grpc.CallOption) (*ClusterListResponse, error) {
	out := new(ClusterListResponse)
	err := grpc.Invoke(ctx, "/appscode.glusterfs.v1beta1.Clusters/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Describe(ctx context.Context, in *ClusterDescribeRequest, opts ...grpc.CallOption) (*ClusterDescribeResponse, error) {
	out := new(ClusterDescribeResponse)
	err := grpc.Invoke(ctx, "/appscode.glusterfs.v1beta1.Clusters/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Create(ctx context.Context, in *ClusterCreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.glusterfs.v1beta1.Clusters/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Delete(ctx context.Context, in *ClusterDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.glusterfs.v1beta1.Clusters/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Clusters service

type ClustersServer interface {
	// Glusterfs cluster list. Needs to work with two modes.
	// First is to list all the glusterfs cluster through the
	// space with out considering the kubernetes cluster. if the
	// cluster_name is provided then list all the glusterfs cluster
	// with respect to the provided kube cluster space.
	List(context.Context, *ClusterListRequest) (*ClusterListResponse, error)
	Describe(context.Context, *ClusterDescribeRequest) (*ClusterDescribeResponse, error)
	Create(context.Context, *ClusterCreateRequest) (*appscode_dtypes.VoidResponse, error)
	Delete(context.Context, *ClusterDeleteRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterClustersServer(s *grpc.Server, srv ClustersServer) {
	s.RegisterService(&_Clusters_serviceDesc, srv)
}

func _Clusters_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.glusterfs.v1beta1.Clusters/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).List(ctx, req.(*ClusterListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.glusterfs.v1beta1.Clusters/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).Describe(ctx, req.(*ClusterDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.glusterfs.v1beta1.Clusters/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).Create(ctx, req.(*ClusterCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.glusterfs.v1beta1.Clusters/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).Delete(ctx, req.(*ClusterDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Clusters_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.glusterfs.v1beta1.Clusters",
	HandlerType: (*ClustersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Clusters_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Clusters_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Clusters_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Clusters_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cluster.proto",
}

func init() { proto.RegisterFile("cluster.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 673 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x55, 0x4f, 0x4f, 0xd4, 0x4e,
	0x18, 0xce, 0xec, 0xff, 0x7d, 0x59, 0x7e, 0x87, 0xf9, 0x19, 0x68, 0x1a, 0xd1, 0xb5, 0xc4, 0x84,
	0x70, 0x68, 0x65, 0xb9, 0x99, 0x18, 0x03, 0x8b, 0xf1, 0x62, 0x14, 0x4b, 0xe2, 0xc1, 0xcb, 0x66,
	0xb6, 0x1d, 0x96, 0x86, 0xa5, 0x53, 0x3b, 0xb3, 0x24, 0x42, 0xb8, 0x70, 0xf1, 0xe4, 0xc9, 0xcf,
	0xe2, 0xc5, 0xaf, 0xe0, 0x81, 0x03, 0x5f, 0xc1, 0x83, 0x7e, 0x0b, 0x33, 0xd3, 0x69, 0xb7, 0xb5,
	0x10, 0x0b, 0x89, 0x5c, 0x36, 0xed, 0xf3, 0xfe, 0x7d, 0xde, 0xf7, 0x79, 0xb7, 0xb0, 0xe8, 0x4d,
	0x67, 0x5c, 0xd0, 0xd8, 0x8e, 0x62, 0x26, 0x18, 0x36, 0x49, 0x14, 0x71, 0x8f, 0xf9, 0xd4, 0x9e,
	0x24, 0xf8, 0x3e, 0xb7, 0x8f, 0x37, 0xc6, 0x54, 0x90, 0x0d, 0xf3, 0xfe, 0x84, 0xb1, 0xc9, 0x94,
	0x3a, 0x24, 0x0a, 0x1c, 0x12, 0x86, 0x4c, 0x10, 0x11, 0xb0, 0x90, 0x27, 0x91, 0xe6, 0x83, 0x34,
	0xf2, 0x1a, 0xfb, 0xc3, 0x82, 0xdd, 0x17, 0x1f, 0x23, 0xca, 0x1d, 0xf5, 0x9b, 0x38, 0x58, 0x3f,
	0x11, 0xb4, 0x87, 0x49, 0x51, 0x8c, 0xa1, 0x11, 0x1d, 0x04, 0xbe, 0x81, 0xfa, 0x68, 0xad, 0xeb,
	0xaa, 0x67, 0x89, 0x85, 0xe4, 0x88, 0x1a, 0xb5, 0x04, 0x93, 0xcf, 0xf8, 0x11, 0xf4, 0x0e, 0x67,
	0x63, 0x3a, 0xd2, 0x24, 0x8c, 0xba, 0xb2, 0x2d, 0x48, 0x2c, 0x4d, 0xf5, 0x18, 0xfe, 0x53, 0x2e,
	0xd2, 0x9f, 0x47, 0xc4, 0xa3, 0x46, 0x43, 0x39, 0x2d, 0x4a, 0xf4, 0x75, 0x0a, 0x62, 0x03, 0xda,
	0x31, 0x8d, 0xa6, 0x81, 0x47, 0x8c, 0x66, 0x1f, 0xad, 0x35, 0xdd, 0xf4, 0x15, 0x9b, 0xd0, 0xa1,
	0xa1, 0x1f, 0xb1, 0x20, 0x14, 0x46, 0x4b, 0x85, 0x66, 0xef, 0x78, 0x05, 0xc0, 0x8b, 0x29, 0x11,
	0xd4, 0x1f, 0x11, 0x61, 0xb4, 0xfb, 0x68, 0xad, 0xee, 0x76, 0x35, 0xb2, 0x25, 0xf0, 0x12, 0xb4,
	0xb8, 0x20, 0x62, 0xc6, 0x8d, 0x8e, 0x0a, 0xd4, 0x6f, 0xd6, 0x1b, 0xc0, 0xba, 0xbd, 0x57, 0x01,
	0x17, 0x2e, 0xfd, 0x30, 0xa3, 0x5c, 0x94, 0xc8, 0xa0, 0x32, 0x99, 0x79, 0xc2, 0x5a, 0xbf, 0x9e,
	0x4b, 0xf8, 0x09, 0xc1, 0xff, 0x85, 0x8c, 0x3c, 0x62, 0x21, 0xa7, 0xd8, 0xc9, 0xfc, 0x65, 0xb2,
	0x85, 0xc1, 0xb2, 0x9d, 0xed, 0x37, 0xd9, 0x80, 0xbd, 0xa7, 0xcc, 0x69, 0x22, 0xfc, 0x1c, 0x3a,
	0xba, 0x7c, 0x52, 0x62, 0x61, 0xb0, 0x6a, 0x5f, 0x2f, 0x09, 0x5b, 0xd7, 0x74, 0xb3, 0x20, 0xeb,
	0x18, 0x96, 0x34, 0xb8, 0x43, 0xb9, 0x17, 0x07, 0x63, 0x7a, 0x03, 0x7a, 0xe5, 0x5d, 0xd5, 0xae,
	0xda, 0x55, 0xaa, 0x84, 0xfa, 0x5c, 0x09, 0xd6, 0x67, 0x04, 0xcb, 0xa5, 0xc2, 0xb7, 0x9d, 0xc2,
	0x16, 0x74, 0x33, 0xae, 0xaa, 0x85, 0x8a, 0x63, 0x98, 0x47, 0x59, 0x17, 0x08, 0xee, 0x69, 0x78,
	0xa8, 0xf4, 0x90, 0x8e, 0x21, 0x6d, 0x1e, 0xe5, 0x64, 0x2c, 0x31, 0xe6, 0x27, 0x6c, 0x9b, 0xae,
	0x7a, 0xc6, 0x7d, 0xe8, 0xf9, 0x01, 0x3f, 0x1c, 0xf1, 0xe0, 0x84, 0x8e, 0x26, 0x63, 0x25, 0xbd,
	0xa6, 0x0b, 0x12, 0xdb, 0x0b, 0x4e, 0xe8, 0xcb, 0x71, 0x69, 0xa0, 0x8d, 0x2a, 0x03, 0x6d, 0x5e,
	0x35, 0xd0, 0x55, 0x58, 0xe4, 0x82, 0xc5, 0x64, 0x22, 0x93, 0x11, 0xce, 0x95, 0x92, 0xbb, 0x6e,
	0x4f, 0x83, 0x43, 0x89, 0x59, 0x22, 0x23, 0xb4, 0x43, 0xa7, 0x54, 0xdc, 0xcd, 0x5e, 0x07, 0x97,
	0x2d, 0xe8, 0xe8, 0x34, 0x1c, 0x7f, 0x43, 0xd0, 0x90, 0xfa, 0xc6, 0x76, 0x85, 0x6d, 0xe4, 0x4e,
	0xcb, 0x74, 0x2a, 0xfb, 0x27, 0x92, 0xb1, 0xde, 0x9e, 0x7f, 0x35, 0x6a, 0x1d, 0x74, 0x7e, 0xf9,
	0xe3, 0x4b, 0xed, 0x05, 0x1e, 0x3a, 0xa3, 0xc2, 0xbf, 0x97, 0x6c, 0x3b, 0x0e, 0xa9, 0xa0, 0xdc,
	0xd1, 0x49, 0x9c, 0xf4, 0x00, 0x9c, 0xd3, 0xfc, 0x30, 0xce, 0x9c, 0xac, 0x1c, 0xfe, 0x85, 0xa0,
	0x93, 0x4a, 0x13, 0x0f, 0x2a, 0x34, 0xf4, 0xc7, 0x01, 0x99, 0x9b, 0x37, 0x8a, 0xd1, 0x44, 0x44,
	0x8e, 0xc8, 0x01, 0xde, 0xbf, 0x3d, 0x91, 0x6c, 0x77, 0xa9, 0x25, 0x03, 0x72, 0x24, 0x9d, 0x53,
	0x89, 0x9e, 0xe1, 0xef, 0x08, 0x5a, 0x89, 0xec, 0xf1, 0x93, 0x0a, 0x5d, 0x17, 0x2e, 0xc4, 0x5c,
	0x29, 0x9d, 0xe7, 0x3b, 0x16, 0xf8, 0x19, 0xa3, 0x28, 0xc7, 0xc8, 0xb7, 0x46, 0xff, 0x98, 0xd1,
	0x53, 0xb4, 0x8e, 0x2f, 0x10, 0xb4, 0x12, 0xcd, 0x57, 0x62, 0x53, 0x38, 0x8f, 0xbf, 0xb1, 0x29,
	0xec, 0x67, 0xfd, 0x8e, 0xf6, 0xb3, 0xfd, 0x0c, 0x2c, 0x8f, 0x1d, 0xcd, 0x3b, 0x23, 0x51, 0x50,
	0xe6, 0xb3, 0xdd, 0xd3, 0x84, 0x76, 0xe5, 0xf7, 0x79, 0x17, 0xbd, 0x6f, 0x6b, 0xc3, 0xb8, 0xa5,
	0xbe, 0xd8, 0x9b, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x46, 0xdf, 0xe8, 0x3b, 0x3d, 0x08, 0x00,
	0x00,
}