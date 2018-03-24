// Code generated by protoc-gen-go. DO NOT EDIT.
// source: streamer.proto

/*
Package btc is a generated protocol buffer package.

It is generated from these files:
	streamer.proto

It has these top-level messages:
	BTCTransaction
	AddSpOut
	BlockHeight
	ReqDeleteSpOut
	MempoolToDelete
	WatchAddress
	MempoolRecord
	Empty
	RawTx
	AddressToResync
	UsersData
	ReplyInfo
*/
package btc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type BTCTransaction struct {
	UserID        string                         `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	TxID          string                         `protobuf:"bytes,2,opt,name=txID" json:"txID,omitempty"`
	TxHash        string                         `protobuf:"bytes,3,opt,name=txHash" json:"txHash,omitempty"`
	TxOutScript   string                         `protobuf:"bytes,4,opt,name=txOutScript" json:"txOutScript,omitempty"`
	TxAddress     []string                       `protobuf:"bytes,5,rep,name=txAddress" json:"txAddress,omitempty"`
	TxStatus      int32                          `protobuf:"varint,6,opt,name=txStatus" json:"txStatus,omitempty"`
	TxOutAmount   int64                          `protobuf:"varint,7,opt,name=txOutAmount" json:"txOutAmount,omitempty"`
	BlockTime     int64                          `protobuf:"varint,8,opt,name=blockTime" json:"blockTime,omitempty"`
	BlockHeight   int64                          `protobuf:"varint,9,opt,name=blockHeight" json:"blockHeight,omitempty"`
	Confirmations int32                          `protobuf:"varint,10,opt,name=confirmations" json:"confirmations,omitempty"`
	TxFee         int64                          `protobuf:"varint,11,opt,name=txFee" json:"txFee,omitempty"`
	MempoolTime   int64                          `protobuf:"varint,12,opt,name=mempoolTime" json:"mempoolTime,omitempty"`
	TxInputs      []*BTCTransaction_AddresAmount `protobuf:"bytes,13,rep,name=txInputs" json:"txInputs,omitempty"`
	TxOutputs     []*BTCTransaction_AddresAmount `protobuf:"bytes,14,rep,name=TxOutputs" json:"TxOutputs,omitempty"`
	WalletsInput  []*BTCTransaction_WalletForTx  `protobuf:"bytes,15,rep,name=WalletsInput" json:"WalletsInput,omitempty"`
	WalletsOutput []*BTCTransaction_WalletForTx  `protobuf:"bytes,16,rep,name=WalletsOutput" json:"WalletsOutput,omitempty"`
}

func (m *BTCTransaction) Reset()                    { *m = BTCTransaction{} }
func (m *BTCTransaction) String() string            { return proto.CompactTextString(m) }
func (*BTCTransaction) ProtoMessage()               {}
func (*BTCTransaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BTCTransaction) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *BTCTransaction) GetTxID() string {
	if m != nil {
		return m.TxID
	}
	return ""
}

func (m *BTCTransaction) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *BTCTransaction) GetTxOutScript() string {
	if m != nil {
		return m.TxOutScript
	}
	return ""
}

func (m *BTCTransaction) GetTxAddress() []string {
	if m != nil {
		return m.TxAddress
	}
	return nil
}

func (m *BTCTransaction) GetTxStatus() int32 {
	if m != nil {
		return m.TxStatus
	}
	return 0
}

func (m *BTCTransaction) GetTxOutAmount() int64 {
	if m != nil {
		return m.TxOutAmount
	}
	return 0
}

func (m *BTCTransaction) GetBlockTime() int64 {
	if m != nil {
		return m.BlockTime
	}
	return 0
}

func (m *BTCTransaction) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *BTCTransaction) GetConfirmations() int32 {
	if m != nil {
		return m.Confirmations
	}
	return 0
}

func (m *BTCTransaction) GetTxFee() int64 {
	if m != nil {
		return m.TxFee
	}
	return 0
}

func (m *BTCTransaction) GetMempoolTime() int64 {
	if m != nil {
		return m.MempoolTime
	}
	return 0
}

func (m *BTCTransaction) GetTxInputs() []*BTCTransaction_AddresAmount {
	if m != nil {
		return m.TxInputs
	}
	return nil
}

func (m *BTCTransaction) GetTxOutputs() []*BTCTransaction_AddresAmount {
	if m != nil {
		return m.TxOutputs
	}
	return nil
}

func (m *BTCTransaction) GetWalletsInput() []*BTCTransaction_WalletForTx {
	if m != nil {
		return m.WalletsInput
	}
	return nil
}

func (m *BTCTransaction) GetWalletsOutput() []*BTCTransaction_WalletForTx {
	if m != nil {
		return m.WalletsOutput
	}
	return nil
}

type BTCTransaction_AddresAmount struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Amount  int64  `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *BTCTransaction_AddresAmount) Reset()                    { *m = BTCTransaction_AddresAmount{} }
func (m *BTCTransaction_AddresAmount) String() string            { return proto.CompactTextString(m) }
func (*BTCTransaction_AddresAmount) ProtoMessage()               {}
func (*BTCTransaction_AddresAmount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *BTCTransaction_AddresAmount) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *BTCTransaction_AddresAmount) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type BTCTransaction_WalletForTx struct {
	Userid     string `protobuf:"bytes,1,opt,name=userid" json:"userid,omitempty"`
	Address    string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	TxOutIndex int32  `protobuf:"varint,3,opt,name=txOutIndex" json:"txOutIndex,omitempty"`
	Amount     int64  `protobuf:"varint,4,opt,name=amount" json:"amount,omitempty"`
}

