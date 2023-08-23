package main

import (
	"flag"
	"log"
	"planner/clients/telegram"
)

const (
	tgHost = "api.telegram.org"
)

func main() {
	//token = flag.Get(token) 		(получаем в виде флага для секретности токена, чтообы нельзя было увидеть его в коде)

	client = telegram.New(tgHost, mustToken())
	//fetcher = fetcher.New() 		[отправка запросов для полоучения новых событий]
	//processore = processor.New() 	[отправка новых сообщений полоьзователю]

	//consumer.Start(fetcher, prrocessor)

}

func mustToken() string {
	tok := flag.String("token-bot-token", "", "token to run bot") //[имя флага, занчение по умолочанию, подсказка при запуске]
	flag.Parse()
	if *tok == "" {
		log.Fatal("uncorrect token")
	}
	return *tok
}
