package create

import (
	"github.com/geekmdio/noted/src/ehrproto"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type NoteFragmentBuilder struct {
	noteFragment *ehrpb.NoteFragment
}

func (nb *NoteFragmentBuilder) InitFromNote(note *ehrpb.Note) *NoteFragmentBuilder {
	nb.noteFragment = &ehrpb.NoteFragment{}
	nb.noteFragment.VisitGuid = note.VisitGuid
	nb.noteFragment.NoteGuid = note.NoteGuid
	nb.noteFragment.AuthorGuid = note.AuthorGuid
	nb.noteFragment.PatientGuid = note.PatientGuid
	nb.noteFragment.NoteFragmentGuid = GenerateGuidString()
	return nb
}

func (nb *NoteFragmentBuilder) SetId(id int32) *NoteFragmentBuilder {
	nb.noteFragment.Id = id
	return nb
}

func (nb *NoteFragmentBuilder) SetDateCreated(ts *timestamp.Timestamp) *NoteFragmentBuilder {
	nb.noteFragment.DateCreated = ts
	return nb
}

func (nb *NoteFragmentBuilder) SetIssue(t ehrpb.MedicalIssue) *NoteFragmentBuilder {
	nb.noteFragment.Issue = t
	return nb
}

func (nb *NoteFragmentBuilder) SetIcd10Code(c string) *NoteFragmentBuilder {
	nb.noteFragment.Icd_10Code = c
	return nb
}

func (nb *NoteFragmentBuilder) SetStatus(s ehrpb.NoteFragmentStatus) *NoteFragmentBuilder {
	nb.noteFragment.Status = s
	return nb
}

func (nb *NoteFragmentBuilder) SetPriority(p ehrpb.FragmentPriority) *NoteFragmentBuilder {
	nb.noteFragment.Priority = p
	return nb
}

func (nb *NoteFragmentBuilder) SetTopic(t ehrpb.FragmentTopic) *NoteFragmentBuilder {
	nb.noteFragment.Topic = t
	return nb
}

func (nb *NoteFragmentBuilder) SetMarkdownContent(c string) *NoteFragmentBuilder {
	nb.noteFragment.MarkdownContent = c
	return nb
}

func (nb *NoteFragmentBuilder) Build() *ehrpb.NoteFragment {
	return nb.noteFragment
}