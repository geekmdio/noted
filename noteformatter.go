package noted

import (
	"github.com/geekmdio/noted/ehrproto"
	"sort"
	)

type noteFormatterError struct {
	Message string
}

func (n noteFormatterError) Error() string {
	return n.Message
}

// NoteFormatter is required to resolve the inherent problem that comes with NoteFragment
// objects. NoteFragment's are stored in a Note as a slice of NoteFragment pointers,
// and the entire body of a note is included in these NoteFragments. As such, there is
// nothing stopping people from putting together notes in wildly different arrangements.
// A medical note generally follows the following format:
// - Subjective Information
// - Medical History
// - Medical Allergies
// - Family (Genetic) History
// - Social History
// - Vital Signs
// - Physical Exam
// - Medical Problems, sorted by priority along with a plan.
// NoteFormatter uses the enumerated values of the FragmentTopic type, which are sorted
// in order to achieve this structure.
func NoteFormatter(n *ehrpb.Note) error {

	if len(n.Fragments) <= 0 {
		return noteFormatterError{ Message: "The note does not have any fragments to sort."}
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
