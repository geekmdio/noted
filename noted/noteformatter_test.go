package noted

import (
	"testing"
	"github.com/geekmdio/noted/ehrproto"
)

func TestNoteFormatter(t *testing.T) {
	nb := NoteBuilder{}
	nfb := NoteFragmentBuilder{}

	note := nb.Init().Build()
	subjFrag := nfb.SetTopic(ehrpb.FragmentTopic_SUBJECTIVE).Build()
	mhx := nfb.SetTopic(ehrpb.FragmentTopic_MEDICAL_HISTORY).Build()
	all := nfb.SetTopic(ehrpb.FragmentTopic_ALLERGIES).Build()
	meds := nfb.SetTopic(ehrpb.FragmentTopic_MEDICATIONS).Build()
	fam := nfb.SetTopic(ehrpb.FragmentTopic_FAMILY_HISTORY).Build()
	soc := nfb.SetTopic(ehrpb.FragmentTopic_SOCIAL_HISTORY).Build()
	vit := nfb.SetTopic(ehrpb.FragmentTopic_VITALS).Build()
	pe := nfb.SetTopic(ehrpb.FragmentTopic_PHYSICAL_EXAM).Build()
	asthma := nfb.SetTopic(ehrpb.FragmentTopic_MEDICAL_PROBLEM).SetIssue(ehrpb.MedicalIssue_ASTHMA).Build()
	cad := nfb.SetTopic(ehrpb.FragmentTopic_MEDICAL_PROBLEM).SetIssue(ehrpb.MedicalIssue_ISCHEMIC_CORONARY_ARTERY_DISEASE).Build()

	note.Fragments = append(note.Fragments, vit, all, meds, subjFrag, cad, asthma, mhx, fam, soc, pe)

	NoteFormatter(note)
}

//type FragmentTopic int32
//
//const (
//	FragmentTopic_NO_TOPIC        FragmentTopic = 0
//	FragmentTopic_SUBJECTIVE      FragmentTopic = 1
//	FragmentTopic_MEDICAL_HISTORY FragmentTopic = 2
//	FragmentTopic_ALLERGIES       FragmentTopic = 3
//	FragmentTopic_MEDICATIONS     FragmentTopic = 4
//	FragmentTopic_FAMILY_HISTORY  FragmentTopic = 5
//	FragmentTopic_SOCIAL_HISTORY  FragmentTopic = 6
//	FragmentTopic_VITALS          FragmentTopic = 7
//	FragmentTopic_PHYSICAL_EXAM   FragmentTopic = 8
//	FragmentTopic_MEDICAL_PROBLEM FragmentTopic = 9
//)