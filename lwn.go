package lwn

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "net/smtp"
    "strings"
    "encoding/base64"
)

func GetLwnLink() string {
    resp, err := http.Get("http://lwn.net/Archives/")
    if err != nil {
        log.Fatal(err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
        log.Fatal(err)
    }
    lines := strings.Split(string(body[:]), "\n")
    link := ""
    index := 0
    for i := 0; i < len(lines); i++ {
        if strings.Contains(lines[i], "One big page") {
            index++
            if index == 2 {
                link = strings.TrimSpace(lines[i])[10:]
                link = "http://lwn.net" + link[:len(link) - 19]
                break
            }
        }
    }
    return link
}

func GetLwnContent() string {
    link := GetLwnLink()
    resp, err := http.Get(link)
    if err != nil {
        log.Fatal(err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
        log.Fatal(err)
    }
    sBody := string(body[:])
    start := strings.Index(sBody, "<div class=\"PageHeadline\">")
    end := strings.Index(sBody, "<!-- ArticleText -->")
    content := sBody[start:end]
    content = strings.Replace(content, "href=\"/", "href=\"http://lwn.net/", -1)
    content = strings.Replace(content, "src=\"/", "src=\"http://lwn.net/", -1)
    return content
}

func SendEmail(content string, receiver string, password string, server string, port int) {
    sender := "noreply@lwn.net"
    subject := content[strings.Index(content, "LWN"):strings.Index(content, "</h1>")]

    template := "Content-Type: text/html; charset=\"utf-8\"\r\nMIME-Version: 1.0\r\nContent-Transfer-Encoding: base64\r\nFrom: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s\r\n.\r\n"
    content = fmt.Sprintf(template, sender, receiver, subject, base64.StdEncoding.EncodeToString([]byte(content)))
    auth := smtp.PlainAuth("", receiver, password, server)
    err := smtp.SendMail(fmt.Sprintf("%s:%d", server, port), auth, sender, []string{receiver}, []byte(content))
    if err != nil {
        log.Fatal(err)
    }
}