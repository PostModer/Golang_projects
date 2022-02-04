package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	host string = "localhost"
	port int    = 8083

	html string = `<html>
		<body>
			<h1>
				Wow, you are genius!!!
				<br>Now look at this cute kitty;)
			</h1>
			<img src="%v">
			</img>
		</body>
	</html>`

	img = []string{"https://avatars.mds.yandex.net/get-zen_doc/1658683/pub_5e4071f3e6e8eb5b95da9dfc_5e4072472e9e63535024c7d2/scale_1200",
		"https://chudo-prirody.com/uploads/posts/2021-08/1628642965_7-p-persidskii-kot-foto-7.jpg"}
)

func main() {

	serverAddres := fmt.Sprintf("%v:%v", host, port)
	fmt.Println("Start Kitty server")
	http.HandleFunc("/siam", Siam)
	http.HandleFunc("/pers", Pers)
	http.HandleFunc("/random", Random)

	fmt.Println(serverAddres)

	log.Fatal(http.ListenAndServe(serverAddres, nil))
}

func Siam(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("User-Agent")
	fmt.Println(user)
	w.Write([]byte(fmt.Sprintf(html, img[0])))
}

func Pers(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("User-Agent")
	fmt.Println(user)
	w.Write([]byte(fmt.Sprintf(html, img[1])))
}

func Random(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("User-Agent")
	fmt.Println(user)
	rand.Seed(time.Now().UnixNano())
	pic := img[rand.Intn(len(img))]
	w.Write([]byte(fmt.Sprintf(html, pic)))
}
