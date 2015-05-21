package version

import "testing"

var tostring = []string{
	"1.2.3",
	"1.2.3-alpha",
	"1.2.3.4",
}

func TestToString(t *testing.T) {
	for _, s := range tostring {
		v, _ := Parse(s)
		vs := v.String()
		if s != vs {
			t.Errorf("Version %v was not correctly formatted as string. Recieved %v", s, vs)
		}
	}
}
