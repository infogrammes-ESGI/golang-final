// Coded by TD, SD & KD the 10/10/22
// Memory & CPU info
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pid struct {
	name string
	number int
	cwd string
	exe string
	memory int
}

func getProcess() []int {

	dir, err : = os.Open("/proc")
	if err != nil {
		panic(err)
	}
	
	defer dir.Close()

	
	onlydir := int[]

	for {
		all, err = dir.Readdirnames(0) // list all ( 0 ) files in the directory.

		for _, name := range all {
			fmt.Println(name)
			onlydir = append(onlydir, name)
		}
		

	}
}


func pid(){
	
	//fmt.Println(pid_max())

//	file, err := os.Open("/proc/")
}


func main(){
	fmt.Println(pid())
}

