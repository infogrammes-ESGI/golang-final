package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
)


//TODO fonction de boucle pour recuperer les informations processeurs CF structure
// NE PAS OUBLIER LES LIENS SYMBOLIQUES
//TODO fonction pour formater les valeurs dans une seule string

func parser(line string) (key string, value string){
	text := strings.ReplaceAll(line[:len(line)], " ", "")
	values := strings.Split(text, ":")
	return values[0], values[1]
}

func getPidInfo(pidArray []string) []string {

	var processConcat []string

	for _, name := range pidArray {
		pidDir := "/proc/"+name
		fileName := "/proc/" + name + "/status"
		file, err := os.Open(fileName)
		if err != nil{
			panic(err)
		} 
		scanner := bufio.NewScanner(file)
		var processInfo []string
		for scanner.Scan(){
			key, value := parser(scanner.Text())
			switch key {
			case "Name":
				processInfo = append(processInfo, "NAME: " + value)
			case "Pid":
				processInfo = append(processInfo, " PID:  " + value)
			case "VmRSS":
				processInfo = append(processInfo, " RAM:  " + value)
			}
		}
		dir, _ := os.Open(pidDir)
		exe, _ := os.Readlink(pidDir+"/exe")
		cwd, err := os.Readlink(pidDir+"/cwd")
		formatString := " exe: "+exe+" cwd:  "+cwd+"\n"

		processConcat = append(processConcat, processInfo...)
		processConcat = append(processConcat, formatString)
		file.Close()
		dir.Close()
	}	
	

	return processConcat
}

func getProcess() []string {

	dir, err := os.Open("/proc")
	if err != nil {
		panic(err)
	}

	defer dir.Close()
	var onlydir []string
	for {
		all, err := dir.Readdirnames(10) // list all files in the directory, slice per slice .
		// error occurs when you arrive at the end of the directory content
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		// stores in the onlydir array files that are dir
		for _, name := range all { // Itteration du dossier contenu dans all
			path := "/proc/" + name
			fileInfo, err := os.Stat(path) // donnees sur le directory en question, si erreur alors panic
			if err != nil {
				panic(err)
			}
			if _, err := strconv.Atoi(name); err == nil { // on ne fait qu'un append si le fichier contient des chiffre
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

func pid() {

	processList := getProcess()

	fmt.Println(getPidInfo(processList))

	//fmt.Println(getProcess())

}

func main() {
	pid()
	// fmt.Println(pid())
}
