// Sorapointa - A server software re-implementation for a certain anime game, and avoid sorapointa.
// Copyright (C) 2022  Sorapointa Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: GCGMsgOpTimer.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GCGMsgOpTimer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeginTime    uint64       `protobuf:"varint,9,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
	Phase        GCGPhaseType `protobuf:"varint,3,opt,name=phase,proto3,enum=proto.GCGPhaseType" json:"phase,omitempty"`
	TimeStamp    uint64       `protobuf:"varint,13,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	ControllerId uint32       `protobuf:"varint,8,opt,name=controller_id,json=controllerId,proto3" json:"controller_id,omitempty"`
}

func (x *GCGMsgOpTimer) Reset() {
	*x = GCGMsgOpTimer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GCGMsgOpTimer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCGMsgOpTimer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCGMsgOpTimer) ProtoMessage() {}

func (x *GCGMsgOpTimer) ProtoReflect() protoreflect.Message {
	mi := &file_GCGMsgOpTimer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCGMsgOpTimer.ProtoReflect.Descriptor instead.
func (*GCGMsgOpTimer) Descriptor() ([]byte, []int) {
	return file_GCGMsgOpTimer_proto_rawDescGZIP(), []int{0}
}

func (x *GCGMsgOpTimer) GetBeginTime() uint64 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

func (x *GCGMsgOpTimer) GetPhase() GCGPhaseType {
	if x != nil {
		return x.Phase
	}
	return GCGPhaseType_GCG_PHASE_TYPE_INVALID
}

func (x *GCGMsgOpTimer) GetTimeStamp() uint64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

func (x *GCGMsgOpTimer) GetControllerId() uint32 {
	if x != nil {
		return x.ControllerId
	}
	return 0
}

var File_GCGMsgOpTimer_proto protoreflect.FileDescriptor

var file_GCGMsgOpTimer_proto_rawDesc = []byte{
	0x0a, 0x13, 0x47, 0x43, 0x47, 0x4d, 0x73, 0x67, 0x4f, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x47, 0x43,
	0x47, 0x50, 0x68, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9d, 0x01, 0x0a, 0x0d, 0x47, 0x43, 0x47, 0x4d, 0x73, 0x67, 0x4f, 0x70, 0x54, 0x69, 0x6d,
	0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x29, 0x0a, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x43, 0x47, 0x50, 0x68, 0x61, 0x73,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GCGMsgOpTimer_proto_rawDescOnce sync.Once
	file_GCGMsgOpTimer_proto_rawDescData = file_GCGMsgOpTimer_proto_rawDesc
)

func file_GCGMsgOpTimer_proto_rawDescGZIP() []byte {
	file_GCGMsgOpTimer_proto_rawDescOnce.Do(func() {
		file_GCGMsgOpTimer_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGMsgOpTimer_proto_rawDescData)
	})
	return file_GCGMsgOpTimer_proto_rawDescData
}

var file_GCGMsgOpTimer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GCGMsgOpTimer_proto_goTypes = []interface{}{
	(*GCGMsgOpTimer)(nil), // 0: proto.GCGMsgOpTimer
	(GCGPhaseType)(0),     // 1: proto.GCGPhaseType
}
var file_GCGMsgOpTimer_proto_depIdxs = []int32{
	1, // 0: proto.GCGMsgOpTimer.phase:type_name -> proto.GCGPhaseType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GCGMsgOpTimer_proto_init() }
func file_GCGMsgOpTimer_proto_init() {
	if File_GCGMsgOpTimer_proto != nil {
		return
	}
	file_GCGPhaseType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GCGMsgOpTimer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCGMsgOpTimer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GCGMsgOpTimer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGMsgOpTimer_proto_goTypes,
		DependencyIndexes: file_GCGMsgOpTimer_proto_depIdxs,
		MessageInfos:      file_GCGMsgOpTimer_proto_msgTypes,
	}.Build()
	File_GCGMsgOpTimer_proto = out.File
	file_GCGMsgOpTimer_proto_rawDesc = nil
	file_GCGMsgOpTimer_proto_goTypes = nil
	file_GCGMsgOpTimer_proto_depIdxs = nil
}