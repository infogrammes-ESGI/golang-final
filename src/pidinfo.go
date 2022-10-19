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
// gets max pid for the system
func pid_max() int{
	file, err := os.Open("/proc/sys/kernel/pid_max")
	// Control that the file was open with no error or exiting
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	fmt.Println("Max system pid : " + scanner.Text())
	valret, _ := strconv.Atoi(scanner.Text())
	return valret
}


func pid(){
	
	pid_max()
	//fmt.Println(pid_max())

//	file, err := os.Open("/proc/")
}


func main(){
//	fmt.Println(pid())
	pid()
}

// Function to open the memory file
func read_mem_info() {
	file, err := os.Open("/proc/meminfo")
	// Control that the file was open with no error or exiting
	if err != nil {
		fmt.Println("Could not open the memory file - Exiting - : ", err)
		os.Exit(1)
	}
	// Used to close the file at the end of main()
	defer file.Close()

	// Define a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	// Scan all the line and retrieve our targets
	array := [13]string{"MemTotal", "MemFree", "MemAvailable", "Buffers", "Cached", "SwapCached", "Active", "Inactive", "SwapTotal", "SwapFree", "Dirty", "Writeback", "Shmem"}
	i := 0
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), array[i]) && i < len(array) {
			fmt.Println(scanner.Text())
			i++
			if i == len(array) {
				break
			}
		}
	}
}

