package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// go calculator()
	// controllerr()

}

func controllerr() {
	r := mux.NewRouter()

	http.HandleFunc("/accounts", func(rw http.ResponseWriter, r *http.Request) {})

	http.ListenAndServe(":4321", r)
}

func calculator() {
	glchan := make(chan string)
	fbchan := make(chan string)
	quit := make(chan bool)

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go google(glchan, quit)
	go facebook(fbchan)

	for {
		select {
		case a := <-glchan:
			log.Println("select a")

			log.Println(a)

		case b := <-fbchan:
			log.Println("select b")

			log.Println(b)

		case <-quit:
			return
		}
	}
}

func google(gl chan string, b chan bool) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 1000)
		log.Println("code was google")

		str1 := fmt.Sprint(i)
		gl <- str1
	}
	b <- true
}

func facebook(fb chan string) {
	for i := 10; ; i++ {
		time.Sleep(time.Millisecond * 1000)
		log.Println("code was facebook")

		str1 := fmt.Sprint(i)
		fb <- str1
	}
}
