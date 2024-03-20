// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proposal.proto

package peer

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// This structure is necessary to sign the proposal which contains the header
// and the payload. Without this structure, we would have to concatenate the
// header and the payload to verify the signature, which could be expensive
// with large payload
//
// When an endorser receives a SignedProposal message, it should verify the
// signature over the proposal bytes. This verification requires the following
// steps:
// 1. Verification of the validity of the certificate that was used to produce
//    the signature.  The certificate will be available once proposalBytes has
//    been unmarshalled to a Proposal message, and Proposal.header has been
//    unmarshalled to a Header message. While this unmarshalling-before-verifying
//    might not be ideal, it is unavoidable because i) the signature needs to also
//    protect the signing certificate; ii) it is desirable that Header is created
//    once by the client and never changed (for the sake of accountability and
//    non-repudiation). Note also that it is actually impossible to conclusively
//    verify the validity of the certificate included in a Proposal, because the
//    proposal needs to first be endorsed and ordered with respect to certificate
//    expiration transactions. Still, it is useful to pre-filter expired
//    certificates at this stage.
// 2. Verification that the certificate is trusted (signed by a trusted CA) and
//    that it is allowed to transact with us (with respect to some ACLs);
// 3. Verification that the signature on proposalBytes is valid;
// 4. Detect replay attacks;
type SignedProposal struct {
	Type int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	// The bytes of Proposal
	ProposalBytes []byte `protobuf:"bytes,2,opt,name=proposal_bytes,json=proposalBytes,proto3" json:"proposal_bytes,omitempty"`
	// Signaure over proposalBytes; this signature is to be verified against
	// the creator identity contained in the header of the Proposal message
	// marshaled as proposalBytes
	Signature            []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedProposal) Reset()         { *m = SignedProposal{} }
func (m *SignedProposal) String() string { return proto.CompactTextString(m) }
func (*SignedProposal) ProtoMessage()    {}
func (*SignedProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3ac5ce23bf32d05, []int{0}
}

func (m *SignedProposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedProposal.Unmarshal(m, b)
}
func (m *SignedProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedProposal.Marshal(b, m, deterministic)
}
func (m *SignedProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedProposal.Merge(m, src)
}
func (m *SignedProposal) XXX_Size() int {
	return xxx_messageInfo_SignedProposal.Size(m)
}
func (m *SignedProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SignedProposal proto.InternalMessageInfo

func (m *SignedProposal) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *SignedProposal) GetProposalBytes() []byte {
	if m != nil {
		return m.ProposalBytes
	}
	return nil
}

func (m *SignedProposal) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type SignatureList struct {
	SignatureList        [][]byte `protobuf:"bytes,1,rep,name=signature_list,json=signatureList,proto3" json:"signature_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignatureList) Reset()         { *m = SignatureList{} }
func (m *SignatureList) String() string { return proto.CompactTextString(m) }
func (*SignatureList) ProtoMessage()    {}
func (*SignatureList) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3ac5ce23bf32d05, []int{1}
}

func (m *SignatureList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureList.Unmarshal(m, b)
}
func (m *SignatureList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureList.Marshal(b, m, deterministic)
}
func (m *SignatureList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureList.Merge(m, src)
}
func (m *SignatureList) XXX_Size() int {
	return xxx_messageInfo_SignatureList.Size(m)
}
func (m *SignatureList) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureList.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureList proto.InternalMessageInfo

func (m *SignatureList) GetSignatureList() [][]byte {
	if m != nil {
		return m.SignatureList
	}
	return nil
}

type ProposalList struct {
	Proposals            []*Proposal `protobuf:"bytes,1,rep,name=proposals,proto3" json:"proposals,omitempty"`
	Graph                *Graph      `protobuf:"bytes,2,opt,name=graph,proto3" json:"graph,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ProposalList) Reset()         { *m = ProposalList{} }
func (m *ProposalList) String() string { return proto.CompactTextString(m) }
func (*ProposalList) ProtoMessage()    {}
func (*ProposalList) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3ac5ce23bf32d05, []int{2}
}

