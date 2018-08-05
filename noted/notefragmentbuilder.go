package noted

import (
	"github.com/geekmdio/noted/ehrproto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/geekmdio/noted/uuidHelper"
)

// NoteFragmentBuilder allows for a fluent means of constructing complex objects.
// Additionally, it handles initialization of a few features such as creation
// of the NoteFragment object.
type NoteFragmentBuilder struct {
	noteFragment *ehrpb.NoteFragment
}

// InitFromNote requires a Note when initializing the fluent build syntax because
// the purpose of the NoteFragment is to be a child to the note. It initializes
// visit GUID, note GUID, author GUID, patient GUID from the Note and also generates
// a GUID for the NoteFragment itself.
// RETURNS: *NoteFragmentBuilder
func (nb *NoteFragmentBuilder) InitFromNote(note *ehrpb.Note) *NoteFragmentBuilder {
	nb.noteFragment = &ehrpb.NoteFragment{}
	nb.noteFragment.VisitGuid = note.VisitGuid
	nb.noteFragment.NoteGuid = note.NoteGuid
	nb.noteFragment.AuthorGuid = note.AuthorGuid
	nb.noteFragment.PatientGuid = note.PatientGuid
	nb.noteFragment.NoteFragmentGuid = uuidHelper.GenerateGuidString()
	return nb
}

// Set the Id of the note.
// RETURNS: *NoteFragmentBuilder
func (nb *NoteFragmentBuilder) SetId(id int32) *NoteFragmentBuilder {
	nb.noteFragment.Id = id
	return nb
}

// Set the date the note was created using the timestamp format.
// RETURNS: *NoteFragmentBuilder
func (nb *NoteFragmentBuilder) SetDateCreated(ts *timestamp.Timestamp) *NoteFragmentBuilder {
	nb.noteFragment.DateCreated = ts
	return nb
}

// Set the relevant medical issue for the NoteFragment. Some NoteFragments may not have
// an associated medical issue because the content of the note will be related to
// gathering or recording information, rather than formulating a response to a medical issue.
// Set the issue as NO_MEDICAL_ISSUE in this case. The medical issue can be one of numerous items.
// RETURNS: *NoteFragmentBuilder
func (nb *NoteFragmentBuilder) SetIssue(t ehrpb.MedicalIssue) *NoteFragmentBuilder {
	nb.noteFragment.Issue = t
	return nb
}
// Set the ICD10 code for the issue. If the fragment is note associated with a medical issue
// leave blank.
// RETURNS: *NoteFragmentBuilder
func (nb *NoteFragmentBuilder) SetIcd10Code(c string) *NoteFragmentBuilder {
	nb.noteFragment.Icd_10Code = c
	return nb
}

// Set the status of the note. Of note, medical documentation should never be deleted.
// To support the archiving, rather than deleting, functionality we have three different
// statuses.
// - Incomplete
// - Active
// - Replaced (Old notes should be set to replaced, allowing the user to filter)
// RETURNS: *NoteFragmentBuilder
func (nb *NoteFragmentBuilder) SetStatus(s ehrpb.NoteFragmentStatus) *NoteFragmentBuilder {
	nb.noteFragment.Status = s
	return nb
}

// Set the priority of a Fragment to allow filtering.
// - High
// - Normal
// - Low
// RETURNS: *NoteFragmentBuilder
func (nb *NoteFragmentBuilder) SetPriority(p ehrpb.FragmentPriority) *NoteFragmentBuilder {
	nb.noteFragment.Priority = p
	return nb
}

// A NoteFragment will have a topic which allows for structuring of the
// Note, currently through NoteFormatter. The following fragment types exist
// and, through NoteFormatter, will be ordered in the below order.
// - Subjective Information
// - Medical History
// - Medical Allergies
// - Family (Genetic) History
// - Social History
// - Vital Signs
// - Physical Exam
// - Medical Problems, sorted by priority along with a plan.
// RETURNS: *NoteFragmentBuilder
func (nb *NoteFragmentBuilder) SetTopic(t ehrpb.FragmentTopic) *NoteFragmentBuilder {
	nb.noteFragment.Topic = t
	return nb
}

// The SetMarkdownContent field can support plain text, of course, but
// the intent is ultimately that this will be serving web applications or
// related applications which can easily support formatting through markdown.
// RETURNS: *NoteFragmentBuilder
func (nb *NoteFragmentBuilder) SetMarkdownContent(c string) *NoteFragmentBuilder {
	nb.noteFragment.MarkdownContent = c
	return nb
}

// Returns a NoteFragment object, indicating that the build process is
// complete.
// RETURNS: *NoteFragment
func (nb *NoteFragmentBuilder) Build() *ehrpb.NoteFragment {
	return nb.noteFragment
}