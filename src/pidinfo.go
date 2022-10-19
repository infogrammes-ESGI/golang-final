// Coded by TD, SD & KD the 10/10/22
// Memory & CPU info
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)


func pid(){
	file, err := os.Open("/proc/")
}


func main(){
	fmt.println(pid())
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

