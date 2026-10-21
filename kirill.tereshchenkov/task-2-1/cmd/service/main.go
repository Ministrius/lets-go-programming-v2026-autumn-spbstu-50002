package main

import (
	"fmt"
)

func main() {
	var n, k, temperature int
	var Symbol string

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
			if _, err := fmt.Scan(&Symbol, &temperature); err != nil {
				return
			}

			switch Symbol {
			case "<=":
				if temperature < maxT {
					maxT = temperature
				}
			case ">=":
				if temperature > minT {
					minT = temperature
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
