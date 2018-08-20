package noted

import (
	"github.com/geekmdio/ehrprotorepo/v1/generated/goproto"
	"github.com/google/uuid"
	"testing"
)

func TestOrganizeNoteFragments(t *testing.T) {
	nb := NoteBuilder{}
	note := nb.Init().Build()

	buildAppendNoteFrags(note)

	_ = OrganizeNoteFragments(note)

	firstTopic := note.Fragments[0].Topic
	expectedFirstTopic := ehrpb.FragmentType(1) // 0 is reserved so index starts at 1

	if firstTopic != expectedFirstTopic {
		t.Errorf("First topic should have been %v and was %v", firstTopic, expectedFirstTopic)
	}
}

func TestOrganizeNoteFragmentsError_Error(t *testing.T) {
	nb := NoteBuilder{}
	note := nb.Init().Build()

	err := OrganizeNoteFragments(note)

	if err == nil {
		t.Errorf("An error should have been thrown!")
	}

	nfe := noteFormatterError{Message: "Test"}
	_, ok := error(nfe).(error)
	if !ok {
		t.Errorf("Should have type asserted to error interface")
	}

	if nfe.Error() != nfe.Message {
		t.Error("Error should be returning message")
	}

}

func buildAppendNoteFrags(note *ehrpb.Note) {
	subjFragBuilder := NoteFragmentBuilder{}
	subjFragBuilder.InitFromNote(note)
	subjFrag := subjFragBuilder.SetTopic(ehrpb.FragmentType_SUBJECTIVE).Build()
	mhxFragBuilder := NoteFragmentBuilder{}
	mhxFragBuilder.InitFromNote(note)
	mhx := mhxFragBuilder.SetTopic(ehrpb.FragmentType_MEDICAL_HISTORY).Build()
	allFragBuilder := NoteFragmentBuilder{}
	allFragBuilder.InitFromNote(note)
	all := allFragBuilder.SetTopic(ehrpb.FragmentType_ALLERGIES).Build()
	medsFragBuilder := NoteFragmentBuilder{}
	medsFragBuilder.InitFromNote(note)
	meds := medsFragBuilder.SetTopic(ehrpb.FragmentType_MEDICATIONS).Build()
	famFragBuilder := NoteFragmentBuilder{}
	famFragBuilder.InitFromNote(note)
	fam := famFragBuilder.SetTopic(ehrpb.FragmentType_FAMILY_HISTORY).Build()
	socFragBuilder := NoteFragmentBuilder{}
	socFragBuilder.InitFromNote(note)
	soc := socFragBuilder.SetTopic(ehrpb.FragmentType_SOCIAL_HISTORY).Build()
	vitFragBuilder := NoteFragmentBuilder{}
	vitFragBuilder.InitFromNote(note)
	vit := vitFragBuilder.SetTopic(ehrpb.FragmentType_VITALS).Build()
	peFragBuilder := NoteFragmentBuilder{}
	peFragBuilder.InitFromNote(note)
	pe := peFragBuilder.SetTopic(ehrpb.FragmentType_PHYSICAL_EXAM).Build()
	asthmaFragBuilder := NoteFragmentBuilder{}
	asthmaFragBuilder.InitFromNote(note)
	asthma := asthmaFragBuilder.SetTopic(ehrpb.FragmentType_PHYSICAL_EXAM).SetIssueGuid(uuid.New().String()).Build()
	cadFragBuilder := NoteFragmentBuilder{}
	cadFragBuilder.InitFromNote(note)
	cad := cadFragBuilder.SetTopic(ehrpb.FragmentType_MEDICAL_PROBLEM).SetIssueGuid(uuid.New().String()).Build()
	note.Fragments = append(note.Fragments, vit, all, meds, subjFrag, cad, asthma, mhx, fam, soc, pe)
}
