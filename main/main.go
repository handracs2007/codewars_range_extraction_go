package main

import (
	"bytes"
	"fmt"
)

func Solution(list []int) string {
	var buffer bytes.Buffer
	var queue = make([]int, 0)
	var prev = -1
	var first = true

	for _, val := range list {
		if first {
			// For first item
			buffer.WriteString(fmt.Sprintf("%v", val))
			first = false
		} else if (val - prev) == 1 {
			// Put to queue
			queue = append(queue, val)
		} else {
			if len(queue) == 1 {
				buffer.WriteString(fmt.Sprintf(",%v", prev))
				queue = queue[:0]
			} else if len(queue) > 1 {
				buffer.WriteString(fmt.Sprintf("-%v", prev))
				queue = queue[:0]
			}

			buffer.WriteString(fmt.Sprintf(",%v", val))
		}

		prev = val
	}

	if len(queue) > 0 {
		if len(queue) == 1 {
			buffer.WriteString(fmt.Sprintf(",%v", prev))
			queue = queue[:0]
		} else if len(queue) > 1 {
			buffer.WriteString(fmt.Sprintf("-%v", prev))
			queue = queue[:0]
		}
	}

	return buffer.String()
}

func main() {
	fmt.Println(Solution([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}))
}
