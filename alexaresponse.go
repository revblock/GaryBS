package garyexcuse

// AlexaResponse JSON from Alexa endpoint
type AlexaResponse struct {
	Version  string
	Response ResponseBody
}

// ResponseBody Body of Alexa response
type ResponseBody struct {
	OutputSpeech     OutputSpeech
	Card             Card
	Reprompt         Reprompt
	ShouldEndSession bool
}

// OutputSpeech Actual speech to output
type OutputSpeech struct {
	Type string
	Text string
}

// Card Alex card
type Card struct {
	Type    string
	Title   string
	Content string
}

// Reprompt Reprompt speech
type Reprompt struct {
	OutputSpeech OutputSpeech
}
