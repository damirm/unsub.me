package subscription

import (
	"testing"
	"time"
)

type testCase struct {
	s Subscription
	f Filter
	r bool
}

func runTestCases(t *testing.T, testCases []testCase) {
	for _, tc := range testCases {
		r := tc.f.passes(tc.s)
		if r != tc.r {
			t.Errorf("Invalid filter passes: tc=%v, r=%v ", tc, r)
		}
	}
}

func TestFilterLastActivityUntil(t *testing.T) {
	makeTestCase := func(la, until time.Time, result bool) testCase {
		return testCase{
			Subscription{LastActivity: la},
			Filter{LastActivityUntil: until},
			result,
		}
	}

	testCases := []testCase{
		makeTestCase(time.Now(), time.Now().Add(-1*time.Hour), false),
		makeTestCase(time.Now().Add(-1*time.Hour), time.Now(), true),
	}

	runTestCases(t, testCases)
}

func TestFilterName(t *testing.T) {
	makeTestCase := func(sName, fName string, result bool) testCase {
		return testCase{
			Subscription{Name: sName},
			Filter{Name: fName},
			result,
		}
	}

	testCases := []testCase{
		makeTestCase("", "123", false),
		makeTestCase("a", "b", false),
		makeTestCase("abc", "b", true),
		makeTestCase("abc", "ba", false),
		makeTestCase("ABC", "abc", true),
		makeTestCase("Hello World", "hello wo", true),
	}

	runTestCases(t, testCases)
}
