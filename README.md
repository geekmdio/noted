# GeekMD:Noted v0.2.2
A library for creating and managing clinical notes.

|Branch|Build Status|Coverage %|
|:---:|:---:|:---:|
|Master| ![master](https://travis-ci.org/geekmdio/noted.svg?branch=master)|[![coverage-master](https://codecov.io/gh/geekmdio/noted/branch/master/graph/badge.svg)](https://codecov.io/gh/geekmdio/noted)|
|Development|![development](https://travis-ci.org/geekmdio/noted.svg?branch=development)|[![coverage-development](https://codecov.io/gh/geekmdio/noted/branch/development/graph/badge.svg)](https://codecov.io/gh/geekmdio/noted)|


PACKAGE DOCUMENTATION

package noted
    import "github.com/geekmdio/noted"

    Noted is a simple, no-nonsense solution to creating structured clinical
    notes. It breaks notes down into their essential elements, making them
    more easily searchable and customizable for the user. It is built upon
    Google's protobuf technology, and therefore is a ready-to-go solution
    for anything from simple, plain-text REST solutions to binary,
    bidirectional streaming RPC's.

    The process is simple. While you can access the relevant protobuf
    generated classes directly, there is a much easier to use fluent builder
    syntax for constructing notes.

    Step 1: Build a note with a NoteBuilder Step 2: Build one or more
    NoteFragment's with a NoteFragmentBuilder Step 3: Organize the Note into
    a structure consistent with what you would see in most clinical notes by
    running the NoteFormatter on the note Step 4: This part is up to you.
    Your object is ready to go.

FUNCTIONS

func NoteFormatter(n *ehrpb.Note) error

    NoteFormatter is required to resolve the inherent problem that comes
    with NoteFragment objects. NoteFragment's are stored in a Note as a
    slice of NoteFragment pointers, and the entire body of a note is
    included in these NoteFragments. As such, there is nothing stopping
    people from putting together notes in wildly different arrangements. A
    medical note generally follows the following format: - Subjective
    Information - Medical History - Medical Allergies - Family (Genetic)
    History - Social History - Vital Signs - Physical Exam - Medical
    Problems, sorted by priority along with a plan. NoteFormatter uses the
    enumerated values of the FragmentTopic type, which are sorted in order
    to achieve this structure.

TYPES

type NoteBuilder struct {
    // contains filtered or unexported fields
}

    NoteBuilder allows for a fluent means of constructing complex objects.
    Additionally, it handles initialization of a few features such as
    creation of the Note object, the Fragments pointer byte slice, and
    generating a GUID for the note. The NoteBuilder must be initialized,
    then the resulting object can call Init first, Build last, and any other
    functions in between to build up the object.

func (nb *NoteBuilder) Build() *ehrpb.Note

    Returns a Note object, indicating that the build process is complete.
    RETURNS: *Note

func (nb *NoteBuilder) Init() *NoteBuilder

    Initializes the NoteBuilder by creating a Note instance, setting the
    Notes GUID, and initializing a NoteFragment byte slice instance in the
    note. This must be called first after instantiation. RETURNS:
    *NoteBuilder

func (nb *NoteBuilder) SetAuthorGuid(guid string) *NoteBuilder

    Associates this object with an author by author GUID RETURNS:
    *NoteBuilder

func (nb *NoteBuilder) SetDateCreated(ts *timestamp.Timestamp) *NoteBuilder

    Set's the date created for the object, using the timestamp format.
    RETURNS: *NoteBuilder

func (nb *NoteBuilder) SetId(id int32) *NoteBuilder

    Set's the Id for the object. RETURNS: *NoteBuilder

func (nb *NoteBuilder) SetPatientGuid(guid string) *NoteBuilder

    Associates this object with a patient by patient GUID RETURNS:
    *NoteBuilder

func (nb *NoteBuilder) SetType(t ehrpb.NoteType) *NoteBuilder

    Sets the Note type. A note can be: - Follow up - Phone call
    documentation - Continued care documentation (e.g. documenting refills,
    or new results not associated with a visit) - Intake note - Procedure
    note - History and physical note (often for new or annual visits)
    RETURNS: *NoteBuilder

func (nb *NoteBuilder) SetVisitGuid(guid string) *NoteBuilder

    Associates this object with a visit by visit GUID RETURNS: *NoteBuilder

type NoteFragmentBuilder struct {
    // contains filtered or unexported fields
}

    NoteFragmentBuilder allows for a fluent means of constructing complex
    objects. Additionally, it handles initialization of a few features such
    as creation of the NoteFragment object.

func (nb *NoteFragmentBuilder) Build() *ehrpb.NoteFragment

    Returns a NoteFragment object, indicating that the build process is
    complete. RETURNS: *NoteFragment

func (nb *NoteFragmentBuilder) InitFromNote(note *ehrpb.Note) *NoteFragmentBuilder

    InitFromNote requires a Note when initializing the fluent build syntax
    because the purpose of the NoteFragment is to be a child to the note. It
    initializes visit GUID, note GUID, author GUID, patient GUID from the
    Note and also generates a GUID for the NoteFragment itself. RETURNS:
    *NoteFragmentBuilder

func (nb *NoteFragmentBuilder) SetDateCreated(ts *timestamp.Timestamp) *NoteFragmentBuilder

    Set the date the note was created using the timestamp format. RETURNS:
    *NoteFragmentBuilder

func (nb *NoteFragmentBuilder) SetIcd10Code(c string) *NoteFragmentBuilder

    Set the ICD10 code for the issue. If the fragment is note associated
    with a medical issue leave blank. RETURNS: *NoteFragmentBuilder

func (nb *NoteFragmentBuilder) SetId(id int32) *NoteFragmentBuilder

    Set the Id of the note. RETURNS: *NoteFragmentBuilder

func (nb *NoteFragmentBuilder) SetIssue(t ehrpb.MedicalIssue) *NoteFragmentBuilder

    Set the relevant medical issue for the NoteFragment. Some NoteFragments
    may not have an associated medical issue because the content of the note
    will be related to gathering or recording information, rather than
    formulating a response to a medical issue. Set the issue as
    NO_MEDICAL_ISSUE in this case. The medical issue can be one of numerous
    items. RETURNS: *NoteFragmentBuilder

func (nb *NoteFragmentBuilder) SetMarkdownContent(c string) *NoteFragmentBuilder

    The SetMarkdownContent field can support plain text, of course, but the
    intent is ultimately that this will be serving web applications or
    related applications which can easily support formatting through
    markdown. RETURNS: *NoteFragmentBuilder

func (nb *NoteFragmentBuilder) SetPriority(p ehrpb.FragmentPriority) *NoteFragmentBuilder

    Set the priority of a Fragment to allow filtering. - High - Normal - Low
    RETURNS: *NoteFragmentBuilder

func (nb *NoteFragmentBuilder) SetStatus(s ehrpb.NoteFragmentStatus) *NoteFragmentBuilder

    Set the status of the note. Of note, medical documentation should never
    be deleted. To support the archiving, rather than deleting,
    functionality we have three different statuses. - Incomplete - Active -
    Replaced (Old notes should be set to replaced, allowing the user to
    filter) RETURNS: *NoteFragmentBuilder

func (nb *NoteFragmentBuilder) SetTopic(t ehrpb.FragmentTopic) *NoteFragmentBuilder

    A NoteFragment will have a topic which allows for structuring of the
    Note, currently through NoteFormatter. The following fragment types
    exist and, through NoteFormatter, will be ordered in the below order. -
    Subjective Information - Medical History - Medical Allergies - Family
    (Genetic) History - Social History - Vital Signs - Physical Exam -
    Medical Problems, sorted by priority along with a plan. RETURNS:
    *NoteFragmentBuilder

SUBDIRECTORIES

        ehrproto
        examples
