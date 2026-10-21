package main

import (
	"fmt"
)

func main() {
	var n, k, t int
	var tp string

	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	for range n {
		if _, err := fmt.Scan(&k); err != nil {
			return
		}

		minT := 15
		maxT := 30

		for range k {
			if _, err := fmt.Scan(&tp, &t); err != nil {
				return
			}

			switch tp {
			case "<=":
				if t < maxT {
					maxT = t
				}
			case ">=":
				if t > minT {
					minT = t
				}
			default:
				fmt.Print("Wrong symbol\n")
			}

			if maxT >= minT {
				fmt.Println(minT)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
