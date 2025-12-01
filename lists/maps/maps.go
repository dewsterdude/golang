package main

import "fmt"

func main() { // string is the key (wihtin brackets), value is outside - key value pairs, key and value

	// websites := []string{"https://google.com", "https://aws.com"}
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])

	// can always add values to a map, always dynamic
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google") // delete in websites the key associated to Google
	fmt.Println(websites)

}
