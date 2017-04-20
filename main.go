package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var homeView []byte
var productView []byte
var contactView []byte
var teamView []byte
var materialView []byte
var astaxanthinView []byte
var astaxanthinHBView []byte
var db *gorm.DB = nil

func main() {

	dbname := flag.String("dbname", DB_NAME, "The DB name this server runs on.")
	dbuname := flag.String("dbuname", DB_UNAME, "The user name of the DB.")
	dbpswd := flag.String("dbpswd", DB_PSWD, "Password of the DB.")
	flag.Parse()

	time := strconv.FormatInt(time.Now().Unix(), 10)
	log.Println("Starting server at epoch time: " + time)
	relativePath := getRelativePath()

	homeView = loadView("home")
	productView = loadView("product")
	contactView = loadView("contact")
	teamView = loadView("team")
	materialView = loadView("material")
	astaxanthinView = loadView("astaxanthin")
	astaxanthinHBView = loadView("astaxanthin-health-benefits")

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
	http.HandleFunc("/astaxanthin-health-benefits", astaxanthinHBHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/subscribe", subscribeHandler)
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/admin/signin", adminSignInHandler)
	http.HandleFunc("/admin/subscriber", adminSubscriberHandler)
	http.HandleFunc("/admin/feedback", adminFeedbackHandler)
	// Initialize database
	args := "host=localhost user=" + *dbuname + " dbname=" + *dbname + " sslmode=disable password=" + *dbpswd
	var err error
	if db, err = gorm.Open("postgres", args); err != nil {
		log.Println(err)
		panic("ERROR: Failed to initialize database")
	}
	createTables()
	defer db.Close()
	// Start session update task
	startAdminSessionTokenUpdateTask()
	go func() {
		log.Println("Starting iconthin.com application on" + PORT)
		// httpErr := http.ListenAndServe(PORT, secureRedirectHandler(http.StatusFound))
		httpErr := http.ListenAndServe(PORT, nil)
		if httpErr != nil {
			panic("ERROR: " + httpErr.Error())
		}
	}()
	log.Println("APP: Securely server HTTPs on port:" + PORT_SECURE)
	if err = http.ListenAndServeTLS(PORT_SECURE, "ssl/cert.pem", "ssl/privkey-rsa.pem", nil); err != nil {
		log.Fatal("ERROR: ListenAndServeTLS:", err)
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

func astaxanthinHBHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write(loader.LoadView("astaxanthin-health-benefits"))
	w.Write(astaxanthinHBView)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(loadView("admin"))
}
