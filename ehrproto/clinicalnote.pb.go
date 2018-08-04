// Code generated by protoc-gen-go. DO NOT EDIT.
// source: clinicalnote.proto

package ehrpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FragmentTopic int32

const (
	FragmentTopic_NO_TOPIC        FragmentTopic = 0
	FragmentTopic_SUBJECTIVE      FragmentTopic = 1
	FragmentTopic_MEDICAL_HISTORY FragmentTopic = 2
	FragmentTopic_ALLERGIES       FragmentTopic = 3
	FragmentTopic_MEDICATIONS     FragmentTopic = 4
	FragmentTopic_FAMILY_HISTORY  FragmentTopic = 5
	FragmentTopic_SOCIAL_HISTORY  FragmentTopic = 6
	FragmentTopic_VITALS          FragmentTopic = 7
	FragmentTopic_PHYSICAL_EXAM   FragmentTopic = 8
	FragmentTopic_MEDICAL_PROBLEM FragmentTopic = 9
)

var FragmentTopic_name = map[int32]string{
	0: "NO_TOPIC",
	1: "SUBJECTIVE",
	2: "MEDICAL_HISTORY",
	3: "ALLERGIES",
	4: "MEDICATIONS",
	5: "FAMILY_HISTORY",
	6: "SOCIAL_HISTORY",
	7: "VITALS",
	8: "PHYSICAL_EXAM",
	9: "MEDICAL_PROBLEM",
}
var FragmentTopic_value = map[string]int32{
	"NO_TOPIC":        0,
	"SUBJECTIVE":      1,
	"MEDICAL_HISTORY": 2,
	"ALLERGIES":       3,
	"MEDICATIONS":     4,
	"FAMILY_HISTORY":  5,
	"SOCIAL_HISTORY":  6,
	"VITALS":          7,
	"PHYSICAL_EXAM":   8,
	"MEDICAL_PROBLEM": 9,
}

func (x FragmentTopic) String() string {
	return proto.EnumName(FragmentTopic_name, int32(x))
}
func (FragmentTopic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_clinicalnote_47f3cf4497da3f58, []int{0}
}

type FragmentPriority int32

const (
	FragmentPriority_NO_PRIORITY FragmentPriority = 0
	FragmentPriority_HIGH        FragmentPriority = 1
	FragmentPriority_NORMAL      FragmentPriority = 2
	FragmentPriority_LOW         FragmentPriority = 3
)

var FragmentPriority_name = map[int32]string{
	0: "NO_PRIORITY",
	1: "HIGH",
	2: "NORMAL",
	3: "LOW",
}
var FragmentPriority_value = map[string]int32{
	"NO_PRIORITY": 0,
	"HIGH":        1,
	"NORMAL":      2,
	"LOW":         3,
}

func (x FragmentPriority) String() string {
	return proto.EnumName(FragmentPriority_name, int32(x))
}
func (FragmentPriority) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_clinicalnote_47f3cf4497da3f58, []int{1}
}

type NoteFragmentStatus int32

const (
	NoteFragmentStatus_NO_STATUS  NoteFragmentStatus = 0
	NoteFragmentStatus_INCOMPLETE NoteFragmentStatus = 1
	NoteFragmentStatus_ACTIVE     NoteFragmentStatus = 2
	NoteFragmentStatus_REPLACED   NoteFragmentStatus = 3
)

var NoteFragmentStatus_name = map[int32]string{
	0: "NO_STATUS",
	1: "INCOMPLETE",
	2: "ACTIVE",
	3: "REPLACED",
}
var NoteFragmentStatus_value = map[string]int32{
	"NO_STATUS":  0,
	"INCOMPLETE": 1,
	"ACTIVE":     2,
	"REPLACED":   3,
}

func (x NoteFragmentStatus) String() string {
	return proto.EnumName(NoteFragmentStatus_name, int32(x))
}
func (NoteFragmentStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_clinicalnote_47f3cf4497da3f58, []int{2}
}

type NoteType int32

