package noted

import (
	"github.com/geekmdio/ehrprotorepo/v1/generated/goproto"
	"github.com/google/uuid"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

// Generate a new Note with essential elements of instantiation handled.
func NewNote() *ehrpb.Note {
	return &ehrpb.Note{
		Id:          0,
		DateCreated: TimestampNow(),
		NoteGuid:    uuid.New().String(),
		Fragments:   make([]*ehrpb.NoteFragment, 0),
		Tags:        make([]string, 0),
	}
}

// Generate a new NoteFragment with essential elements of instantiation handled.
func NewNoteFragment() *ehrpb.NoteFragment {
	return &ehrpb.NoteFragment{
		Id:               0,
		DateCreated:      TimestampNow(),
		NoteFragmentGuid: uuid.New().String(),
		IssueGuid:        "",
		Icd_10Code:       "",
		Icd_10Long:       "",
		Description:      "",
		Status:           ehrpb.RecordStatus_INCOMPLETE,
		Priority:         ehrpb.RecordPriority_NO_PRIORITY,
		Topic:            ehrpb.FragmentType_NO_TOPIC,
		Content:          "",
		Tags:             make([]string, 0),
	}
}

// Generate a timestamp for now.
func TimestampNow() *timestamp.Timestamp {
	now := time.Now()
	ts := &timestamp.Timestamp{
		Seconds: now.Unix(),
		Nanos:   int32(now.UnixNano()),
	}
	return ts
}