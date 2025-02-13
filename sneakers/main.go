package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net"

	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// ::: VARS :::::
var VACANT_INDEXES = make([]int, 0)


var SNEAKERS = getSneakerMockData(10)
var SNEAKERS_FAVORITES  = make([]*SneakerFavorite, 0) 

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
type SneakerFavorite struct {
	ID    int `json:"id"`
	  ParentId  uuid.UUID `json:"parentId"`
}


type Order struct {
	ID    uuid.UUID `json:"id"`
	Items []Sneaker `json:"items"`
	Price int       `json:"price"`
}

func SetupEndpoints() {

	http.HandleFunc("/sneakers", getSneakers())
	http.HandleFunc("/sneakers/order", makeOrder())
	http.HandleFunc("/refresh", refreshSnickers())
	http.HandleFunc("/sneakers/favorites", getSneakersFavorites())
	http.HandleFunc("/sneakers/favorites/add", addToFavorites())
	http.HandleFunc("/sneakers/favorites/delete", deleteFromFavorites())
}

func deleteFromFavorites()func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

id:= r.URL.Query().Get("id")
println("delete from favorites", id)
w.Header().Add("Access-Control-Allow-Origin", "*")
w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
w.WriteHeader(200)

for i, item := range SNEAKERS_FAVORITES {
	if strconv.Itoa(item.ID )== id {
		SNEAKERS_FAVORITES = append(SNEAKERS_FAVORITES[:i], SNEAKERS_FAVORITES[i+1:]...)
		VACANT_INDEXES = append(VACANT_INDEXES, item.ID)
		break
	}
}
w.Write([]byte("deleted from favorites"))
	}}

func addToFavorites()  func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		println("add to favorites")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.WriteHeader(200)

		var item  *SneakerFavorite
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	
	
	log.Println("item", item)
	favorite := SneakerFavorite{
		ID: getVacantIndex(),
		ParentId: item.ParentId,
	}
	SNEAKERS_FAVORITES = append(SNEAKERS_FAVORITES, &favorite)
	// favoriteBytes, _ := json.Marshal(favorite)
	w.Header().Set("Content-Type", "application/json")
	if err:= json.NewEncoder(w).Encode(favorite); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	}
	}

func getVacantIndex() int {


	if len(VACANT_INDEXES) == 0 {
		return len(SNEAKERS_FAVORITES)
	}else{

		id := VACANT_INDEXES[0]
		VACANT_INDEXES = VACANT_INDEXES[1:]
		return id
	}
}

func getSneakersFavorites() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(SNEAKERS_FAVORITES)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		sneakersDataBytes, _ := json.Marshal(SNEAKERS_FAVORITES)
		reader := bytes.NewReader(sneakersDataBytes)
		_, err := io.Copy(w, reader)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func refreshSnickers() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)

		w.Write([]byte("refreshing snickers OK!"))
		SNEAKERS = getSneakerMockData(len(SNEAKERS))
	}
}
func makeOrder() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.WriteHeader(200)
		var order  *Order
		if err:= json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
		log.Println("order", order)
		w.Write([]byte("order created"))
	}
}

func getSneakers() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// GET "sort by" paraameter

		sortBy := r.URL.Query().Get("sortBy")
		queryByTitle := r.URL.Query().Get("title")

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		sneakersData := getSneakersBy(sortBy, queryByTitle)
		sneakersDataBytes, _ := json.Marshal(sneakersData)
		reader := bytes.NewReader(sneakersDataBytes)
		_, err := io.Copy(w, reader)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

}

func getSneakersBy(by string, queryByTitle string) []*Sneaker {

	sneakers := make([]*Sneaker, 0)
	if queryByTitle != "" {
		for _, sneaker := range SNEAKERS {
			if strings.Contains(strings.ToLower(sneaker.Title), strings.ToLower(queryByTitle)) {
				sneakers = append(sneakers, sneaker)
			}
		}

	} else {
		sneakers = SNEAKERS
	}
	switch by {
	case "title":
		sortByTitle(sneakers)
	case "price":

		sortByPrice(sneakers)
	case "-price":

		sortByPriceDesc(sneakers)
	case "id":
		sortByID(sneakers)

	default:
		sortByTitle(sneakers)
	}

	return sneakers

}

func sortByTitle(s []*Sneaker) {
	sort.Slice(s, func(i, j int) bool {
		return s[i].Title < s[j].Title
	})

}
func sortByPrice(s []*Sneaker) {
	sort.Slice(s, func(i, j int) bool {
		return s[i].Price < s[j].Price
	})

}
func sortByPriceDesc(s []*Sneaker) {
	sort.Slice(s, func(i, j int) bool {
		return s[i].Price > s[j].Price
	})

}
func sortByID(s []*Sneaker) {
	sort.Slice(s, func(i, j int) bool {
		return s[i].ID.String() < s[j].ID.String()
	})

}

func getSneakerMockData(dataAmount int) []*Sneaker {
	sneakers := make([]*Sneaker, 0)
	for i := 0; i < dataAmount; i++ {
		time.Sleep(100 * time.Millisecond)
		sneaker := randomSneaker()
		log.Println(sneaker)
		sneakers = append(sneakers, sneaker)
	}
	sneakers = append(sneakers)
	return sneakers
}

func randomSneaker() *Sneaker {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	log.Println(r)
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
