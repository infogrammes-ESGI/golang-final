// Coded by TD, SD & KD the 10/10/22
// Memory & CPU info
package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
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
		// error occurs when you arrive at the end of the directory content
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		// stores in the onlydir array files that are dir
		for _, name := range all { // Itteration du dossier contenu dans all
			fmt.Println(name) 
			
			fileInfo, err := name.Stat() // donnees sur le directory en question, si erreur alors panic
			if err != nil{
				panic(err)
			}
			if _, err := strconv.Atoi(name); err == nil { // on ne fait qu'un append si le fichier est un chiffre
				if fileInfo.IsDir() { // et si c'est un dossier
					onlydir = append(onlydir, name) // append du name dans only dir
				} else {
					continue // sinon on next
				}
			}
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

