package examples

import (
	"fmt"
	"github.com/geekmdio/noted/pkg/create"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/geekmdio/noted/third_party/ehrproto"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	fmt.Println("This will be a demo soon...")
	noteBuilder := create.NoteBuilder{}
	note := noteBuilder.
		Init().
		SetAuthorGuid(create.GenerateGuidString()).
		SetDateCreated(&timestamp.Timestamp{}).
		SetId(0).
		SetPatientGuid(create.GenerateGuidString()).
		SetType(ehrpb.NoteType_CONTINUED_CARE_DOCUMENTATION).
		SetVisitGuid(create.GenerateGuidString()).
		Build()

	fragBuilder := create.NoteFragmentBuilder{}
	asthmaFrag := fragBuilder.
		InitFromNote(note).
		SetId(0).
		SetDateCreated(&timestamp.Timestamp{}).
		SetIssue(ehrpb.MedicalIssue_ASTHMA).
		SetIcd10Code("J45.902").
		SetStatus(ehrpb.NoteFragmentStatus_INCOMPLETE).
		SetPriority(ehrpb.FragmentPriority_HIGH).
		SetTopic(ehrpb.FragmentTopic_MEDICAL_PROBLEM).
		SetMarkdownContent("#Asthma\nSub-title\n##Heading2\nSub-title").
		Build()

	cadFrag := fragBuilder.
		InitFromNote(note).
		SetId(0).
		SetDateCreated(&timestamp.Timestamp{}).
		SetIssue(ehrpb.MedicalIssue_ISCHEMIC_CORONARY_ARTERY_DISEASE).
		SetIcd10Code("I25.9").
		SetStatus(ehrpb.NoteFragmentStatus_INCOMPLETE).
		SetPriority(ehrpb.FragmentPriority_HIGH).
		SetTopic(ehrpb.FragmentTopic_MEDICAL_PROBLEM).
		SetMarkdownContent("#Coronary Artery Disease\nSub-title\n##Heading2\nSub-title").
		Build()

	note.Fragments = append(note.Fragments, asthmaFrag, cadFrag)

	bytes, err := proto.Marshal(note)
	if err != nil {
		log.Fatalf("Error creating binary from note: %v", err)
	}
	fmt.Println(bytes)

	myReceivedNote := ehrpb.Note{}
	proto.Unmarshal(bytes, &myReceivedNote)
	fmt.Println(myReceivedNote)
}