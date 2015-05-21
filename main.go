package main

import (
	"errors"
	"flag"
	"log"
	"os"

	"github.com/ColinScott/version/version"
)

type fileModifier func(f *os.File, v *version.Version) (*version.Version, error)

var initType, label string
var incrementParts [4]bool

func init() {
	flag.StringVar(&initType, "init", "", "Initialse the VERSION file. Valid values: semver, 4part.")
	flag.BoolVar(&incrementParts[version.Major], "major", false, "Increment the major version")
	flag.BoolVar(&incrementParts[version.Minor], "minor", false, "Increment the major version")
	flag.BoolVar(&incrementParts[version.Patch], "patch", false, "Increment the major version")
	flag.BoolVar(&incrementParts[version.Revision], "revision", false, "Increment the major version")
	flag.StringVar(&label, "label", "", "Sets the version of a label")
}

func main() {
	flag.Parse()

	validateFlags()

	if initType != "" {
		intialiseVersion()
		return
	}

	for i, p := range incrementParts {
		if p {
			incrementVersion(version.IncrementField(i))
			return
		}
	}

	if label != "" {
		setLabel()
		return
	}
	flag.PrintDefaults()
}

func validateFlags() {
	var count int
	if initType != "" {
		count++
	}
	for _, p := range incrementParts {
		if p {
			count++
		}
	}
	if label != "" {
		count++
	}
	if count > 1 {
		log.Fatal("Only one operation at a time is supported.")
	}
}

func intialiseVersion() {
	var v *version.Version
	switch initType {
	case "semver":
		v = version.NewSemVer()
	case "4part":
		v = version.New4Part()
	default:
		log.Fatal("Init only supports the types 'semver' and '4part'")
	}
	err := writeVersion(v)
	if err != nil {
		log.Fatal(err)
	}
}

func writeVersion(v *version.Version) error {
	if _, err := os.Stat("VERSION"); !os.IsNotExist(err) {
		return errors.New("Cannot initialise as VERSION file already exists.")
	}
	f, err := os.Create("VERSION")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(v.String())

	return err
}

func incrementVersion(i version.IncrementField) {
	updateVersionFile(func(f *os.File, v *version.Version) (*version.Version, error) {
		return v.Increment(i)
	})
}

func setLabel() {
	updateVersionFile(func(f *os.File, v *version.Version) (*version.Version, error) {
		return v.SetLabel(label)
	})
}

func updateVersionFile(m fileModifier) {
	f, err := os.OpenFile("VERSION", os.O_RDWR, 0660)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b := make([]byte, 16)

	_, err = f.Read(b)
	if err != nil {
		log.Fatal(err)
	}

	v, err := version.Parse(string(b))
	if err != nil {
		log.Fatal(err)
	}

	v, err = m(f, v)
	if err != nil {
		log.Fatal(err)
	}

	f.Seek(0, 0)
	err = f.Truncate(0)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.WriteString(v.String())
	if err != nil {
		log.Fatal(err)
	}
}
