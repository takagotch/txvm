// Code generated by protoc-gen-go.
// source: rawsnapshot.proto
// DO NOT EDIT!

/*
Package state is a generated protocol buffer package.

It is generated from these files:
	rawsnapshot.proto

It has these top-level messages:
	RawSnapshot
*/
package state

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import bc "github.com/chain/txvm/protocol/bc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Snapshot represents a snapshot of the blockchain, including the contracts
// tree and issuance memory.
type RawSnapshot struct {
	// ContractNodes contains every leaf node within the contracts tree.
	// The nodes are ordered according to a pre-order traversal.
	ContractNodes  [][]byte        `protobuf:"bytes,1,rep,name=contract_nodes,json=contractNodes,proto3" json:"contract_nodes,omitempty"`
	NonceNodes     [][]byte        `protobuf:"bytes,2,rep,name=nonce_nodes,json=nonceNodes,proto3" json:"nonce_nodes,omitempty"`
	Header         *bc.BlockHeader `protobuf:"bytes,3,opt,name=header" json:"header,omitempty"`
	InitialBlockId *bc.Hash        `protobuf:"bytes,4,opt,name=initial_block_id,json=initialBlockId" json:"initial_block_id,omitempty"`
	RefIds         []*bc.Hash      `protobuf:"bytes,5,rep,name=ref_ids,json=refIds" json:"ref_ids,omitempty"`
}

func (m *RawSnapshot) Reset()                    { *m = RawSnapshot{} }
func (m *RawSnapshot) String() string            { return proto.CompactTextString(m) }
func (*RawSnapshot) ProtoMessage()               {}
func (*RawSnapshot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RawSnapshot) GetContractNodes() [][]byte {
	if m != nil {
		return m.ContractNodes
	}
	return nil
}

func (m *RawSnapshot) GetNonceNodes() [][]byte {
	if m != nil {
		return m.NonceNodes
	}
	return nil
}

func (m *RawSnapshot) GetHeader() *bc.BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *RawSnapshot) GetInitialBlockId() *bc.Hash {
	if m != nil {
		return m.InitialBlockId
	}
	return nil
}

func (m *RawSnapshot) GetRefIds() []*bc.Hash {
	if m != nil {
		return m.RefIds
	}
	return nil
}

func init() {
	proto.RegisterType((*RawSnapshot)(nil), "chain.protocol.state.RawSnapshot")
}

func init() { proto.RegisterFile("rawsnapshot.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xa9, 0x75, 0xbb, 0x92, 0xea, 0xaa, 0xc1, 0x43, 0x11, 0xc1, 0x2a, 0x88, 0x7b, 0x4a,
	0x75, 0x7d, 0x83, 0x3d, 0xed, 0x5e, 0x3c, 0xd4, 0x9b, 0x97, 0x32, 0x99, 0x64, 0x69, 0xb0, 0x64,
	0x96, 0x24, 0xb0, 0x4f, 0xea, 0xfb, 0x48, 0xd3, 0xf4, 0xe0, 0x71, 0xbe, 0xff, 0xfb, 0x07, 0x7e,
	0x76, 0xeb, 0xe0, 0xe4, 0x2d, 0x1c, 0x7d, 0x4f, 0x41, 0x1c, 0x1d, 0x05, 0xe2, 0x77, 0xd8, 0x83,
	0xb1, 0xd3, 0x81, 0x34, 0x08, 0x1f, 0x20, 0xe8, 0xfb, 0x07, 0xf3, 0xfe, 0xe6, 0x84, 0xa1, 0x66,
	0xe6, 0x8d, 0xc4, 0x46, 0xe2, 0xa4, 0x3d, 0xff, 0x66, 0xac, 0x6c, 0xe1, 0xf4, 0x95, 0x3e, 0xf1,
	0x17, 0xb6, 0x42, 0xb2, 0xc1, 0x01, 0x86, 0xce, 0x92, 0xd2, 0xbe, 0xca, 0xea, 0x7c, 0x7d, 0xd9,
	0x5e, 0xcd, 0xf4, 0x73, 0x84, 0xfc, 0x91, 0x95, 0x96, 0x2c, 0xea, 0xe4, 0x9c, 0x45, 0x87, 0x45,
	0x34, 0x09, 0xaf, 0xac, 0xe8, 0x35, 0x28, 0xed, 0xaa, 0xbc, 0xce, 0xd6, 0xe5, 0xe6, 0x5a, 0x48,
	0x14, 0xdb, 0x81, 0xf0, 0x67, 0x17, 0x71, 0x9b, 0x62, 0xbe, 0x61, 0x37, 0xc6, 0x9a, 0x60, 0x60,
	0xe8, 0xe4, 0x18, 0x77, 0x46, 0x55, 0xe7, 0xb1, 0x72, 0x31, 0x56, 0x76, 0xe0, 0xfb, 0x76, 0x95,
	0x8c, 0xd8, 0xdf, 0x2b, 0xfe, 0xc4, 0x96, 0x4e, 0x1f, 0x3a, 0xa3, 0x7c, 0xb5, 0xa8, 0xf3, 0x7f,
	0x6a, 0xe1, 0xf4, 0x61, 0xaf, 0xfc, 0x76, 0xf9, 0xbd, 0x88, 0xf3, 0x65, 0x11, 0x77, 0x7e, 0xfc,
	0x05, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xf6, 0x71, 0x84, 0x30, 0x01, 0x00, 0x00,
}
