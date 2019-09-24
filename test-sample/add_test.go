package add

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		I      int
		J      int
		Expect int
	}{
		// tc1
		{
			I: 2, J: 3,
			Expect: 5,
		},
		// tc2
		{
			I: 3, J: 4,
			Expect: 7,
		},
	}

	for _, tc := range cases {
		actual := Add(tc.I, tc.J)
		if actual != tc.Expect {
			t.Errorf("expect: %d, actual: %d", tc.Expect, actual)
		}
	}
}
