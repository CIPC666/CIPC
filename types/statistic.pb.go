// Code generated by protoc-gen-go. DO NOT EDIT.
// source: statistic.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 手续费
type TotalFee struct {
	Fee     int64 `protobuf:"varint,1,opt,name=fee" json:"fee,omitempty"`
	TxCount int64 `protobuf:"varint,2,opt,name=txCount" json:"txCount,omitempty"`
}

func (m *TotalFee) Reset()                    { *m = TotalFee{} }
func (m *TotalFee) String() string            { return proto.CompactTextString(m) }
func (*TotalFee) ProtoMessage()               {}
func (*TotalFee) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *TotalFee) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *TotalFee) GetTxCount() int64 {
	if m != nil {
		return m.TxCount
	}
	return 0
}

// 查询symbol代币总额
type ReqGetTotalCoins struct {
	Symbol    string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	StateHash []byte `protobuf:"bytes,2,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	StartKey  []byte `protobuf:"bytes,3,opt,name=startKey,proto3" json:"startKey,omitempty"`
	Count     int64  `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
	Execer    string `protobuf:"bytes,5,opt,name=execer" json:"execer,omitempty"`
}

func (m *ReqGetTotalCoins) Reset()                    { *m = ReqGetTotalCoins{} }
func (m *ReqGetTotalCoins) String() string            { return proto.CompactTextString(m) }
func (*ReqGetTotalCoins) ProtoMessage()               {}
func (*ReqGetTotalCoins) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *ReqGetTotalCoins) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ReqGetTotalCoins) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *ReqGetTotalCoins) GetStartKey() []byte {
	if m != nil {
		return m.StartKey
	}
	return nil
}

func (m *ReqGetTotalCoins) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqGetTotalCoins) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

// 查询symbol代币总额应答
type ReplyGetTotalCoins struct {
	Count   int64  `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	Num     int64  `protobuf:"varint,2,opt,name=num" json:"num,omitempty"`
	Amount  int64  `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
	NextKey []byte `protobuf:"bytes,4,opt,name=nextKey,proto3" json:"nextKey,omitempty"`
}

func (m *ReplyGetTotalCoins) Reset()                    { *m = ReplyGetTotalCoins{} }
func (m *ReplyGetTotalCoins) String() string            { return proto.CompactTextString(m) }
func (*ReplyGetTotalCoins) ProtoMessage()               {}
func (*ReplyGetTotalCoins) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

func (m *ReplyGetTotalCoins) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReplyGetTotalCoins) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *ReplyGetTotalCoins) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ReplyGetTotalCoins) GetNextKey() []byte {
	if m != nil {
		return m.NextKey
	}
	return nil
}

// 迭代查询symbol代币总额
type IterateRangeByStateHash struct {
	StateHash []byte `protobuf:"bytes,1,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Start     []byte `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End       []byte `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
	Count     int64  `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
}

func (m *IterateRangeByStateHash) Reset()                    { *m = IterateRangeByStateHash{} }
func (m *IterateRangeByStateHash) String() string            { return proto.CompactTextString(m) }
func (*IterateRangeByStateHash) ProtoMessage()               {}
func (*IterateRangeByStateHash) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

func (m *IterateRangeByStateHash) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *IterateRangeByStateHash) GetStart() []byte {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *IterateRangeByStateHash) GetEnd() []byte {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *IterateRangeByStateHash) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type TotalAmount struct {
	// 统计的总数
	Total int64 `protobuf:"varint,1,opt,name=total" json:"total,omitempty"`
}

func (m *TotalAmount) Reset()                    { *m = TotalAmount{} }
func (m *TotalAmount) String() string            { return proto.CompactTextString(m) }
func (*TotalAmount) ProtoMessage()               {}
func (*TotalAmount) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{4} }

func (m *TotalAmount) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

// 查询symbol在合约中的代币总额，如果execAddr为空，则为查询symbol在所有合约中的代币总额
type ReqGetExecBalance struct {
	Symbol    string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	StateHash []byte `protobuf:"bytes,2,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Addr      []byte `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	ExecAddr  []byte `protobuf:"bytes,4,opt,name=execAddr,proto3" json:"execAddr,omitempty"`
	Execer    string `protobuf:"bytes,5,opt,name=execer" json:"execer,omitempty"`
	Count     int64  `protobuf:"varint,6,opt,name=count" json:"count,omitempty"`
	NextKey   []byte `protobuf:"bytes,7,opt,name=nextKey,proto3" json:"nextKey,omitempty"`
}

func (m *ReqGetExecBalance) Reset()                    { *m = ReqGetExecBalance{} }
func (m *ReqGetExecBalance) String() string            { return proto.CompactTextString(m) }
func (*ReqGetExecBalance) ProtoMessage()               {}
func (*ReqGetExecBalance) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{5} }

func (m *ReqGetExecBalance) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ReqGetExecBalance) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *ReqGetExecBalance) GetAddr() []byte {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *ReqGetExecBalance) GetExecAddr() []byte {
	if m != nil {
		return m.ExecAddr
	}
	return nil
}

