package main

import (
    "lwn"
)

func main() {
    config, err := NewIniConfig()
    receiver := config.Get("receiver")
    password := config.Get("password")
    server := config.Get("server")
    port := config.Get("port")
    content := lwn.GetLwnContent()
    lwn.SendEmail(content, receiver, password, server, port)
}