const (
	NoteType_NO_NOTE_TYPE                 NoteType = 0
	NoteType_FOLLOW_UP                    NoteType = 1
	NoteType_PHONE_CALL                   NoteType = 2
	NoteType_CONTINUED_CARE_DOCUMENTATION NoteType = 4
	NoteType_REFILL_DOCUMENTATION         NoteType = 5
	NoteType_INTAKE                       NoteType = 6
	NoteType_PROCEDURE                    NoteType = 7
	NoteType_HISTORY_AND_PHYSICAL         NoteType = 8
)

var NoteType_name = map[int32]string{
	0: "NO_NOTE_TYPE",
	1: "FOLLOW_UP",
	2: "PHONE_CALL",
	4: "CONTINUED_CARE_DOCUMENTATION",
	5: "REFILL_DOCUMENTATION",
	6: "INTAKE",
	7: "PROCEDURE",
	8: "HISTORY_AND_PHYSICAL",
}
var NoteType_value = map[string]int32{
	"NO_NOTE_TYPE":                 0,
	"FOLLOW_UP":                    1,
	"PHONE_CALL":                   2,
	"CONTINUED_CARE_DOCUMENTATION": 4,
	"REFILL_DOCUMENTATION":         5,
	"INTAKE":                       6,
	"PROCEDURE":                    7,
	"HISTORY_AND_PHYSICAL":         8,
}

func (x NoteType) String() string {
	return proto.EnumName(NoteType_name, int32(x))
}
func (NoteType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_clinicalnote_47f3cf4497da3f58, []int{3}
}

type Note struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DateCreated          *timestamp.Timestamp `protobuf:"bytes,2,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	NoteGuid             string               `protobuf:"bytes,3,opt,name=note_guid,json=noteGuid,proto3" json:"note_guid,omitempty"`
	VisitGuid            string               `protobuf:"bytes,4,opt,name=visit_guid,json=visitGuid,proto3" json:"visit_guid,omitempty"`
	AuthorGuid           string               `protobuf:"bytes,5,opt,name=author_guid,json=authorGuid,proto3" json:"author_guid,omitempty"`
	PatientGuid          string               `protobuf:"bytes,6,opt,name=patient_guid,json=patientGuid,proto3" json:"patient_guid,omitempty"`
	Type                 NoteType             `protobuf:"varint,7,opt,name=type,proto3,enum=ehr.clinicalnote.NoteType" json:"type,omitempty"`
	Fragments            []*NoteFragment      `protobuf:"bytes,8,rep,name=fragments,proto3" json:"fragments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Note) Reset()         { *m = Note{} }
func (m *Note) String() string { return proto.CompactTextString(m) }
func (*Note) ProtoMessage()    {}
func (*Note) Descriptor() ([]byte, []int) {
	return fileDescriptor_clinicalnote_47f3cf4497da3f58, []int{0}
}
func (m *Note) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Note.Unmarshal(m, b)
}
func (m *Note) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Note.Marshal(b, m, deterministic)
}
func (dst *Note) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Note.Merge(dst, src)
}
func (m *Note) XXX_Size() int {
	return xxx_messageInfo_Note.Size(m)
}
func (m *Note) XXX_DiscardUnknown() {
	xxx_messageInfo_Note.DiscardUnknown(m)
}

var xxx_messageInfo_Note proto.InternalMessageInfo

func (m *Note) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Note) GetDateCreated() *timestamp.Timestamp {
	if m != nil {
		return m.DateCreated
	}
	return nil
}

func (m *Note) GetNoteGuid() string {
	if m != nil {
		return m.NoteGuid
	}
	return ""
}

func (m *Note) GetVisitGuid() string {
	if m != nil {
		return m.VisitGuid
	}
	return ""
}

func (m *Note) GetAuthorGuid() string {
	if m != nil {
		return m.AuthorGuid
	}
	return ""
}

func (m *Note) GetPatientGuid() string {
	if m != nil {
		return m.PatientGuid
	}
	return ""
}

func (m *Note) GetType() NoteType {
	if m != nil {
		return m.Type
	}
	return NoteType_NO_NOTE_TYPE
}

func (m *Note) GetFragments() []*NoteFragment {
	if m != nil {
		return m.Fragments
	}
	return nil
}

