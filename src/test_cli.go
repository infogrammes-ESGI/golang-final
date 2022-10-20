package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "sync"
)

func connection(host string, port string) {

    var wg sync.WaitGroup

    // Connexion au serveur
    conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))

    if err != nil {

        panic(err)

    }

    wg.Add(2)

    go func() { // goroutine entré user

        defer wg.Done()

        for {

            reader := bufio.NewReader(os.Stdin)

            text, err := reader.ReadString('\n')

            if err != nil {

                panic(err)

            }

            conn.Write([]byte(text))
        }

    }()

    go func() { // goroutine messages

        defer wg.Done()

        for {

            message, err := bufio.NewReader(conn).ReadString('\n')

            if err != nil {
                panic(err)
            }

            fmt.Print("serveur : " + message)
        }
    }()

    wg.Wait()

}

func main(){

    host := "127.0.0.1"
    port := "2568"

    connection(host,port)

}