package main

import (
	"fmt"
)

func main() {
	mat := "(0123)"
	fmt.Printf("result %+v\n", ambiguousCoordinates(mat))
}

func ambiguousCoordinates(s string) []string {
	w := s[1 : len(s)-1]
	res := []string{}
	fmt.Println(w[:1] + "." + w[1:2])
	for i := 1; i < len(w); i++ {
		lefts := coordinate(w, 0, i)
		rights := coordinate(w, i, len(w))

		for _, left := range lefts {
			for _, right := range rights {
				val := "(" + left + ", " + right + ")"
				res = append(res, val)
			}
		}
	}
	return res
}

func coordinate(s string, start, end int) []string {
	res := []string{}
	for p := 1; p <= end-start; p++ {
		w := s[start : start+p]
		d := s[start+p : end]
		if (w == "0" || w[0] != '0') && (len(d) == 0 || d[len(d)-1] != '0') {
			if len(d) == 0 {
				res = append(res, w)
			} else {
				res = append(res, w+"."+d)
			}
		}
	}

	return res
}