type NoteFragment struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DateCreated          *timestamp.Timestamp `protobuf:"bytes,2,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	NoteFragmentGuid     string               `protobuf:"bytes,3,opt,name=note_fragment_guid,json=noteFragmentGuid,proto3" json:"note_fragment_guid,omitempty"`
	NoteGuid             string               `protobuf:"bytes,4,opt,name=note_guid,json=noteGuid,proto3" json:"note_guid,omitempty"`
	VisitGuid            string               `protobuf:"bytes,5,opt,name=visit_guid,json=visitGuid,proto3" json:"visit_guid,omitempty"`
	AuthorGuid           string               `protobuf:"bytes,6,opt,name=author_guid,json=authorGuid,proto3" json:"author_guid,omitempty"`
	PatientGuid          string               `protobuf:"bytes,7,opt,name=patient_guid,json=patientGuid,proto3" json:"patient_guid,omitempty"`
	Issue                MedicalIssue         `protobuf:"varint,8,opt,name=issue,proto3,enum=ehr.medicalissues.MedicalIssue" json:"issue,omitempty"`
	Icd_10Code           string               `protobuf:"bytes,9,opt,name=icd_10_code,json=icd10Code,proto3" json:"icd_10_code,omitempty"`
	Status               NoteFragmentStatus   `protobuf:"varint,10,opt,name=status,proto3,enum=ehr.clinicalnote.NoteFragmentStatus" json:"status,omitempty"`
	Priority             FragmentPriority     `protobuf:"varint,11,opt,name=priority,proto3,enum=ehr.clinicalnote.FragmentPriority" json:"priority,omitempty"`
	Topic                FragmentTopic        `protobuf:"varint,12,opt,name=topic,proto3,enum=ehr.clinicalnote.FragmentTopic" json:"topic,omitempty"`
	MarkdownContent      string               `protobuf:"bytes,13,opt,name=markdown_content,json=markdownContent,proto3" json:"markdown_content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *NoteFragment) Reset()         { *m = NoteFragment{} }
func (m *NoteFragment) String() string { return proto.CompactTextString(m) }
func (*NoteFragment) ProtoMessage()    {}
func (*NoteFragment) Descriptor() ([]byte, []int) {
	return fileDescriptor_clinicalnote_47f3cf4497da3f58, []int{1}
}
func (m *NoteFragment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoteFragment.Unmarshal(m, b)
}
func (m *NoteFragment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoteFragment.Marshal(b, m, deterministic)
}
func (dst *NoteFragment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoteFragment.Merge(dst, src)
}
func (m *NoteFragment) XXX_Size() int {
	return xxx_messageInfo_NoteFragment.Size(m)
}
func (m *NoteFragment) XXX_DiscardUnknown() {
	xxx_messageInfo_NoteFragment.DiscardUnknown(m)
}

var xxx_messageInfo_NoteFragment proto.InternalMessageInfo

func (m *NoteFragment) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NoteFragment) GetDateCreated() *timestamp.Timestamp {
	if m != nil {
		return m.DateCreated
	}
	return nil
}

func (m *NoteFragment) GetNoteFragmentGuid() string {
	if m != nil {
		return m.NoteFragmentGuid
	}
	return ""
}

func (m *NoteFragment) GetNoteGuid() string {
	if m != nil {
		return m.NoteGuid
	}
	return ""
}

func (m *NoteFragment) GetVisitGuid() string {
	if m != nil {
		return m.VisitGuid
	}
	return ""
}

func (m *NoteFragment) GetAuthorGuid() string {
	if m != nil {
		return m.AuthorGuid
	}
	return ""
}

func (m *NoteFragment) GetPatientGuid() string {
	if m != nil {
		return m.PatientGuid
	}
	return ""
}

func (m *NoteFragment) GetIssue() MedicalIssue {
	if m != nil {
		return m.Issue
	}
	return MedicalIssue_NO_MEDICAL_ISSUE
}

func (m *NoteFragment) GetIcd_10Code() string {
	if m != nil {
		return m.Icd_10Code
	}
	return ""
}

func (m *NoteFragment) GetStatus() NoteFragmentStatus {
	if m != nil {
		return m.Status
	}
	return NoteFragmentStatus_NO_STATUS
}