func (m *ReqGetExecBalance) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

func (m *ReqGetExecBalance) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqGetExecBalance) GetNextKey() []byte {
	if m != nil {
		return m.NextKey
	}
	return nil
}

type ExecBalanceItem struct {
	ExecAddr []byte `protobuf:"bytes,1,opt,name=execAddr,proto3" json:"execAddr,omitempty"`
	Frozen   int64  `protobuf:"varint,2,opt,name=frozen" json:"frozen,omitempty"`
	Active   int64  `protobuf:"varint,3,opt,name=active" json:"active,omitempty"`
}

func (m *ExecBalanceItem) Reset()                    { *m = ExecBalanceItem{} }
func (m *ExecBalanceItem) String() string            { return proto.CompactTextString(m) }
func (*ExecBalanceItem) ProtoMessage()               {}
func (*ExecBalanceItem) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{6} }

func (m *ExecBalanceItem) GetExecAddr() []byte {
	if m != nil {
		return m.ExecAddr
	}
	return nil
}

func (m *ExecBalanceItem) GetFrozen() int64 {
	if m != nil {
		return m.Frozen
	}
	return 0
}

func (m *ExecBalanceItem) GetActive() int64 {
	if m != nil {
		return m.Active
	}
	return 0
}

// 查询symbol在合约中的代币总额应答
type ReplyGetExecBalance struct {
	Amount       int64              `protobuf:"varint,1,opt,name=amount" json:"amount,omitempty"`
	AmountFrozen int64              `protobuf:"varint,2,opt,name=amountFrozen" json:"amountFrozen,omitempty"`
	AmountActive int64              `protobuf:"varint,3,opt,name=amountActive" json:"amountActive,omitempty"`
	NextKey      []byte             `protobuf:"bytes,4,opt,name=nextKey,proto3" json:"nextKey,omitempty"`
	Items        []*ExecBalanceItem `protobuf:"bytes,5,rep,name=items" json:"items,omitempty"`
}

func (m *ReplyGetExecBalance) Reset()                    { *m = ReplyGetExecBalance{} }
func (m *ReplyGetExecBalance) String() string            { return proto.CompactTextString(m) }
func (*ReplyGetExecBalance) ProtoMessage()               {}
func (*ReplyGetExecBalance) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{7} }

func (m *ReplyGetExecBalance) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ReplyGetExecBalance) GetAmountFrozen() int64 {
	if m != nil {
		return m.AmountFrozen
	}
	return 0
}

func (m *ReplyGetExecBalance) GetAmountActive() int64 {
	if m != nil {
		return m.AmountActive
	}
	return 0
}

func (m *ReplyGetExecBalance) GetNextKey() []byte {
	if m != nil {
		return m.NextKey
	}
	return nil
}

func (m *ReplyGetExecBalance) GetItems() []*ExecBalanceItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*TotalFee)(nil), "types.TotalFee")
	proto.RegisterType((*ReqGetTotalCoins)(nil), "types.ReqGetTotalCoins")
	proto.RegisterType((*ReplyGetTotalCoins)(nil), "types.ReplyGetTotalCoins")
	proto.RegisterType((*IterateRangeByStateHash)(nil), "types.IterateRangeByStateHash")
	proto.RegisterType((*TotalAmount)(nil), "types.TotalAmount")
	proto.RegisterType((*ReqGetExecBalance)(nil), "types.ReqGetExecBalance")
	proto.RegisterType((*ExecBalanceItem)(nil), "types.ExecBalanceItem")
	proto.RegisterType((*ReplyGetExecBalance)(nil), "types.ReplyGetExecBalance")
}

