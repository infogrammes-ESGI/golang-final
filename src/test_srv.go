package main

import (
    "bufio"
    "fmt"
    "strings"
    "net"
    "strconv"
)

func verif_user() string {

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

            if (x == 0 && tmp <= 0 || tmp > 256) || (tmp < 0 || tmp > 255) {

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

    return addr
}

func connection(host string, port string) {

    fmt.Println("Lancement du serveur ...")

    ln, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))

    if err != nil {

        panic(err)

    }

    var clients []net.Conn // tableau de clients

    for {

        conn, err := ln.Accept()

        if err == nil {

            clients = append(clients, conn) //rajout d'un client à notre tableau

        }

        if err != nil {

            panic(err)

        }

        fmt.Println("Un client est connecté depuis", conn.RemoteAddr())

        go func() { // goroutine de connexion

            buf := bufio.NewReader(conn)

            for {

                name, err := buf.ReadString('\n')

                if err != nil {

                    fmt.Printf("Client disconnected.\n")
                    break

                }

                for _, c := range clients {

                    c.Write([]byte(name)) // broadcast de message
                }
            }
        }()
    }
}

func main(){

    infos := strings.Split(verif_user(), ":")
    host := infos[0]
    port := infos[1]

    connection(host,port)

}