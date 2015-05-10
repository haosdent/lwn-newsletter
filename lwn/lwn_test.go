package lwn

import (
    "log"
    "strings"
    "testing"
)

func TestGetLwnContent(t *testing.T) {
    category := "Security"
    _, err := GetLwnContent(category)
    if err != nil {
        log.Fatal(err)
    }
}

func TestSendEmail(t *testing.T) {
    category := "Kernel"
    receiver := "xxx"
    password := "xxx"
    server := "smtp.gmail.com"
    port := 587
    content, err := GetLwnContent(category)
    if err != nil {
        log.Fatal(err)
    }
    err = SendEmail(category, content, receiver, password, server, port)
    if err == nil {
        return
    }
    if !strings.Contains(err.Error(), "Username and Password not accepted.") && !strings.Contains(err.Error(), "connection refused") && !strings.Contains(err.Error(), "operation timed out") {
        log.Fatal(err)
    }
}