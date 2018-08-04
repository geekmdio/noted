package noted

import (
	"testing"
	"github.com/geekmdio/noted/ehrproto"
	"fmt"
)

func TestNoteFormatter(t *testing.T) {
	nb := NoteBuilder{}
	note := nb.Init().Build()

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
	cad := subjFragBuilder.SetTopic(ehrpb.FragmentTopic_MEDICAL_PROBLEM).SetIssue(ehrpb.MedicalIssue_ISCHEMIC_CORONARY_ARTERY_DISEASE).Build()

	note.Fragments = append(note.Fragments, vit, all, meds, subjFrag, cad, asthma, mhx, fam, soc, pe)

	NoteFormatter(note)

	fmt.Println(note)
}
