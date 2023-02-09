package main

import (
	"fmt"
	"time"
)

//buffered channeller 1dan kop data saklap bilyarler

// Birden çok veri almaktadır ancak standart bloklama konusunda herhangi bir değişiklik yoktur. Kanala veri girişi ve çıkışı esnasında bloklama işlemi gerçekleşir.
// Kanala veri girişi esnasında kanalın tamamı dolunca yalnızca yeni veri girişi bloklanır. Kanal boşsa yalnızca okuma yani veri çıkışı işlemi bloklanır.
// FIFO (First-In-First-Out) yaklaşımı burada geçerlidir yani kanala ilk giren veri, okuma tarafında ilk çıkan veridir.

func main() {
	bufferedKanal := make(chan string, 3)
	go func() {
		bufferedKanal <- "birinci"
		fmt.Println("1. Gönderildi")
		time.Sleep(time.Second * 1)
		bufferedKanal <- "ikinci"
		fmt.Println("2. Gönderildi")
		bufferedKanal <- "üçüncü"
		fmt.Println("3. Gönderildi")
	}()
	// <-time.After(time.Second * 1)
	go func() {
		ilkOkunan := <-bufferedKanal
		fmt.Println("Aliniyor...")
		fmt.Println(ilkOkunan)
		ikinciOkunan := <-bufferedKanal
		fmt.Println(ikinciOkunan)
		ucuncuOkunan := <-bufferedKanal
		fmt.Println(ucuncuOkunan)
	}()
	<-time.After(time.Second * 5)
}
