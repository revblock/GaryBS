package garyexcuse

// AlexaResponse JSON from Alexa endpoint
type AlexaResponse struct {
	Version  string       `json:"version"`
	Response ResponseBody `json:"response"`
}

// ResponseBody Body of Alexa response
type ResponseBody struct {
	OutputSpeech     OutputSpeech `json:"outputSpeech"`
	ShouldEndSession bool         `json:"shouldEndSession"`
}

// OutputSpeech Actual speech to output
type OutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
