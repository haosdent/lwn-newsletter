package main

import (
    "os"
    "log"
    "fmt"
    "lwn"
    "strconv"
    "github.com/haosdent/commons-configuration"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Printf("[Usage]:\n\t %s [config file location]\n", os.Args[0])
        os.Exit(-1)
    }

    path := os.Args[1]
    var conf config.Configer = config.NewIniConfig(path)
    receiver, err := conf.Get("receiver")
    if err != nil {
        log.Fatal(err)
    }
    password, err := conf.Get("password")
    if err != nil {
        log.Fatal(err)
    }
    server, err := conf.Get("server")
    if err != nil {
        log.Fatal(err)
    }
    sPort, err := conf.Get("port")
    if err != nil {
        log.Fatal(err)
    }
    port, err := strconv.Atoi(sPort)
    if err != nil {
        log.Fatal(err)
    }
    content := lwn.GetLwnContent()
    lwn.SendEmail(content, receiver, password, server, port)
    fmt.Printf("Send lwn weekly to %s success!\n", receiver)
}