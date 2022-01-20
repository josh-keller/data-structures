package llstack

import (
	"reflect"
	"testing"
)

// NOT IMPLEMENTED YET...JUST CUT AND PASTE...
func TestPop(t *testing.T) {
	testCases := []struct {
		desc       string
		actions    []string
		pushVals   []int
		expLength  int
		expTopVal  int
		expPopVals []int
	}{
		{
			desc:       "Pushing and popping, results in length 0 and returning that value",
			actions:    []string{"push", "pop"},
			pushVals:   []int{5},
			expLength:  0,
			expTopVal:  0,
			expPopVals: []int{5},
		},
		{
			desc:       "Pushing two and popping one results in length 1, popping second push and first push at top",
			actions:    []string{"push", "push", "pop"},
			pushVals:   []int{5, 6},
			expLength:  1,
			expTopVal:  5,
			expPopVals: []int{6},
		},
		{
			desc:       "Alternating pushing and popping yields correct length and popped vals",
			actions:    []string{"push", "push", "pop", "push", "pop", "pop"},
			pushVals:   []int{1, 2, 3},
			expLength:  0,
			expTopVal:  0,
			expPopVals: []int{2, 3, 1},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			s := NewStack()
			i := 0
			poppedVals := []int{}

			for _, action := range tC.actions {
				if action == "push" {
					val := tC.pushVals[i]
					i++
					s.Push(val)
				} else if action == "pop" {
					poppedVals = append(poppedVals, s.Pop())
				}
			}

			if length := s.Length(); length != tC.expLength {
				t.Errorf("Expected length of %d, got %d", tC.expLength, length)
			}

			if tC.expLength != 0 {
				if val := s.Peek(); val != tC.expTopVal {
					t.Errorf("Expected top value of %d, got %d", tC.expTopVal, val)
				}
			}

			if !reflect.DeepEqual(tC.expPopVals, poppedVals) {
				t.Errorf("Expected popped values of %v, got %b", tC.expPopVals, poppedVals)
			}
		})
	}
}
func TestPush(t *testing.T) {
	testCases := []struct {
		desc      string
		pushVals  []int
		expLength int
		expTopVal int
	}{
		{
			desc:      "Pushing one value onto empty stack puts value at head, length 1",
			pushVals:  []int{5},
			expLength: 1,
			expTopVal: 5,
		},
		{
			desc:      "Pushing two values onto empty stack leave second value pushed at head and length 2",
			pushVals:  []int{5, 6},
			expLength: 2,
			expTopVal: 6,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			s := NewStack()
			for _, val := range tC.pushVals {
				s.Push(val)
			}

			if length := s.Length(); length != tC.expLength {
				t.Errorf("Expected length of %d, got %d", tC.expLength, length)
			}

			if val := s.Peek(); val != tC.expTopVal {
				t.Errorf("Expected top value of %d, got %d", tC.expTopVal, val)
			}
		})
	}
}
