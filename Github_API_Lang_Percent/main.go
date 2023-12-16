package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type RepoResponse struct{
	LanguagePercent map[string]float32 `json:"language_percent"`
}

type RepoInfo struct{
	Id int `json:"id"`
	Name string `json:"name"`
	IsFork bool `json:"fork"`
	Url string `json:"url"`
	Language string `json:"language"`
}



func githubHandle(w http.ResponseWriter, r *http.Request){
	// fetch username from <url>/github/<username>
	path:=r.URL.Path
	path= strings.Replace(path, "/github/", "", 1)
	// trim away succeeding '/' if exists 
	path= strings.ReplaceAll(path, "/", "")

	url:="https://api.github.com/users/"+path+"/repos"

	// Fetch the json response from github's API
	responseJSON,err:=http.Get(url)
	if err!=nil{
		log.Println("Error on fetching Github API: ",err)
		http.Error(w, "Github API request failed", 500)
		return 
	}

	defer responseJSON.Body.Close()

	// make an object of RepoList struct
	var repolist []RepoInfo	

	// Decode JSON body to struct
	if err := json.NewDecoder(responseJSON.Body).Decode(&repolist); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
		return
	}

	

	// Making a new map to save the language and number of repos that uses it
	map_of_lang:=make(map[string]float32)

	// Number of non null lang usages
	var num_of_langs int
	
	// Traverse through entire language and makes a map for the usage
	for i,_:=range repolist{
		if repolist[i].Language !="" && repolist[i].IsFork==false{
			num_of_langs++
			map_of_lang[repolist[i].Language]++
		}
	}

	// Its time to calculate the percentage usage of these langs
	for key, value := range map_of_lang {
		map_of_lang[key] = (value * 100.0) / float32(num_of_langs)
	}

	/*
	// This part directly converts map as a json

	// Convert the map to JSON
	jsonData, err := json.Marshal(map_of_lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the http.ResponseWriter
	w.Write(jsonData)
	*/

	// Preparing response
	// Creating an object for struct RepoResponse
	var reporesponse RepoResponse
	//Assigning LanguagePercent field to the previously created map
	reporesponse.LanguagePercent=map_of_lang

	// Converting struct object to json
	jsonResponse,err:=json.Marshal(reporesponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// setting header to let browser know that the response is json
	w.Header().Set("Content-Type", "application/json")
	// writing response
	w.Write(jsonResponse)



}

func main(){
	http.HandleFunc("/github/",githubHandle)

	fmt.Println("Starting server at 8080..")
	http.ListenAndServe(":8080",nil)
}