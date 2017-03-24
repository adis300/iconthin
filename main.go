package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

var homeView []byte
var productView []byte
var contactView []byte
var teamView []byte
var materialView []byte
var astaxanthinView []byte

func main() {
	time := strconv.FormatInt(time.Now().Unix(), 10)
	log.Println("Starting server at epoch time: " + time)
	relativePath := getRelativePath()

	homeView = loadView("home")
	productView = loadView("product")
	contactView = loadView("contact")
	teamView = loadView("team")
	materialView = loadView("material")
	astaxanthinView = loadView("astaxanthin")

	http.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir(relativePath+"/public/js"))))
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir(relativePath+"/public/css"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts", http.FileServer(http.Dir(relativePath+"/public/fonts"))))
	http.Handle("/video/", http.StripPrefix("/video", http.FileServer(http.Dir(relativePath+"/public/video"))))
	http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir(relativePath+"/public/img"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/product", productHandler)
	http.HandleFunc("/team", teamHandler)
	http.HandleFunc("/material", materialHandler)
	http.HandleFunc("/astaxanthin", astaxanthinHandler)
	http.HandleFunc("/contact", contactHandler)

	log.Println("Starting iconthin.com application on " + PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal("Fatal error happened server on port" + PORT)
	}
}

// ======== All handlers implementation
func homeHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write(loader.LoadView("home"))
	w.Write(homeView)
}
func productHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write(loader.LoadView("product"))
	w.Write(productView)
}
func teamHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write(loader.LoadView("team"))
	w.Write(teamView)
}
func materialHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write(loader.LoadView("material"))
	w.Write(materialView)
}
func astaxanthinHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write(loader.LoadView("astaxanthin"))
	w.Write(astaxanthinView)
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write(loader.LoadView("contact"))
	w.Write(contactView)
}
