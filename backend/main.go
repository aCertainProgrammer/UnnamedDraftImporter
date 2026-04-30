package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

const (
	ADDR string = ":8080"
)

func sendJsonResponse(w http.ResponseWriter, v any) {
	response, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")

	allowedOrigins := []string{
		"https://acertainprogrammer.github.io",
		"http://localhost:8080",
		"http://127.0.0.1:8080",
	}

	for _, allowed := range allowedOrigins {
		if origin == allowed {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			break
		}
	}

	if strings.Contains(r.URL.Path, "favicon") {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if strings.Contains(r.URL.Path, "sw.js") {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	url := r.URL.Query().Get("url")
	mode := r.URL.Query().Get("mode")

	if strings.Contains(url, "drafter.lol") {
		switch mode {
		case "draft":
			log.Println("Getting draft from URL:", url)
			draft, err := GetDrafterDraftByURL(url)
			if err != nil {
				log.Println(err)
				return
			}
			sendJsonResponse(w, draft)

		case "series":
			picksAndBans, err := GetDrafterSeriesByURL(url)
			if err != nil {
				log.Println(err)
				return
			}

			sendJsonResponse(w, picksAndBans)
		default:
			log.Println("import mode not provided to draftlol handler")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		return
	} else {
		log.Println("Could not match a drafting host link, aborting...")
		type StatusResponse struct {
			Status      string `json:"status"`
			Usage       string `json:"usage"`
			ExampleCall string `json:"example_call"`
			LegalModes  string `json:"legal_modes"`
		}

		sendJsonResponse(w,
			StatusResponse{
				Status:      "The scraper is online",
				Usage:       "Provide a drafter.lol link and a mode to get a draft from under the link",
				ExampleCall: "https://frog03-20729.wykr.es/?url=https://drafter.lol/draft/G1fVTicg?game=2&mode=draft",
				LegalModes:  "draft, series",
			})
		return
	}
}

func main() {
	log.Println("Setting up http handlers...")
	http.HandleFunc("/", indexHandler)

	log.Println("Starting server...")
	err := http.ListenAndServe(ADDR, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Closing server...")
}
