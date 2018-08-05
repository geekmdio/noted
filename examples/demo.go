package main

import (
		"github.com/geekmdio/noted"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/geekmdio/noted/ehrproto"
	"github.com/golang/protobuf/proto"
	"log"
	"github.com/google/uuid"
	"github.com/golang/protobuf/jsonpb"
	"os"
	"fmt"
)
// This is a demo of how to use the GeekMD Noted library to
// build and up format a note.

func main() {
	note := buildNote()

	buildNoteFragments(note)

	fmtErr := noted.NoteFormatter(note)
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
	authorUuid, err := uuid.NewUUID()
	if err != nil {
		log.Fatalf("Error making UUID: %v", err)
	}
	patientGuid, err := uuid.NewUUID()
	if err != nil {
		log.Fatalf("Error making UUID: %v", err)
	}
	visitUuid, err := uuid.NewUUID()
	if err != nil {
		log.Fatalf("Error making UUID: %v", err)
	}
	noteBuilder := noted.NoteBuilder{}
	note := noteBuilder.
		Init().
		SetAuthorGuid(authorUuid.String()).
		SetDateCreated(&timestamp.Timestamp{}).
		SetId(0).
		SetPatientGuid(patientGuid.String()).
		SetType(ehrpb.NoteType_CONTINUED_CARE_DOCUMENTATION).
		SetVisitGuid(visitUuid.String()).
		Build()
	return note
}

func buildNoteFragments(note *ehrpb.Note) {
	fragBuilder := noted.NoteFragmentBuilder{}
	asthmaFrag := fragBuilder.
		InitFromNote(note).
		SetId(0).
		SetDateCreated(&timestamp.Timestamp{}).
		SetIssue(ehrpb.MedicalIssue_ASTHMA).
		SetIcd10Code("J45.902").
		SetStatus(ehrpb.NoteFragmentStatus_INCOMPLETE).
		SetPriority(ehrpb.FragmentPriority_HIGH).
		SetTopic(ehrpb.FragmentTopic_MEDICAL_PROBLEM).
		SetMarkdownContent("Asthma: ...").
		Build()
	cadFrag := fragBuilder.
		InitFromNote(note).
		SetId(1).
		SetDateCreated(&timestamp.Timestamp{}).
		SetIssue(ehrpb.MedicalIssue_ISCHEMIC_CORONARY_ARTERY_DISEASE).
		SetIcd10Code("I25.9").
		SetStatus(ehrpb.NoteFragmentStatus_INCOMPLETE).
		SetPriority(ehrpb.FragmentPriority_HIGH).
		SetTopic(ehrpb.FragmentTopic_MEDICAL_PROBLEM).
		SetMarkdownContent("Coronary arterial disease: ...").
		Build()
	subjectiveFrag := fragBuilder.
		InitFromNote(note).
		SetId(0).
		SetDateCreated(&timestamp.Timestamp{}).
		SetIssue(ehrpb.MedicalIssue_NO_MEDICAL_ISSUE).
		SetIcd10Code("").
		SetStatus(ehrpb.NoteFragmentStatus_INCOMPLETE).
		SetPriority(ehrpb.FragmentPriority_HIGH).
		SetTopic(ehrpb.FragmentTopic_SUBJECTIVE).
		SetMarkdownContent("The reason for the visit today is...").
		Build()
	allergiesFrag := fragBuilder.
		InitFromNote(note).
		SetId(0).
		SetDateCreated(&timestamp.Timestamp{}).
		SetIssue(ehrpb.MedicalIssue_NO_MEDICAL_ISSUE).
		SetIcd10Code("").
		SetStatus(ehrpb.NoteFragmentStatus_INCOMPLETE).
		SetPriority(ehrpb.FragmentPriority_HIGH).
		SetTopic(ehrpb.FragmentTopic_ALLERGIES).
		SetMarkdownContent("Penicillin -> Rash").
		Build()
	vitalsFrag := fragBuilder.
		InitFromNote(note).
		SetId(0).
		SetDateCreated(&timestamp.Timestamp{}).
		SetIssue(ehrpb.MedicalIssue_NO_MEDICAL_ISSUE).
		SetIcd10Code("").
		SetStatus(ehrpb.NoteFragmentStatus_INCOMPLETE).
		SetPriority(ehrpb.FragmentPriority_HIGH).
		SetTopic(ehrpb.FragmentTopic_VITALS).
		SetMarkdownContent("120/83 mmHg").
		Build()
	socialFrag := fragBuilder.
		InitFromNote(note).
		SetId(0).
		SetDateCreated(&timestamp.Timestamp{}).
		SetIssue(ehrpb.MedicalIssue_NO_MEDICAL_ISSUE).
		SetIcd10Code("").
		SetStatus(ehrpb.NoteFragmentStatus_INCOMPLETE).
		SetPriority(ehrpb.FragmentPriority_HIGH).
		SetTopic(ehrpb.FragmentTopic_SOCIAL_HISTORY).
		SetMarkdownContent("Smokes, drinks, does drugs...").
		Build()
	familyFrag := fragBuilder.
		InitFromNote(note).
		SetId(0).
		SetDateCreated(&timestamp.Timestamp{}).
		SetIssue(ehrpb.MedicalIssue_NO_MEDICAL_ISSUE).
		SetIcd10Code("").
		SetStatus(ehrpb.NoteFragmentStatus_INCOMPLETE).
		SetPriority(ehrpb.FragmentPriority_HIGH).
		SetTopic(ehrpb.FragmentTopic_FAMILY_HISTORY).
		SetMarkdownContent("Dad had 4V CABG at age 43").
		Build()
	medsFrag := fragBuilder.
		InitFromNote(note).
		SetId(0).
		SetDateCreated(&timestamp.Timestamp{}).
		SetIssue(ehrpb.MedicalIssue_NO_MEDICAL_ISSUE).
		SetIcd10Code("").
		SetStatus(ehrpb.NoteFragmentStatus_INCOMPLETE).
		SetPriority(ehrpb.FragmentPriority_HIGH).
		SetTopic(ehrpb.FragmentTopic_MEDICATIONS).
		SetMarkdownContent("Amphetamine salts").
		Build()
	peFrag := fragBuilder.
		InitFromNote(note).
		SetId(0).
		SetDateCreated(&timestamp.Timestamp{}).
		SetIssue(ehrpb.MedicalIssue_NO_MEDICAL_ISSUE).
		SetIcd10Code("").
		SetStatus(ehrpb.NoteFragmentStatus_INCOMPLETE).
		SetPriority(ehrpb.FragmentPriority_HIGH).
		SetTopic(ehrpb.FragmentTopic_PHYSICAL_EXAM).
		SetMarkdownContent("Normal exam").
		Build()
	note.Fragments = append(note.Fragments, asthmaFrag, cadFrag, subjectiveFrag, allergiesFrag, vitalsFrag, socialFrag, familyFrag, medsFrag, peFrag)
}