package binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
  testCases := []struct {
    desc  string
    arr []int
    target int
    expected int
  }{
    {
      desc: "finds a number in the middle",
      arr: []int{1, 2, 4, 7, 9, 13, 27},
      target: 7,
      expected: 3,
    },
    {
      desc: "finds a number at the end",
      arr: []int{1, 2, 4, 7, 9, 13, 27},
      target: 27,
      expected: 6,
    },
    {
      desc: "finds a number at the beginning",
      arr: []int{1, 2, 4, 7, 9, 13, 27},
      target: 1,
      expected: 0,
    },
    {
      desc: "finds a number between the beginning and middle",
      arr: []int{1, 2, 4, 7, 9, 13, 27},
      target: 2,
      expected: 1,
    },
    {
      desc: "finds a number between the middle and the end",
      arr: []int{1, 2, 4, 7, 9, 13, 27},
      target: 13,
      expected: 5,
    },
    {
      desc: "correctly returns -1 when a number isn't present (too large)",
      arr: []int{1, 2, 4, 7, 9, 13, 27},
      target: 28,
      expected: -1,
    },
    {
      desc: "correctly returns -1 when a number isn't present (too small)",
      arr: []int{1, 2, 4, 7, 9, 13, 27},
      target: 0,
      expected: -1,
    },
    {
      desc: "correctly returns -1 when a number isn't present (not present)",
      arr: []int{1, 2, 4, 7, 9, 13, 27},
      target: 5,
      expected: -1,
    },
  }
  for _, tC := range testCases {
    t.Run(tC.desc, func(t *testing.T) {
      if output := BinarySearch(tC.arr, tC.target); output != tC.expected {
        t.Errorf("%q expected: %d, received: %d", tC.desc, tC.expected, output)
      }
    })
  }
}
