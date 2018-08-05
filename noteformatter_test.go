package noted

import (
	"testing"
	"github.com/geekmdio/noted/ehrproto"
	)

func TestNoteFormatter(t *testing.T) {
	nb := NoteBuilder{}
	note := nb.Init().Build()

	buildAppendNoteFrags(note)

	_ = NoteFormatter(note)

	firstTopic := note.Fragments[0].Topic
	expectedFirstTopic := ehrpb.FragmentTopic(1) // 0 is reserved so index starts at 1

	if firstTopic != expectedFirstTopic {
		t.Errorf("First topic should have been %v and was %v", firstTopic, expectedFirstTopic)
	}
}

func TestNoteFormatterError_Error(t *testing.T) {
	nb := NoteBuilder{}
	note := nb.Init().Build()

	err := NoteFormatter(note)

	if err == nil {
		t.Errorf("An error should have been thrown!")
	}

	nfe := noteFormatterError{ Message: "Test"}
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
	subjFrag := subjFragBuilder.SetTopic(ehrpb.FragmentTopic_SUBJECTIVE).Build()
	mhxFragBuilder := NoteFragmentBuilder{}
	mhxFragBuilder.InitFromNote(note)
	mhx := mhxFragBuilder.SetTopic(ehrpb.FragmentTopic_MEDICAL_HISTORY).Build()
	allFragBuilder := NoteFragmentBuilder{}
	allFragBuilder.InitFromNote(note)
	all := allFragBuilder.SetTopic(ehrpb.FragmentTopic_ALLERGIES).Build()
	medsFragBuilder := NoteFragmentBuilder{}
	medsFragBuilder.InitFromNote(note)
	meds := medsFragBuilder.SetTopic(ehrpb.FragmentTopic_MEDICATIONS).Build()
	famFragBuilder := NoteFragmentBuilder{}
	famFragBuilder.InitFromNote(note)
	fam := famFragBuilder.SetTopic(ehrpb.FragmentTopic_FAMILY_HISTORY).Build()
	socFragBuilder := NoteFragmentBuilder{}
	socFragBuilder.InitFromNote(note)
	soc := socFragBuilder.SetTopic(ehrpb.FragmentTopic_SOCIAL_HISTORY).Build()
	vitFragBuilder := NoteFragmentBuilder{}
	vitFragBuilder.InitFromNote(note)
	vit := vitFragBuilder.SetTopic(ehrpb.FragmentTopic_VITALS).Build()
	peFragBuilder := NoteFragmentBuilder{}
	peFragBuilder.InitFromNote(note)
	pe := peFragBuilder.SetTopic(ehrpb.FragmentTopic_PHYSICAL_EXAM).Build()
	asthmaFragBuilder := NoteFragmentBuilder{}
	asthmaFragBuilder.InitFromNote(note)
	asthma := asthmaFragBuilder.SetTopic(ehrpb.FragmentTopic_MEDICAL_PROBLEM).SetIssue(ehrpb.MedicalIssue_ASTHMA).Build()
	cadFragBuilder := NoteFragmentBuilder{}
	cadFragBuilder.InitFromNote(note)
	cad := cadFragBuilder.SetTopic(ehrpb.FragmentTopic_MEDICAL_PROBLEM).SetIssue(ehrpb.MedicalIssue_ISCHEMIC_CORONARY_ARTERY_DISEASE).Build()
	note.Fragments = append(note.Fragments, vit, all, meds, subjFrag, cad, asthma, mhx, fam, soc, pe)
}