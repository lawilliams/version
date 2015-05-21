package version

import (
	"reflect"
	"testing"
)

func TestNewSemVer(t *testing.T) {
	expected := &Version{
		versionType: semVer,
		major:       0,
		minor:       0,
		patch:       0,
		label:       "",
	}
	v := NewSemVer()

	if !reflect.DeepEqual(expected, v) {
		t.Errorf("Expected %+v, got %+v", expected, v)
	}
}

func TestNew4Part(t *testing.T) {
	expected := &Version{
		versionType: fourPartNumeric,
		major:       0,
		minor:       0,
		patch:       0,
		revision:    0,
	}
	v := New4Part()

	if !reflect.DeepEqual(expected, v) {
		t.Errorf("Expected %+v, got %+v", expected, v)
	}
}
