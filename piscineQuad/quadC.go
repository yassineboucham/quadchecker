package main

import (
	"os"

	"github.com/01-edu/z01"
)

func atoi(n string) int {
	nb := 0
	signe := 1
	if n[0] == '-' {
		signe = -1
	}
	for i := 0; i < len(n); i++ {
		if n[i] >= '0' && n[i] <= '9' {
			nb = (nb * 10) + int(n[i]-'0')
		}
	}
	return nb * signe
}

func QuadC(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 {
				if j == 1 || j == x {
					z01.PrintRune('A')
				} else {
					z01.PrintRune('B')
				}
			} else if i == y {
				if j == 1 || j == x {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			} else {
				if j == 1 || j == x {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
func main() {
	args := os.Args
	QuadC(atoi(args[1]), atoi(args[2]))
}
