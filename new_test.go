package noted

import "testing"

func TestNewNote(t *testing.T) {
	n := NewNote()
	if n.Fragments == nil {
		t.Fatal("Fragments are nil, and should be initialized to a slice of NoteFragment.")
	}
	if n.Tags == nil {
		t.Fatal("Tags are nil, and should be initialized to a slice of string.")
	}
	if n.DateCreated == nil {
		t.Fatal("DateCreated is nil, and should be initialized to timestamp of now.")
	}
}

func TestNewNoteFragment(t *testing.T) {
	nf := NewNoteFragment()
	if nf.Tags == nil {
		t.Fatal("Tags are nil, and should be initialized to a slice of string.")
	}
	if nf.DateCreated == nil {
		t.Fatal("DateCreated is nil, and should be initialized to a timestamp of now.")
	}
}