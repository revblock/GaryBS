package garyexcuse

// AlexaResponse JSON from Alexa endpoint
type AlexaResponse struct {
	Version  string       `json:"version"`
	Response ResponseBody `json:"response"`
}

// ResponseBody Body of Alexa response
type ResponseBody struct {
	OutputSpeech     OutputSpeech `json:"outputSpeech"`
	Card             Card         `json:"card"`
	Reprompt         Reprompt     `json:"reprompt"`
	ShouldEndSession bool         `json:"shouldEndSession"`
}

// OutputSpeech Actual speech to output
type OutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// Card Alex card
type Card struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Reprompt Reprompt speech
type Reprompt struct {
	OutputSpeech OutputSpeech `json:"outputSpeech"`
}
