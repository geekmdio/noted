package noted

import (
	"testing"
	"github.com/geekmdio/ehrprotorepo/goproto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
	"github.com/google/uuid"
)

func TestNoteBuilder_Init(t *testing.T) {
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

func TestNoteFragmentBuilder_SetId(t *testing.T) {
	var id int32 = 3
	nfb := NoteFragmentBuilder{}
	nf := nfb.InitFromNote(&ehrpb.Note{}).SetId(id).Build()

	if nf.GetId() != id {
		t.Errorf("Expected %v, but got %v", id, nf.GetId())
	}
}

func TestNoteFragmentBuilder_SetIssue(t *testing.T) {
	issueGuid := uuid.New().String()
	nfb := NoteFragmentBuilder{}
	nf := nfb.InitFromNote(&ehrpb.Note{}).SetIssueGuid(issueGuid).Build()

	if nf.GetIssueGuid() != issueGuid {
		t.Errorf("Expected %v, but got %v", issueGuid, nf.GetIssueGuid())
	}
}

func TestNoteFragmentBuilder_SetIcd10Code(t *testing.T) {
	icd10Code := "Z.011"
	nfb := NoteFragmentBuilder{}
	nf := nfb.InitFromNote(&ehrpb.Note{}).SetIcd10Code(icd10Code).Build()

	if nf.GetIcd_10Code() != icd10Code {
		t.Errorf("Expected %v, but got %v", icd10Code, nf.GetIcd_10Code())
	}
}

func TestNoteFragmentBuilder_SetIcd10LongDescription(t *testing.T) {
	icd10Long := "Diabetes mellitus, type 2"
	nfb := NoteFragmentBuilder{}
	nf := nfb.InitFromNote(&ehrpb.Note{}).SetIcd10LongDescription(icd10Long).Build()

	if nf.GetIcd_10Long() != icd10Long {
		t.Errorf("Expected %v, but got %v", icd10Long, nf.GetIcd_10Long())
	}
}

func TestNoteFragmentBuilder_SetDescription(t *testing.T) {
	desc := "Diabetes Type 2"
	nfb := NoteFragmentBuilder{}
	nf := nfb.InitFromNote(&ehrpb.Note{}).SetDescription(desc).Build()

	if nf.GetDescription() != desc {
		t.Errorf("Expected %v, but got %v", desc, nf.GetDescription())
	}
}

func TestNoteFragmentBuilder_SetStatus(t *testing.T) {
	status := ehrpb.NoteFragmentStatus_REPLACED
	nfb := NoteFragmentBuilder{}
	nf := nfb.InitFromNote(&ehrpb.Note{}).SetStatus(status).Build()

	if nf.GetStatus() != status {
		t.Errorf("Expected %v, but got %v", status, nf.GetStatus())
	}
}

func TestNoteFragmentBuilder_SetPriority(t *testing.T) {
	priority := ehrpb.FragmentPriority_HIGH
	nfb := NoteFragmentBuilder{}
	nf := nfb.InitFromNote(&ehrpb.Note{}).SetPriority(priority).Build()

	if nf.GetPriority() != priority {
		t.Errorf("Expected %v, but got %v", priority, nf.GetPriority())
	}
}

func TestNoteFragmentBuilder_SetTopic(t *testing.T) {
	topic := ehrpb.FragmentTopic_ALLERGIES
	nfb := NoteFragmentBuilder{}
	nf := nfb.InitFromNote(&ehrpb.Note{}).SetTopic(topic).Build()

	if nf.GetTopic() != topic {
		t.Errorf("Expected %v, but got %v", topic, nf.GetTopic())
	}
}

func TestNoteFragmentBuilder_SetMarkdownContent(t *testing.T) {
	mdContent := "# Heading1"
	nfb := NoteFragmentBuilder{}
	nf := nfb.InitFromNote(&ehrpb.Note{}).SetMarkdownContent(mdContent).Build()

	if nf.GetMarkdownContent() != mdContent {
		t.Errorf("Expected %v, but got %v", mdContent, nf.GetMarkdownContent())
	}
}

func TestNoteFragmentBuilder_SetDateCreated(t *testing.T) {
	dateCreated := timestamp.Timestamp{ Seconds: time.Now().Unix() }
	nfb := NoteFragmentBuilder{}
	nf := nfb.InitFromNote(&ehrpb.Note{}).SetDateCreated(&dateCreated).Build()

	if nf.GetDateCreated().Seconds != dateCreated.Seconds {
		t.Errorf("Expected %v, but got %v", dateCreated, nf.GetTopic())
	}
}