package pack

import (
	"testing"
)

type testData struct {
	list []string
	want string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{ // create a slice of test data
		testData{list: []string{}, want: ""},
		testData{list: []string{"apple"}, want: "apple"},                                     // test for one element
		testData{list: []string{"apple", "orange"}, want: "apple and orange"},                // test for two elements
		testData{list: []string{"apple", "orange", "pear"}, want: "apple, orange, and pear"}, // test for three elements
	}
	for _, test := range tests { // Process each testData value in the slice
		got := JoinWithCommas(test.list) // Pass the slice JoinWithCommas.
		if got != test.want {            // If the return value we got doesn't eaqual the value we want...
			t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", test.list, got, test.want)

		}
	}

}