func init() { proto.RegisterFile("statistic.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xd5, 0xe2, 0x38, 0x6d, 0xa7, 0x95, 0x1a, 0x96, 0xa8, 0x58, 0x08, 0x44, 0x65, 0x2e, 0x3d,
	0xa0, 0x44, 0xc2, 0x12, 0xf7, 0xa4, 0xa2, 0x50, 0x71, 0x33, 0x9c, 0x90, 0x38, 0x6c, 0x36, 0xd3,
	0xc6, 0x52, 0xbc, 0x0e, 0xf6, 0x04, 0xc5, 0x7c, 0x06, 0xff, 0xc3, 0x81, 0x3f, 0x43, 0x3b, 0xbb,
	0x76, 0xec, 0x40, 0x2f, 0xdc, 0xe6, 0x8d, 0xd7, 0xf3, 0xe6, 0xbd, 0xb7, 0x0b, 0xe7, 0x15, 0x29,
	0xca, 0x2a, 0xca, 0xf4, 0x64, 0x53, 0x16, 0x54, 0xc8, 0x90, 0xea, 0x0d, 0x56, 0xf1, 0x5b, 0x38,
	0xfe, 0x5c, 0x90, 0x5a, 0xdf, 0x20, 0xca, 0x11, 0x04, 0x77, 0x88, 0x91, 0xb8, 0x14, 0x57, 0x41,
	0x6a, 0x4b, 0x19, 0xc1, 0x11, 0xed, 0xae, 0x8b, 0xad, 0xa1, 0xe8, 0x11, 0x77, 0x1b, 0x18, 0xff,
	0x14, 0x30, 0x4a, 0xf1, 0xdb, 0x7b, 0x24, 0xfe, 0xfd, 0xba, 0xc8, 0x4c, 0x25, 0x2f, 0x60, 0x58,
	0xd5, 0xf9, 0xa2, 0x58, 0xf3, 0x8c, 0x93, 0xd4, 0x23, 0xf9, 0x1c, 0x4e, 0x2c, 0x3d, 0x7e, 0x50,
	0xd5, 0x8a, 0x07, 0x9d, 0xa5, 0xfb, 0x86, 0x7c, 0x06, 0xc7, 0x15, 0xa9, 0x92, 0x3e, 0x62, 0x1d,
	0x05, 0xfc, 0xb1, 0xc5, 0x72, 0x0c, 0xa1, 0x66, 0xfa, 0x01, 0xd3, 0x3b, 0x60, 0x79, 0x70, 0x87,
	0x1a, 0xcb, 0x28, 0x74, 0x3c, 0x0e, 0xc5, 0x06, 0x64, 0x8a, 0x9b, 0x75, 0xdd, 0xdf, 0xaa, 0x9d,
	0x21, 0xba, 0x33, 0x46, 0x10, 0x98, 0x6d, 0xee, 0x65, 0xd9, 0xd2, 0x4e, 0x55, 0x39, 0x1f, 0x0c,
	0xb8, 0xe9, 0x91, 0x35, 0xc1, 0xe0, 0x8e, 0xd7, 0x1b, 0xf0, 0x7a, 0x0d, 0x8c, 0xb7, 0xf0, 0xf4,
	0x96, 0xb0, 0x54, 0x84, 0xa9, 0x32, 0xf7, 0x38, 0xaf, 0x3f, 0xb5, 0xa2, 0x7a, 0x92, 0xc5, 0xa1,
	0xe4, 0x31, 0x84, 0x2c, 0xd1, 0x9b, 0xe1, 0x80, 0x5d, 0x09, 0xcd, 0xd2, 0x7b, 0x60, 0xcb, 0x7f,
	0xcb, 0x8f, 0x5f, 0xc1, 0x29, 0xcb, 0x9b, 0xb9, 0xfd, 0xc6, 0x10, 0x92, 0x85, 0x8d, 0x3e, 0x06,
	0xf1, 0x6f, 0x01, 0x8f, 0x5d, 0x40, 0xef, 0x76, 0xa8, 0xe7, 0x6a, 0xad, 0x8c, 0xc6, 0xff, 0x4c,
	0x48, 0xc2, 0x40, 0x2d, 0x97, 0xa5, 0xdf, 0x8c, 0x6b, 0x9b, 0x9a, 0x75, 0x7d, 0x66, 0xfb, 0xce,
	0x96, 0x16, 0x3f, 0x94, 0xcf, 0x5e, 0xce, 0xb0, 0x9b, 0x44, 0xc7, 0xdf, 0xa3, 0xbe, 0xbf, 0x5f,
	0xe1, 0xbc, 0xb3, 0xfc, 0x2d, 0x61, 0xde, 0xa3, 0x15, 0x7f, 0xd3, 0xde, 0x95, 0xc5, 0x0f, 0x34,
	0x3e, 0x55, 0x8f, 0x38, 0x58, 0x4d, 0xd9, 0x77, 0x6c, 0x83, 0x65, 0x14, 0xff, 0x12, 0xf0, 0xa4,
	0xb9, 0x2f, 0x07, 0x26, 0xf9, 0x8b, 0x20, 0x7a, 0x17, 0x21, 0x86, 0x33, 0x57, 0xdd, 0x74, 0x59,
	0x7a, 0xbd, 0xfd, 0x99, 0x59, 0x97, 0xb1, 0xd7, 0x7b, 0xf8, 0x42, 0xc9, 0xd7, 0x10, 0x66, 0x84,
	0x79, 0x15, 0x85, 0x97, 0xc1, 0xd5, 0xe9, 0x9b, 0x8b, 0x09, 0x3f, 0xd2, 0xc9, 0x81, 0x09, 0xa9,
	0x3b, 0x34, 0x7f, 0xf9, 0xe5, 0xc5, 0x7d, 0x46, 0xab, 0xed, 0x62, 0xa2, 0x8b, 0x7c, 0x9a, 0x24,
	0xda, 0x4c, 0xf5, 0x4a, 0x65, 0x26, 0x49, 0xa6, 0xfc, 0xdf, 0x62, 0xc8, 0x4f, 0x3d, 0xf9, 0x13,
	0x00, 0x00, 0xff, 0xff, 0x5a, 0x92, 0xa1, 0xd5, 0xfd, 0x03, 0x00, 0x00,
}
