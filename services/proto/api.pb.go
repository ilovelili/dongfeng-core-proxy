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
	// 694 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x96, 0xed, 0x6a, 0xd4, 0x4e,
	0x14, 0xc6, 0xff, 0xe5, 0x8f, 0x4a, 0x4f, 0x5b, 0x5f, 0xc6, 0x0f, 0x6a, 0xfd, 0xa0, 0xb6, 0xda,
	0xaa, 0x65, 0xb3, 0xd0, 0x6d, 0x2f, 0xa0, 0x5a, 0x88, 0x85, 0xae, 0x94, 0xb5, 0xb5, 0x82, 0x62,
	0xcd, 0xcb, 0xd9, 0x64, 0x60, 0x37, 0x13, 0x33, 0x93, 0xc5, 0xde, 0x80, 0x20, 0x78, 0x4d, 0x5e,
	0x9b, 0x24, 0xd9, 0x99, 0x4d, 0xb2, 0x99, 0x24, 0x0b, 0xee, 0x97, 0x52, 0xf6, 0xf9, 0x9d, 0xe7,
	0x39, 0x39, 0xc9, 0x1c, 0x06, 0x8e, 0x3c, 0x2a, 0xfc, 0xd8, 0x36, 0x1c, 0x36, 0xee, 0xd2, 0x11,
	0x9b, 0xe0, 0x88, 0x8e, 0x68, 0xd7, 0x65, 0x81, 0x37, 0xc4, 0xc0, 0xeb, 0x38, 0x2c, 0xc2, 0x4e,
	0x18, 0xb1, 0x1f, 0xd7, 0x5d, 0x8e, 0xd1, 0x84, 0x3a, 0xc8, 0xbb, 0x61, 0xc4, 0x04, 0xeb, 0x5a,
	0x21, 0x35, 0xd2, 0xff, 0xc8, 0x03, 0x49, 0x1b, 0x7c, 0xe2, 0x18, 0x49, 0x85, 0x91, 0x56, 0x6c,
	0x1e, 0xd6, 0x7b, 0xa7, 0xd5, 0x76, 0x3c, 0x54, 0xbf, 0x64, 0x7e, 0xfb, 0x7f, 0x1e, 0xc2, 0xff,
	0x56, 0x48, 0xc9, 0x29, 0xdc, 0x38, 0x65, 0x1e, 0x0d, 0xc8, 0x13, 0xa3, 0x48, 0xd8, 0xf1, 0xd0,
	0x48, 0x95, 0x01, 0x7e, 0x8f, 0x91, 0x8b, 0xcd, 0xa7, 0x7a, 0x80, 0x87, 0x2c, 0xe0, 0xb8, 0xf5,
	0x1f, 0xf9, 0x04, 0xab, 0xc7, 0x16, 0xf7, 0x6d, 0x66, 0x45, 0x2e, 0xd9, 0xae, 0x28, 0x50, 0xaa,
	0x74, 0x7d, 0x5e, 0x0f, 0x29, 0xe7, 0xcf, 0x00, 0x17, 0xa1, 0x6b, 0x09, 0xbc, 0xe0, 0x18, 0x91,
	0xaa, 0xaa, 0x99, 0x2c, 0xbd, 0x5f, 0x34, 0x50, 0xca, 0x1c, 0xe1, 0xb6, 0x89, 0xe2, 0x48, 0x08,
	0x0c, 0x5c, 0x2b, 0x70, 0x90, 0x93, 0xdd, 0x8a, 0xd2, 0x02, 0x22, 0x33, 0x5e, 0x36, 0x83, 0x2a,
	0x66, 0x0c, 0x77, 0xb3, 0xf8, 0x99, 0x4a, 0x5e, 0x6b, 0x7b, 0x9c, 0xcf, 0xda, 0x6b, 0xc5, 0xaa,
	0xb8, 0x00, 0xee, 0x95, 0x55, 0xbe, 0xcc, 0xbc, 0x4b, 0x00, 0x13, 0xc5, 0xdb, 0x91, 0xc5, 0x39,
	0x72, 0xb2, 0x55, 0x3d, 0x98, 0x54, 0x96, 0x01, 0xdb, 0xb5, 0x8c, 0x32, 0xfe, 0x0a, 0x6b, 0x59,
	0x6c, 0x2a, 0x10, 0xfd, 0x6b, 0x2d, 0x98, 0xef, 0x34, 0x61, 0xca, 0xff, 0x1b, 0x6c, 0xe4, 0x04,
	0x5c, 0x42, 0xc2, 0x47, 0x58, 0x35, 0x51, 0x9c, 0xc5, 0x21, 0x1d, 0x69, 0x27, 0x93, 0xaa, 0x0d,
	0x93, 0x99, 0x32, 0xf3, 0x93, 0x49, 0x85, 0x9a, 0xbe, 0x0b, 0xe6, 0x3b, 0x4d, 0x98, 0xf2, 0xbf,
	0x82, 0xf5, 0x9c, 0xc0, 0xff, 0x7d, 0xc0, 0x17, 0x58, 0x33, 0x51, 0x9c, 0xa3, 0xe5, 0xf8, 0x18,
	0xf1, 0xca, 0x73, 0x3d, 0xd3, 0xeb, 0xce, 0x75, 0x9e, 0x52, 0xee, 0xae, 0x7c, 0xb1, 0x53, 0xa9,
	0xf2, 0x58, 0x17, 0x88, 0xba, 0x63, 0x5d, 0x02, 0xf3, 0xdb, 0xa3, 0x20, 0xf1, 0xe5, 0xc4, 0x5c,
	0xc1, 0x7a, 0xf2, 0x05, 0xf8, 0xd7, 0x9c, 0x26, 0xc5, 0x44, 0x33, 0x05, 0x09, 0xd4, 0xbd, 0x8b,
	0x02, 0xa6, 0x02, 0x3c, 0xf9, 0x1c, 0x52, 0x23, 0xfa, 0xf6, 0xca, 0x29, 0xaf, 0x5a, 0x90, 0x2a,
	0xc8, 0x87, 0x3b, 0x45, 0x8d, 0x2f, 0x2b, 0xe9, 0xd7, 0x0a, 0x6c, 0x26, 0xdb, 0xd8, 0xc3, 0x77,
	0x48, 0x3d, 0x5f, 0x5c, 0xa6, 0x7f, 0xcf, 0xfa, 0x16, 0x17, 0xc9, 0x7b, 0xea, 0x69, 0x96, 0x77,
	0x25, 0x2e, 0x1b, 0x38, 0x58, 0xac, 0x48, 0xf5, 0xf2, 0x7b, 0x05, 0x1e, 0xcf, 0x73, 0x1f, 0x8e,
	0x65, 0x33, 0xed, 0x7c, 0x25, 0x2f, 0xbb, 0x39, 0x5c, 0xb0, 0x4a, 0xb5, 0x63, 0xc3, 0x86, 0x89,
	0xe2, 0x4d, 0xff, 0x44, 0xe6, 0x6b, 0x3e, 0x14, 0x45, 0xc8, 0xc4, 0xdd, 0x46, 0x4e, 0x65, 0xfc,
	0x5c, 0x81, 0x47, 0x26, 0x8a, 0xac, 0x93, 0x73, 0x56, 0x9a, 0xfe, 0x7e, 0xb5, 0x51, 0x25, 0x2d,
	0xc3, 0x7b, 0x0b, 0xd5, 0x94, 0xbf, 0x83, 0x22, 0x36, 0x1b, 0x7d, 0x2b, 0xd7, 0xf2, 0xe4, 0x0f,
	0x16, 0x2b, 0xca, 0xdf, 0x91, 0x4c, 0x14, 0x03, 0x74, 0x68, 0x88, 0x44, 0xb3, 0xe7, 0x33, 0xb5,
	0xee, 0x8e, 0x94, 0x83, 0x94, 0xb3, 0x25, 0xb7, 0xf5, 0xd4, 0x5c, 0xbf, 0x86, 0x8b, 0xfe, 0xbb,
	0x8d, 0x5c, 0x7e, 0xa3, 0x9a, 0x28, 0x4e, 0x02, 0x2f, 0x42, 0x97, 0x62, 0x20, 0x74, 0x17, 0xa5,
	0x19, 0xd1, 0x70, 0x51, 0xca, 0x83, 0xf3, 0x17, 0xa5, 0x5c, 0x90, 0xfe, 0xe2, 0x32, 0x9f, 0xb5,
	0xd7, 0x8a, 0x55, 0x71, 0x03, 0xb8, 0x65, 0xa2, 0xe8, 0x63, 0x10, 0x93, 0x67, 0xd5, 0x5d, 0x26,
	0x9a, 0x34, 0xdf, 0xaa, 0x43, 0xf2, 0xcb, 0x34, 0xd9, 0xb2, 0x11, 0x73, 0xe2, 0x08, 0xc7, 0xc9,
	0x03, 0x68, 0x06, 0x90, 0x43, 0xea, 0x56, 0x5c, 0x99, 0x54, 0x41, 0x13, 0xb8, 0x9f, 0x3d, 0xda,
	0x7b, 0x26, 0xe8, 0x90, 0x3a, 0x96, 0xa0, 0x2c, 0xe0, 0xa4, 0xa3, 0x1d, 0x41, 0x81, 0x93, 0x91,
	0x46, 0x5b, 0x5c, 0xe6, 0xda, 0x37, 0x53, 0xae, 0xf7, 0x37, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x52,
	0x17, 0x74, 0xdc, 0x0c, 0x00, 0x00,
}
