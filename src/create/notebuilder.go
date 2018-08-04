package create

import (
	"github.com/geekmdio/noted/src/ehrproto"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type NoteBuilder struct {
	note *ehrpb.Note
}

func (nb *NoteBuilder) Init() *NoteBuilder {
	nb.note = &ehrpb.Note{}
	nb.note.NoteGuid = GenerateGuidString()
	nb.note.Fragments = []*ehrpb.NoteFragment{}
	return nb
}


func (nb *NoteBuilder) SetId(id int32) *NoteBuilder {
	nb.note.Id = id
	return nb
}

func (nb *NoteBuilder) SetDateCreated(ts *timestamp.Timestamp) *NoteBuilder {
	nb.note.DateCreated = ts
	return nb
}

func (nb *NoteBuilder) SetVisitGuid(guid string) *NoteBuilder {
	nb.note.VisitGuid = guid
	return nb
}

func (nb *NoteBuilder) SetAuthorGuid(guid string) *NoteBuilder {
	nb.note.AuthorGuid = guid
	return nb
}

func (nb *NoteBuilder) SetPatientGuid(guid string) *NoteBuilder {
	nb.note.PatientGuid = guid
	return nb
}

func (nb *NoteBuilder) SetType(t ehrpb.NoteType) *NoteBuilder {
	nb.note.Type = t
	return nb
}

func (nb *NoteBuilder) Build() *ehrpb.Note {
	return nb.note
}