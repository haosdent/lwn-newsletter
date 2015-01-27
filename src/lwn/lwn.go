package lwn

import (
    "fmt"
    "time"
    "io/ioutil"
    "net/http"
    "net/smtp"
    "strings"
    "encoding/base64"
)

func GetLwnLink(category string) (string, error) {
    resp, err := http.Get(fmt.Sprintf("http://lwn.net/%s/", category))
    if err != nil {
        return "", err
    }
    body, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
        return "", err
    }
    lines := strings.Split(string(body[:]), "\n")
    link := ""
    index := 0
    for i := 0; i < len(lines); i++ {
        if strings.Contains(lines[i], "<td valign=\"top\"><a href=\"/Articles/") {
            index++
            if index == 2 {
                link = strings.TrimSpace(lines[i])[26:]
                link = "http://lwn.net" + link[:len(link) - 33]
                break
            }
        }
    }
    return link, nil
}

func GetLwnContent(category string) (string, error) {
    link, err := GetLwnLink(category)
    if err != nil {
        return "", err
    }
    resp, err := http.Get(link)
    if err != nil {
        return "", err
    }
    body, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
        return "", err
    }
    sBody := string(body[:])
    start := strings.Index(sBody, "<div class=\"PageHeadline\">")
    end := strings.Index(sBody, "<!-- ArticleText -->")
    content := sBody[start:end]
    content = strings.Replace(content, "href=\"/", "href=\"http://lwn.net/", -1)
    content = strings.Replace(content, "src=\"/", "src=\"http://lwn.net/", -1)
    return content, nil
}

func SendEmail(category string, content string, receiver string, password string, server string, port int) error {
    sender := "noreply@lwn.net"
    t := time.Now().Local()
    subject := fmt.Sprintf("[LWN] %s Newsletter %s %02d, %04d", category, t.Month(), t.Day(), t.Year())

    template := "Content-Type: text/html; charset=\"utf-8\"\r\nMIME-Version: 1.0\r\nContent-Transfer-Encoding: base64\r\nFrom: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s\r\n.\r\n"
    content = fmt.Sprintf(template, sender, receiver, subject, base64.StdEncoding.EncodeToString([]byte(content)))
    auth := smtp.PlainAuth("", receiver, password, server)
    err := smtp.SendMail(fmt.Sprintf("%s:%d", server, port), auth, sender, []string{receiver}, []byte(content))
    return err
}