func (m *NoteFragment) GetPriority() FragmentPriority {
	if m != nil {
		return m.Priority
	}
	return FragmentPriority_NO_PRIORITY
}

func (m *NoteFragment) GetTopic() FragmentTopic {
	if m != nil {
		return m.Topic
	}
	return FragmentTopic_NO_TOPIC
}

func (m *NoteFragment) GetMarkdownContent() string {
	if m != nil {
		return m.MarkdownContent
	}
	return ""
}

func init() {
	proto.RegisterType((*Note)(nil), "ehr.clinicalnote.Note")
	proto.RegisterType((*NoteFragment)(nil), "ehr.clinicalnote.NoteFragment")
	proto.RegisterEnum("ehr.clinicalnote.FragmentTopic", FragmentTopic_name, FragmentTopic_value)
	proto.RegisterEnum("ehr.clinicalnote.FragmentPriority", FragmentPriority_name, FragmentPriority_value)
	proto.RegisterEnum("ehr.clinicalnote.NoteFragmentStatus", NoteFragmentStatus_name, NoteFragmentStatus_value)
	proto.RegisterEnum("ehr.clinicalnote.NoteType", NoteType_name, NoteType_value)
}

func init() { proto.RegisterFile("clinicalnote.proto", fileDescriptor_clinicalnote_47f3cf4497da3f58) }

