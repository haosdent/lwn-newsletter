package lwn

import (
    "testing"
)

func TestGetLwnContent(t *testing.T) {
    GetLwnContent()
}

func TestSendEmail(t *testing.T) {
    receiver := "xxx"
    password := "xxx"
    server := "smtp.gmail.com"
    port := 587
    SendEmail(GetLwnContent(), receiver, password, server, port)
}