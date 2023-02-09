package main

import "fmt"

//Go-da channeller bilen ishlaninde in kop gabat gelinyan error "deadlock" errorydyr
//Eger bir kanala nace sany maglumat yazyan(sender) bar bolsa shoncada
//reader yok bolsa shonda "deadlock" erroy doreyar.
//Deadlock kanalyn capacity problemasydyr. Bu errorda cap() we len() funksiyalaryndan
//peydalanyp bileris.
//cap() funksiyasy channel bashda doredilende nace capacity berilenini gaytaryp beryar
//len() funksiyasy channelin icindake nace maglumat bardygyny gaytaryp beryar
//channel-den maglumat okaldygyca len azalyar

func main() {
	teams := make(chan string, 2)
	teams <- "Chelsea"
	teams <- "Real Madrid"
	// teams <- "Man City"
	// fmt.Println(<-teams)
	// fmt.Println(<-teams)

	fmt.Println("Capacity of channel: ", cap(teams))
	fmt.Println("Total data: ", len(teams))
	fmt.Println("Readed data: ", <-teams)
	fmt.Println("New total data: ", len(teams))
}
