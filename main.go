package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

func main() {
	data, _ := ioutil.ReadAll(os.Stdin)
	fmt.Print(data)
}
