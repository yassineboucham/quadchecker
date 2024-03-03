package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	progs := []string{"./quadA", "./quadB", "./quadC", "./quadD", "./quadE"}
	validfunc := []string{}
	data, _ := ioutil.ReadAll(os.Stdin) 

	width := 0
	height := 0
	countWidth := true

	for _, v := range data {

		if v == 10 {
			height++
			countWidth = false
		} else if width != 10 && countWidth {
			width++
		}

	}

	for i := 0; i < len(progs); i++ {
		cmd := exec.Command(progs[i], strconv.Itoa(width), strconv.Itoa(height))
		b, _ := cmd.CombinedOutput()
 

		if string(data) == string(b) { 
			validfunc = append(validfunc, string(progs[i]))
		}
	}
	fmt.Print(validfunc)
}
