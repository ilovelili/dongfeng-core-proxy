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
	// 771 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x96, 0xfb, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x99, 0x10, 0x88, 0x9d, 0x6d, 0x5c, 0xcc, 0x1f, 0x88, 0x81, 0x04, 0x6c, 0xb0, 0x2b,
	0x4b, 0xa5, 0x5d, 0x1e, 0x60, 0x30, 0x29, 0x4c, 0x5a, 0xa7, 0x69, 0x17, 0x06, 0x02, 0x31, 0xd2,
	0xf4, 0x34, 0x35, 0x6a, 0xe3, 0x2c, 0x76, 0x2a, 0xf6, 0x02, 0x48, 0x48, 0x3c, 0x1b, 0xcf, 0x84,
	0xe2, 0xd4, 0x6e, 0x92, 0xc6, 0x49, 0x2a, 0xd1, 0x7f, 0xa6, 0xa9, 0xdf, 0xef, 0x7c, 0xdf, 0xc9,
	0x71, 0x62, 0x1b, 0xf6, 0x3d, 0x2a, 0xba, 0x51, 0xcb, 0x72, 0x59, 0xbf, 0x41, 0x7b, 0x6c, 0x80,
	0x3d, 0xda, 0xa3, 0x8d, 0x36, 0xf3, 0xbd, 0x0e, 0xfa, 0xde, 0x96, 0xcb, 0x42, 0xdc, 0x0a, 0x42,
	0xf6, 0xf3, 0xa6, 0xc1, 0x31, 0x1c, 0x50, 0x17, 0x79, 0x23, 0x08, 0x99, 0x60, 0x0d, 0x27, 0xa0,
	0x96, 0xfc, 0x8f, 0x3c, 0x51, 0xb4, 0xc5, 0x07, 0xae, 0x15, 0x57, 0x58, 0xb2, 0x62, 0x71, 0xaf,
	0xdc, 0x5b, 0x56, 0xb7, 0xa2, 0x8e, 0xfe, 0x25, 0xf1, 0xdb, 0xfe, 0xfb, 0x1c, 0x6e, 0x3b, 0x01,
	0x25, 0x47, 0x70, 0xe7, 0x88, 0x79, 0xd4, 0x27, 0x2f, 0xac, 0x2c, 0xd1, 0x8a, 0x3a, 0x96, 0x54,
	0x4e, 0xf1, 0x3a, 0x42, 0x2e, 0x16, 0x5f, 0x9a, 0x01, 0x1e, 0x30, 0x9f, 0xe3, 0xd2, 0x2d, 0xf2,
	0x09, 0x66, 0x0f, 0x1c, 0xde, 0x6d, 0x31, 0x27, 0x6c, 0x93, 0xe5, 0x82, 0x02, 0xad, 0x2a, 0xd7,
	0xd7, 0xe5, 0x90, 0x76, 0xfe, 0x02, 0x70, 0x11, 0xb4, 0x1d, 0x81, 0x17, 0x1c, 0x43, 0x52, 0x54,
	0x35, 0x92, 0x95, 0xf7, 0x9b, 0x0a, 0x4a, 0x9b, 0x23, 0xdc, 0xb7, 0x51, 0xec, 0x0b, 0x81, 0x7e,
	0xdb, 0xf1, 0x5d, 0xe4, 0x64, 0xb5, 0xa0, 0x34, 0x83, 0xa8, 0x8c, 0xb5, 0x6a, 0x50, 0xc7, 0xf4,
	0xe1, 0x61, 0x12, 0x3f, 0x52, 0xc9, 0x86, 0xb1, 0xc7, 0xf1, 0xac, 0xcd, 0x5a, 0xac, 0x8e, 0xf3,
	0xe1, 0x51, 0x5e, 0xe5, 0xd3, 0xcc, 0xbb, 0x04, 0xb0, 0x51, 0xbc, 0xef, 0x39, 0x9c, 0x23, 0x27,
	0x4b, 0xc5, 0x83, 0x91, 0xb2, 0x0a, 0x58, 0x2e, 0x65, 0xb4, 0xf1, 0x37, 0x98, 0x4b, 0x62, 0xa5,
	0x40, 0xcc, 0xcb, 0x9a, 0x31, 0x5f, 0xa9, 0xc2, 0xb4, 0xff, 0x77, 0x58, 0x48, 0x09, 0x38, 0x85,
	0x84, 0x8f, 0x30, 0x6b, 0xa3, 0x38, 0x89, 0x02, 0xda, 0x33, 0x4e, 0x46, 0xaa, 0x15, 0x93, 0x19,
	0x32, 0xe3, 0x93, 0x91, 0x42, 0x49, 0xdf, 0x19, 0xf3, 0x95, 0x2a, 0x4c, 0xfb, 0x5f, 0xc1, 0x7c,
	0x4a, 0xe0, 0xff, 0x3f, 0xe0, 0x2b, 0xcc, 0xd9, 0x28, 0xce, 0xd1, 0x71, 0xbb, 0x18, 0xf2, 0xc2,
	0xef, 0x7a, 0xa4, 0x97, 0x7d, 0xd7, 0x69, 0x4a, 0xbb, 0xb7, 0xd5, 0xc2, 0x0e, 0xa5, 0xc2, 0xcf,
	0x3a, 0x43, 0x94, 0x7d, 0xd6, 0x39, 0x30, 0xbd, 0x7b, 0x64, 0x24, 0x3e, 0x9d, 0x98, 0x2b, 0x98,
	0x8f, 0xdf, 0x80, 0xee, 0x0d, 0xa7, 0x71, 0x31, 0x31, 0x4c, 0x41, 0x01, 0x65, 0x6b, 0x91, 0xc1,
	0x74, 0x80, 0xa7, 0x9e, 0x43, 0x69, 0xc4, 0xdc, 0x5e, 0x3e, 0x65, 0xbd, 0x06, 0xa9, 0x83, 0xba,
	0xf0, 0x20, 0xab, 0xf1, 0x69, 0x25, 0xfd, 0x9e, 0x81, 0xc5, 0x78, 0x37, 0xf6, 0xf0, 0x03, 0x52,
	0xaf, 0x2b, 0x2e, 0xe5, 0xdf, 0x93, 0xa6, 0xc3, 0x45, 0xbc, 0x4e, 0x3b, 0x86, 0xcd, 0xbb, 0x10,
	0x57, 0x0d, 0xec, 0x4e, 0x56, 0xa4, 0x7b, 0xf9, 0x33, 0x03, 0xcf, 0xc6, 0xb9, 0xb3, 0x03, 0xd5,
	0x4c, 0x3d, 0x5f, 0xc5, 0xab, 0x6e, 0xf6, 0x26, 0xac, 0xd2, 0xed, 0xb4, 0x60, 0xc1, 0x46, 0xf1,
	0xae, 0x79, 0xa8, 0xf2, 0x0d, 0x2f, 0x8a, 0x26, 0x54, 0xe2, 0x6a, 0x25, 0xa7, 0x33, 0x7e, 0xcd,
	0xc0, 0x53, 0x1b, 0x45, 0xd2, 0xc9, 0x39, 0xcb, 0x4d, 0x7f, 0xbb, 0xd8, 0xa8, 0x90, 0x56, 0xe1,
	0x3b, 0x13, 0xd5, 0xe4, 0xdf, 0x83, 0x2c, 0x36, 0x1a, 0x7d, 0x2d, 0xd7, 0xfc, 0xe4, 0x77, 0x27,
	0x2b, 0xd2, 0xbd, 0x9c, 0xc1, 0x3d, 0x1b, 0x45, 0x13, 0xfd, 0x88, 0x93, 0x57, 0xc5, 0x1e, 0xb1,
	0xa8, 0x62, 0x96, 0xca, 0x10, 0x6d, 0xfa, 0x59, 0x9e, 0xbd, 0xa7, 0xe8, 0xd2, 0x00, 0x39, 0x31,
	0x9c, 0x1e, 0x89, 0x5c, 0x76, 0xf3, 0x4a, 0x41, 0xb9, 0xcb, 0xd1, 0xa1, 0xef, 0x85, 0xd8, 0xa6,
	0xe8, 0x0b, 0xe3, 0xe5, 0x68, 0x84, 0x54, 0x5c, 0x8e, 0xd2, 0xa0, 0x8e, 0xb9, 0x06, 0x92, 0x91,
	0x8e, 0x9d, 0x3e, 0x72, 0xb2, 0x59, 0xe5, 0x10, 0x63, 0x2a, 0xee, 0x6d, 0x3d, 0x78, 0xfc, 0x82,
	0x94, 0x7e, 0x38, 0xf3, 0x05, 0x69, 0xfc, 0xf9, 0x36, 0x6b, 0xb1, 0xe9, 0x7d, 0x2f, 0xde, 0x79,
	0x43, 0xe6, 0x46, 0x21, 0xf6, 0x65, 0x9a, 0x61, 0x42, 0x29, 0xa6, 0x6c, 0xdf, 0xcb, 0x93, 0xe9,
	0xdb, 0x72, 0xa2, 0x75, 0x68, 0x0f, 0x4d, 0xa7, 0xea, 0x50, 0xae, 0x38, 0x55, 0x35, 0x95, 0xbe,
	0x74, 0x8c, 0x7e, 0x37, 0x9f, 0x43, 0x43, 0xbd, 0xea, 0x1c, 0xd2, 0xd8, 0xf8, 0xa9, 0xad, 0xfa,
	0x37, 0x1f, 0xa7, 0xb9, 0x47, 0x58, 0xab, 0x06, 0x75, 0xca, 0x0f, 0xb5, 0x18, 0xf1, 0xef, 0xc9,
	0xcb, 0xb6, 0x5e, 0xda, 0xa2, 0x64, 0x54, 0xd2, 0x46, 0x1d, 0xb4, 0x38, 0xeb, 0xc0, 0x11, 0x95,
	0x59, 0x92, 0xa9, 0x97, 0x35, 0x44, 0x75, 0xd6, 0x00, 0x1e, 0x27, 0x8f, 0x7c, 0xcc, 0x04, 0xed,
	0x50, 0xd7, 0x11, 0x94, 0xf9, 0x9c, 0x6c, 0x19, 0x47, 0x93, 0xe1, 0x54, 0xa6, 0x55, 0x17, 0x57,
	0xb9, 0xad, 0xbb, 0x92, 0xdb, 0xf9, 0x17, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xb2, 0xb1, 0xa8, 0xec,
	0x0e, 0x00, 0x00,
}
