package main

import (
    "io"
    "log"
    "net"
    "sync"

    "github.com/mathhaug/is105sem03_REP03/mycrypt"
)

func main() {
    var wg sync.WaitGroup

    server, err := net.Listen("tcp", "172.17.0.2:5002")
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("bundet til %s", server.Addr().String())

    wg.Add(1)
    go func() {
        defer wg.Done()

        for {
            log.Println("f  r server.Accept() kallet")
            conn, err := server.Accept()
            if err != nil {
                return
            }
            go func(c net.Conn) {
                defer c.Close()

                for {
                    buf := make([]byte, 1024)
                    n, err := c.Read(buf)
                    if err != nil {
                        if err != io.EOF {
                            log.Println(err)
                        }
                        return // fra for l  kke
                    }

                    // Dekrypterer melding
                    dekryptertMelding := mycrypt.Krypter([]rune(string(buf[:n])), mycrypt.ALF_SEM03, len(mycrypt.ALF_SEM03)-4)
                    log.Println("Dekryptert melding:", string(dekryptertMelding))

                    // Sender svar
                    _, err = c.Write([]byte("Server: " + string(dekryptertMelding)))
                    if err != nil {
                        if err != io.EOF {
                            log.Println(err)
                        }
                        return // fra for l  kke
                    }
                }
            }(conn)
        }
    }()

    wg.Wait()
}
