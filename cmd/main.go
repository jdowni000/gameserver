package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type GameInfo []struct {
	Game          string `json:"game"`
	Description   string `json:"description"`
	ID            string `json:"id"`
	CurrentPrice  int    `json:"currentPrice"`
	SellerName    string `json:"sellerName"`
	DeveloperName string `json:"developerName"`
	PublisherName string `json:"publisherName"`
	ThumbnailURL  string `json:"thumbnailURL"`
}

type RootGameInfo []struct {
	Game        string `json:"game"`
	Description string `json:"description"`
	ID          string `json:"id"`
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/game", getGame)

	err := http.ListenAndServe(":8080", nil)
	fmt.Print(err)
}

// Gets basic information regarding all games
func getRoot(w http.ResponseWriter, r *http.Request) {
	games := jsonRootGameInfo("games.json")
	fmt.Printf("got / request\n")

	for i := 0; i < len(games); i++ {
		game := fmt.Sprint("Game: " + games[i].Game + "\n")
		description := fmt.Sprint("Description: " + games[i].Description + "\n")
		id := fmt.Sprint("ID: " + games[i].ID + "\n")
		io.WriteString(w, game)
		io.WriteString(w, description)
		io.WriteString(w, id)

	}
}

// Gets all information from URL search
func getGame(w http.ResponseWriter, r *http.Request) {
	resp := strings.Split(r.URL.RawQuery, "{")
	respSplit := strings.Split(resp[1], "}")
	id := respSplit[0]

	games := jsonGameInfo("games.json")

	for _, g := range games {
		// Get specific information matching ID of game
		if g.ID == id {

			game := fmt.Sprint("Game: " + g.Game + "\n")
			description := fmt.Sprint("Description: " + g.Description + "\n")
			id := fmt.Sprint("ID: " + g.ID + "\n")
			currentPrice := fmt.Sprint("CurrentPrice: " + strconv.Itoa(g.CurrentPrice) + "\n")
			sellerName := fmt.Sprint("SellerName: " + g.SellerName + "\n")
			developerName := fmt.Sprint("DeveloperName: " + g.DeveloperName + "\n")
			publisherName := fmt.Sprint("PublisherName: " + g.PublisherName + "\n")
			thumbnailURL := fmt.Sprint("ThumbnailURL: " + g.ThumbnailURL + "\n")

			io.WriteString(w, game)
			io.WriteString(w, description)
			io.WriteString(w, id)
			io.WriteString(w, currentPrice)
			io.WriteString(w, sellerName)
			io.WriteString(w, developerName)
			io.WriteString(w, publisherName)
			io.WriteString(w, thumbnailURL)
		}
	}
}

// Retrieve json information and Unmarshal into RootGameInfo struct
func jsonRootGameInfo(file string) RootGameInfo {
	b, err := os.Open(file)
	if err != nil {
		log.Fatalf("Failed to open file: %v\n", err)
	}

	byteValue, _ := io.ReadAll(b)

	var games RootGameInfo

	json.Unmarshal(byteValue, &games)

	return games
}

// Retrieve json information and Unmarshal into GameInfo struct
func jsonGameInfo(file string) GameInfo {
	b, err := os.Open(file)
	if err != nil {
		log.Fatalf("Failed to open file: %v\n", err)
	}

	byteValue, _ := io.ReadAll(b)

	var games GameInfo

	json.Unmarshal(byteValue, &games)

	return games
}
