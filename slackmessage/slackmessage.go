package slackmessage

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"../utilities"
)

// Slack Message struct for decoding the json into a struct
type SlackMessage struct {
	Type string `json:"type"`
	Text string `json:"text"`
	User string `json:"user"`
}

// ParseAndFilterSlackMessages Given a filepath for a slack message json and a userid, it will parse through and return the messages filtered by the user
func ParseAndFilterSlackMessages(filename string, userID string) []string {

	var _slackMessages []SlackMessage

	var messages []string

	//Open the file for reading
	f, err := os.Open(filename)
	utilities.Check(err)

	//Decode the incoming json
	jsonParser := json.NewDecoder(f)
	if err = jsonParser.Decode(&_slackMessages); err != nil {
		panic(err)
	}

	// for each message in the array of slackmessges filter by user id and if it's not empty or if its not a link
	for _, message := range _slackMessages {
		if message.User == userID && message.Text != "" && !strings.HasPrefix(message.Text, "<http") {
			fmt.Println(userID, " said ", message.Text)
			messages = append(messages, message.Text)
		}
	}
	return messages
}
