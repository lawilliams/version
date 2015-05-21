package version

import "testing"

var labels = []struct {
	initial  string
	label    string
	expected string
}{
	{"1.2.3", "test", "1.2.3-test"},
	{"1.2.3-initial", "test", "1.2.3-test"},
}

func TestSetValidLabel(t *testing.T) {
	for _, l := range labels {
		v, _ := Parse(l.initial)

		updated, _ := v.SetLabel(l.label)

		versionAsString := updated.String()

		if versionAsString != l.expected {
			t.Errorf("Expected %v after assigning label %v but got %v", l.expected, l.label, versionAsString)
		}
	}
}

func TestSetValidLabelDoesNotReturnError(t *testing.T) {
	for _, l := range labels {
		v, _ := Parse(l.initial)

		_, err := v.SetLabel(l.label)

		if err != nil {
			t.Errorf("Received error %v when applying label %v to %v", err, l.label, l.initial)
		}
	}
}

func TestCannotSetLabelOn4Part(t *testing.T) {
	v, _ := Parse("1.2.3.4")

	_, err := v.SetLabel("test")

	if err == nil {
		t.Fatal("No error trying to set label on four part version")
	}
}
