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
	// 680 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x96, 0xed, 0x6a, 0xd4, 0x40,
	0x14, 0x86, 0x2d, 0xa2, 0xd8, 0xd3, 0xd6, 0x8f, 0xf1, 0x47, 0xb1, 0xfe, 0x50, 0xb7, 0xda, 0xaa,
	0x65, 0xb3, 0xd0, 0x6d, 0x2f, 0xa0, 0x5a, 0x88, 0x85, 0xae, 0x94, 0x7e, 0x58, 0x41, 0xb1, 0xe6,
	0xe3, 0x6c, 0x32, 0xb0, 0x9b, 0x89, 0x99, 0xc9, 0x62, 0x6f, 0x40, 0x10, 0xbc, 0x44, 0x2f, 0x46,
	0x92, 0x74, 0x66, 0x93, 0x6c, 0x26, 0xc9, 0x82, 0xfb, 0xa7, 0x94, 0x7d, 0x9f, 0xf3, 0xbe, 0x67,
	0x4f, 0x32, 0x67, 0x07, 0x0e, 0x3c, 0x2a, 0xfc, 0xd8, 0x36, 0x1c, 0x36, 0xee, 0xd1, 0x11, 0x9b,
	0xe0, 0x88, 0x8e, 0x68, 0xcf, 0x65, 0x81, 0x37, 0xc4, 0xc0, 0xeb, 0x3a, 0x2c, 0xc2, 0x6e, 0x18,
	0xb1, 0x9f, 0xd7, 0x3d, 0x8e, 0xd1, 0x84, 0x3a, 0xc8, 0x7b, 0x61, 0xc4, 0x04, 0xeb, 0x59, 0x21,
	0x35, 0xd2, 0xff, 0xc8, 0xba, 0xa4, 0x0d, 0x3e, 0x71, 0x8c, 0xa4, 0xc2, 0x48, 0x2b, 0x36, 0xf6,
	0xeb, 0xbd, 0xd3, 0x6a, 0x3b, 0x1e, 0xaa, 0x4f, 0x32, 0xbf, 0xdd, 0xbf, 0xeb, 0x70, 0xdb, 0x0a,
	0x29, 0x39, 0x86, 0x3b, 0xc7, 0xcc, 0xa3, 0x01, 0x79, 0x66, 0x14, 0x09, 0x3b, 0x1e, 0x1a, 0xa9,
	0x72, 0x8a, 0x3f, 0x62, 0xe4, 0x62, 0xe3, 0xb9, 0x1e, 0xe0, 0x21, 0x0b, 0x38, 0x76, 0x6e, 0x91,
	0xcf, 0xb0, 0x7c, 0x68, 0x71, 0xdf, 0x66, 0x56, 0xe4, 0x92, 0xcd, 0x8a, 0x02, 0xa5, 0x4a, 0xd7,
	0x97, 0xf5, 0x90, 0x72, 0xfe, 0x02, 0x70, 0x11, 0xba, 0x96, 0xc0, 0x0b, 0x8e, 0x11, 0xa9, 0xaa,
	0x9a, 0xca, 0xd2, 0xfb, 0x55, 0x03, 0xa5, 0xcc, 0x11, 0xee, 0x9b, 0x28, 0x0e, 0x84, 0xc0, 0xc0,
	0xb5, 0x02, 0x07, 0x39, 0xd9, 0xae, 0x28, 0x2d, 0x20, 0x32, 0xe3, 0x75, 0x33, 0xa8, 0x62, 0xc6,
	0xf0, 0x30, 0x8b, 0x9f, 0xaa, 0xe4, 0xad, 0xb6, 0xc7, 0xd9, 0xac, 0x9d, 0x56, 0xac, 0x8a, 0x0b,
	0xe0, 0x51, 0x59, 0xe5, 0x8b, 0xcc, 0xbb, 0x04, 0x30, 0x51, 0xbc, 0x1f, 0x59, 0x9c, 0x23, 0x27,
	0x9d, 0xea, 0xc1, 0xa4, 0xb2, 0x0c, 0xd8, 0xac, 0x65, 0x94, 0xf1, 0x37, 0x58, 0xc9, 0x62, 0x53,
	0x81, 0xe8, 0x1f, 0x6b, 0xc1, 0x7c, 0xab, 0x09, 0x53, 0xfe, 0xdf, 0x61, 0x2d, 0x27, 0xe0, 0x02,
	0x12, 0x3e, 0xc1, 0xb2, 0x89, 0xe2, 0x24, 0x0e, 0xe9, 0x48, 0x3b, 0x99, 0x54, 0x6d, 0x98, 0xcc,
	0x0d, 0x33, 0x3b, 0x99, 0x54, 0xa8, 0xe9, 0xbb, 0x60, 0xbe, 0xd5, 0x84, 0x29, 0xff, 0x2b, 0x58,
	0xcd, 0x09, 0xfc, 0xff, 0x07, 0x7c, 0x85, 0x15, 0x13, 0xc5, 0x39, 0x5a, 0x8e, 0x8f, 0x11, 0xaf,
	0x3c, 0xd7, 0x53, 0xbd, 0xee, 0x5c, 0xe7, 0x29, 0xe5, 0xee, 0xca, 0x07, 0x7b, 0x23, 0x55, 0x1e,
	0xeb, 0x02, 0x51, 0x77, 0xac, 0x4b, 0x60, 0x7e, 0x7b, 0x14, 0x24, 0xbe, 0x98, 0x98, 0x2b, 0x58,
	0x4d, 0xde, 0x00, 0xff, 0x9a, 0xd3, 0xa4, 0x98, 0x68, 0xa6, 0x20, 0x81, 0xba, 0x67, 0x51, 0xc0,
	0x54, 0x80, 0x27, 0xbf, 0x87, 0xd4, 0x88, 0xbe, 0xbd, 0x72, 0xca, 0x9b, 0x16, 0xa4, 0x0a, 0xf2,
	0xe1, 0x41, 0x51, 0xe3, 0x8b, 0x4a, 0xfa, 0xbd, 0x04, 0x1b, 0xc9, 0x36, 0xf6, 0xf0, 0x03, 0x52,
	0xcf, 0x17, 0x97, 0xe9, 0xdf, 0x93, 0x81, 0xc5, 0x45, 0xf2, 0x9c, 0xfa, 0x9a, 0xe5, 0x5d, 0x89,
	0xcb, 0x06, 0xf6, 0xe6, 0x2b, 0x52, 0xbd, 0xfc, 0x59, 0x82, 0xa7, 0xb3, 0xdc, 0xd9, 0xa1, 0x6c,
	0xa6, 0x9d, 0xaf, 0xe4, 0x65, 0x37, 0xfb, 0x73, 0x56, 0xa9, 0x76, 0x6c, 0x58, 0x33, 0x51, 0xbc,
	0x1b, 0x1c, 0xc9, 0x7c, 0xcd, 0x8b, 0xa2, 0x08, 0x99, 0xb8, 0xdd, 0xc8, 0xa9, 0x8c, 0x5f, 0x4b,
	0xf0, 0xc4, 0x44, 0x91, 0x75, 0x72, 0xce, 0x4a, 0xd3, 0xdf, 0xad, 0x36, 0xaa, 0xa4, 0x65, 0x78,
	0x7f, 0xae, 0x9a, 0xf2, 0x7b, 0x50, 0xc4, 0xa6, 0xa3, 0x6f, 0xe5, 0x5a, 0x9e, 0xfc, 0xde, 0x7c,
	0x45, 0xaa, 0x97, 0x33, 0xb8, 0x67, 0xa2, 0x18, 0x60, 0x10, 0x73, 0xf2, 0xa2, 0xda, 0x23, 0x11,
	0x65, 0x4c, 0xa7, 0x0e, 0xc9, 0x5f, 0xbc, 0x4c, 0x14, 0xa7, 0xe8, 0xd0, 0x10, 0x89, 0xe6, 0xc7,
	0x23, 0x53, 0xeb, 0x2e, 0x5e, 0x39, 0x28, 0xbf, 0x43, 0x4d, 0x14, 0x47, 0x81, 0x17, 0xa1, 0x4b,
	0x31, 0x10, 0xba, 0xab, 0xd1, 0x94, 0x68, 0xb8, 0x1a, 0xe5, 0xc1, 0xd9, 0xab, 0x51, 0x2e, 0x48,
	0x7f, 0x55, 0x99, 0xcd, 0xda, 0x69, 0xc5, 0xe6, 0x57, 0x5d, 0xb2, 0x03, 0x23, 0xe6, 0xc4, 0x11,
	0x8e, 0x93, 0x30, 0x4d, 0xb3, 0x39, 0xa4, 0x6e, 0x01, 0x95, 0x49, 0x15, 0x34, 0x81, 0xc7, 0x59,
	0x1b, 0x1f, 0x99, 0xa0, 0x43, 0xea, 0x58, 0x82, 0xb2, 0x80, 0x93, 0xae, 0xb6, 0xdd, 0x02, 0x27,
	0x23, 0x8d, 0xb6, 0xb8, 0xcc, 0xb5, 0xef, 0xa6, 0x5c, 0xff, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xbd, 0x87, 0x81, 0xae, 0x7a, 0x0c, 0x00, 0x00,
}
