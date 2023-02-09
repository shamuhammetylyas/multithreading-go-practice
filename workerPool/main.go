package main

import (
	"fmt"
	"time"
)

// Worker Pool bütün dillerde yaygın bir şekilde kullanılan yazılım dizayn paternidir.
// Genel olarak işçi havuzu, görevlendirilmek üzere bekleyen iş parçacıklarından (işçi de diyebiliriz) oluşur.
// Uygulama esnasında üzerlerine görev alırlar, bu görevi tamamlar ve yeni görev alabilecek şekilde hazır bekleyişe geçerler.
// Worker pool dizaynı Go diline go routine ve kanallar ile uyarlanır.
// Aşağıda Go By Example’daki örneği paylaşarak konuyu özetleyelim.
// Yorum satırlarında Türkçe olarak workerların çalışma şeklini anlatmaya çalıştım.

// Birden fazla eş zamanlı görevi vereceğimiz işçilerimizi
// burada tanımlıyoruz. Bu işçiler jobs kanaldaki işleri alıp,
// sonuçlarını işe karşılık gelen results kanalına iletmekteler.
// Yapılan işi tanımlamak için 1 saniyelik bekleme koyulmuştur.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		// Yapılacak iş vs...
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// İşçilere görev vermek ve sonuçlarını almak için
	// jobs ve results isimli iki kanal oluşturuyoruz.
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Uygulamada 3 işçi çalıştırmaktayız. Başlangıçta
	// bu işçiler bloklanmış durumda çünkü jobs kanalında
	// herhangi bir iş bulunmamaktadır.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Burada jobs kanalına 5 görev gönderiyor ve sonrasında
	// kanalı kapatıyoruz.
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Son olarak sonuçları yazdırıyoruz.
	for a := 1; a <= 5; a++ {
		<-results
	}
}