func (m *ProposalList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposalList.Unmarshal(m, b)
}
func (m *ProposalList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposalList.Marshal(b, m, deterministic)
}
func (m *ProposalList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalList.Merge(m, src)
}
func (m *ProposalList) XXX_Size() int {
	return xxx_messageInfo_ProposalList.Size(m)
}
func (m *ProposalList) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalList.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalList proto.InternalMessageInfo

func (m *ProposalList) GetProposals() []*Proposal {
	if m != nil {
		return m.Proposals
	}
	return nil
}

func (m *ProposalList) GetGraph() *Graph {
	if m != nil {
		return m.Graph
	}
	return nil
}

type Graph struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Graph) Reset()         { *m = Graph{} }
func (m *Graph) String() string { return proto.CompactTextString(m) }
func (*Graph) ProtoMessage()    {}
func (*Graph) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3ac5ce23bf32d05, []int{3}
}

func (m *Graph) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Graph.Unmarshal(m, b)
}
func (m *Graph) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Graph.Marshal(b, m, deterministic)
}
func (m *Graph) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Graph.Merge(m, src)
}
func (m *Graph) XXX_Size() int {
	return xxx_messageInfo_Graph.Size(m)
}
func (m *Graph) XXX_DiscardUnknown() {
	xxx_messageInfo_Graph.DiscardUnknown(m)
}

var xxx_messageInfo_Graph proto.InternalMessageInfo

func (m *Graph) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Node struct {
	Offset               int32    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	TxId                 string   `protobuf:"bytes,2,opt,name=txId,proto3" json:"txId,omitempty"`
	InDegree             int32    `protobuf:"varint,3,opt,name=inDegree,proto3" json:"inDegree,omitempty"`
	Edges                []*Edge  `protobuf:"bytes,4,rep,name=edges,proto3" json:"edges,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3ac5ce23bf32d05, []int{4}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Node) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *Node) GetInDegree() int32 {
	if m != nil {
		return m.InDegree
	}
	return 0
}

func (m *Node) GetEdges() []*Edge {
	if m != nil {
		return m.Edges
	}
	return nil
}

type Edge struct {
	TargetNodeId         int32    `protobuf:"varint,1,opt,name=targetNodeId,proto3" json:"targetNodeId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Edge) Reset()         { *m = Edge{} }
func (m *Edge) String() string { return proto.CompactTextString(m) }
func (*Edge) ProtoMessage()    {}
func (*Edge) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3ac5ce23bf32d05, []int{5}
}

func (m *Edge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Edge.Unmarshal(m, b)
}
func (m *Edge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Edge.Marshal(b, m, deterministic)
}
func (m *Edge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Edge.Merge(m, src)
}
func (m *Edge) XXX_Size() int {
	return xxx_messageInfo_Edge.Size(m)
}
func (m *Edge) XXX_DiscardUnknown() {
	xxx_messageInfo_Edge.DiscardUnknown(m)
}

var xxx_messageInfo_Edge proto.InternalMessageInfo

func (m *Edge) GetTargetNodeId() int32 {
	if m != nil {
		return m.TargetNodeId
	}
	return 0
}

