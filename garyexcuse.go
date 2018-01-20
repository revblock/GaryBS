package garyexcuse

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func init() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/alexa/", alexaHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	leads, perps, delays := loadParts()
	excuse := buildExcuse(leads, perps, delays)

	fmt.Fprintf(w, excuse)
}

func alexaHandler(w http.ResponseWriter, r *http.Request) {
	leads, perps, delays := loadParts()
	excuse := buildExcuse(leads, perps, delays)

	repromptOutputSpeech := OutputSpeech{Type: "PlainText", Text: "You can say: What would Gary say"}
	reprompt := Reprompt{OutputSpeech: repromptOutputSpeech}
	card := Card{Type: "Simple", Title: "SessionSpeechlet - Gary Excuse", Content: "SessionSpeechley - " + excuse}
	outputSpeech := OutputSpeech{Type: "PlainText", Text: excuse}
	responseBody := ResponseBody{Card: card, OutputSpeech: outputSpeech, Reprompt: reprompt, ShouldEndSession: true}
	alexaResponse := AlexaResponse{Version: "1.0", Response: responseBody}

	js, err := json.Marshal(alexaResponse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func readFile(path string, done chan []string) {
	f, _ := os.Open(path)
	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	f.Close()

	done <- lines
}

func chooseRandom(options []string, done chan string) {
	rand.Seed(time.Now().Unix())
	option := options[rand.Intn(len(options))]
	done <- option
}

func loadParts() ([]string, []string, []string) {

	leadsDone := make(chan []string)
	perpsDone := make(chan []string)
	delaysDone := make(chan []string)

	go readFile("./assets/lead.txt", leadsDone)
	go readFile("./assets/perp.txt", perpsDone)
	go readFile("./assets/delay.txt", delaysDone)

	return <-leadsDone, <-perpsDone, <-delaysDone
}

func buildExcuse(leads []string, perps []string, delays []string) string {
	leadDone := make(chan string)
	perpDone := make(chan string)
	delayDone := make(chan string)

	go chooseRandom(leads, leadDone)
	go chooseRandom(perps, perpDone)
	go chooseRandom(delays, delayDone)

	lead, delay, perp := <-leadDone, <-delayDone, <-perpDone

	return lead + " " + perp + " " + delay
}
