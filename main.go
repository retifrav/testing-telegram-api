package main

import (
    "strings"
    "bytes"
    "fmt"
    "log"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "decovar.dev/testing-telegram-api/config"
)

var telegramAPI = fmt.Sprintf("https://api.telegram.org/bot%s", config.TelegramBotToken)

func main() {
    //sendGET()
    sendPOST()
}

func sendGET() {
    resp, err := http.Get(fmt.Sprintf("%s/getMe", telegramAPI))
    if err != nil {
        log.Fatalln(err)
    }

    log.Printf("Status: %s", resp.Status)
    // for key, value := range resp.Header {
    //     log.Printf("%v: %v", key, value)
    // }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    log.Printf("Response: %s", string(body))
}

func sendPOST() {
    postBody, _ := json.Marshal(map[string]string{
        "parse_mode": "MarkdownV2",
        "chat_id": config.TelegramChatID,
        "disable_web_page_preview": "true",
        "text": fmt.Sprintf(
            "Here goes [%s](%s) link",
            escapeText("some [testing] long"),
            "ya.ru",
        ),
        //"text": "Here goes [some \\[testing\\] long](ya.ru) link",
    })

    resp, err := http.Post(
        fmt.Sprintf("%s/sendMessage", telegramAPI),
        "application/json",
        bytes.NewBuffer(postBody),
    )
    if err != nil {
        log.Fatalf("Error %v", err)
    }
    defer resp.Body.Close()

    log.Printf("Status: %s", resp.Status)
    // for key, value := range resp.Header {
    //     log.Printf("%v: %v", key, value)
    // }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    log.Printf("Response: %s", string(body))

}

func escapeText(title string) string {
    escSymbols := []string{"_", "*", "[", "]", "(", ")", "~", "`", ">", "#", "+", "-", "=", "|", "{", "}", "!"}
    res := title
    for _, esc := range escSymbols {
        res = strings.Replace(res, esc, "\\"+esc, -1)
    }
    return res
}
