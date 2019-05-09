package main

import (
	"bufio"
	"fmt"
	"os"

	"./slackmessage"
	"./utilities"
)

var (
	root         string
	files        []string
	err          error
	userID       string
	userMessages []string
)

// function takes 2 args full path to the folder containing all the slack messages dump and the slack userid of the user usually in a format of UXXXXX
func main() {

	args := os.Args[1:]

	if len(args) > 0 {
		root := args[0]
		userID := args[1]

		//Open file for writing the output
		outFile, err := os.Create("./" + userID + ".txt")
		utilities.Check(err)
		w := bufio.NewWriter(outFile)

		//build list of files within the root directory
		files, err = utilities.FilePathWalkDir(root)
		utilities.Check(err)

		//Loop through each file and append the slice of strings returned from the ParseAndFilterSlackMessages function
		for _, file := range files {
			fmt.Println("Working on ", file)
			userMessages = append(userMessages, slackmessage.ParseAndFilterSlackMessages(file, userID)...)
		}

		//write out the file of the collective slice of string containing the messages from the filtered user
		for _, message := range userMessages {
			outString := message + "\n"
			_, err = w.WriteString(outString)
			utilities.Check(err)
		}

		w.Flush()
	} else {
		fmt.Println("Not enough args provided.")
	}

}
