package problem_statement_5

import (
	"log"
	"testing"
)

type TestCase struct {
	arg1, arg2, expected int
}

var additionTestCases = []TestCase{
	TestCase{0, 0, 0},
	TestCase{7, 9, 16},
	TestCase{1, 2, 3},
	TestCase{79, 97, 176},
}

var substractionTestCases = []TestCase{
	TestCase{0, 0, 0},
	TestCase{7, 9, -2},
	TestCase{1, 2, -1},
	TestCase{79, 97, -18},
}

func TestAdd(t *testing.T) {
	// Unit test with single test cases
	// got := Add(7, 9)
	// want := 16
	// if got != want {
	//	 t.Errorf("error")
	// }

	// Table driven test
	for _, testCase := range additionTestCases {
		if got := Add(testCase.arg1, testCase.arg2); got != testCase.expected {
			t.Errorf("error")
		}
	}
}

func TestSubtract(t *testing.T) {
	// Unit test cases for single test cases
	//	got := Substract(7, 9)
	//	want := -2
	//	if got != want {
	//		t.Errorf("error")
	//		log.Println("error log")
	//		fmt.Println("error in fmt")
	//	}

	// Table driven test
	for _, testCase := range substractionTestCases {
		if got := Substract(testCase.arg1, testCase.arg2); got != testCase.expected {
			t.Errorf("error")
			log.Println(testCase.arg2)
		}
	}
}
