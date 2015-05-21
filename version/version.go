package version

import (
	"errors"
	"fmt"
	"strings"
)

type versionType int16

const (
	unknown         versionType = iota
	semVer          versionType = iota
	fourPartNumeric versionType = iota
)

// The IncrementField type identifies which field is to be incremented by the
// Increment function
type IncrementField int16

const (
	// Major field
	Major IncrementField = iota
	// Minor field
	Minor IncrementField = iota
	// Patch field
	Patch IncrementField = iota
	// Revision field (four part only)
	Revision IncrementField = iota
)

// Version managed a SemVer or four part version number
type Version struct {
	versionType versionType
	major       int32
	minor       int32
	patch       int32
	revision    int32
	label       string
}

// NewSemVer creates a new SemVer Version instance
func NewSemVer() *Version {
	return &Version{
		versionType: semVer,
		major:       0,
		minor:       0,
		patch:       0,
		label:       "",
	}
}

// New4Part creates a new four part Version instance
func New4Part() *Version {
	return &Version{
		versionType: fourPartNumeric,
		major:       0,
		minor:       0,
		patch:       0,
		revision:    0,
	}
}

// Parse creates a Version from the provided version string. It will detect the
// relevant version type from the version string.
func Parse(versionString string) (*Version, error) {
	if separators := strings.Count(versionString, "."); separators < 2 || separators > 3 {
		return nil, errors.New("Version only supports 3 part SemVer and 4 part versions")
	}

	var parts [4]int32
	_, err := fmt.Sscanf(versionString, "%d.%d.%d.%d", &parts[0], &parts[1], &parts[2], &parts[3])
	if err == nil {
		return &Version{
				versionType: fourPartNumeric,
				major:       parts[0],
				minor:       parts[1],
				patch:       parts[2],
				revision:    parts[3],
			},
			nil
	}

	var label string
	c, err := fmt.Sscanf(versionString, "%d.%d.%d-%s", &parts[0], &parts[1], &parts[2], &label)

	if c != 3 && err != nil {
		return nil, errors.New("Version only supports 3 part SemVer and 4 part versions")
	}

	return &Version{
			versionType: semVer,
			major:       parts[0],
			minor:       parts[1],
			patch:       parts[2],
			label:       label,
		},
		nil
}

// String formats the version for human consumption
func (v *Version) String() string {
	switch v.versionType {
	case semVer:
		if v.label == "" {
			return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)
		}
		return fmt.Sprintf("%d.%d.%d-%v", v.major, v.minor, v.patch, v.label)
	case fourPartNumeric:
		return fmt.Sprintf("%d.%d.%d.%d", v.major, v.minor, v.patch, v.revision)
	}
	return ""
}

// Increment increments the specified field, reseting the lesser fields as
// appropriate. It returns a new instance.
func (v *Version) Increment(f IncrementField) (*Version, error) {
	switch f {
	case Major:
		return &Version{
				versionType: v.versionType,
				major:       v.major + 1,
				minor:       0,
				patch:       0,
				revision:    0,
				label:       "",
			},
			nil
	case Minor:
		return &Version{
				versionType: v.versionType,
				major:       v.major,
				minor:       v.minor + 1,
				patch:       0,
				revision:    0,
				label:       "",
			},
			nil
	case Patch:
		return &Version{
				versionType: v.versionType,
				major:       v.major,
				minor:       v.minor,
				patch:       v.patch + 1,
				revision:    0,
				label:       "",
			},
			nil
	case Revision:
		if v.versionType == semVer {
			return nil, errors.New("Cannot increment a SemVer label as the meaning is unclear.")
		}
		return &Version{
				versionType: v.versionType,
				major:       v.major,
				minor:       v.minor,
				patch:       v.patch,
				revision:    v.revision + 1,
			},
			nil
	}

	return nil, fmt.Errorf("Increment field type %v is unknown", f)
}

// SetLabel assigns the optional label field to a SemVer version.
func (v *Version) SetLabel(label string) (*Version, error) {
	if v.versionType == fourPartNumeric {
		return nil, errors.New("Cannot set label on a four part version.")
	}

	return &Version{
			versionType: semVer,
			major:       v.major,
			minor:       v.minor,
			patch:       v.patch,
			label:       label,
		},
		nil
}
