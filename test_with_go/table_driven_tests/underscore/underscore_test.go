package underscore

import "testing"

func TestToCamelCase(t *testing.T) {
	testCases := map[string]struct {
		arg  string
		want string
	}{
		"pure camel case":         {"thisShouldReturn", "this-should-return"},
		"string with space":       {"with a space", "with a space"},
		"uppercase letter at end": {"returnA", "return-a"},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := ToCamelCase(testCase.arg)

			if got != testCase.want {
				t.Fatalf("got %q, wanted %q", got, testCase.want)
			}
		})
	}
}
