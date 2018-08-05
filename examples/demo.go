package main

import (
	"fmt"
	"github.com/geekmdio/noted/noted"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/geekmdio/noted/ehrproto"
	"github.com/golang/protobuf/proto"
	"log"
	"github.com/geekmdio/noted/guidHelper"
)

// This is a demo of how to use the GeekMD Noted library to
// build and up format a note.
func main() {

	note := buildNote()

	buildNoteFragments(note)

	fmtErr := noted.NoteFormatter(note)
	if fmtErr != nil {
		log.Fatalf("Error formatting the note")
	}

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

func buildNote() *ehrpb.Note {
	noteBuilder := noted.NoteBuilder{}
	note := noteBuilder.
		Init().
		SetAuthorGuid(guidHelper.GenerateGuidString()).
		SetDateCreated(&timestamp.Timestamp{}).
		SetId(0).
		SetPatientGuid(guidHelper.GenerateGuidString()).
		SetType(ehrpb.NoteType_CONTINUED_CARE_DOCUMENTATION).
		SetVisitGuid(guidHelper.GenerateGuidString()).
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
		SetMarkdownContent("120/83 mmHg").
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
		SetMarkdownContent("120/83 mmHg").
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
		SetMarkdownContent("120/83 mmHg").
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
		SetMarkdownContent("120/83 mmHg").
		Build()
	note.Fragments = append(note.Fragments, asthmaFrag, cadFrag, subjectiveFrag, allergiesFrag, vitalsFrag, socialFrag, familyFrag, medsFrag, peFrag)
}