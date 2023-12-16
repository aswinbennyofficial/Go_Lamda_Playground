package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type RepoResponse struct {
	LanguagePercent map[string]float32 `json:"language_percent"`
}

type RepoInfo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	IsFork   bool   `json:"fork"`
	URL      string `json:"url"`
	Language string `json:"language"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// BUG 
	username:=request.Path

	// Log the GitHub username for debugging purposes
	log.Printf("GitHub Username: %s", username)

	if username==""{
		username="aswinbennyofficial"
	}

	// Construct the GitHub API URL
	url := "https://api.github.com/users/" + username + "/repos"

	// Log the GitHub API URL for debugging purposes
	log.Printf("GitHub API URL: %s", url)

	// Fetch the JSON response from GitHub's API
	responseJSON, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching GitHub API: %s", err)
		return events.APIGatewayProxyResponse{}, err
	}
	defer responseJSON.Body.Close()

	// Check the HTTP status code in the response
	if responseJSON.StatusCode < 200 || responseJSON.StatusCode >= 300 {
		log.Printf("GitHub API returned non-2xx status code: %d", responseJSON.StatusCode)
		return events.APIGatewayProxyResponse{}, fmt.Errorf("GitHub API returned non-2xx status code: %d", responseJSON.StatusCode)
	}

	// Make an object of RepoList struct
	var repolist []RepoInfo

	// Decode JSON body to struct
	if err := json.NewDecoder(responseJSON.Body).Decode(&repolist); err != nil {
		log.Printf("Error decoding JSON to struct: %s", err)
		return events.APIGatewayProxyResponse{}, err
	}

	// Log the number of repositories fetched from GitHub
	log.Printf("Number of Repositories Fetched: %d", len(repolist))

	// Making a new map to save the language and the number of repos that use it
	mapOfLang := make(map[string]float32)

	// Number of non-null lang usages
	var numOfLangs int

	// Traverse through the entire language and make a map for the usage
	for i := range repolist {
		if repolist[i].Language != "" && !repolist[i].IsFork {
			numOfLangs++
			mapOfLang[repolist[i].Language]++
		}
	}

	// Calculate the percentage usage of these langs
	for key, value := range mapOfLang {
		mapOfLang[key] = (value * 100.0) / float32(numOfLangs)
	}

	// Preparing the response
	// Creating an object for struct RepoResponse
	var reporesponse RepoResponse
	// Assigning LanguagePercent field to the previously created map
	reporesponse.LanguagePercent = mapOfLang

	// Log the final language percentages for debugging purposes
	log.Printf("Language Percentages: %v", mapOfLang)

	// Converting struct object to JSON
	jsonResponse, err := json.Marshal(reporesponse)
	if err != nil {
		log.Printf("Error converting struct to JSON object: %s", err)
		return events.APIGatewayProxyResponse{}, err
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(jsonResponse),
	}

	return response, nil
}

func main() {
	lambda.Start(handler)
}
