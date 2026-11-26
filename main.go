package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// sending GET request to LoLalytics Annie URL with the params I need (objective: grab annie.html)
	resp, err := http.Get(
		"https://lolalytics.com/lol/annie/build/?tier=grandmaster_plus&patch=16.23",
	)
	fmt.Println(resp.Status)
	if err != nil {
		fmt.Println(err) // too lazy to treat errors
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err) // too lazy to treat errors
	}

	os.WriteFile("annie.html", body, 0644)
}
