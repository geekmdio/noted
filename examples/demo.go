package main

import (
	"fmt"
	"github.com/geekmdio/ehrprotorepo/v1/generated/goproto"
	"github.com/geekmdio/noted"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"log"
	"os"
)

// This is a demo of how to use the GeekMD Noted library to
// build and up format a note.

func main() {
	note := buildNote()

	buildNoteFragments(note)

	fmtErr := noted.OrganizeNoteFragments(note)
	if fmtErr != nil {
		log.Fatalf("Error formatting the note: %f", fmtErr)
	}

	bytes, marshalErr := proto.Marshal(note)
	if marshalErr != nil {
		log.Fatalf("Error creating binary from note: %v", marshalErr)
	}

	fmt.Println("Printing the protobuf encoded byte stream.")
	fmt.Println(bytes)

	myReceivedNote := &ehrpb.Note{}
	proto.Unmarshal(bytes, myReceivedNote)

	fmt.Println("Printing JSON-ified data.")
	m := &jsonpb.Marshaler{}
	m.Marshal(os.Stdout, myReceivedNote)
}

func buildNote() *ehrpb.Note {
	noteBuilder := noted.NoteBuilder{}
	note := noteBuilder.
		Init().
		SetAuthorGuid(uuid.New().String()).
		SetDateCreated(&timestamp.Timestamp{}).
		SetId(0).
		SetPatientGuid(uuid.New().String()).
		SetType(ehrpb.NoteType_CONTINUED_CARE_DOCUMENTATION).
		SetVisitGuid(uuid.New().String()).
		Build()
	return note
}

func buildNoteFragments(note *ehrpb.Note) {
	fragBuilder := noted.NoteFragmentBuilder{}
	asthmaFrag := fragBuilder.
		InitFromNote(note).
		SetId(0).
		SetDateCreated(&timestamp.Timestamp{}).
		SetIssueGuid(uuid.New().String()).
		SetIcd10Code("J45.902").
		SetStatus(ehrpb.RecordStatus_INCOMPLETE).
		SetPriority(ehrpb.RecordPriority_HIGH).
		SetTopic(ehrpb.FragmentType_MEDICAL_PROBLEM).
		SetMarkdownContent("Asthma: ...").
		Build()
	cadFrag := fragBuilder.
		InitFromNote(note).
		SetId(1).
		SetDateCreated(&timestamp.Timestamp{}).
		SetIssueGuid(uuid.New().String()).
		SetIcd10Code("I25.9").
		SetStatus(ehrpb.RecordStatus_INCOMPLETE).
		SetPriority(ehrpb.RecordPriority_HIGH).
		SetTopic(ehrpb.FragmentType_MEDICAL_PROBLEM).
		SetMarkdownContent("Coronary arterial disease: ...").
		Build()
	subjectiveFrag := fragBuilder.
		InitFromNote(note).
		SetId(0).
		SetDateCreated(&timestamp.Timestamp{}).
		SetIssueGuid(uuid.New().String()).
		SetIcd10Code("").
		SetStatus(ehrpb.RecordStatus_INCOMPLETE).
		SetPriority(ehrpb.RecordPriority_HIGH).
		SetTopic(ehrpb.FragmentType_SUBJECTIVE).
		SetMarkdownContent("The reason for the visit today is...").
		Build()
	allergiesFrag := fragBuilder.
		InitFromNote(note).
		SetId(0).
		SetDateCreated(&timestamp.Timestamp{}).
		SetIssueGuid(uuid.New().String()).
		SetIcd10Code("").
		SetStatus(ehrpb.RecordStatus_INCOMPLETE).
		SetPriority(ehrpb.RecordPriority_HIGH).
		SetTopic(ehrpb.FragmentType_ALLERGIES).
		SetMarkdownContent("Penicillin -> Rash").
		Build()
	vitalsFrag := fragBuilder.
		InitFromNote(note).
		SetId(0).
		SetDateCreated(&timestamp.Timestamp{}).
		SetIssueGuid(uuid.New().String()).
		SetIcd10Code("").
		SetStatus(ehrpb.RecordStatus_INCOMPLETE).
		SetPriority(ehrpb.RecordPriority_HIGH).
		SetTopic(ehrpb.FragmentType_VITALS).
		SetMarkdownContent("120/83 mmHg").
		Build()
	socialFrag := fragBuilder.
		InitFromNote(note).
		SetId(0).
		SetDateCreated(&timestamp.Timestamp{}).
		SetIssueGuid(uuid.New().String()).
		SetIcd10Code("").
		SetStatus(ehrpb.RecordStatus_INCOMPLETE).
		SetPriority(ehrpb.RecordPriority_HIGH).
		SetTopic(ehrpb.FragmentType_SOCIAL_HISTORY).
		SetMarkdownContent("Smokes, drinks, does drugs...").
		Build()
	familyFrag := fragBuilder.
		InitFromNote(note).
		SetId(0).
		SetDateCreated(&timestamp.Timestamp{}).
		SetIssueGuid(uuid.New().String()).
		SetIcd10Code("").
		SetStatus(ehrpb.RecordStatus_INCOMPLETE).
		SetPriority(ehrpb.RecordPriority_HIGH).
		SetTopic(ehrpb.FragmentType_FAMILY_HISTORY).
		SetMarkdownContent("Dad had 4V CABG at age 43").
		Build()
	medsFrag := fragBuilder.
		InitFromNote(note).
		SetId(0).
		SetDateCreated(&timestamp.Timestamp{}).
		SetIssueGuid(uuid.New().String()).
		SetIcd10Code("").
		SetStatus(ehrpb.RecordStatus_INCOMPLETE).
		SetPriority(ehrpb.RecordPriority_HIGH).
		SetTopic(ehrpb.FragmentType_MEDICATIONS).
		SetMarkdownContent("Amphetamine salts").
		Build()
	peFrag := fragBuilder.
		InitFromNote(note).
		SetId(0).
		SetDateCreated(&timestamp.Timestamp{}).
		SetIssueGuid(uuid.New().String()).
		SetIcd10Code("").
		SetStatus(ehrpb.RecordStatus_INCOMPLETE).
		SetPriority(ehrpb.RecordPriority_HIGH).
		SetTopic(ehrpb.FragmentType_PHYSICAL_EXAM).
		SetMarkdownContent("Normal exam").
		Build()
	note.Fragments = append(note.Fragments, asthmaFrag, cadFrag, subjectiveFrag, allergiesFrag, vitalsFrag, socialFrag, familyFrag, medsFrag, peFrag)
}
