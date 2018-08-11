package alexa

type Response struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	ResponseBody      responseBody           `json:"response"`
}

type responseBody struct {
	OutputSpeech     outputSpeech `json:"outputSpeech,omitempty"`
	Card             card         `json:"card,omitempty"`
	Reprompt         outputSpeech `json:"reprompt,omitempty"`
	ShouldEndSession bool         `json:"shouldEndSession,omitempty"`
}

type outputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text"`
	SSML string `json:"ssml"`
}

type card struct {
	Type    string `json:"type"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Text    string `json:"text,omitempty"`
}

type responder interface {
	newResponse() *Response
}

type SimpleResponse struct {
	OutputSpeechText     string
	CardTitle            string
	CardContent          string
	ShouldEndSessionBool bool
}

func (res *SimpleResponse) newResponse() *Response {
	return &Response{
		Version: "1.0",
		ResponseBody: responseBody{
			OutputSpeech: outputSpeech{
				Type: "PlainText",
				Text: res.OutputSpeechText,
			},
			Card: card{
				Type:    "Simple",
				Title:   res.CardTitle,
				Content: res.CardContent,
			},
			ShouldEndSession: res.ShouldEndSessionBool,
		},
	}
}

func NewResponse(r responder) *Response {
	return r.newResponse()
}
