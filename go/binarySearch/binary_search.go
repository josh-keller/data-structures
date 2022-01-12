package binary_search

func BinarySearch(arr []int, target int) int {
  start := 0
  end := len(arr) - 1

  for start <= end {
    idx := (end - start) / 2 + start
    if arr[idx] == target {
      return idx
    } else if target < arr[idx] {
      end = idx - 1
    } else {
      start = idx + 1
    }
  }

  return -1
}