// A Proposal is sent to an endorser for endorsement.  The proposal contains:
// 1. A header which should be unmarshaled to a Header message.  Note that
//    Header is both the header of a Proposal and of a Transaction, in that i)
//    both headers should be unmarshaled to this message; and ii) it is used to
//    compute cryptographic hashes and signatures.  The header has fields common
//    to all proposals/transactions.  In addition it has a type field for
//    additional customization. An example of this is the ChaincodeHeaderExtension
//    message used to extend the Header for type CHAINCODE.
// 2. A payload whose type depends on the header's type field.
// 3. An extension whose type depends on the header's type field.
//
// Let us see an example. For type CHAINCODE (see the Header message),
// we have the following:
// 1. The header is a Header message whose extensions field is a
//    ChaincodeHeaderExtension message.
// 2. The payload is a ChaincodeProposalPayload message.
// 3. The extension is a ChaincodeAction that might be used to ask the
//    endorsers to endorse a specific ChaincodeAction, thus emulating the
//    submitting peer model.
type Proposal struct {
	// The header of the proposal. It is the bytes of the Header
	Header []byte `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The payload of the proposal as defined by the type in the proposal
	// header.
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// Optional extensions to the proposal. Its content depends on the Header's
	// type field.  For the type CHAINCODE, it might be the bytes of a
	// ChaincodeAction message.
	Extension            []byte   `protobuf:"bytes,3,opt,name=extension,proto3" json:"extension,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Proposal) Reset()         { *m = Proposal{} }
func (m *Proposal) String() string { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()    {}
func (*Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3ac5ce23bf32d05, []int{6}
}

func (m *Proposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proposal.Unmarshal(m, b)
}
func (m *Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proposal.Marshal(b, m, deterministic)
}
func (m *Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposal.Merge(m, src)
}
func (m *Proposal) XXX_Size() int {
	return xxx_messageInfo_Proposal.Size(m)
}
func (m *Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_Proposal proto.InternalMessageInfo

func (m *Proposal) GetHeader() []byte {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Proposal) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Proposal) GetExtension() []byte {
	if m != nil {
		return m.Extension
	}
	return nil
}

// ChaincodeHeaderExtension is the Header's extentions message to be used when
// the Header's type is CHAINCODE.  This extensions is used to specify which
// chaincode to invoke and what should appear on the ledger.
type ChaincodeHeaderExtension struct {
	// The ID of the chaincode to target.
	ChaincodeId          *ChaincodeID `protobuf:"bytes,2,opt,name=chaincode_id,json=chaincodeId,proto3" json:"chaincode_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ChaincodeHeaderExtension) Reset()         { *m = ChaincodeHeaderExtension{} }
func (m *ChaincodeHeaderExtension) String() string { return proto.CompactTextString(m) }
func (*ChaincodeHeaderExtension) ProtoMessage()    {}
func (*ChaincodeHeaderExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3ac5ce23bf32d05, []int{7}
}

func (m *ChaincodeHeaderExtension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeHeaderExtension.Unmarshal(m, b)
}
func (m *ChaincodeHeaderExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeHeaderExtension.Marshal(b, m, deterministic)
}
func (m *ChaincodeHeaderExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeHeaderExtension.Merge(m, src)
}
func (m *ChaincodeHeaderExtension) XXX_Size() int {
	return xxx_messageInfo_ChaincodeHeaderExtension.Size(m)
}
func (m *ChaincodeHeaderExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeHeaderExtension.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeHeaderExtension proto.InternalMessageInfo

func (m *ChaincodeHeaderExtension) GetChaincodeId() *ChaincodeID {
	if m != nil {
		return m.ChaincodeId
	}
	return nil
}

// ChaincodeProposalPayload is the Proposal's payload message to be used when
// the Header's type is CHAINCODE.  It contains the arguments for this
// invocation.
type ChaincodeProposalPayload struct {
	// Input contains the arguments for this invocation. If this invocation
	// deploys a new chaincode, ESCC/VSCC are part of this field.
	// This is usually a marshaled ChaincodeInvocationSpec
	Input []byte `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	// TransientMap contains data (e.g. cryptographic material) that might be used
	// to implement some form of application-level confidentiality. The contents
	// of this field are supposed to always be omitted from the transaction and
	// excluded from the ledger.
	TransientMap         map[string][]byte `protobuf:"bytes,2,rep,name=TransientMap,proto3" json:"TransientMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ChaincodeProposalPayload) Reset()         { *m = ChaincodeProposalPayload{} }
func (m *ChaincodeProposalPayload) String() string { return proto.CompactTextString(m) }
func (*ChaincodeProposalPayload) ProtoMessage()    {}
func (*ChaincodeProposalPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3ac5ce23bf32d05, []int{8}
}

func (m *ChaincodeProposalPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeProposalPayload.Unmarshal(m, b)
}
func (m *ChaincodeProposalPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeProposalPayload.Marshal(b, m, deterministic)
}
func (m *ChaincodeProposalPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeProposalPayload.Merge(m, src)
}
func (m *ChaincodeProposalPayload) XXX_Size() int {
	return xxx_messageInfo_ChaincodeProposalPayload.Size(m)
}
func (m *ChaincodeProposalPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeProposalPayload.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeProposalPayload proto.InternalMessageInfo

func (m *ChaincodeProposalPayload) GetInput() []byte {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *ChaincodeProposalPayload) GetTransientMap() map[string][]byte {
	if m != nil {
		return m.TransientMap
	}
	return nil
}

// ChaincodeAction contains the executed chaincode results, response, and event.
type ChaincodeAction struct {
	// This field contains the read set and the write set produced by the
	// chaincode executing this invocation.
	Results []byte `protobuf:"bytes,1,opt,name=results,proto3" json:"results,omitempty"`
	// This field contains the event generated by the chaincode.
	// Only a single marshaled ChaincodeEvent is included.
	Events []byte `protobuf:"bytes,2,opt,name=events,proto3" json:"events,omitempty"`
	// This field contains the result of executing this invocation.
	Response *Response `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
	// This field contains the ChaincodeID of executing this invocation. Endorser
	// will set it with the ChaincodeID called by endorser while simulating proposal.
	// Committer will validate the version matching with latest chaincode version.
	// Adding ChaincodeID to keep version opens up the possibility of multiple
	// ChaincodeAction per transaction.
	ChaincodeId          *ChaincodeID `protobuf:"bytes,4,opt,name=chaincode_id,json=chaincodeId,proto3" json:"chaincode_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ChaincodeAction) Reset()         { *m = ChaincodeAction{} }
func (m *ChaincodeAction) String() string { return proto.CompactTextString(m) }
func (*ChaincodeAction) ProtoMessage()    {}
func (*ChaincodeAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3ac5ce23bf32d05, []int{9}
}

func (m *ChaincodeAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeAction.Unmarshal(m, b)
}
func (m *ChaincodeAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeAction.Marshal(b, m, deterministic)
}
func (m *ChaincodeAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeAction.Merge(m, src)
}
func (m *ChaincodeAction) XXX_Size() int {
	return xxx_messageInfo_ChaincodeAction.Size(m)
}
func (m *ChaincodeAction) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeAction.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeAction proto.InternalMessageInfo

func (m *ChaincodeAction) GetResults() []byte {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ChaincodeAction) GetEvents() []byte {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *ChaincodeAction) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *ChaincodeAction) GetChaincodeId() *ChaincodeID {
	if m != nil {
		return m.ChaincodeId
	}
	return nil
}

func init() {
	proto.RegisterType((*SignedProposal)(nil), "protos.SignedProposal")
	proto.RegisterType((*SignatureList)(nil), "protos.SignatureList")
	proto.RegisterType((*ProposalList)(nil), "protos.ProposalList")
	proto.RegisterType((*Graph)(nil), "protos.Graph")
	proto.RegisterType((*Node)(nil), "protos.Node")
	proto.RegisterType((*Edge)(nil), "protos.Edge")
	proto.RegisterType((*Proposal)(nil), "protos.Proposal")
	proto.RegisterType((*ChaincodeHeaderExtension)(nil), "protos.ChaincodeHeaderExtension")
	proto.RegisterType((*ChaincodeProposalPayload)(nil), "protos.ChaincodeProposalPayload")
	proto.RegisterMapType((map[string][]byte)(nil), "protos.ChaincodeProposalPayload.TransientMapEntry")
	proto.RegisterType((*ChaincodeAction)(nil), "protos.ChaincodeAction")
}

func init() { proto.RegisterFile("proposal.proto", fileDescriptor_c3ac5ce23bf32d05) }

var fileDescriptor_c3ac5ce23bf32d05 = []byte{
	// 627 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x41, 0x4f, 0x1b, 0x3b,
	0x10, 0xd6, 0x86, 0x24, 0x24, 0x43, 0x80, 0xe0, 0x87, 0xd0, 0x2a, 0xe2, 0x80, 0xf6, 0xe9, 0x49,
	0xe8, 0x3d, 0x48, 0x24, 0x9e, 0x84, 0xaa, 0x5e, 0xaa, 0x52, 0x50, 0x9b, 0xaa, 0xad, 0x90, 0xa9,
	0x7a, 0xe0, 0x12, 0x39, 0xbb, 0xc3, 0xc6, 0xca, 0xd6, 0x5e, 0xd9, 0x0e, 0x62, 0x7f, 0x5e, 0x8f,
	0xfd, 0x57, 0x95, 0xbd, 0xf6, 0x26, 0x94, 0x4b, 0x4f, 0xeb, 0x6f, 0xe6, 0x9b, 0x99, 0x6f, 0x66,
	0xbc, 0x86, 0xbd, 0x52, 0xc9, 0x52, 0x6a, 0x56, 0x8c, 0x4b, 0x25, 0x8d, 0x24, 0x5d, 0xf7, 0xd1,
	0xa3, 0xc3, 0x12, 0x51, 0x4d, 0xd2, 0x05, 0xe3, 0x22, 0x95, 0x19, 0xd6, 0xde, 0xd1, 0xb1, 0xb3,
	0x86, 0x90, 0x99, 0x42, 0x5d, 0x4a, 0xa1, 0xbd, 0x37, 0xe1, 0xb0, 0x77, 0xc7, 0x73, 0x81, 0xd9,
	0xad, 0x27, 0x10, 0x02, 0x6d, 0x53, 0x95, 0x18, 0x47, 0x27, 0xd1, 0x69, 0x87, 0xba, 0x33, 0xf9,
	0x67, 0x5d, 0x73, 0x36, 0xaf, 0x0c, 0xea, 0xb8, 0x75, 0x12, 0x9d, 0x0e, 0xe8, 0x6e, 0xb0, 0x5e,
	0x59, 0x23, 0x39, 0x86, 0xbe, 0xe6, 0xb9, 0x60, 0x66, 0xa5, 0x30, 0xde, 0x72, 0x8c, 0xb5, 0x21,
	0xb9, 0x84, 0xdd, 0xbb, 0x00, 0x3e, 0x71, 0x6d, 0x6c, 0xd6, 0xc6, 0x3b, 0x2b, 0xb8, 0x36, 0x71,
	0x74, 0xb2, 0x65, 0xb3, 0xea, 0x4d, 0x5a, 0x92, 0xc2, 0x20, 0x88, 0x73, 0x61, 0x63, 0xe8, 0x87,
	0xb2, 0xda, 0x45, 0xec, 0x5c, 0x0c, 0xeb, 0x6e, 0xf4, 0x38, 0x10, 0xe9, 0x9a, 0x42, 0xfe, 0x86,
	0x4e, 0xae, 0x58, 0xb9, 0x70, 0x9a, 0x77, 0x2e, 0x76, 0x03, 0xf7, 0xbd, 0x35, 0xd2, 0xda, 0x97,
	0xfc, 0x07, 0x1d, 0x87, 0x49, 0x02, 0x1d, 0x21, 0x33, 0x0c, 0x99, 0x07, 0x81, 0xfd, 0x45, 0x66,
	0x48, 0x6b, 0x57, 0xa2, 0xa0, 0x6d, 0x21, 0x39, 0x82, 0xae, 0x7c, 0x78, 0xd0, 0x68, 0xfc, 0xb0,
	0x3c, 0x72, 0x23, 0x7c, 0x9a, 0x66, 0xae, 0x60, 0x9f, 0xba, 0x33, 0x19, 0x41, 0x8f, 0x8b, 0x6b,
	0xcc, 0x15, 0xd6, 0xa3, 0xe9, 0xd0, 0x06, 0xdb, 0x9a, 0x98, 0xe5, 0xa8, 0xe3, 0xf6, 0xf3, 0x9a,
	0x37, 0x59, 0x8e, 0xb4, 0x76, 0x25, 0xff, 0x42, 0xdb, 0x42, 0x92, 0xc0, 0xc0, 0x30, 0x95, 0xa3,
	0xb1, 0x0a, 0xa6, 0x99, 0xaf, 0xfc, 0xcc, 0x96, 0xdc, 0x43, 0xaf, 0x59, 0xe7, 0x11, 0x74, 0x17,
	0xc8, 0x32, 0x54, 0x8e, 0x39, 0xa0, 0x1e, 0x91, 0x18, 0xb6, 0x4b, 0x56, 0x15, 0x92, 0x65, 0x7e,
	0x97, 0x01, 0xda, 0x2d, 0xe2, 0x93, 0x41, 0xa1, 0xb9, 0x14, 0x61, 0x8b, 0x8d, 0x21, 0x59, 0x42,
	0xfc, 0x2e, 0xdc, 0xb0, 0x0f, 0x2e, 0xd5, 0x4d, 0xf0, 0x91, 0x4b, 0x18, 0x34, 0xb7, 0x6f, 0xc6,
	0x33, 0x3f, 0xf0, 0xbf, 0x42, 0x3b, 0x4d, 0xdc, 0xf4, 0x9a, 0xee, 0x34, 0xc4, 0x69, 0xf6, 0xb1,
	0xdd, 0x8b, 0x86, 0x2d, 0x7a, 0xe0, 0x05, 0xcc, 0x1e, 0xb9, 0x9e, 0xf3, 0x82, 0x9b, 0x2a, 0xf9,
	0x19, 0x6d, 0x54, 0x0b, 0x2d, 0xdd, 0x7a, 0x9d, 0x87, 0xd0, 0xe1, 0xa2, 0x5c, 0x19, 0xdf, 0x58,
	0x0d, 0xc8, 0x37, 0x18, 0x7c, 0x55, 0x4c, 0x68, 0x8e, 0xc2, 0x7c, 0x66, 0x65, 0xdc, 0x72, 0x23,
	0xbd, 0x78, 0xa1, 0xe1, 0xb7, 0x6c, 0xe3, 0xcd, 0xa0, 0x1b, 0x61, 0x54, 0x45, 0x9f, 0xe5, 0x19,
	0xbd, 0x81, 0x83, 0x17, 0x14, 0x32, 0x84, 0xad, 0x25, 0x56, 0x4e, 0x40, 0x9f, 0xda, 0xa3, 0x15,
	0xf5, 0xc8, 0x8a, 0x15, 0xfa, 0xa1, 0xd6, 0xe0, 0x75, 0xeb, 0x55, 0x94, 0xfc, 0x88, 0x60, 0xbf,
	0xa9, 0xfe, 0x36, 0x35, 0x76, 0x60, 0x31, 0x6c, 0x2b, 0xd4, 0xab, 0xc2, 0x68, 0xdf, 0x44, 0x80,
	0x76, 0x6d, 0xf8, 0x88, 0xc2, 0x84, 0x3f, 0xcd, 0x23, 0x72, 0x06, 0xbd, 0xf0, 0x07, 0xbb, 0xdd,
	0x6c, 0xdc, 0x7d, 0xea, 0xed, 0xb4, 0x61, 0xbc, 0x58, 0x48, 0xfb, 0x8f, 0x17, 0xd2, 0x19, 0x76,
	0xe9, 0xd0, 0xc8, 0x25, 0x8a, 0x99, 0x2c, 0x51, 0x31, 0x2b, 0x57, 0x5f, 0xa5, 0x90, 0x48, 0x95,
	0x8f, 0x17, 0x55, 0x89, 0xaa, 0xb0, 0xf7, 0x52, 0x8d, 0x1f, 0xd8, 0x5c, 0xf1, 0x34, 0x64, 0xb4,
	0x6f, 0xcd, 0xd5, 0xfe, 0x7a, 0xb6, 0xe9, 0x92, 0xe5, 0x78, 0x7f, 0x96, 0x73, 0xb3, 0x58, 0xcd,
	0xc7, 0xa9, 0xfc, 0x3e, 0xd9, 0x88, 0x9d, 0xd4, 0xb1, 0xe7, 0x75, 0xec, 0x79, 0x2e, 0x27, 0x36,
	0x7c, 0x5e, 0x3f, 0x67, 0xff, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x55, 0x22, 0xc4, 0xe7,
	0x04, 0x00, 0x00,
}
