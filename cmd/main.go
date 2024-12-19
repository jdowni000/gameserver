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
	http.HandleFunc("/", GetRoot)
	http.HandleFunc("/game", GetGame)

	err := http.ListenAndServe(":8080", nil)
	fmt.Print(err)
}

func Writer(w http.ResponseWriter, data string) error {
	_, error := io.WriteString(w, data)
	return error
}

// Gets basic information regarding all games
func GetRoot(w http.ResponseWriter, r *http.Request) {
	games, err := JsonRootGameInfo("games.json")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("got / request\n")

	for i := 0; i < len(games); i++ {
		game := fmt.Sprint("Game: " + games[i].Game + "\n")
		description := fmt.Sprint("Description: " + games[i].Description + "\n")
		id := fmt.Sprint("ID: " + games[i].ID + "\n")
		err := Writer(w, game)
		if err != nil {
			log.Println(err)
		}
		err = Writer(w, description)
		if err != nil {
			log.Println(err)
		}
		err = Writer(w, id)
		if err != nil {
			log.Println(err)
		}
	}
}

// Gets all information from URL search
func GetGame(w http.ResponseWriter, r *http.Request) {
	resp := strings.Split(r.URL.RawQuery, "{")
	respSplit := strings.Split(resp[1], "}")
	id := respSplit[0]

	games, err := JsonGameInfo("games.json")
	if err != nil {
		log.Fatalln(err)
	}

	for _, g := range games {
		// Get specific information matching ID of game
		if g.ID == id {

			var gameInfo []string

			gameInfo = append(gameInfo, fmt.Sprint("Game: "+g.Game+"\n"))
			gameInfo = append(gameInfo, fmt.Sprint("Description: "+g.Description+"\n"))
			gameInfo = append(gameInfo, fmt.Sprint("ID: "+g.ID+"\n"))
			gameInfo = append(gameInfo, fmt.Sprint("CurrentPrice: "+strconv.Itoa(g.CurrentPrice)+"\n"))
			gameInfo = append(gameInfo, fmt.Sprint("SellerName: "+g.SellerName+"\n"))
			gameInfo = append(gameInfo, fmt.Sprint("DeveloperName: "+g.DeveloperName+"\n"))
			gameInfo = append(gameInfo, fmt.Sprint("PublisherName: "+g.PublisherName+"\n"))
			gameInfo = append(gameInfo, fmt.Sprint("ThumbnailURL: "+g.ThumbnailURL+"\n"))

			for _, v := range gameInfo {
				err := Writer(w, v)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}

// Retrieve json information and Unmarshal into RootGameInfo struct
func JsonRootGameInfo(file string) (RootGameInfo, error) {
	b, err := os.Open(file)
	if err != nil {
		log.Printf("Failed to open file: %v\n", err)
		return RootGameInfo{}, err
	}

	byteValue, _ := io.ReadAll(b)

	var games RootGameInfo

	json.Unmarshal(byteValue, &games)

	return games, nil
}

// Retrieve json information and Unmarshal into GameInfo struct
func JsonGameInfo(file string) (GameInfo, error) {
	b, err := os.Open(file)
	if err != nil {
		log.Printf("Failed to open file: %v\n", err)
		return GameInfo{}, err
	}

	byteValue, _ := io.ReadAll(b)

	var games GameInfo

	json.Unmarshal(byteValue, &games)

	return games, nil
}
