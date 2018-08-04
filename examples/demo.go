package main

import (
	"fmt"
	"github.com/geekmdio/noted/pkg/create"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/geekmdio/noted/pkg/ehrproto"
	"github.com/golang/protobuf/proto"
	"log"
	"github.com/geekmdio/noted/pkg/uuidHelper"
)

func main() {

	noteBuilder := create.NoteBuilder{}
	note := noteBuilder.
		Init().
		SetAuthorGuid(uuidHelper.GenerateGuidString()).
		SetDateCreated(&timestamp.Timestamp{}).
		SetId(0).
		SetPatientGuid(uuidHelper.GenerateGuidString()).
		SetType(ehrpb.NoteType_CONTINUED_CARE_DOCUMENTATION).
		SetVisitGuid(uuidHelper.GenerateGuidString()).
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
	fmt.Println("Printing the protobuf encoded byte stream.")
	fmt.Println(bytes)

	myReceivedNote := ehrpb.Note{}
	proto.Unmarshal(bytes, &myReceivedNote)
	fmt.Println("Printing the decoded note")
	fmt.Println(myReceivedNote)
}