var fileDescriptor_clinicalnote_47f3cf4497da3f58 = []byte{
	// 791 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xdf, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0xeb, 0xfc, 0x75, 0x4e, 0xd2, 0x76, 0x98, 0xe5, 0xc2, 0x2a, 0xb0, 0x0d, 0x15, 0x17,
	0xa1, 0x42, 0xd9, 0xdd, 0x22, 0xee, 0x00, 0xc9, 0x75, 0xa6, 0x8d, 0xc1, 0xf6, 0x58, 0xe3, 0xc9,
	0x2e, 0xe1, 0x66, 0xe4, 0xc6, 0xb3, 0xa9, 0x45, 0x13, 0x47, 0xce, 0x04, 0xd4, 0x77, 0xe2, 0x19,
	0xb8, 0x80, 0x17, 0x43, 0x33, 0x4e, 0x96, 0xb4, 0x5d, 0xed, 0xde, 0x70, 0xe9, 0x73, 0xbe, 0xef,
	0x7c, 0xe7, 0xfc, 0xec, 0x04, 0xf0, 0xec, 0x2e, 0x5f, 0xe6, 0xb3, 0xf4, 0x6e, 0x59, 0x28, 0x39,
	0x5c, 0x95, 0x85, 0x2a, 0x30, 0x92, 0xb7, 0xe5, 0x70, 0xbf, 0x7e, 0x72, 0x3a, 0x2f, 0x8a, 0xf9,
	0x9d, 0x7c, 0x61, 0xfa, 0x37, 0x9b, 0xb7, 0x2f, 0x54, 0xbe, 0x90, 0x6b, 0x95, 0x2e, 0x56, 0x95,
	0xe5, 0xe4, 0xd9, 0x42, 0x66, 0x5a, 0x9d, 0xaf, 0xd7, 0x1b, 0xb9, 0xae, 0x8a, 0x67, 0xff, 0xd4,
	0xa0, 0x11, 0x15, 0x4a, 0xe2, 0x23, 0xa8, 0xe5, 0x99, 0x63, 0xf5, 0xad, 0x41, 0x93, 0xd5, 0xf2,
	0x0c, 0xff, 0x00, 0xbd, 0x2c, 0x55, 0x52, 0xcc, 0x4a, 0x99, 0x2a, 0x99, 0x39, 0xb5, 0xbe, 0x35,
	0xe8, 0x5e, 0x9c, 0x0c, 0xab, 0x94, 0xe1, 0x2e, 0x65, 0xc8, 0x77, 0x29, 0xac, 0xab, 0xf5, 0x5e,
	0x25, 0xc7, 0x9f, 0x41, 0x47, 0x6f, 0x25, 0xe6, 0x9b, 0x3c, 0x73, 0xea, 0x7d, 0x6b, 0xd0, 0x61,
	0xb6, 0x2e, 0x5c, 0x6f, 0xf2, 0x0c, 0x7f, 0x01, 0xf0, 0x7b, 0xbe, 0xce, 0x55, 0xd5, 0x6d, 0x98,
	0x6e, 0xc7, 0x54, 0x4c, 0xfb, 0x14, 0xba, 0xe9, 0x46, 0xdd, 0x16, 0x65, 0xd5, 0x6f, 0x9a, 0x3e,
	0x54, 0x25, 0x23, 0xf8, 0x12, 0x7a, 0xab, 0x54, 0xe5, 0x72, 0xb9, 0x9d, 0xd0, 0x32, 0x8a, 0xee,
	0xb6, 0x66, 0x24, 0x43, 0x68, 0xa8, 0xfb, 0x95, 0x74, 0xda, 0x7d, 0x6b, 0x70, 0x74, 0x71, 0x32,
	0x7c, 0x8c, 0x6b, 0xa8, 0x8f, 0xe6, 0xf7, 0x2b, 0xc9, 0x8c, 0x0e, 0x7f, 0x0f, 0x9d, 0xb7, 0x65,
	0x3a, 0x5f, 0xc8, 0xa5, 0x5a, 0x3b, 0x76, 0xbf, 0x3e, 0xe8, 0x5e, 0x3c, 0x7f, 0xbf, 0xe9, 0x6a,
	0x2b, 0x63, 0xff, 0x19, 0xce, 0xfe, 0x6e, 0x40, 0x6f, 0xbf, 0xf7, 0x7f, 0xd3, 0xfc, 0x06, 0xb0,
	0xa1, 0xb9, 0x4b, 0xdc, 0xc7, 0x8a, 0x96, 0x7b, 0xc1, 0xe6, 0xf6, 0x07, 0xec, 0x1b, 0x1f, 0x64,
	0xdf, 0xfc, 0x08, 0xfb, 0xd6, 0x47, 0xd9, 0xb7, 0x9f, 0xb2, 0xff, 0x0e, 0x9a, 0xe6, 0x1b, 0x73,
	0x6c, 0x03, 0xff, 0xd4, 0x70, 0x7c, 0xf8, 0xf1, 0x85, 0xd5, 0x93, 0xaf, 0x9f, 0x58, 0xa5, 0xc6,
	0xcf, 0xa1, 0x9b, 0xcf, 0x32, 0xf1, 0xea, 0xa5, 0x98, 0x15, 0x99, 0x74, 0x3a, 0xd5, 0x6a, 0xf9,
	0x2c, 0x7b, 0xf5, 0xd2, 0x2b, 0x32, 0xfd, 0x8a, 0x5a, 0x6b, 0x95, 0xaa, 0xcd, 0xda, 0x01, 0x33,
	0xf7, 0xab, 0x0f, 0xbf, 0x9f, 0xc4, 0x68, 0xd9, 0xd6, 0x83, 0x7f, 0x04, 0x7b, 0x55, 0xe6, 0x45,
	0x99, 0xab, 0x7b, 0xa7, 0x6b, 0xfc, 0x67, 0x4f, 0xfd, 0x3b, 0x6f, 0xbc, 0x55, 0xb2, 0x77, 0x1e,
	0x7d, 0x94, 0x2a, 0x56, 0xf9, 0xcc, 0xe9, 0xed, 0x1d, 0xf5, 0x5e, 0x33, 0xd7, 0x32, 0x56, 0xa9,
	0xf1, 0xd7, 0x80, 0x16, 0x69, 0xf9, 0x5b, 0x56, 0xfc, 0xb1, 0x14, 0xb3, 0x62, 0xa9, 0xe4, 0x52,
	0x39, 0x87, 0xe6, 0xb2, 0xe3, 0x5d, 0xdd, 0xab, 0xca, 0xe7, 0x7f, 0x59, 0x70, 0xf8, 0x60, 0x06,
	0xee, 0x81, 0x1d, 0x51, 0xc1, 0x69, 0xec, 0x7b, 0xe8, 0x00, 0x1f, 0x01, 0x24, 0x93, 0xcb, 0x9f,
	0x88, 0xc7, 0xfd, 0xd7, 0x04, 0x59, 0xf8, 0x19, 0x1c, 0x87, 0x64, 0xe4, 0x7b, 0x6e, 0x20, 0xc6,
	0x7e, 0xc2, 0x29, 0x9b, 0xa2, 0x1a, 0x3e, 0x84, 0x8e, 0x1b, 0x04, 0x84, 0x5d, 0xfb, 0x24, 0x41,
	0x75, 0x7c, 0x0c, 0xdd, 0x4a, 0xc3, 0x7d, 0x1a, 0x25, 0xa8, 0x81, 0x31, 0x1c, 0x5d, 0xb9, 0xa1,
	0x1f, 0x4c, 0xdf, 0x79, 0x9a, 0xba, 0x96, 0x50, 0xcf, 0xdf, 0x9b, 0xd3, 0xc2, 0x00, 0xad, 0xd7,
	0x3e, 0x77, 0x83, 0x04, 0xb5, 0xf1, 0x27, 0x70, 0x18, 0x8f, 0xa7, 0x89, 0x49, 0x22, 0xbf, 0xb8,
	0x21, 0xb2, 0xf7, 0xb3, 0x63, 0x46, 0x2f, 0x03, 0x12, 0xa2, 0xce, 0xf9, 0x25, 0xa0, 0xc7, 0x00,
	0xf5, 0x02, 0x11, 0x15, 0x31, 0xf3, 0x29, 0xf3, 0xf9, 0x14, 0x1d, 0x60, 0x1b, 0x1a, 0x63, 0xff,
	0x7a, 0x8c, 0x2c, 0x1d, 0x11, 0x51, 0x16, 0xba, 0x01, 0xaa, 0xe1, 0x36, 0xd4, 0x03, 0xfa, 0x06,
	0xd5, 0xcf, 0x43, 0xc0, 0x4f, 0x5f, 0xa2, 0xbe, 0x2a, 0xa2, 0x22, 0xe1, 0x2e, 0x9f, 0x24, 0x15,
	0x09, 0x3f, 0xf2, 0x68, 0x18, 0x07, 0x84, 0x93, 0x6a, 0x92, 0x5b, 0x51, 0xa9, 0x69, 0x66, 0x8c,
	0xc4, 0x81, 0xeb, 0x91, 0x11, 0xaa, 0x9f, 0xff, 0x69, 0x81, 0xbd, 0xfb, 0xa5, 0x63, 0x04, 0xbd,
	0x88, 0x8a, 0x88, 0x72, 0x22, 0xf8, 0x34, 0x26, 0xe8, 0x40, 0xcf, 0xbd, 0xa2, 0x41, 0x40, 0xdf,
	0x88, 0x49, 0x8c, 0x2c, 0x3d, 0x37, 0x1e, 0xd3, 0x88, 0x08, 0xcf, 0x0d, 0xf4, 0x56, 0x7d, 0xf8,
	0xdc, 0xa3, 0x11, 0xf7, 0xa3, 0x09, 0x19, 0x09, 0xcf, 0x65, 0x44, 0x8c, 0xa8, 0x37, 0x09, 0x49,
	0xc4, 0x0d, 0x4f, 0xd4, 0xc0, 0x0e, 0x7c, 0xca, 0xc8, 0x95, 0x1f, 0x04, 0x8f, 0x3a, 0x4d, 0xbd,
	0x93, 0x1f, 0x71, 0xf7, 0x67, 0x82, 0x5a, 0x3a, 0x26, 0x66, 0xd4, 0x23, 0xa3, 0x09, 0x23, 0xa8,
	0xad, 0x4d, 0x5b, 0xd0, 0xc2, 0x8d, 0x46, 0x62, 0xc7, 0x16, 0xd9, 0x97, 0xed, 0x5f, 0x9b, 0xf2,
	0xb6, 0x5c, 0xdd, 0xdc, 0xb4, 0xcc, 0x3f, 0xc2, 0xb7, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe3,
	0xe4, 0x16, 0xe0, 0xfb, 0x05, 0x00, 0x00,
}