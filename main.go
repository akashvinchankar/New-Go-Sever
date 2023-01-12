package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Order struct {
	Id   string `json:"id"`
	Name string `json:"Order"`
}

type Schedule struct {
	Id   int    `json:"id"`
	Name string `json:"schedule"`
}

type Screen struct {
	Id   string `json:"code"`
	Name string `json:"screen"`
}

type Show struct {
	Id   int    `json:"id"`
	Name string `json:"show"`
}

var Orders []Order
var Schedules []Schedule
var Screens []Screen
var Shows []Show

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TestData homepage endpoint hit")
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getOrders")

	// Get the "id" parameter from the request
	params := r.URL.Query()
	id := params.Get("id")

	// Check if the "id" parameter is provided in the request
	if id == "" {
		// If not provided, return the whole list of orders
		json.NewEncoder(w).Encode(Orders)
	} else {
		// If provided, filter the list of orders based on the provided id
		var order Order
		for _, o := range Orders {
			if o.Id == id {
				order = o
				break
			}
		}

		// Return the filtered order
		json.NewEncoder(w).Encode(order)
	}
}

func getSchedules(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getSchedules")
	// Get the "id" parameter from the request
	params := r.URL.Query()
	id := params.Get("id")

	// Check if the "id" parameter is provided in the request
	if id == "" {
		// If not provided, return the whole list of schedules
		json.NewEncoder(w).Encode(Schedules)
	} else {
		// If provided, filter the list of schedules based on the provided id
		var schedule Schedule
		for _, s := range Schedules {
			if strconv.Itoa(s.Id) == id {
				schedule = s
				break
			}
		}

		// Return the filtered schedule
		json.NewEncoder(w).Encode(schedule)
	}
}

func getScreens(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getScreens")

	// Get the "code" parameter from the request
	params := r.URL.Query()
	code := params.Get("code")

	// Check if the "code" parameter is provided in the request
	if code == "" {
		// If not provided, return the whole list of screens
		json.NewEncoder(w).Encode(Screens)
	} else {
		// If provided, filter the list of screens based on the provided code
		var screen Screen
		for _, s := range Screens {
			if s.Id == code {
				screen = s
				break
			}
		}

		// Return the filtered screen
		json.NewEncoder(w).Encode(screen)
	}
}

func getShows(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getShows")
	// Get the "id" parameter from the request
	params := r.URL.Query()
	id := params.Get("id")

	// Check if the "id" parameter is provided in the request
	if id == "" {
		// If not provided, return the whole list of shows
		json.NewEncoder(w).Encode(Shows)
	} else {
		// If provided, filter the list of shows based on the provided id
		var show Show
		for _, s := range Shows {
			if strconv.Itoa(s.Id) == id {
				show = s
				break
			}
		}

		// Return the filtered show
		json.NewEncoder(w).Encode(show)
	}
}

func handleRequests() {
	http.HandleFunc("/", homePage)

	http.HandleFunc("/getOrders", getOrders)

	http.HandleFunc("/getSchedules", getSchedules)

	http.HandleFunc("/getScreens", getScreens)

	http.HandleFunc("/getShows", getShows)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	Orders = []Order{
		Order{Id: "Order1", Name: "Order Name 1"},
		Order{Id: "Order2", Name: "Order Name 2"},
	}

	Schedules = []Schedule{
		Schedule{Id: 1, Name: "Schedule Name 1"},
		Schedule{Id: 2, Name: "Schedule Name 2"},
	}

	Screens = []Screen{
		Screen{Id: "ScreenCode1", Name: "Screen Name 1"},
		Screen{Id: "ScreenCode2", Name: "Screen Name 2"},
	}

	Shows = []Show{
		Show{Id: 1, Name: "Morning"},
		Show{Id: 2, Name: "Afternoon"},
		Show{Id: 3, Name: "Evening"},
		Show{Id: 3, Name: "Night"},
	}

	handleRequests()
}
