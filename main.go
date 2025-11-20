package main

import(
	"fmt"
	"encoding/json"
	// "regexp"
)


type bookmark struct {
	Url string `json:"url"`
	Name string `json:"name"`
}


func main () {
	// Create a Go struct instance
	p:= bookmark{
		Url: "https://www.duckduckgo.com/innerdocument.docx",
		Name: "innerdocument.docx",
	}

	// Encode (marshal) the struct to JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
	  fmt.Println("Error encoding JSON", err)
	  return
	}
	fmt.Println("Encoded JSON:", string(jsonData))


	// Decode (unmarshal) JSON back into a struct
	var decodedBookmark bookmark
	err = json.Unmarshal(jsonData, &decodedBookmark)
	if err != nil {
	  fmt.Println("Error decoding JSON", err)
	  return
	}
	fmt.Println("Decoded struct:", decodedBookmark)
}
