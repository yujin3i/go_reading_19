package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yujin3i/alexa/ygo/alexa"
)

func fetchParticipantsInfo() string {
	res, err := http.Get("https://connpass.com/search/?q=横浜Go読書会")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	numberofParticipants := doc.Find(".amount > span").First().Text()

	return numberofParticipants
}

func makeYgoMessage() string {
	var msg = "今回の参加人数は、" + fetchParticipantsInfo() + "人です。"
	return msg
}

// START OMIT
func alexaDispatchIntentHandler(req alexa.Request) (*alexa.Response, error) {
	switch req.RequestBody.Intent.Name {
	case "LaunchRequest":
		return alexafirstHandler(req)
	case "Participants":
		return alexaYgoHandler(req)
	default:
		return alexaHelpHandler(req)
	}
}

func alexaYgoHandler(req alexa.Request) (*alexa.Response, error) {
	var output string
	output = makeYgoMessage()
	ygoResponse := &alexa.SimpleResponse{
		OutputSpeechText:     output,
		CardTitle:            "参加人数",
		CardContent:          output,
		ShouldEndSessionBool: true,
	}
	return alexa.NewResponse(ygoResponse), nil
}

// END OMIT

func alexafirstHandler(req alexa.Request) (*alexa.Response, error) {
	welcomeResponse := &alexa.SimpleResponse{
		OutputSpeechText:     "横浜Go読書会へようこそ。今回の参加人数をお届けします。",
		CardTitle:            "横浜Go読書会へようこそ。",
		CardContent:          "横浜Go読書会へよここそ。今回の参加人数をお届けします。",
		ShouldEndSessionBool: true,
	}
	return alexa.NewResponse(welcomeResponse), nil
}

func alexaHelpHandler(req alexa.Request) (*alexa.Response, error) {
	helpResponse := &alexa.SimpleResponse{
		OutputSpeechText:     "横浜Go読書会です。今回の参加人数をお届けします。",
		CardTitle:            "参加人数",
		CardContent:          "参加人数をお届けします。",
		ShouldEndSessionBool: true,
	}
	return alexa.NewResponse(helpResponse), nil
}

func main() {
	lambda.Start(alexaDispatchIntentHandler)
}
