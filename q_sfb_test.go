package tc

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSfb(t *testing.T) {
	tests := map[string]struct {
		val  Sfb
		err1 error
		err2 error
	}{
		"empty":  {err1: fmt.Errorf("Sfb options are missing")},
		"simple": {val: Sfb{Parms: &SfbQopt{Max: 0xFF}}},
	}

	for name, testcase := range tests {
		t.Run(name, func(t *testing.T) {
			data, err1 := marshalSfb(&testcase.val)
			if err1 != nil {
				if testcase.err1 != nil && testcase.err1.Error() == err1.Error() {
					return
				}
				t.Fatalf("Unexpected error: %v", err1)
			}
			val := Sfb{}
			err2 := unmarshalSfb(data, &val)
			if err2 != nil {
				if testcase.err2 != nil && testcase.err2.Error() == err2.Error() {
					return
				}
				t.Fatalf("Unexpected error: %v", err2)

			}
			if diff := cmp.Diff(val, testcase.val); diff != "" {
				t.Fatalf("Sfb missmatch (want +got):\n%s", diff)
			}
		})
	}
}