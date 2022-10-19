package main

import (
//	"bufio"
	"fmt"
	"os"
	"io"
	"strconv"
//	"strings"
)

type Pid struct {
	name string
	number int
	cwd string
	exe string
	memory int
}

func getProcess() []string {

	dir, err := os.Open("/proc")
	if err != nil {
		panic(err)
	}
	
	defer dir.Close()

	var onlydir []string
	for {
		all, err := dir.Readdirnames(0) 	// list all ( 0 ) files in the directory.
							// error occurs when you arrive at the end of the directory content
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
							// stores in the onlydir array files that are dir
		for _, name := range all { 		// Itteration du dossier contenu dans all
			path := "/proc/" + name
			fileInfo, err := os.Stat(path) 	// donnees sur le directory en question, si erreur alors panic
			if err != nil{
				panic(err)
			}
			if _ , err := strconv.Atoi(name); err != nil { // on ne fait qu'un append si le fichier contient des chiffre
				if fileInfo.IsDir() { // et si c'est un dossier
					onlydir = append(onlydir, name) // append du name dans only dir
				} else {
					continue // sinon on next
				}
			}
		}
	}
	return onlydir
}


func pid(){
	
	fmt.Println(getProcess())

//	file, err := os.Open("/proc/")
}


func main(){
	pid()
//	fmt.Println(pid())
}

