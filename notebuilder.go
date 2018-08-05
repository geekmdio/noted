// Noted is a simple, no-nonsense solution to creating structured
// clinical notes. It breaks notes down into their essential elements,
// making them more easily searchable and customizable for the user.
// It is built upon Google's protobuf technology, and therefore
// is a ready-to-go solution for anything from simple, plain-text REST solutions
// to binary, bidirectional streaming RPC's.
//
// The process is simple. While you can access the relevant protobuf generated
// classes directly, there is a much easier to use fluent builder syntax
// for constructing notes.
//
// Step 1: Build a note with a NoteBuilder
// Step 2: Build one or more NoteFragment's with a NoteFragmentBuilder
// Step 3: Organize the Note into a structure consistent with what you
// would see in most clinical notes by running the NoteFormatter on the note
// Step 4: This part is up to you. Your object is ready to go.
package noted

import (
	"github.com/geekmdio/ehrprotorepo/goproto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	)


// NoteBuilder allows for a fluent means of constructing complex objects.
// Additionally, it handles initialization of a few features such as creation
// of the Note object, the Fragments pointer byte slice, and generating
// a GUID for the note. The NoteBuilder must be initialized, then the
// resulting object can call Init first, Build last, and any other
// functions in between to build up the object.
type NoteBuilder struct {
	note *ehrpb.Note
}

// Initializes the NoteBuilder by creating a Note instance, setting the
// Notes GUID, and initializing a NoteFragment byte slice instance
// in the note. This must be called first after instantiation.
// RETURNS: *NoteBuilder
func (nb *NoteBuilder) Init() *NoteBuilder {
	nb.note = &ehrpb.Note{}
	nb.note.Fragments = []*ehrpb.NoteFragment{}
	nb.note.NoteGuid = uuid.New().String()
	return nb
}

// Set's the Id for the object.
// RETURNS: *NoteBuilder
func (nb *NoteBuilder) SetId(id int32) *NoteBuilder {
	nb.note.Id = id
	return nb
}

// Set's the date created for the object, using the timestamp format.
// RETURNS: *NoteBuilder
func (nb *NoteBuilder) SetDateCreated(ts *timestamp.Timestamp) *NoteBuilder {
	nb.note.DateCreated = ts
	return nb
}

// Associates this object with a visit by visit GUID
// RETURNS: *NoteBuilder
func (nb *NoteBuilder) SetVisitGuid(guid string) *NoteBuilder {
	nb.note.VisitGuid = guid
	return nb
}

// Associates this object with an author by author GUID
// RETURNS: *NoteBuilder
func (nb *NoteBuilder) SetAuthorGuid(guid string) *NoteBuilder {
	nb.note.AuthorGuid = guid
	return nb
}

// Associates this object with a patient by patient GUID
// RETURNS: *NoteBuilder
func (nb *NoteBuilder) SetPatientGuid(guid string) *NoteBuilder {
	nb.note.PatientGuid = guid
	return nb
}

// Sets the Note type. A note can be:
// - Follow up
// - Phone call documentation
// - Continued care documentation (e.g. documenting refills, or
// new results not associated with a visit)
// - Intake note
// - Procedure note
// - History and physical note (often for new or annual visits)
// RETURNS: *NoteBuilder
func (nb *NoteBuilder) SetType(t ehrpb.NoteType) *NoteBuilder {
	nb.note.Type = t
	return nb
}

// Returns a Note object, indicating that the build process is
// complete.
// RETURNS: *Note
func (nb *NoteBuilder) Build() *ehrpb.Note {
	return nb.note
}