func (m *BTCTransaction_WalletForTx) Reset()                    { *m = BTCTransaction_WalletForTx{} }
func (m *BTCTransaction_WalletForTx) String() string            { return proto.CompactTextString(m) }
func (*BTCTransaction_WalletForTx) ProtoMessage()               {}
func (*BTCTransaction_WalletForTx) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *BTCTransaction_WalletForTx) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

func (m *BTCTransaction_WalletForTx) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *BTCTransaction_WalletForTx) GetTxOutIndex() int32 {
	if m != nil {
		return m.TxOutIndex
	}
	return 0
}

func (m *BTCTransaction_WalletForTx) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type AddSpOut struct {
	TxID        string `protobuf:"bytes,1,opt,name=txID" json:"txID,omitempty"`
	TxOutID     int32  `protobuf:"varint,2,opt,name=txOutID" json:"txOutID,omitempty"`
	TxOutAmount int64  `protobuf:"varint,3,opt,name=txOutAmount" json:"txOutAmount,omitempty"`
	TxOutScript string `protobuf:"bytes,4,opt,name=txOutScript" json:"txOutScript,omitempty"`
	Address     string `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	UserID      string `protobuf:"bytes,6,opt,name=userID" json:"userID,omitempty"`
	TxStatus    int32  `protobuf:"varint,7,opt,name=txStatus" json:"txStatus,omitempty"`
}

func (m *AddSpOut) Reset()                    { *m = AddSpOut{} }
func (m *AddSpOut) String() string            { return proto.CompactTextString(m) }
func (*AddSpOut) ProtoMessage()               {}
func (*AddSpOut) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddSpOut) GetTxID() string {
	if m != nil {
		return m.TxID
	}
	return ""
}

func (m *AddSpOut) GetTxOutID() int32 {
	if m != nil {
		return m.TxOutID
	}
	return 0
}

func (m *AddSpOut) GetTxOutAmount() int64 {
	if m != nil {
		return m.TxOutAmount
	}
	return 0
}

func (m *AddSpOut) GetTxOutScript() string {
	if m != nil {
		return m.TxOutScript
	}
	return ""
}

func (m *AddSpOut) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AddSpOut) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *AddSpOut) GetTxStatus() int32 {
	if m != nil {
		return m.TxStatus
	}
	return 0
}

type BlockHeight struct {
	Height int64 `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
}

func (m *BlockHeight) Reset()                    { *m = BlockHeight{} }
func (m *BlockHeight) String() string            { return proto.CompactTextString(m) }
func (*BlockHeight) ProtoMessage()               {}
func (*BlockHeight) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BlockHeight) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type ReqDeleteSpOut struct {
	UserID  string `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	TxID    string `protobuf:"bytes,2,opt,name=txID" json:"txID,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
}

func (m *ReqDeleteSpOut) Reset()                    { *m = ReqDeleteSpOut{} }
func (m *ReqDeleteSpOut) String() string            { return proto.CompactTextString(m) }
func (*ReqDeleteSpOut) ProtoMessage()               {}
func (*ReqDeleteSpOut) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReqDeleteSpOut) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *ReqDeleteSpOut) GetTxID() string {
	if m != nil {
		return m.TxID
	}
	return ""
}

func (m *ReqDeleteSpOut) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type MempoolToDelete struct {
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
}

func (m *MempoolToDelete) Reset()                    { *m = MempoolToDelete{} }
func (m *MempoolToDelete) String() string            { return proto.CompactTextString(m) }
func (*MempoolToDelete) ProtoMessage()               {}
func (*MempoolToDelete) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MempoolToDelete) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type WatchAddress struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	UserID  string `protobuf:"bytes,2,opt,name=userID" json:"userID,omitempty"`
}

func (m *WatchAddress) Reset()                    { *m = WatchAddress{} }
func (m *WatchAddress) String() string            { return proto.CompactTextString(m) }
func (*WatchAddress) ProtoMessage()               {}
func (*WatchAddress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *WatchAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *WatchAddress) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type MempoolRecord struct {
	Category int32  `protobuf:"varint,1,opt,name=category" json:"category,omitempty"`
	HashTX   string `protobuf:"bytes,2,opt,name=hashTX" json:"hashTX,omitempty"`
}

