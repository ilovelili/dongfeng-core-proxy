// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/ilovelili/dongfeng-core-proxy/services/proto/api.proto

package dongfeng_svc_core_proxy

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/ilovelili/dongfeng-protobuf"
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

func init() {
	proto.RegisterFile("github.com/ilovelili/dongfeng-core-proxy/services/proto/api.proto", fileDescriptor_b5a8ae25ec7233ee)
}

var fileDescriptor_b5a8ae25ec7233ee = []byte{
	// 725 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x96, 0x5b, 0x4f, 0xd4, 0x40,
	0x14, 0xc7, 0x25, 0x46, 0x23, 0x07, 0xf0, 0x32, 0x3e, 0x18, 0xd7, 0x07, 0x15, 0x14, 0x50, 0xa4,
	0x9b, 0x70, 0xf9, 0x00, 0x28, 0x49, 0x25, 0x01, 0x42, 0xb8, 0x88, 0x46, 0x23, 0x76, 0xbb, 0x67,
	0xbb, 0x93, 0xec, 0x76, 0x4a, 0x67, 0xba, 0x91, 0x2f, 0x60, 0x62, 0xe2, 0xb3, 0x9f, 0xd7, 0x74,
	0xba, 0x33, 0xdb, 0xdb, 0xb4, 0xdd, 0xc4, 0x7d, 0x21, 0x64, 0xff, 0xbf, 0xf3, 0xff, 0x9f, 0x9e,
	0x76, 0x2e, 0xb0, 0xe7, 0x51, 0xd1, 0x8f, 0x3a, 0x96, 0xcb, 0x86, 0x6d, 0x3a, 0x60, 0x23, 0x1c,
	0xd0, 0x01, 0x6d, 0x77, 0x99, 0xef, 0xf5, 0xd0, 0xf7, 0x36, 0x5d, 0x16, 0xe2, 0x66, 0x10, 0xb2,
	0x9f, 0x37, 0x6d, 0x8e, 0xe1, 0x88, 0xba, 0xc8, 0xdb, 0x41, 0xc8, 0x04, 0x6b, 0x3b, 0x01, 0xb5,
	0xe4, 0x7f, 0xe4, 0x89, 0xa2, 0x2d, 0x3e, 0x72, 0xad, 0xb8, 0xc2, 0x92, 0x15, 0xad, 0xdd, 0x6a,
	0x6f, 0x59, 0xdd, 0x89, 0x7a, 0xfa, 0x97, 0xc4, 0x6f, 0xeb, 0x6f, 0x0b, 0x6e, 0x3b, 0x01, 0x25,
	0x87, 0x70, 0xe7, 0x90, 0x79, 0xd4, 0x27, 0xcf, 0xad, 0x2c, 0xd1, 0x89, 0x7a, 0x96, 0x54, 0x4e,
	0xf1, 0x3a, 0x42, 0x2e, 0x5a, 0x2f, 0xcc, 0x00, 0x0f, 0x98, 0xcf, 0x71, 0xf9, 0x16, 0xf9, 0x0c,
	0xf3, 0xfb, 0x0e, 0xef, 0x77, 0x98, 0x13, 0x76, 0xc9, 0x4a, 0x49, 0x81, 0x56, 0x95, 0xeb, 0xab,
	0x6a, 0x48, 0x3b, 0x7f, 0x05, 0xb8, 0x08, 0xba, 0x8e, 0xc0, 0x0b, 0x8e, 0x21, 0x29, 0xab, 0x9a,
	0xc8, 0xca, 0xfb, 0x75, 0x0d, 0xa5, 0xcd, 0x11, 0xee, 0xdb, 0x28, 0xf6, 0x84, 0x40, 0xbf, 0xeb,
	0xf8, 0x2e, 0x72, 0xb2, 0x56, 0x52, 0x9a, 0x41, 0x54, 0xc6, 0x7a, 0x3d, 0xa8, 0x63, 0x86, 0xf0,
	0x30, 0x89, 0x9f, 0xa8, 0xe4, 0xad, 0xb1, 0xc7, 0x62, 0xd6, 0x46, 0x23, 0x56, 0xc7, 0xf9, 0xf0,
	0x28, 0xaf, 0xf2, 0x59, 0xe6, 0x5d, 0x02, 0xd8, 0x28, 0x3e, 0x0c, 0x1c, 0xce, 0x91, 0x93, 0xe5,
	0xf2, 0xc1, 0x48, 0x59, 0x05, 0xac, 0x54, 0x32, 0xda, 0xf8, 0x3b, 0x2c, 0x24, 0xb1, 0x52, 0x20,
	0xe6, 0xd7, 0x9a, 0x31, 0x5f, 0xad, 0xc3, 0xb4, 0xff, 0x0f, 0x58, 0x4a, 0x09, 0x38, 0x83, 0x84,
	0x4f, 0x30, 0x6f, 0xa3, 0x38, 0x89, 0x02, 0x3a, 0x30, 0x4e, 0x46, 0xaa, 0x35, 0x93, 0x19, 0x33,
	0xc5, 0xc9, 0x48, 0xa1, 0xa2, 0xef, 0x8c, 0xf9, 0x6a, 0x1d, 0xa6, 0xfd, 0xaf, 0x60, 0x31, 0x25,
	0xf0, 0xff, 0x1f, 0xf0, 0x0d, 0x16, 0x6c, 0x14, 0xe7, 0xe8, 0xb8, 0x7d, 0x0c, 0x79, 0xe9, 0xba,
	0x9e, 0xe8, 0x55, 0xeb, 0x3a, 0x4d, 0x69, 0xf7, 0xae, 0x7a, 0xb1, 0x63, 0xa9, 0x74, 0x59, 0x67,
	0x88, 0xaa, 0x65, 0x9d, 0x03, 0xd3, 0xbb, 0x47, 0x46, 0xe2, 0xb3, 0x89, 0xb9, 0x82, 0xc5, 0xf8,
	0x0b, 0xe8, 0xdf, 0x70, 0x1a, 0x17, 0x13, 0xc3, 0x14, 0x14, 0x50, 0xf5, 0x2e, 0x32, 0x98, 0x0e,
	0xf0, 0xd4, 0x73, 0x28, 0x8d, 0x98, 0xdb, 0xcb, 0xa7, 0xbc, 0x69, 0x40, 0xea, 0xa0, 0x3e, 0x3c,
	0xc8, 0x6a, 0x7c, 0x56, 0x49, 0xbf, 0xe7, 0xa0, 0x15, 0xef, 0xc6, 0x1e, 0x7e, 0x44, 0xea, 0xf5,
	0xc5, 0xa5, 0xfc, 0x7b, 0x72, 0xe4, 0x70, 0x11, 0xbf, 0xa7, 0x6d, 0xc3, 0xe6, 0x5d, 0x8a, 0xab,
	0x06, 0x76, 0xa6, 0x2b, 0xd2, 0xbd, 0xfc, 0x99, 0x83, 0x67, 0x45, 0xee, 0x6c, 0x5f, 0x35, 0xd3,
	0xcc, 0x57, 0xf1, 0xaa, 0x9b, 0xdd, 0x29, 0xab, 0x74, 0x3b, 0x1d, 0x58, 0xb2, 0x51, 0xbc, 0x3f,
	0x3a, 0x50, 0xf9, 0x86, 0x0f, 0x45, 0x13, 0x2a, 0x71, 0xad, 0x96, 0xd3, 0x19, 0xbf, 0xe6, 0xe0,
	0xa9, 0x8d, 0x22, 0xe9, 0xe4, 0x9c, 0xe5, 0xa6, 0xbf, 0x55, 0x6e, 0x54, 0x4a, 0xab, 0xf0, 0xed,
	0xa9, 0x6a, 0xf2, 0xdf, 0x41, 0x16, 0x9b, 0x8c, 0xbe, 0x91, 0x6b, 0x7e, 0xf2, 0x3b, 0xd3, 0x15,
	0xe9, 0x5e, 0xce, 0xe0, 0x9e, 0x8d, 0xe2, 0x08, 0xfd, 0x88, 0x93, 0x97, 0xe5, 0x1e, 0xb1, 0xa8,
	0x62, 0x96, 0xab, 0x10, 0x6d, 0xfa, 0x45, 0x9e, 0xbd, 0xa7, 0xe8, 0xd2, 0x00, 0x39, 0x31, 0x9c,
	0x1e, 0x89, 0x5c, 0x75, 0xf3, 0x4a, 0x41, 0xb9, 0xcb, 0xd1, 0x81, 0xef, 0x85, 0xd8, 0xa5, 0xe8,
	0x0b, 0xe3, 0xe5, 0x68, 0x82, 0xd4, 0x5c, 0x8e, 0xd2, 0xa0, 0x8e, 0xb9, 0x06, 0x92, 0x91, 0x8e,
	0x9d, 0x21, 0x72, 0xb2, 0x51, 0xe7, 0x10, 0x63, 0x2a, 0xee, 0x5d, 0x33, 0xb8, 0x78, 0x41, 0x4a,
	0x3f, 0x9c, 0xf9, 0x82, 0x54, 0x7c, 0xbe, 0x8d, 0x46, 0x6c, 0x7a, 0xdf, 0x8b, 0x77, 0xde, 0x90,
	0xb9, 0x51, 0x88, 0x43, 0x99, 0x66, 0x98, 0x50, 0x8a, 0xa9, 0xda, 0xf7, 0xf2, 0x64, 0xfa, 0xb6,
	0x9c, 0x68, 0x3d, 0x3a, 0x40, 0xd3, 0xa9, 0x3a, 0x96, 0x6b, 0x4e, 0x55, 0x4d, 0x15, 0x4f, 0x55,
	0xe5, 0x6f, 0x3e, 0xee, 0x72, 0x11, 0xeb, 0xf5, 0xa0, 0x4e, 0x19, 0xc1, 0xe3, 0x44, 0x3a, 0x66,
	0x82, 0xf6, 0xa8, 0xeb, 0x08, 0xca, 0x7c, 0x4e, 0x36, 0x8d, 0x16, 0x19, 0x4e, 0x25, 0x5a, 0x4d,
	0x71, 0x95, 0xdb, 0xb9, 0x2b, 0xb9, 0xed, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xd7, 0x20,
	0x22, 0xb4, 0x0d, 0x00, 0x00,
}
