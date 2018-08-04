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

	newFragments := generateNewFragmentsProperlyOrdered(n)

	n.Fragments = newFragments
	return nil
}

func generateNewFragmentsProperlyOrdered(n *ehrpb.Note) []*ehrpb.NoteFragment {
	topicVals := getSortedTopicIntegerValues(n)

	var newFragments []*ehrpb.NoteFragment
	for v := range topicVals {
		for f := range n.Fragments {
			if topicVals[v] == int(n.Fragments[f].Topic) {
				newFragments = append(newFragments, n.Fragments[f])
			}
		}
	}
	return newFragments
}

func getSortedTopicIntegerValues(n *ehrpb.Note) []int {
	var topicVals []int
	for t := range n.Fragments {
		topicVals = append(topicVals, int(n.Fragments[t].Topic))
	}
	sort.Ints(topicVals)
	return topicVals
}
