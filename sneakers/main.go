package main

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"math/rand"
	"net"
	"net/http"
	"sort"
	"strings"
	"time"
)

// ::: VARS :::::

var SNEAKERS = getSneakerMockData(3)
var TITLE = strings.Split("nike adidas belkelme toto sigomoto jira boosty", " ")
var IMG = []string{
	"/sneakers/sneakers-1.jpg",
	"/sneakers/sneakers-2.jpg",
	"/sneakers/sneakers-3.jpg",
	"/sneakers/sneakers-4.jpg",
	"/sneakers/sneakers-5.jpg",
}

type Sneaker struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	Price int       `json:"price"`
	Img   string    `json:"img"`
}

func SetupEndpoints() {

	http.HandleFunc("/sneakers", getSneakers())
	http.HandleFunc("/refresh", refreshSnickers())
}

func refreshSnickers() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("refreshing snickers OK!"))
		SNEAKERS = getSneakerMockData(len(SNEAKERS))
	}
}

func getSneakers() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// GET "sort by" paraameter

		sortBy := r.URL.Query().Get("sortBy")

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		sneakersData := getSneakersByFilter(sortBy)
		sneakersDataBytes, _ := json.Marshal(sneakersData)
		reader := bytes.NewReader(sneakersDataBytes)
		_, err := io.Copy(w, reader)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

}

func getSneakersByFilter(by string) []*Sneaker {

	switch by {
	case "title":
		sortByTitle()
	case "price":

		sortByPrice()
	case "id":
		sortByID()

	default:
		sortByTitle()
	}

	return SNEAKERS

}

func sortByTitle() {
	sort.Slice(SNEAKERS, func(i, j int) bool {
		return SNEAKERS[i].Title < SNEAKERS[j].Title
	})

}
func sortByPrice() {
	sort.Slice(SNEAKERS, func(i, j int) bool {
		return SNEAKERS[i].Price < SNEAKERS[j].Price
	})

}
func sortByID() {
	sort.Slice(SNEAKERS, func(i, j int) bool {
		return SNEAKERS[i].ID.String() < SNEAKERS[j].ID.String()
	})

}

func getSneakerMockData(dataAmount int) []*Sneaker {
	sneakers := make([]*Sneaker, 0)
	for i := 0; i < dataAmount; i++ {
		sneaker := randomSneaker()
		sneakers = append(sneakers, sneaker)
	}
	sneakers = append(sneakers)
	return sneakers
}

func randomSneaker() *Sneaker {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &Sneaker{
		ID:    uuid.New(),
		Title: TITLE[r.Intn(len(TITLE))],
		Img:   IMG[r.Intn(len(IMG))],
		Price: r.Intn(10) * 100,
	}

}
func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("Error while listening on port 8080" + err.Error())
	}
	SetupEndpoints()
	if err := http.Serve(listener, nil); err != nil {
		panic("Error while serving on port 8080" + err.Error())
	}

}