func (m *MempoolRecord) Reset()                    { *m = MempoolRecord{} }
func (m *MempoolRecord) String() string            { return proto.CompactTextString(m) }
func (*MempoolRecord) ProtoMessage()               {}
func (*MempoolRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MempoolRecord) GetCategory() int32 {
	if m != nil {
		return m.Category
	}
	return 0
}

func (m *MempoolRecord) GetHashTX() string {
	if m != nil {
		return m.HashTX
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type RawTx struct {
	Transaction string `protobuf:"bytes,1,opt,name=transaction" json:"transaction,omitempty"`
}

func (m *RawTx) Reset()                    { *m = RawTx{} }
func (m *RawTx) String() string            { return proto.CompactTextString(m) }
func (*RawTx) ProtoMessage()               {}
func (*RawTx) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RawTx) GetTransaction() string {
	if m != nil {
		return m.Transaction
	}
	return ""
}

type AddressToResync struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

func (m *AddressToResync) Reset()                    { *m = AddressToResync{} }
func (m *AddressToResync) String() string            { return proto.CompactTextString(m) }
func (*AddressToResync) ProtoMessage()               {}
func (*AddressToResync) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AddressToResync) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type UsersData struct {
	Map map[string]string `protobuf:"bytes,1,rep,name=map" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *UsersData) Reset()                    { *m = UsersData{} }
func (m *UsersData) String() string            { return proto.CompactTextString(m) }
func (*UsersData) ProtoMessage()               {}
func (*UsersData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *UsersData) GetMap() map[string]string {
	if m != nil {
		return m.Map
	}
	return nil
}

type ReplyInfo struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *ReplyInfo) Reset()                    { *m = ReplyInfo{} }
func (m *ReplyInfo) String() string            { return proto.CompactTextString(m) }
func (*ReplyInfo) ProtoMessage()               {}
func (*ReplyInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ReplyInfo) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*BTCTransaction)(nil), "btc.BTCTransaction")
	proto.RegisterType((*BTCTransaction_AddresAmount)(nil), "btc.BTCTransaction.AddresAmount")
	proto.RegisterType((*BTCTransaction_WalletForTx)(nil), "btc.BTCTransaction.WalletForTx")
	proto.RegisterType((*AddSpOut)(nil), "btc.AddSpOut")
	proto.RegisterType((*BlockHeight)(nil), "btc.BlockHeight")
	proto.RegisterType((*ReqDeleteSpOut)(nil), "btc.ReqDeleteSpOut")
	proto.RegisterType((*MempoolToDelete)(nil), "btc.MempoolToDelete")
	proto.RegisterType((*WatchAddress)(nil), "btc.WatchAddress")
	proto.RegisterType((*MempoolRecord)(nil), "btc.MempoolRecord")
	proto.RegisterType((*Empty)(nil), "btc.Empty")
	proto.RegisterType((*RawTx)(nil), "btc.RawTx")
	proto.RegisterType((*AddressToResync)(nil), "btc.AddressToResync")
	proto.RegisterType((*UsersData)(nil), "btc.UsersData")
	proto.RegisterType((*ReplyInfo)(nil), "btc.ReplyInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NodeCommuunications service

type NodeCommuunicationsClient interface {
	EventInitialAdd(ctx context.Context, in *UsersData, opts ...grpc.CallOption) (*ReplyInfo, error)
	EventAddNewAddress(ctx context.Context, in *WatchAddress, opts ...grpc.CallOption) (*ReplyInfo, error)
	EventGetBlockHeight(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BlockHeight, error)
	EventGetAllMempool(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_EventGetAllMempoolClient, error)
	EventAddMempoolRecord(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_EventAddMempoolRecordClient, error)
	EventDeleteMempool(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_EventDeleteMempoolClient, error)
	EventResyncAddress(ctx context.Context, in *AddressToResync, opts ...grpc.CallOption) (*ReplyInfo, error)
	EventSendRawTx(ctx context.Context, in *RawTx, opts ...grpc.CallOption) (*ReplyInfo, error)
	EventDeleteSpendableOut(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_EventDeleteSpendableOutClient, error)
	EventAddSpendableOut(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_EventAddSpendableOutClient, error)
	NewTx(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_NewTxClient, error)
}

type nodeCommuunicationsClient struct {
	cc *grpc.ClientConn
}

func NewNodeCommuunicationsClient(cc *grpc.ClientConn) NodeCommuunicationsClient {
	return &nodeCommuunicationsClient{cc}
}

func (c *nodeCommuunicationsClient) EventInitialAdd(ctx context.Context, in *UsersData, opts ...grpc.CallOption) (*ReplyInfo, error) {
	out := new(ReplyInfo)
	err := grpc.Invoke(ctx, "/btc.NodeCommuunications/EventInitialAdd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommuunicationsClient) EventAddNewAddress(ctx context.Context, in *WatchAddress, opts ...grpc.CallOption) (*ReplyInfo, error) {
	out := new(ReplyInfo)
	err := grpc.Invoke(ctx, "/btc.NodeCommuunications/EventAddNewAddress", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommuunicationsClient) EventGetBlockHeight(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BlockHeight, error) {
	out := new(BlockHeight)
	err := grpc.Invoke(ctx, "/btc.NodeCommuunications/EventGetBlockHeight", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommuunicationsClient) EventGetAllMempool(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_EventGetAllMempoolClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_NodeCommuunications_serviceDesc.Streams[0], c.cc, "/btc.NodeCommuunications/EventGetAllMempool", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeCommuunicationsEventGetAllMempoolClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeCommuunications_EventGetAllMempoolClient interface {
	Recv() (*MempoolRecord, error)
	grpc.ClientStream
}

type nodeCommuunicationsEventGetAllMempoolClient struct {
	grpc.ClientStream
}

func (x *nodeCommuunicationsEventGetAllMempoolClient) Recv() (*MempoolRecord, error) {
	m := new(MempoolRecord)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeCommuunicationsClient) EventAddMempoolRecord(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_EventAddMempoolRecordClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_NodeCommuunications_serviceDesc.Streams[1], c.cc, "/btc.NodeCommuunications/EventAddMempoolRecord", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeCommuunicationsEventAddMempoolRecordClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeCommuunications_EventAddMempoolRecordClient interface {
	Recv() (*MempoolRecord, error)
	grpc.ClientStream
}

type nodeCommuunicationsEventAddMempoolRecordClient struct {
	grpc.ClientStream
}

func (x *nodeCommuunicationsEventAddMempoolRecordClient) Recv() (*MempoolRecord, error) {
	m := new(MempoolRecord)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeCommuunicationsClient) EventDeleteMempool(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_EventDeleteMempoolClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_NodeCommuunications_serviceDesc.Streams[2], c.cc, "/btc.NodeCommuunications/EventDeleteMempool", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeCommuunicationsEventDeleteMempoolClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeCommuunications_EventDeleteMempoolClient interface {
	Recv() (*MempoolToDelete, error)
	grpc.ClientStream
}

type nodeCommuunicationsEventDeleteMempoolClient struct {
	grpc.ClientStream
}

func (x *nodeCommuunicationsEventDeleteMempoolClient) Recv() (*MempoolToDelete, error) {
	m := new(MempoolToDelete)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeCommuunicationsClient) EventResyncAddress(ctx context.Context, in *AddressToResync, opts ...grpc.CallOption) (*ReplyInfo, error) {
	out := new(ReplyInfo)
	err := grpc.Invoke(ctx, "/btc.NodeCommuunications/EventResyncAddress", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommuunicationsClient) EventSendRawTx(ctx context.Context, in *RawTx, opts ...grpc.CallOption) (*ReplyInfo, error) {
	out := new(ReplyInfo)
	err := grpc.Invoke(ctx, "/btc.NodeCommuunications/EventSendRawTx", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommuunicationsClient) EventDeleteSpendableOut(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_EventDeleteSpendableOutClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_NodeCommuunications_serviceDesc.Streams[3], c.cc, "/btc.NodeCommuunications/EventDeleteSpendableOut", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeCommuunicationsEventDeleteSpendableOutClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeCommuunications_EventDeleteSpendableOutClient interface {
	Recv() (*ReqDeleteSpOut, error)
	grpc.ClientStream
}

type nodeCommuunicationsEventDeleteSpendableOutClient struct {
	grpc.ClientStream
}

func (x *nodeCommuunicationsEventDeleteSpendableOutClient) Recv() (*ReqDeleteSpOut, error) {
	m := new(ReqDeleteSpOut)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeCommuunicationsClient) EventAddSpendableOut(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_EventAddSpendableOutClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_NodeCommuunications_serviceDesc.Streams[4], c.cc, "/btc.NodeCommuunications/EventAddSpendableOut", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeCommuunicationsEventAddSpendableOutClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeCommuunications_EventAddSpendableOutClient interface {
	Recv() (*AddSpOut, error)
	grpc.ClientStream
}

type nodeCommuunicationsEventAddSpendableOutClient struct {
	grpc.ClientStream
}

func (x *nodeCommuunicationsEventAddSpendableOutClient) Recv() (*AddSpOut, error) {
	m := new(AddSpOut)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeCommuunicationsClient) NewTx(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_NewTxClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_NodeCommuunications_serviceDesc.Streams[5], c.cc, "/btc.NodeCommuunications/NewTx", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeCommuunicationsNewTxClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeCommuunications_NewTxClient interface {
	Recv() (*BTCTransaction, error)
	grpc.ClientStream
}

type nodeCommuunicationsNewTxClient struct {
	grpc.ClientStream
}

func (x *nodeCommuunicationsNewTxClient) Recv() (*BTCTransaction, error) {
	m := new(BTCTransaction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for NodeCommuunications service

type NodeCommuunicationsServer interface {
	EventInitialAdd(context.Context, *UsersData) (*ReplyInfo, error)
	EventAddNewAddress(context.Context, *WatchAddress) (*ReplyInfo, error)
	EventGetBlockHeight(context.Context, *Empty) (*BlockHeight, error)
	EventGetAllMempool(*Empty, NodeCommuunications_EventGetAllMempoolServer) error
	EventAddMempoolRecord(*Empty, NodeCommuunications_EventAddMempoolRecordServer) error
	EventDeleteMempool(*Empty, NodeCommuunications_EventDeleteMempoolServer) error
	EventResyncAddress(context.Context, *AddressToResync) (*ReplyInfo, error)
	EventSendRawTx(context.Context, *RawTx) (*ReplyInfo, error)
	EventDeleteSpendableOut(*Empty, NodeCommuunications_EventDeleteSpendableOutServer) error
	EventAddSpendableOut(*Empty, NodeCommuunications_EventAddSpendableOutServer) error
	NewTx(*Empty, NodeCommuunications_NewTxServer) error
}

func RegisterNodeCommuunicationsServer(s *grpc.Server, srv NodeCommuunicationsServer) {
	s.RegisterService(&_NodeCommuunications_serviceDesc, srv)
}

func _NodeCommuunications_EventInitialAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommuunicationsServer).EventInitialAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/btc.NodeCommuunications/EventInitialAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommuunicationsServer).EventInitialAdd(ctx, req.(*UsersData))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommuunications_EventAddNewAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WatchAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommuunicationsServer).EventAddNewAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/btc.NodeCommuunications/EventAddNewAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommuunicationsServer).EventAddNewAddress(ctx, req.(*WatchAddress))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommuunications_EventGetBlockHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommuunicationsServer).EventGetBlockHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/btc.NodeCommuunications/EventGetBlockHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommuunicationsServer).EventGetBlockHeight(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommuunications_EventGetAllMempool_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeCommuunicationsServer).EventGetAllMempool(m, &nodeCommuunicationsEventGetAllMempoolServer{stream})
}

type NodeCommuunications_EventGetAllMempoolServer interface {
	Send(*MempoolRecord) error
	grpc.ServerStream
}

type nodeCommuunicationsEventGetAllMempoolServer struct {
	grpc.ServerStream
}

func (x *nodeCommuunicationsEventGetAllMempoolServer) Send(m *MempoolRecord) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeCommuunications_EventAddMempoolRecord_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeCommuunicationsServer).EventAddMempoolRecord(m, &nodeCommuunicationsEventAddMempoolRecordServer{stream})
}

type NodeCommuunications_EventAddMempoolRecordServer interface {
	Send(*MempoolRecord) error
	grpc.ServerStream
}

type nodeCommuunicationsEventAddMempoolRecordServer struct {
	grpc.ServerStream
}

func (x *nodeCommuunicationsEventAddMempoolRecordServer) Send(m *MempoolRecord) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeCommuunications_EventDeleteMempool_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeCommuunicationsServer).EventDeleteMempool(m, &nodeCommuunicationsEventDeleteMempoolServer{stream})
}

type NodeCommuunications_EventDeleteMempoolServer interface {
	Send(*MempoolToDelete) error
	grpc.ServerStream
}

type nodeCommuunicationsEventDeleteMempoolServer struct {
	grpc.ServerStream
}

func (x *nodeCommuunicationsEventDeleteMempoolServer) Send(m *MempoolToDelete) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeCommuunications_EventResyncAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressToResync)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommuunicationsServer).EventResyncAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/btc.NodeCommuunications/EventResyncAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommuunicationsServer).EventResyncAddress(ctx, req.(*AddressToResync))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommuunications_EventSendRawTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommuunicationsServer).EventSendRawTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/btc.NodeCommuunications/EventSendRawTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommuunicationsServer).EventSendRawTx(ctx, req.(*RawTx))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommuunications_EventDeleteSpendableOut_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeCommuunicationsServer).EventDeleteSpendableOut(m, &nodeCommuunicationsEventDeleteSpendableOutServer{stream})
}

type NodeCommuunications_EventDeleteSpendableOutServer interface {
	Send(*ReqDeleteSpOut) error
	grpc.ServerStream
}

type nodeCommuunicationsEventDeleteSpendableOutServer struct {
	grpc.ServerStream
}

func (x *nodeCommuunicationsEventDeleteSpendableOutServer) Send(m *ReqDeleteSpOut) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeCommuunications_EventAddSpendableOut_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeCommuunicationsServer).EventAddSpendableOut(m, &nodeCommuunicationsEventAddSpendableOutServer{stream})
}

type NodeCommuunications_EventAddSpendableOutServer interface {
	Send(*AddSpOut) error
	grpc.ServerStream
}

type nodeCommuunicationsEventAddSpendableOutServer struct {
	grpc.ServerStream
}

func (x *nodeCommuunicationsEventAddSpendableOutServer) Send(m *AddSpOut) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeCommuunications_NewTx_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeCommuunicationsServer).NewTx(m, &nodeCommuunicationsNewTxServer{stream})
}

type NodeCommuunications_NewTxServer interface {
	Send(*BTCTransaction) error
	grpc.ServerStream
}

type nodeCommuunicationsNewTxServer struct {
	grpc.ServerStream
}

func (x *nodeCommuunicationsNewTxServer) Send(m *BTCTransaction) error {
	return x.ServerStream.SendMsg(m)
}

var _NodeCommuunications_serviceDesc = grpc.ServiceDesc{
	ServiceName: "btc.NodeCommuunications",
	HandlerType: (*NodeCommuunicationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EventInitialAdd",
			Handler:    _NodeCommuunications_EventInitialAdd_Handler,
		},
		{
			MethodName: "EventAddNewAddress",
			Handler:    _NodeCommuunications_EventAddNewAddress_Handler,
		},
		{
			MethodName: "EventGetBlockHeight",
			Handler:    _NodeCommuunications_EventGetBlockHeight_Handler,
		},
		{
			MethodName: "EventResyncAddress",
			Handler:    _NodeCommuunications_EventResyncAddress_Handler,
		},
		{
			MethodName: "EventSendRawTx",
			Handler:    _NodeCommuunications_EventSendRawTx_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EventGetAllMempool",
			Handler:       _NodeCommuunications_EventGetAllMempool_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "EventAddMempoolRecord",
			Handler:       _NodeCommuunications_EventAddMempoolRecord_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "EventDeleteMempool",
			Handler:       _NodeCommuunications_EventDeleteMempool_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "EventDeleteSpendableOut",
			Handler:       _NodeCommuunications_EventDeleteSpendableOut_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "EventAddSpendableOut",
			Handler:       _NodeCommuunications_EventAddSpendableOut_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "NewTx",
			Handler:       _NodeCommuunications_NewTx_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "streamer.proto",
}

func init() { proto.RegisterFile("streamer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 899 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xdb, 0x6e, 0xdb, 0x46,
	0x10, 0x35, 0x45, 0x53, 0xb2, 0x46, 0x96, 0xec, 0xae, 0xdd, 0x64, 0x21, 0x14, 0xad, 0x40, 0xd4,
	0x80, 0x82, 0x16, 0x42, 0x91, 0x20, 0x45, 0x9a, 0x06, 0x41, 0x14, 0xdb, 0x49, 0xf4, 0x10, 0x07,
	0xa0, 0xd4, 0xcb, 0xeb, 0x8a, 0x9c, 0x58, 0x44, 0x78, 0x2b, 0xb9, 0x8c, 0xa9, 0xd7, 0xfe, 0x53,
	0x7f, 0xa2, 0x5f, 0x55, 0xec, 0x85, 0xd2, 0x32, 0x91, 0x5b, 0xf7, 0x6d, 0x67, 0xb8, 0xe7, 0xcc,
	0xd9, 0xd9, 0x33, 0x2b, 0xc1, 0xa0, 0xe0, 0x39, 0xb2, 0x18, 0xf3, 0x49, 0x96, 0xa7, 0x3c, 0x25,
	0xf6, 0x92, 0xfb, 0xee, 0x5f, 0x6d, 0x18, 0xbc, 0x5c, 0x9c, 0x2f, 0x72, 0x96, 0x14, 0xcc, 0xe7,
	0x61, 0x9a, 0x90, 0x7b, 0xd0, 0x2e, 0x0b, 0xcc, 0x67, 0x17, 0xd4, 0x1a, 0x59, 0xe3, 0xae, 0xa7,
	0x23, 0x42, 0x60, 0x9f, 0x57, 0xb3, 0x0b, 0xda, 0x92, 0x59, 0xb9, 0x16, 0x7b, 0x79, 0xf5, 0x86,
	0x15, 0x2b, 0x6a, 0xab, 0xbd, 0x2a, 0x22, 0x23, 0xe8, 0xf1, 0xea, 0x5d, 0xc9, 0xe7, 0x7e, 0x1e,
	0x66, 0x9c, 0xee, 0xcb, 0x8f, 0x66, 0x8a, 0x7c, 0x05, 0x5d, 0x5e, 0x4d, 0x83, 0x20, 0xc7, 0xa2,
	0xa0, 0xce, 0xc8, 0x1e, 0x77, 0xbd, 0x6d, 0x82, 0x0c, 0xe1, 0x80, 0x57, 0x73, 0xce, 0x78, 0x59,
	0xd0, 0xf6, 0xc8, 0x1a, 0x3b, 0xde, 0x26, 0xde, 0x70, 0x4f, 0xe3, 0xb4, 0x4c, 0x38, 0xed, 0x8c,
	0xac, 0xb1, 0xed, 0x99, 0x29, 0xc1, 0xbd, 0x8c, 0x52, 0xff, 0xc3, 0x22, 0x8c, 0x91, 0x1e, 0xc8,
	0xef, 0xdb, 0x84, 0xc0, 0xcb, 0xe0, 0x0d, 0x86, 0xd7, 0x2b, 0x4e, 0xbb, 0x0a, 0x6f, 0xa4, 0xc8,
	0xb7, 0xd0, 0xf7, 0xd3, 0xe4, 0x7d, 0x98, 0xc7, 0x4c, 0x74, 0xa4, 0xa0, 0x20, 0x25, 0x34, 0x93,
	0xe4, 0x14, 0x1c, 0x5e, 0xbd, 0x42, 0xa4, 0x3d, 0xc9, 0xa0, 0x02, 0xc1, 0x1e, 0x63, 0x9c, 0xa5,
	0x69, 0x24, 0xab, 0x1f, 0x2a, 0x76, 0x23, 0x45, 0x9e, 0x89, 0xb3, 0xcd, 0x92, 0xac, 0xe4, 0x05,
	0xed, 0x8f, 0xec, 0x71, 0xef, 0xe1, 0x68, 0xb2, 0xe4, 0xfe, 0xa4, 0x79, 0x0d, 0x13, 0xd5, 0x0a,
	0x75, 0x22, 0x6f, 0x83, 0x20, 0xcf, 0xa1, 0xbb, 0x10, 0x47, 0x95, 0xf0, 0xc1, 0x1d, 0xe1, 0x5b,
	0x08, 0x39, 0x87, 0xc3, 0xdf, 0x58, 0x14, 0x21, 0x2f, 0x24, 0x21, 0x3d, 0x92, 0x14, 0xdf, 0xec,
	0xa2, 0x50, 0xfb, 0x5e, 0xa5, 0xf9, 0xa2, 0xf2, 0x1a, 0x20, 0x72, 0x09, 0x7d, 0x1d, 0x2b, 0x5a,
	0x7a, 0x7c, 0x37, 0x96, 0x26, 0x6a, 0xf8, 0x02, 0x0e, 0x4d, 0x99, 0x84, 0x42, 0x87, 0x69, 0x47,
	0x28, 0xeb, 0xd5, 0xa1, 0xf0, 0x19, 0x53, 0xd7, 0xdd, 0x92, 0x0d, 0xd5, 0xd1, 0xf0, 0x06, 0x7a,
	0x06, 0x7f, 0x6d, 0xdd, 0x30, 0x30, 0xad, 0x1b, 0x06, 0x26, 0x71, 0xab, 0x49, 0xfc, 0x35, 0x80,
	0x74, 0xce, 0x2c, 0x09, 0xb0, 0x92, 0x26, 0x76, 0x3c, 0x23, 0x63, 0x14, 0xde, 0x37, 0x0b, 0xbb,
	0x7f, 0x5b, 0x70, 0x30, 0x0d, 0x82, 0x79, 0xf6, 0xae, 0xe4, 0x9b, 0xc9, 0xb0, 0x8c, 0xc9, 0xa0,
	0xd0, 0x51, 0x34, 0x6a, 0x60, 0x1c, 0xaf, 0x0e, 0x3f, 0xf5, 0xaf, 0xfd, 0xb9, 0x7f, 0xff, 0x7b,
	0x7a, 0x8c, 0x03, 0x39, 0x9f, 0x75, 0x4a, 0x4f, 0x6f, 0xbb, 0x31, 0xbd, 0xe6, 0x44, 0x75, 0x9a,
	0x13, 0xe5, 0x9e, 0x41, 0xef, 0xa5, 0x61, 0xff, 0x7b, 0xd0, 0x5e, 0xa9, 0xd9, 0xb0, 0xd4, 0x99,
	0x55, 0xe4, 0xfe, 0x0a, 0x03, 0x0f, 0xff, 0xb8, 0xc0, 0x08, 0x39, 0xaa, 0x83, 0xff, 0x9f, 0xa7,
	0xc2, 0x90, 0x6c, 0x37, 0x24, 0xbb, 0x67, 0x70, 0xf4, 0x56, 0xcf, 0x47, 0xaa, 0xd8, 0x05, 0xc1,
	0x4a, 0xbc, 0x2a, 0xba, 0xa3, 0x62, 0xed, 0xbe, 0x10, 0xce, 0xe5, 0xfe, 0xaa, 0x7e, 0x23, 0xfe,
	0xd5, 0x2d, 0x5a, 0x56, 0xcb, 0x94, 0xe5, 0x9e, 0x43, 0x5f, 0x17, 0xf2, 0xd0, 0x4f, 0xf3, 0x40,
	0x34, 0xc5, 0x67, 0x1c, 0xaf, 0xd3, 0x7c, 0x2d, 0x39, 0x1c, 0x6f, 0x13, 0xcb, 0x2e, 0xb0, 0x62,
	0xb5, 0xf8, 0xbd, 0x26, 0x51, 0x91, 0xdb, 0x01, 0xe7, 0x32, 0xce, 0xf8, 0xda, 0x7d, 0x00, 0x8e,
	0xc7, 0x6e, 0x16, 0x95, 0xbc, 0xae, 0xad, 0xe1, 0xb5, 0x18, 0x33, 0xe5, 0x7e, 0x07, 0x47, 0x5a,
	0xf5, 0x22, 0xf5, 0xb0, 0x58, 0x27, 0xfe, 0xed, 0xea, 0xdd, 0x04, 0xba, 0xbf, 0x14, 0x98, 0x17,
	0x17, 0x8c, 0x33, 0xf2, 0x00, 0xec, 0x98, 0x65, 0xd4, 0x92, 0xf3, 0x75, 0x5f, 0xce, 0xd7, 0xe6,
	0xe3, 0xe4, 0x2d, 0xcb, 0x2e, 0x13, 0x9e, 0xaf, 0x3d, 0xb1, 0x67, 0xf8, 0x23, 0x1c, 0xd4, 0x09,
	0x72, 0x0c, 0xf6, 0x07, 0x5c, 0x6b, 0x66, 0xb1, 0x14, 0xaf, 0xd5, 0x47, 0x16, 0x95, 0xa8, 0x4f,
	0xa3, 0x82, 0xa7, 0xad, 0x27, 0x96, 0x7b, 0x06, 0x5d, 0x0f, 0xb3, 0x68, 0x3d, 0x4b, 0xde, 0xa7,
	0x42, 0x56, 0x8c, 0x45, 0xc1, 0xae, 0xb1, 0x96, 0xa5, 0xc3, 0x87, 0x7f, 0x3a, 0x70, 0x72, 0x95,
	0x06, 0x78, 0x9e, 0xc6, 0x71, 0x59, 0x26, 0xa1, 0xaf, 0x9f, 0xc1, 0x47, 0x70, 0x74, 0xf9, 0x11,
	0x13, 0x3e, 0x4b, 0x42, 0x1e, 0xb2, 0x68, 0x1a, 0x04, 0x64, 0xd0, 0xd4, 0x39, 0x54, 0xf1, 0xa6,
	0x88, 0xbb, 0x47, 0x7e, 0x02, 0x22, 0x41, 0xd3, 0x20, 0xb8, 0xc2, 0x9b, 0xfa, 0x46, 0xbf, 0x90,
	0xfb, 0xcc, 0x4b, 0xde, 0x01, 0x7d, 0x0c, 0x27, 0x12, 0xfa, 0x1a, 0xb9, 0x69, 0x5a, 0x90, 0x1b,
	0xe5, 0xcd, 0x0c, 0x8f, 0xd5, 0x3b, 0xb4, 0xfd, 0xea, 0xee, 0x91, 0x27, 0xba, 0xe2, 0x6b, 0xe4,
	0xd3, 0x28, 0xd2, 0x36, 0x68, 0xa0, 0x88, 0x5c, 0x37, 0x0c, 0xe2, 0xee, 0xfd, 0x60, 0x91, 0x9f,
	0xe1, 0xcb, 0x5a, 0x6b, 0xd3, 0x3d, 0x77, 0x01, 0x3f, 0xd5, 0x65, 0x95, 0xaf, 0x77, 0x95, 0x3d,
	0x35, 0x91, 0xf5, 0x00, 0x48, 0xec, 0x33, 0x8d, 0x55, 0x8e, 0xa9, 0x9b, 0xa4, 0xf6, 0x7f, 0x62,
	0xa7, 0x1d, 0x7d, 0x9a, 0xc0, 0x40, 0xa2, 0xe7, 0x98, 0x04, 0xca, 0xa7, 0xaa, 0xaa, 0x5c, 0xef,
	0xd8, 0xff, 0x1c, 0xee, 0x1b, 0x4a, 0xe7, 0x19, 0x26, 0x01, 0x5b, 0x46, 0x28, 0xc6, 0xdc, 0x94,
	0x7b, 0xa2, 0x81, 0xe6, 0x3b, 0x20, 0xd5, 0x3e, 0x86, 0xd3, 0xba, 0x4d, 0xb7, 0x82, 0xfb, 0xb5,
	0xf6, 0x2d, 0xec, 0x7b, 0x70, 0xae, 0x70, 0xab, 0xce, 0x2c, 0xd2, 0xfc, 0x21, 0x11, 0xbb, 0x97,
	0x6d, 0xf9, 0xd7, 0xe5, 0xd1, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x02, 0x3b, 0x5c, 0x9c, 0xcc,
	0x08, 0x00, 0x00,
}
