package noted

import (
	"github.com/geekmdio/ehrprotorepo/v1/generated/goproto"
	"github.com/golang/protobuf/ptypes/timestamp"
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
// a GUID for the NoteFragment itself. For the ICD10 codes, we highly recommend
// leveraging a service like the Watson API for natural language processing, which
// can be trained to predict ICD10 codes from Descriptions.
// RETURNS: *NoteFragmentBuilder
func (nb *NoteFragmentBuilder) InitFromNote(note *ehrpb.Note) *NoteFragmentBuilder {
	nb.noteFragment = NewNoteFragment()
	nb.noteFragment.NoteGuid = note.GetNoteGuid()
	return nb
}

// Set the Id of the note.
// RETURNS: *NoteFragmentBuilder
func (nb *NoteFragmentBuilder) SetId(id int64) *NoteFragmentBuilder {
	nb.noteFragment.Id = id
	return nb
}

// Set the date the note was created using the timestamp format.
// RETURNS: *NoteFragmentBuilder
func (nb *NoteFragmentBuilder) SetDateCreated(ts *timestamp.Timestamp) *NoteFragmentBuilder {
	nb.noteFragment.DateCreated = ts
	return nb
}

// Set the user- and clinician-friendly description of the problem. This should use
// sufficient wording such that predictive programs, such as the watson API,
// can return ICD codes with high confidence allowing the structuring of data.
// RETURNS: *NoteFragmentBuilder
func (nb *NoteFragmentBuilder) SetDescription(s string) *NoteFragmentBuilder {
	nb.noteFragment.Description = s
	return nb
}

// Set the relevant medical issue for the NoteFragment. Some NoteFragments may not have
// an associated medical issue because the content of the note will be related to
// gathering or recording information, rather than formulating a response to a medical issue.
// Set the issue as NO_MEDICAL_ISSUE in this case. The medical issue can be one of numerous items.
// RETURNS: *NoteFragmentBuilder
func (nb *NoteFragmentBuilder) SetIssueGuid(i string) *NoteFragmentBuilder {
	nb.noteFragment.IssueGuid = i
	return nb
}

// Set the ICD10 code for the issue. If the fragment is note associated with a medical issue
// leave blank.
// RETURNS: *NoteFragmentBuilder
func (nb *NoteFragmentBuilder) SetIcd10Code(c string) *NoteFragmentBuilder {
	nb.noteFragment.Icd_10Code = c
	return nb
}

// Set the ICD10 long description for the issue. If the fragment is note associated with a medical issue
// leave blank.
// RETURNS: *NoteFragmentBuilder
func (nb *NoteFragmentBuilder) SetIcd10LongDescription(c string) *NoteFragmentBuilder {
	nb.noteFragment.Icd_10Long = c
	return nb
}

// Set the status of the note. Of note, medical documentation should never be deleted.
// To support the archiving, rather than deleting, functionality we have three different
// statuses.
// - Incomplete
// - Active
// - Replaced (Old notes should be set to replaced, allowing the user to filter)
// RETURNS: *NoteFragmentBuilder
func (nb *NoteFragmentBuilder) SetStatus(s ehrpb.RecordStatus) *NoteFragmentBuilder {
	nb.noteFragment.Status = s
	return nb
}

// Set the priority of a Fragment to allow filtering.
// - High
// - Normal
// - Low
// RETURNS: *NoteFragmentBuilder
func (nb *NoteFragmentBuilder) SetPriority(p ehrpb.RecordPriority) *NoteFragmentBuilder {
	nb.noteFragment.Priority = p
	return nb
}

// A NoteFragment will have a topic which allows for structuring of the
// Note, currently through OrganizeNoteFragments. The following fragment types exist
// and, through OrganizeNoteFragments, will be ordered in the below order.
// - Subjective Information
// - Medical History
// - Medical Allergies
// - Family (Genetic) History
// - Social History
// - Vital Signs
// - Physical Exam
// - Medical Problems, sorted by priority along with a plan.
// RETURNS: *NoteFragmentBuilder
func (nb *NoteFragmentBuilder) SetTopic(t ehrpb.FragmentType) *NoteFragmentBuilder {
	nb.noteFragment.Topic = t
	return nb
}

// The SetContent field can support plain text, of course, but
// the intent is ultimately that this will be serving web applications or
// related applications which can easily support formatting through markdown.
// RETURNS: *NoteFragmentBuilder
func (nb *NoteFragmentBuilder) SetContent(c string) *NoteFragmentBuilder {
	nb.noteFragment.Content = c
	return nb
}

// Returns a NoteFragment object, indicating that the build process is
// complete.
// RETURNS: *NoteFragment
func (nb *NoteFragmentBuilder) Build() *ehrpb.NoteFragment {
	return nb.noteFragment
}
