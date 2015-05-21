package version

import (
	"reflect"
	"testing"
)

var semvertests = []struct {
	formatted string
	major     int32
	minor     int32
	patch     int32
	label     string
}{
	{"1.2.3", 1, 2, 3, ""},
	{"1.2.3-alpha", 1, 2, 3, "alpha"},
}

func TestSemVerParse(t *testing.T) {
	for _, s := range semvertests {
		v, err := Parse(s.formatted)
		if err != nil {
			t.Errorf("Format %v produced error %v", s.formatted, err)
			continue
		}

		expected := &Version{
			versionType: semVer,
			major:       s.major,
			minor:       s.minor,
			patch:       s.patch,
			label:       s.label,
		}

		if !reflect.DeepEqual(expected, v) {
			t.Errorf("Expected %+v, got %+v", expected, v)
		}
	}
}

func Test4PartParse(t *testing.T) {
	v, err := Parse("1.2.3.4")
	if err != nil {
		t.Fatalf("Format 1.2.3.4 produced error %v", err)
	}

	expected := &Version{
		versionType: fourPartNumeric,
		major:       1,
		minor:       2,
		patch:       3,
		revision:    4,
	}

	if !reflect.DeepEqual(expected, v) {
		t.Fatalf("Expected %+v, got %+v", expected, v)
	}
}

var invalidformat = []string{
	"",
	"1.2",
	"1.2.3.4.5",
	"a.2.3",
	"1.b.3",
	"1.2.c",
	"..3",
	"1..",
	"1..3",
	"..",
	"...",
	"1..3.4",
	"1..3.alpha",
}

func TestInvalidFormats(t *testing.T) {
	for _, i := range invalidformat {
		v, err := Parse(i)
		if err == nil {
			t.Errorf("Expected %v to have parse error but it produced %+v", i, v)
		}
	}
}
