package timesince

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type test struct {
	Delta    time.Duration
	Expected string
	Short    bool
}

var testMatrix = map[string]map[string]test{
	"Second": {
		"Plural": {
			Delta:    time.Duration(2.5 * float64(time.Second)),
			Expected: "2.5 seconds ago",
		},
		"Singular": {
			Delta:    time.Second,
			Expected: "1.0 second ago",
		},
		"Short": {
			Delta:    time.Second,
			Expected: "1.0s ago",
			Short:    true,
		},
	},
	"Minute": {
		"Plural": {
			Delta:    time.Duration(2.5 * float64(time.Minute)),
			Expected: "2.5 minutes ago",
		},
		"Singular": {
			Delta:    time.Minute,
			Expected: "1.0 minute ago",
		},
		"Short": {
			Delta:    time.Minute,
			Expected: "1.0m ago",
			Short:    true,
		},
	},
	"Hour": {
		"Plural": {
			Delta:    time.Duration(2.5 * float64(time.Hour)),
			Expected: "2.5 hours ago",
		},
		"Singular": {
			Delta:    time.Hour,
			Expected: "1.0 hour ago",
		},
		"Short": {
			Delta:    time.Hour,
			Expected: "1.0h ago",
			Short:    true,
		},
	},
	"Day": {
		"Plural": {
			Delta:    time.Duration(2.5 * 24 * float64(time.Hour)),
			Expected: "2.5 days ago",
		},
		"Singular": {
			Delta:    time.Duration(24 * float64(time.Hour)),
			Expected: "1.0 day ago",
		},
		"Short": {
			Delta:    time.Duration(24 * float64(time.Hour)),
			Expected: "1.0d ago",
			Short:    true,
		},
	},
}

func formatTimeDeltaTest(testCase test) func(*testing.T) {
	return func(t *testing.T) {
		comp := time.Now()
		ref := comp.Add(testCase.Delta) // ref is delta ahead of comp
		actual := FormatTimeDelta(ref, comp, testCase.Short)
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestFormatTimeDelta(t *testing.T) {
	for format, m := range testMatrix {
		t.Run(format, func(t *testing.T) {
			for plurality, test := range m {
				t.Run(plurality, formatTimeDeltaTest(test))
			}
		})
	}
}
