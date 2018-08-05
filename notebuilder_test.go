package noted

import (
	"testing"
	"github.com/google/uuid"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
	"github.com/geekmdio/noted/ehrproto"
	"log"
)

func TestNoteBuilderInitSetsNewGuid(t *testing.T) {
	b := NoteBuilder{}
	note := b.Init().Build()

	_, err := uuid.Parse(note.NoteGuid)

	if err != nil {
		t.Errorf("Could not parse note guid and it was therefore likely not set properly: %v", err)
	}
}

func TestNoteBuilderInitCreatesNewNoteFragmentSliceWithZeroLength(t *testing.T) {
	b := NoteBuilder{}
	note := b.Init().Build()

	if note.Fragments == nil || len(note.Fragments) != 0 {
		t.Errorf("The Note Fragments were not properly initialized.")
	}
}

func TestNoteBuilderSetIdSetsProperFieldAndReturnsProperValue(t *testing.T) {
	b := NoteBuilder{}
	var id int32 = 43
	note := b.Init().SetId(id).Build()

	if note.GetId() != id {
		t.Errorf("Expecting %v and recieved %v", id, note.GetId())
	}
}

func TestSetDateCreatedSetsTimeStampAndReturnsProperValue(t *testing.T) {
	b := NoteBuilder{}
	now := timestamp.Timestamp{Seconds:time.Now().Unix()}
	note := b.Init().SetDateCreated(&now).Build()

	// Test made in Aug 2018, therefore the value for seconds since UNIX epoch
	// must be greater than January 2018, which is about 40 years.
	// 40 yr after UNIX epoch * 365 day/yr * 24 hrs/day * 60 min/hr * 60 sec/min
	var jan2018 int64 = 1261440000
	if note.GetDateCreated().Seconds <= jan2018 {
		t.Errorf("Time for date created is less than or equal to zero; it was not properly set.")
	}
}

func TestSetXGuidSetsProperFields(t *testing.T) {
	newUuid, err := uuid.NewUUID()
	if err != nil {
		log.Fatalf("Error making UUID: %v", err)
	}
	visitGuid := newUuid
	newUuid, err = uuid.NewUUID()
	if err != nil {
		log.Fatalf("Error making UUID: %v", err)
	}
	authorGuid := newUuid
	newUuid, err = uuid.NewUUID()
	if err != nil {
		log.Fatalf("Error making UUID: %v", err)
	}
	patientGuid := newUuid

	b := NoteBuilder{}
	note := b.Init().
		SetVisitGuid(visitGuid.String()).
		SetAuthorGuid(authorGuid.String()).
		SetPatientGuid(patientGuid.String()).
		Build()

	fieldsSetProperly := note.GetVisitGuid() == visitGuid.String() &&
						note.GetAuthorGuid() == authorGuid.String() &&
						note.GetPatientGuid() == patientGuid.String()

	if !fieldsSetProperly {
		t.Errorf("One or more of the fields containing Guid's were improperly set.")
	}
}

func TestSetTypeSetsTypeProperly(t *testing.T) {
	b := NoteBuilder{}
	note := b.Init().
		SetType(ehrpb.NoteType_CONTINUED_CARE_DOCUMENTATION).
		Build()

	if note.GetType() != ehrpb.NoteType_CONTINUED_CARE_DOCUMENTATION {
		t.Errorf("The note types do not match up and there is likely an error.")
	}
}