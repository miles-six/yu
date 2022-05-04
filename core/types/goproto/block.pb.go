// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: block.proto

package goproto

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

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompactBlock *CompactBlock `protobuf:"bytes,1,opt,name=compact_block,json=compactBlock,proto3" json:"compact_block,omitempty"`
	Txns         *SignedTxns   `protobuf:"bytes,2,opt,name=txns,proto3" json:"txns,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetCompactBlock() *CompactBlock {
	if x != nil {
		return x.CompactBlock
	}
	return nil
}

func (x *Block) GetTxns() *SignedTxns {
	if x != nil {
		return x.Txns
	}
	return nil
}

type CompactBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header     *Header  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	TxnsHashes [][]byte `protobuf:"bytes,2,rep,name=txns_hashes,json=txnsHashes,proto3" json:"txns_hashes,omitempty"`
}

func (x *CompactBlock) Reset() {
	*x = CompactBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompactBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactBlock) ProtoMessage() {}

func (x *CompactBlock) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactBlock.ProtoReflect.Descriptor instead.
func (*CompactBlock) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{1}
}

func (x *CompactBlock) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *CompactBlock) GetTxnsHashes() [][]byte {
	if x != nil {
		return x.TxnsHashes
	}
	return nil
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash           []byte      `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	PrevHash       []byte      `protobuf:"bytes,2,opt,name=prev_hash,json=prevHash,proto3" json:"prev_hash,omitempty"`
	Height         uint64      `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	ChainId        uint64      `protobuf:"varint,4,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	TxnRoot        []byte      `protobuf:"bytes,5,opt,name=txn_root,json=txnRoot,proto3" json:"txn_root,omitempty"`
	StateRoot      []byte      `protobuf:"bytes,6,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	ReceiptRoot    []byte      `protobuf:"bytes,7,opt,name=receipt_root,json=receiptRoot,proto3" json:"receipt_root,omitempty"`
	Timestamp      uint64      `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	PeerId         string      `protobuf:"bytes,9,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	LeiLimit       uint64      `protobuf:"varint,10,opt,name=lei_limit,json=leiLimit,proto3" json:"lei_limit,omitempty"`
	LeiUsed        uint64      `protobuf:"varint,11,opt,name=lei_used,json=leiUsed,proto3" json:"lei_used,omitempty"`
	MinerPubkey    []byte      `protobuf:"bytes,12,opt,name=miner_pubkey,json=minerPubkey,proto3" json:"miner_pubkey,omitempty"`
	MinerSignature []byte      `protobuf:"bytes,13,opt,name=miner_signature,json=minerSignature,proto3" json:"miner_signature,omitempty"`
	Validators     *Validators `protobuf:"bytes,14,opt,name=validators,proto3" json:"validators,omitempty"`
	ProofBlockHash []byte      `protobuf:"bytes,15,opt,name=proof_block_hash,json=proofBlockHash,proto3" json:"proof_block_hash,omitempty"`
	ProofHeight    uint64      `protobuf:"varint,16,opt,name=proof_height,json=proofHeight,proto3" json:"proof_height,omitempty"`
	Proof          []byte      `protobuf:"bytes,17,opt,name=proof,proto3" json:"proof,omitempty"`
	Nonce          uint64      `protobuf:"varint,18,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Difficulty     uint64      `protobuf:"varint,19,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	Extra          []byte      `protobuf:"bytes,20,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{2}
}

func (x *Header) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *Header) GetPrevHash() []byte {
	if x != nil {
		return x.PrevHash
	}
	return nil
}

func (x *Header) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Header) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *Header) GetTxnRoot() []byte {
	if x != nil {
		return x.TxnRoot
	}
	return nil
}

func (x *Header) GetStateRoot() []byte {
	if x != nil {
		return x.StateRoot
	}
	return nil
}

func (x *Header) GetReceiptRoot() []byte {
	if x != nil {
		return x.ReceiptRoot
	}
	return nil
}

func (x *Header) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Header) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *Header) GetLeiLimit() uint64 {
	if x != nil {
		return x.LeiLimit
	}
	return 0
}

func (x *Header) GetLeiUsed() uint64 {
	if x != nil {
		return x.LeiUsed
	}
	return 0
}

func (x *Header) GetMinerPubkey() []byte {
	if x != nil {
		return x.MinerPubkey
	}
	return nil
}

func (x *Header) GetMinerSignature() []byte {
	if x != nil {
		return x.MinerSignature
	}
	return nil
}

func (x *Header) GetValidators() *Validators {
	if x != nil {
		return x.Validators
	}
	return nil
}

func (x *Header) GetProofBlockHash() []byte {
	if x != nil {
		return x.ProofBlockHash
	}
	return nil
}

func (x *Header) GetProofHeight() uint64 {
	if x != nil {
		return x.ProofHeight
	}
	return 0
}

func (x *Header) GetProof() []byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

func (x *Header) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Header) GetDifficulty() uint64 {
	if x != nil {
		return x.Difficulty
	}
	return 0
}

func (x *Header) GetExtra() []byte {
	if x != nil {
		return x.Extra
	}
	return nil
}

type Validators struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Validators []*Validator `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators,omitempty"`
}

