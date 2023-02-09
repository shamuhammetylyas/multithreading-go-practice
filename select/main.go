package main

import (
	"fmt"
	"time"
)

//unbuffered channeller read we write ishlemlerinde kody bloklayar diyipdik
//eger read yerinde bloklanmasyny islemesek name etmeli bolar?
//bu yagdayda Select/Case ulanyp bileris.
//n sany goroutine ishledyan bolsak channel ve select-leri ulanyp dine
//ishini gutaranlaryn sonuclaryny alyp bileris

//meselem ashakdaky kod ishletsek goroutine ishlap dowam edende
//sho selectin icindaki 1nji case eger msg kanalyna bir zat gelen bolsa
//ishleyar, gelmedik bolsa default-yn icindaki fmt.Println ishleyar

func main() {
	kanalim := make(chan string)

	go func() {
		kanalim <- "mesaj"
	}()

	select {
	case msg := <-kanalim:
		fmt.Println(msg)
	default:
		fmt.Println("mesaj yok")
	}
	time.Sleep(time.Second * 1)
	select {
	case msg := <-kanalim:
		fmt.Println(msg)
	default:
		fmt.Println("mesaj yok")
	}
}
