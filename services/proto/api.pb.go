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
	// 857 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0xed, 0x4e, 0xd4, 0x4a,
	0x18, 0xc7, 0x0f, 0x39, 0x39, 0x27, 0x87, 0x01, 0x8e, 0x3a, 0x7e, 0x30, 0xe2, 0x07, 0x05, 0x14,
	0x50, 0xd8, 0x6e, 0xc2, 0xcb, 0x05, 0x20, 0x98, 0x95, 0x04, 0x70, 0xc3, 0x8b, 0x60, 0x34, 0x62,
	0xdb, 0x7d, 0xb6, 0x3b, 0xb1, 0xdb, 0x29, 0x9d, 0xd9, 0x0d, 0xdc, 0x80, 0x89, 0x89, 0xb7, 0xe6,
	0x3d, 0x99, 0x4e, 0x77, 0x66, 0xa7, 0x2f, 0xd3, 0x76, 0x13, 0xd7, 0x2f, 0x84, 0xec, 0xff, 0xf7,
	0xfc, 0xff, 0x4f, 0x67, 0xda, 0x99, 0x07, 0xed, 0x79, 0x84, 0xf7, 0x06, 0x8e, 0xe5, 0xd2, 0x7e,
	0x93, 0xf8, 0x74, 0x08, 0x3e, 0xf1, 0x49, 0xb3, 0x43, 0x03, 0xaf, 0x0b, 0x81, 0xd7, 0x70, 0x69,
	0x04, 0x8d, 0x30, 0xa2, 0xb7, 0x77, 0x4d, 0x06, 0xd1, 0x90, 0xb8, 0xc0, 0x9a, 0x61, 0x44, 0x39,
	0x6d, 0xda, 0x21, 0xb1, 0xc4, 0x7f, 0xf8, 0x91, 0xa4, 0x2d, 0x36, 0x74, 0xad, 0xb8, 0xc2, 0x12,
	0x15, 0x8b, 0xbb, 0xe5, 0xde, 0xa2, 0xda, 0x19, 0x74, 0xd5, 0x2f, 0x89, 0xdf, 0xd6, 0xcf, 0x25,
	0xf4, 0xb7, 0x1d, 0x12, 0x7c, 0x84, 0xfe, 0x39, 0xa2, 0x1e, 0x09, 0xf0, 0x53, 0x2b, 0x4d, 0x38,
	0x83, 0xae, 0x25, 0x94, 0x53, 0xb8, 0x19, 0x00, 0xe3, 0x8b, 0xcf, 0xcc, 0x00, 0x0b, 0x69, 0xc0,
	0x60, 0xf9, 0x2f, 0x7c, 0x85, 0x66, 0x0f, 0x6c, 0xd6, 0x73, 0xa8, 0x1d, 0x75, 0xf0, 0x4a, 0x41,
	0x81, 0x52, 0xa5, 0xeb, 0xf3, 0x72, 0x48, 0x39, 0x7f, 0x44, 0xe8, 0x22, 0xec, 0xd8, 0x1c, 0x2e,
	0x18, 0x44, 0xb8, 0xa8, 0x6a, 0x2c, 0x4b, 0xef, 0x17, 0x15, 0x94, 0x32, 0x07, 0xf4, 0x7f, 0x0b,
	0xf8, 0x1e, 0xe7, 0x10, 0x74, 0xec, 0xc0, 0x05, 0x86, 0xd7, 0x0a, 0x4a, 0x53, 0x88, 0xcc, 0x58,
	0xaf, 0x06, 0x55, 0x4c, 0x1f, 0xdd, 0x4f, 0xe2, 0xc7, 0x2a, 0x7e, 0x65, 0xec, 0x31, 0x9f, 0xb5,
	0x51, 0x8b, 0x55, 0x71, 0x01, 0x7a, 0x90, 0x55, 0xd9, 0x34, 0xf3, 0x2e, 0x11, 0x6a, 0x01, 0xdf,
	0xf7, 0x6d, 0xc6, 0x80, 0xe1, 0xe5, 0xe2, 0x85, 0x11, 0xb2, 0x0c, 0x58, 0x29, 0x65, 0x94, 0xf1,
	0x67, 0x34, 0x97, 0xc4, 0x0a, 0x01, 0x9b, 0xb7, 0x35, 0x65, 0xbe, 0x5a, 0x85, 0x29, 0xff, 0x2f,
	0x68, 0x41, 0x13, 0x60, 0x0a, 0x09, 0xef, 0xd1, 0x6c, 0x0b, 0x78, 0x7b, 0x10, 0x12, 0xdf, 0xb8,
	0x32, 0x42, 0xad, 0x58, 0x99, 0x11, 0x93, 0x5f, 0x19, 0x21, 0x94, 0xf4, 0x9d, 0x32, 0x5f, 0xad,
	0xc2, 0x94, 0xff, 0x35, 0x9a, 0xd7, 0x04, 0xf6, 0xfb, 0x03, 0x3e, 0xa1, 0xb9, 0x16, 0xf0, 0x73,
	0xb0, 0xdd, 0x1e, 0x44, 0xac, 0xf0, 0xbb, 0x1e, 0xeb, 0x65, 0xdf, 0xb5, 0x4e, 0x29, 0xf7, 0x8e,
	0xdc, 0xd8, 0x91, 0x54, 0xf8, 0x59, 0xa7, 0x88, 0xb2, 0xcf, 0x3a, 0x03, 0xea, 0xa7, 0x47, 0x4a,
	0x62, 0xd3, 0x89, 0xb9, 0x46, 0xf3, 0xf1, 0x1b, 0xd0, 0xbb, 0x63, 0x24, 0x2e, 0xc6, 0x86, 0x55,
	0x90, 0x40, 0xd9, 0x5e, 0xa4, 0x30, 0x15, 0xe0, 0xc9, 0xe7, 0x90, 0x1a, 0x36, 0xb7, 0x97, 0x4d,
	0x79, 0x59, 0x83, 0x54, 0x41, 0x3d, 0x74, 0x2f, 0xad, 0xb1, 0x69, 0x25, 0x7d, 0x9f, 0x41, 0x8b,
	0xf1, 0x69, 0xec, 0xc1, 0x5b, 0x20, 0x5e, 0x8f, 0x5f, 0x8a, 0xbf, 0xed, 0x63, 0x9b, 0xf1, 0x78,
	0x9f, 0xb6, 0x0d, 0x87, 0x77, 0x21, 0x2e, 0x1b, 0xd8, 0x99, 0xac, 0x48, 0xf5, 0xf2, 0x63, 0x06,
	0x3d, 0xc9, 0x73, 0x67, 0x07, 0xb2, 0x99, 0x7a, 0xbe, 0x92, 0x97, 0xdd, 0xec, 0x4e, 0x58, 0xa5,
	0xda, 0x71, 0xd0, 0x42, 0x0b, 0xf8, 0xeb, 0xe3, 0x43, 0x99, 0x6f, 0x78, 0x51, 0x14, 0x21, 0x13,
	0xd7, 0x2a, 0x39, 0x95, 0xf1, 0x6d, 0x06, 0x3d, 0x6e, 0x01, 0x4f, 0x3a, 0x39, 0xa7, 0x99, 0xd5,
	0xdf, 0x2a, 0x36, 0x2a, 0xa4, 0x65, 0xf8, 0xf6, 0x44, 0x35, 0xd9, 0xf7, 0x20, 0x8d, 0x8d, 0x97,
	0xbe, 0x96, 0x6b, 0x76, 0xe5, 0x77, 0x26, 0x2b, 0x52, 0xbd, 0x9c, 0xa1, 0xff, 0x5a, 0xc0, 0x8f,
	0x21, 0x18, 0x30, 0xbc, 0x54, 0xec, 0x11, 0x8b, 0x32, 0x66, 0xb9, 0x0c, 0x51, 0xa6, 0x1f, 0xc4,
	0xdd, 0x7b, 0x0a, 0x2e, 0x09, 0x81, 0x61, 0xc3, 0xed, 0x91, 0xc8, 0x65, 0x93, 0x97, 0x06, 0x65,
	0x86, 0xa3, 0xc3, 0xc0, 0x8b, 0xa0, 0x43, 0x20, 0xe0, 0xc6, 0xe1, 0x68, 0x8c, 0x54, 0x0c, 0x47,
	0x3a, 0xa8, 0x62, 0x6e, 0x10, 0x4e, 0x49, 0x27, 0x76, 0x1f, 0x18, 0xde, 0xa8, 0x72, 0x88, 0x31,
	0x19, 0xb7, 0x59, 0x0f, 0xce, 0x0f, 0x48, 0xfa, 0xc3, 0x99, 0x07, 0xa4, 0xfc, 0xf3, 0x6d, 0xd4,
	0x62, 0xf5, 0x73, 0x2f, 0x3e, 0x79, 0x23, 0xea, 0x0e, 0x22, 0xe8, 0x8b, 0x34, 0xc3, 0x0a, 0x69,
	0x4c, 0xd9, 0xb9, 0x97, 0x25, 0x55, 0x52, 0x28, 0x9f, 0x4c, 0x93, 0xb1, 0xb9, 0xdb, 0x82, 0xb8,
	0xcd, 0x7a, 0xb0, 0x3e, 0x9f, 0x27, 0xdd, 0x74, 0x89, 0x0f, 0xa6, 0x7b, 0x7c, 0x24, 0x57, 0xdc,
	0xe3, 0x8a, 0xd2, 0xc7, 0x9c, 0xf1, 0xef, 0xe6, 0x9b, 0x6f, 0xa4, 0x57, 0xdd, 0x7c, 0x0a, 0x53,
	0xfe, 0x4c, 0xbc, 0xe2, 0xed, 0x08, 0x86, 0xf2, 0x01, 0x2c, 0x53, 0x2d, 0x0c, 0xdf, 0x45, 0x27,
	0x70, 0x9b, 0x7d, 0x94, 0x66, 0x6d, 0x3e, 0x13, 0xaa, 0x69, 0x7f, 0x22, 0xb4, 0x83, 0x16, 0xf6,
	0x23, 0x48, 0x76, 0x51, 0x64, 0xae, 0x95, 0xed, 0xb3, 0x1e, 0xb6, 0x5e, 0x0d, 0xe6, 0xe7, 0xae,
	0x69, 0xa7, 0x1c, 0x80, 0x0f, 0x53, 0x4e, 0xb9, 0x12, 0xa3, 0xfb, 0x1b, 0x87, 0xd2, 0xaf, 0xc6,
	0x83, 0x35, 0x51, 0x2b, 0x0e, 0x56, 0x09, 0xe5, 0x87, 0x77, 0xa1, 0x94, 0xcc, 0xd6, 0x42, 0xaf,
	0x9e, 0xad, 0x47, 0x98, 0xf2, 0x1f, 0xa2, 0x87, 0x89, 0x70, 0x42, 0x39, 0xe9, 0x12, 0xd7, 0xe6,
	0x84, 0x06, 0x0c, 0x37, 0x8c, 0x06, 0x29, 0x4e, 0xe6, 0x59, 0x75, 0x71, 0xed, 0xc5, 0x8e, 0x4f,
	0xf2, 0x3d, 0xd7, 0x05, 0xc6, 0x88, 0xe3, 0x43, 0xdb, 0xe6, 0x3d, 0x86, 0x0d, 0x87, 0x73, 0x06,
	0x93, 0xa9, 0x8d, 0x9a, 0xb4, 0x0c, 0x75, 0xfe, 0x15, 0xd8, 0xf6, 0xaf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xfd, 0x5b, 0x4b, 0x71, 0x6b, 0x11, 0x00, 0x00,
}
