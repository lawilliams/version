package version

import (
	"reflect"
	"testing"
)

var validincrement = []struct {
	initial  string
	field    IncrementField
	expected string
}{
	{"10.2.3.4", Major, "11.0.0.0"},
	{"1.234.3.4", Minor, "1.235.0.0"},
	{"1.2.5939.4", Patch, "1.2.5940.0"},
	{"1.2.3.4345", Revision, "1.2.3.4346"},
	{"47.2.3", Major, "48.0.0"},
	{"47.2.3-alpha", Major, "48.0.0"},
	{"1.92.3", Minor, "1.93.0"},
	{"1.92.3-alpha", Minor, "1.93.0"},
	{"1.2.493", Patch, "1.2.494"},
	{"1.2.493-alpha", Patch, "1.2.494"},
}

func TestValidIncrement(t *testing.T) {
	for _, vi := range validincrement {
		initial, _ := Parse(vi.initial)
		expected, _ := Parse(vi.expected)

		incremented, _ := initial.Increment(vi.field)

		if !reflect.DeepEqual(expected, incremented) {
			t.Errorf("Incremented %v expecting %+v but got %+v", vi.field, expected, incremented)
		}
	}
}

func TestValidIncrementDoesNotError(t *testing.T) {
	for _, vi := range validincrement {
		initial, _ := Parse(vi.initial)

		_, err := initial.Increment(vi.field)

		if err != nil {
			t.Errorf("Expected no error incrementing %v on %+v but got %v", vi.field, initial, err)
		}
	}
}

func TestInvalidIncrementField(t *testing.T) {
	v, _ := Parse("1.2.3.4")

	_, err := v.Increment(IncrementField(6))

	if err == nil {
		t.Fatal("Does not error for increment of invalid field type")
	}
}

func TestInvalidIncrementOfLabelOnSemVer(t *testing.T) {
	v, _ := Parse("1.2.3-alpha")

	_, err := v.Increment(Revision)

	if err == nil {
		t.Fatal("Should produce error when attempting to increment the label on a SemVer version")
	}
}
