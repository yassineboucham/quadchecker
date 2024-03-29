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

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 1; i <= y; i++ {
		if i == 1 {
			z01.PrintRune('A')
		} else if i == y {
			z01.PrintRune('C')
		} else {
			z01.PrintRune('B')
		}
		if x != 1 {
			for j := 1; j <= x-2; j++ {
				if i == 1 || i == y {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
			if i == 1 {
				z01.PrintRune('C')
			} else if i == y {
				z01.PrintRune('A')
			} else {
				z01.PrintRune('B')
			}
		}
		z01.PrintRune('\n')
	}
}
func main() {
	args := os.Args
	QuadE(atoi(args[1]), atoi(args[2]))
}
