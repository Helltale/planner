package main

import (
	"flag"
	"log"
)

//ТОКЕН НЕЛЬЗЯ ПЕРЕДАВАТЬ В КАЧЕСТВЕ КОНСТАНТЫ, ПОТОМУ ЧТО БОТА ПОО ТОКЕНУ МОГУТ УГНАТЬ
//ЛУЧШЕ ПЕРЕДАВАТЬ ТОКЕ Н ЧЕРЕЗ ФЛАГИ

func main() {

	token := mustToken()
	//tgclient = telegram.New(token)
	//fetchher = fetchher.New()
	//processor = priocessor.New()

	//consumer.Start(fetcher, processor)
}

func mustToken() string {
	toke := flag.String("token-bot-token", //name
		"",                 //default
		"token for tg bot", //podskazka
	)
	flag.Parse()

	if *toke == "" {
		log.Fatal("token is broken")
	}

	return *toke
}