func (x *Validators) Reset() {
	*x = Validators{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Validators) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Validators) ProtoMessage() {}

func (x *Validators) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Validators.ProtoReflect.Descriptor instead.
func (*Validators) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{3}
}

func (x *Validators) GetValidators() []*Validator {
	if x != nil {
		return x.Validators
	}
	return nil
}

type Validator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PubKey        []byte `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	ProposeWeight uint64 `protobuf:"varint,2,opt,name=propose_weight,json=proposeWeight,proto3" json:"propose_weight,omitempty"`
	VoteWeight    uint64 `protobuf:"varint,3,opt,name=vote_weight,json=voteWeight,proto3" json:"vote_weight,omitempty"`
}

func (x *Validator) Reset() {
	*x = Validator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Validator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Validator) ProtoMessage() {}

func (x *Validator) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Validator.ProtoReflect.Descriptor instead.
func (*Validator) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{4}
}

func (x *Validator) GetPubKey() []byte {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *Validator) GetProposeWeight() uint64 {
	if x != nil {
		return x.ProposeWeight
	}
	return 0
}

func (x *Validator) GetVoteWeight() uint64 {
	if x != nil {
		return x.VoteWeight
	}
	return 0
}

var File_block_proto protoreflect.FileDescriptor

var file_block_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x74,
	0x78, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x32, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x5f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x63, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1f, 0x0a, 0x04, 0x74, 0x78, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x78, 0x6e, 0x73,
	0x52, 0x04, 0x74, 0x78, 0x6e, 0x73, 0x22, 0x50, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1f, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x78, 0x6e, 0x73, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a, 0x74, 0x78,
	0x6e, 0x73, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x22, 0xe0, 0x04, 0x0a, 0x06, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x76, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x72, 0x65, 0x76,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x6e, 0x5f, 0x72,
	0x6f, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x78, 0x6e, 0x52, 0x6f,
	0x6f, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x5f, 0x72, 0x6f, 0x6f,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74,
	0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x65, 0x69, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x6c, 0x65, 0x69, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x69, 0x5f,
	0x75, 0x73, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6c, 0x65, 0x69, 0x55,
	0x73, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62,
	0x6b, 0x65, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x65, 0x72,
	0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x5f,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0e, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x2b, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x28, 0x0a, 0x10,
	0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x70, 0x72,
	0x6f, 0x6f, 0x66, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f,
	0x6f, 0x66, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x12,
	0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75,
	0x6c, 0x74, 0x79, 0x18, 0x13, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69,
	0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x22, 0x38, 0x0a, 0x0a, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x0a, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x6c, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x70,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x57, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x76, 0x6f, 0x74, 0x65, 0x57, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_block_proto_rawDescOnce sync.Once
	file_block_proto_rawDescData = file_block_proto_rawDesc
)

func file_block_proto_rawDescGZIP() []byte {
	file_block_proto_rawDescOnce.Do(func() {
		file_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_block_proto_rawDescData)
	})
	return file_block_proto_rawDescData
}

var file_block_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_block_proto_goTypes = []interface{}{
	(*Block)(nil),        // 0: Block
	(*CompactBlock)(nil), // 1: CompactBlock
	(*Header)(nil),       // 2: Header
	(*Validators)(nil),   // 3: Validators
	(*Validator)(nil),    // 4: Validator
	(*SignedTxns)(nil),   // 5: SignedTxns
}
var file_block_proto_depIdxs = []int32{
	1, // 0: Block.compact_block:type_name -> CompactBlock
	5, // 1: Block.txns:type_name -> SignedTxns
	2, // 2: CompactBlock.header:type_name -> Header
	3, // 3: Header.validators:type_name -> Validators
	4, // 4: Validators.validators:type_name -> Validator
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_block_proto_init() }
func file_block_proto_init() {
	if File_block_proto != nil {
		return
	}
	file_txn_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompactBlock); i {
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
		file_block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_block_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Validators); i {
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
		file_block_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Validator); i {
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
			RawDescriptor: file_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_block_proto_goTypes,
		DependencyIndexes: file_block_proto_depIdxs,
		MessageInfos:      file_block_proto_msgTypes,
	}.Build()
	File_block_proto = out.File
	file_block_proto_rawDesc = nil
	file_block_proto_goTypes = nil
	file_block_proto_depIdxs = nil
}
