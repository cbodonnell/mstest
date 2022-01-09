package mstest

import "testing"

func TestSolution(t *testing.T) {
	testCases := [][]string{
		{"A2Le", "2pL1"},
		{"a10", "10a"},
		{"ba1", "1Ad"},
		{"3x2x", "8"},
		{"a3b", "5"},
		{"a33b", "35"},
		{"0", ""},
		{"", "0"},
		{"", ""},
		{"1", "0"},
	}
	expected := []bool{
		true,
		true,
		false,
		false,
		true,
		true,
		true,
		true,
		true,
		false,
	}

	for i, testCase := range testCases {
		result := Solution(testCase[0], testCase[1])
		if result != expected[i] {
			t.Fatalf(
				"result was %v for inputs %s and %s, expected %v",
				result,
				testCase[0],
				testCase[1],
				expected[i],
			)
		}
	}
}
