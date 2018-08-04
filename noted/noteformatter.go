package noted

import (
	"github.com/geekmdio/noted/ehrproto"
	"sort"
)

type NoteFormatterError struct {
	Message string
}

func (n NoteFormatterError) Error() string {
	return n.Message
}

func NoteFormatter(n *ehrpb.Note) error {

	if len(n.Fragments) <= 0 {
		return NoteFormatterError{ Message: "The note does not have any fragments to sort."}
	}
	var topicVals []int
	for t := range n.Fragments {
		topicVals = append(topicVals, int(n.Fragments[t].Topic))
	}
	sort.Ints(topicVals)

	var newFragments []*ehrpb.NoteFragment
	for v := range topicVals {
		newFragments = append(newFragments, n.Fragments[v])
	}

	n.Fragments = newFragments
	return nil
}
