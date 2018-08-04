package main

import (
	"fmt"
	"github.com/geekmdio/noted/pkg/create"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/geekmdio/noted/pkg/ehrproto"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	fmt.Println("Building a note.")
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

	fmt.Println("Building a note fragment for asthma.")
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

	fmt.Println("Building a note fragment for ischemic CAD")
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

	fmt.Println("Appending the fragments to the note")
	note.Fragments = append(note.Fragments, asthmaFrag, cadFrag)

	fmt.Println("Marshaling the message into a protobuf byte stream.")
	bytes, err := proto.Marshal(note)
	if err != nil {
		log.Fatalf("Error creating binary from note: %v", err)
	}
	fmt.Println("Printing the protobuf byte stream.")
	fmt.Println(bytes)

	fmt.Println("Unmarshaling the byte stream")
	myReceivedNote := ehrpb.Note{}
	proto.Unmarshal(bytes, &myReceivedNote)
	fmt.Println("Printing the note")
	fmt.Println(myReceivedNote)
}