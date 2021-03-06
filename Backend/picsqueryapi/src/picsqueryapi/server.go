/*
	picsqueryapi REST API (Version)
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	//"strconv"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"gopkg.in/mong-go/query.v2/paginate"
	//"github.com/aws/aws-sdk-go/service/s3"
)

// MongoDB Config

//var mongodb_server = "10.0.0.117"
//var mongodb_server = "dockerhost"
var cloudfront_endpoint = "http://d2krh5h0ip6hb6.cloudfront.net"

var mongodb_server = "52.11.201.189" //os.Getenv("Server")
var mongodb_database = "cmpe281"     //os.Getenv("Database") // pics
var mongodb_collection = "Picture"   //os.Getenv("Collection") // picassa
var mongo_user = "cmpe281"           //os.Getenv("User") // masea
var mongo_pass = "cmpe281"           //os.Getenv("Pass") // cmpe281

var last_id string

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	router := mux.NewRouter()
	initRoutes(router, formatter)
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "OPTIONS"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	n.UseHandler(handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(router))
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/pictures", pictureQueryByPageNumberAndCount(formatter)).Methods("GET")
	mx.HandleFunc("/pictures/{pictureId}", pictureQueryByPictureId(formatter)).Methods("GET")
	mx.HandleFunc("/users/{userid}/pictures", pictureQueryByUserId(formatter)).Methods("GET")
}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func getSystemIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr).String()
	address := strings.Split(localAddr, ":")
	fmt.Println("address: ", address[0])
	return address[0]
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		message := "Picture Query API Server Working on machine: " + getSystemIp()
		formatter.JSON(w, http.StatusOK, struct{ Test string }{message})
	}
}

// API Get Pictures by Page Number and Count Handler
func pictureQueryByPageNumberAndCount(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//var request Request
		//_ = json.NewDecoder(req.Body).Decode(&request)
		session, err := mgo.Dial(mongodb_server)
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		/**err = session.DB("cmpe281").Login(mongo_user, mongo_pass)
		if err != nil {
			log.Println(err)
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}**/
		c := session.DB(mongodb_database).C(mongodb_collection)
		params := mux.Vars(req)
		log.Println(params)
		var howMany int = 30
		i, err := strconv.Atoi(req.URL.Query().Get("pagenumber"))
		if err != nil {
			i = 1
		}
		var pagenumber int = i
		result := make([]Payload, 10, 10)
		err = c.Find(nil).Skip(10 * (pagenumber - 1)).Limit(howMany).All(&result)

		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		if len(result) == 0 {
			formatter.JSON(w, http.StatusNotFound, "No Picture Found.")
			return
		}
		_ = json.NewDecoder(req.Body).Decode(&result)
		//mResult, err := json.Marshal(result)
		// if err != nil {
		//     // do something with the error (log it or write it in the response)
		// }
		fmt.Println("Pictures:", result[0])
		formatter.JSON(w, http.StatusOK, result)
	}
}

// API Get Picture by Picture Id Handler
func pictureQueryByPictureId(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//var request Request
		//_ = json.NewDecoder(req.Body).Decode(&request)
		session, _ := mgo.Dial(mongodb_server)
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		err := session.DB("admin").Login(mongo_user, mongo_pass)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		c := session.DB(mongodb_database).C(mongodb_collection)
		params := mux.Vars(req)
		var picid string = params["pictureId"]
		//var picid string = request.PictureId
		fmt.Println("pictureId: ", picid)
		result := make([]Payload, 0, 10)
		err = c.Find(bson.M{"PictureId": picid}).All(&result)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, " Internal Server Error")
			return
		}
		if len(result) == 0 {
			formatter.JSON(w, http.StatusNotFound, "This picture Id does not exist.")
			return
		}
		_ = json.NewDecoder(req.Body).Decode(&result)
		fmt.Println("Pictures: ", result)
		formatter.JSON(w, http.StatusOK, result)
	}
}

// API Get Pictures by User Id
func pictureQueryByUserId(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, _ := mgo.Dial(mongodb_server)
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		err := session.DB("admin").Login(mongo_user, mongo_pass)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		c := session.DB(mongodb_database).C(mongodb_collection)
		params := mux.Vars(req)
		//var qryid string = params["queryId"]
		var userid string = params["userid"]
		fmt.Println("UserId: ", userid)
		result := make([]Payload, 10, 10)
		err = c.Find(bson.M{"userid": userid}).All(&result)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		if len(result) == 0 {
			formatter.JSON(w, http.StatusNotFound, "No results found for this user.")
			return
		}
		_ = json.NewDecoder(req.Body).Decode(&result)
		fmt.Println("Pictures: ", result)
		formatter.JSON(w, http.StatusOK, result)
	}
}
