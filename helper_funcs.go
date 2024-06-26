package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/projectdiscovery/gologger"
)

func HttpRequest(url string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		gologger.Fatal().Msg(fmt.Sprintf("Error When requesting to: %s Error : %s", url, err))
	}

	resp, err := client.Do(req)
	if err != nil {
		gologger.Fatal().Msg(err.Error())
	}
	return resp
}

func ChangeUrlToTelegramWebUrl(input string) string {
	// Check if the input URL already contains "/s/", if not, add it
	if !strings.Contains(input, "/s/") {
		// Find the position of "/t.me/" in the URL
		index := strings.Index(input, "/t.me/")
		if index != -1 {
			// Insert "/s/" after "/t.me/"
			modifiedURL := input[:index+len("/t.me/")] + "s/" + input[index+len("/t.me/"):]
			return modifiedURL
		}
	}
	// If "/s/" already exists or "/t.me/" is not found, return the input as is
	return input
}

func readFileContent(filePath string) (string, error) {
	// Read the entire file content
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Convert the content to a string and return
	return string(content), nil
}

func reverse(lines []string) []string {
	for i := 0; i < len(lines)/2; i++ {
		j := len(lines) - i - 1
		lines[i], lines[j] = lines[j], lines[i]
	}
	return lines
}

func RemoveDuplicate(config string) string {
	lines := strings.Split(config, "\n")

	// Use a map to keep track of unique lines
	uniqueLines := make(map[string]bool)

	// Loop over lines and add unique lines to map
	for _, line := range lines {
		if len(line) > 0 {
			uniqueLines[line] = true
		}
	}

	// Join unique lines into a string
	uniqueString := strings.Join(getKeys(uniqueLines), "\n")

	return uniqueString
}

func getKeys(m map[string]bool) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func WriteToFile(fileContent string, filePath string) {

	// Check if the file exists
	if _, err := os.Stat(filePath); err == nil {
		// If the file exists, clear its content
		err = ioutil.WriteFile(filePath, []byte{}, 0644)
		if err != nil {
			fmt.Println("Error clearing file:", err)
			return
		}
	} else if os.IsNotExist(err) {
		// If the file does not exist, create it
		_, err = os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
	} else {
		// If there was some other error, print it and return
		fmt.Println("Error checking file:", err)
		return
	}

	// Write the new content to the file
	err := ioutil.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("File written successfully")
}
