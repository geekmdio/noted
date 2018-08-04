package create

import (
	"testing"
	"github.com/geekmdio/noted/ehrproto"
)

func TestInitFromNoteReturnsObject(t *testing.T) {
	nb := NoteBuilder{}
	nfb := NoteFragmentBuilder{}
	n := nb.Init().Build()
	nf := nfb.InitFromNote(n).Build()

	// A note has an automatically generated Guid and the object will therefore
	// have a guid length of greater than 0 if it's successfully instantiated.
	if nf == nil || len(nf.GetNoteGuid()) <= 0 {
		t.Errorf("Object was not created")
	}
}

func TestSetIdProperlySetsId(t *testing.T) {
	var id int32 = 3
	nfb := NoteFragmentBuilder{}
	nf := nfb.InitFromNote(&ehrpb.Note{}).SetId(id).Build()

	if nf.GetId() != id {
		t.Errorf("Expected %v, but got %v", id, nf.GetId())
	}
}

func TestSetIssueProperlySetsIssue(t *testing.T) {
	issue := ehrpb.MedicalIssue_ASTHMA
	nfb := NoteFragmentBuilder{}
	nf := nfb.InitFromNote(&ehrpb.Note{}).SetIssue(issue).Build()

	if nf.GetIssue() != issue {
		t.Errorf("Expected %v, but got %v", issue, nf.GetIssue())
	}
}

func TestSetIcd10CodeProperlySetsIcd10Code(t *testing.T) {
	icd10Code := "Z.011"
	nfb := NoteFragmentBuilder{}
	nf := nfb.InitFromNote(&ehrpb.Note{}).SetIcd10Code(icd10Code).Build()

	if nf.GetIcd_10Code() != icd10Code {
		t.Errorf("Expected %v, but got %v", icd10Code, nf.GetIcd_10Code())
	}
}

func TestSetStatusProperlySetsStatus(t *testing.T) {
	status := ehrpb.NoteFragmentStatus_REPLACED
	nfb := NoteFragmentBuilder{}
	nf := nfb.InitFromNote(&ehrpb.Note{}).SetStatus(status).Build()

	if nf.GetStatus() != status {
		t.Errorf("Expected %v, but got %v", status, nf.GetStatus())
	}
}

func TestSetPriorityProperlySetsPriority(t *testing.T) {
	priority := ehrpb.FragmentPriority_HIGH
	nfb := NoteFragmentBuilder{}
	nf := nfb.InitFromNote(&ehrpb.Note{}).SetPriority(priority).Build()

	if nf.GetPriority() != priority {
		t.Errorf("Expected %v, but got %v", priority, nf.GetPriority())
	}
}

func TestSetTopicProperlySetsTopic(t *testing.T) {
	topic := ehrpb.FragmentTopic_ALLERGIES
	nfb := NoteFragmentBuilder{}
	nf := nfb.InitFromNote(&ehrpb.Note{}).SetTopic(topic).Build()

	if nf.GetTopic() != topic {
		t.Errorf("Expected %v, but got %v", topic, nf.GetTopic())
	}
}

func TestSetMarkdownContentProperlySetsMarkdownContent(t *testing.T) {
	mdContent := "# Heading1"
	nfb := NoteFragmentBuilder{}
	nf := nfb.InitFromNote(&ehrpb.Note{}).SetMarkdownContent(mdContent).Build()

	if nf.GetMarkdownContent() != mdContent {
		t.Errorf("Expected %v, but got %v", mdContent, nf.GetMarkdownContent())
	}
}