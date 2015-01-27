package lwn

import (
    "log"
    "strings"
    "testing"
)

func TestGetLwnContent(t *testing.T) {
    _, err := GetLwnContent()
    if err != nil {
        log.Fatal(err)
    }
}

func TestSendEmail(t *testing.T) {
    receiver := "xxx"
    password := "xxx"
    server := "smtp.gmail.com"
    port := 587
    content, err := GetLwnContent()
    if err != nil {
        log.Fatal(err)
    }
    err = SendEmail(content, receiver, password, server, port)
    if !strings.Contains(err.Error(), "Username and Password not accepted.")
    || !strings.Contains(err.Error(), "587: connection refused") {
        log.Fatal(err)
    }
}