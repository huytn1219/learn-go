// Hashtasg Finder
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter tweet: ")
		tweet, _ := reader.ReadString('\n')

		if strings.EqualFold(strings.TrimSpace(tweet), "quit") {
			fmt.Print("Exiting...")
			break
		}

		hashtags := findHashtags(tweet)

		for _, hashtags := range hashtags {
			fmt.Printf("Hashtag: %s\n", hashtags)
		}

		if len(hashtags) == 0 {
			fmt.Printf("No hashtags")
		}
	}
}

func findHashtags(tweet string) []string {
	var hashtags []string
	words := strings.Split(tweet, " ")

	for _, word := range words {
		if word[0] == '#' {
			hashtags = append(hashtags, word)
		}
	}

	return hashtags
}
