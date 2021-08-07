# testing-telegram-api

Testing escaping for Telegram API in Go. Created for <https://github.com/umputun/remark42/issues/839>.

## The goal

Send the following message:

``` md
Here goes [some \\[testing\\] long](ya.ru) link
```

so recipient in Telegram would receive the following:

![](/img/telegram-link-escaped-brackets.png?raw=true)

## How to run

Enter your Telegram bot API and chat ID in `config/config.go`. Then:

``` sh
$ cd testing-telegram-api
$ go run .
```
