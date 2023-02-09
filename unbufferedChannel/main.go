package main

import (
	"fmt"
	"time"
)

//unbuffered channeller-e dine bir data berip bolyar
//bu channlerde sender ve receiver bir birne ihtiyac duyyar
//meselem sender channele baha berende shony receiver okamasa bloklanyar shondan ashakdaky kod ishlanok
//edil sholar yaly hem receiver channelden baha aljak bolanda sender yok bolsa onda bloklanyar, shondan ashakdaky kod ishlanok
//funksiyanyn onune "go" programma ishlap bashlanda yokardan ashak gidende onunde "go" yazylan funksiyany ishledip gecip gidyar
//yagny funksiyanyn ishini bitirmegine garashanok ashaklygyna dowam edyar, son 2nji funksiya gelyar edil ol hem sholar yaly
//onunde "go" yazylan bolsa funskiyanyn ishini gutarmagyna garashanok dowam edyar

//indi ashakdaky koda seredip 2 funksiyanyn hem shol birwagtda ishine bashlandygyny goz onune getirelin
//we channele data ugradylmaya diyip pikir edelin
//bu yagdayda 2nji funksiyamyz for-yn icine gelyar we channel-den data almalydygyny goryar
//emma bizin birinji funksiyamyz channele data ugradanok dine print edyar
//shu yagdayda 2nji funksiyamyzdaky channelden data okalyan yerinde blok bolyar
//shondan ashakdaky kod ishlanok sebabi channele data gelenine garashyar, gelenok
//emma 2nji funksiya channelden data garashyp durka 1nji funksiya oz ishini dowam edyar, ishini bitirip cykyp gidyar

func main() {
	kahramanlar := []string{"Marvel", "Flash", "Thanos", "Flash"}
	ilkKanalim := make(chan string)

	// Bulucu
	go func(dizi []string) {
		for _, kahraman := range dizi {
			ilkKanalim <- kahraman // Kanala gönderiliyor
			fmt.Println(kahraman)
			time.Sleep(time.Second)
		}
	}(kahramanlar)

	// Alıcı
	go func() {
		for i := 0; i < 4; i++ {
			// bulunan := <-ilkKanalim // Kanaldan alınıyor
			// fmt.Println("Alici: Bulucudan " + bulunan + " alindi")
			fmt.Println("hello")
		}
	}()
	<-time.After(time.Second * 5)
}
