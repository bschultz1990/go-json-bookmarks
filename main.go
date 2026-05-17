package main

import (
	"encoding/json"
	"fmt"
	"os"
	// "regexp"
)


type bookmarks struct {
	Url string `json:"url"`
	Name string `json:"name"`
}


// func importTextLinks () {
//
// }


// func extractLinkName () {
//
// }


func main () {
	// Create a Go struct instance
	p:= bookmarks{
		Url: "https://www.duckduckgo.com/innerdocument.docx",
		Name: "innerdocument.docx",
	}


	// Encode (marshal) the struct to JSON
	jsonData, err := json.MarshalIndent(p,"","	")
	if err != nil {
	  fmt.Println("Error encoding JSON", err)
	  return
	}
	fmt.Println("Encoded JSON:", string(jsonData))
	os.WriteFile("bookmarks.json", []byte(jsonData), 0666)


	// Decode (unmarshal) JSON back into a struct
	var decodedBookmark bookmarks
	err = json.Unmarshal(jsonData, &decodedBookmark)
	if err != nil {
	  fmt.Println("Error decoding JSON", err)
	  return
	}
	fmt.Println("Decoded struct:", decodedBookmark)
}
