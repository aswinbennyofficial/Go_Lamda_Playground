package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Struct for the response
type RepoResponse struct{
	LanguagePercent map[string]float32 `json:"language_percent"`
}


// Struct to save the repo info for the recent commit feature
type Repo struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Url string `json:"Url"`
}

// Struct to store the payload field in the recent commit feature
type Payload struct{

}

// struct used to save the latest commit and related from an account
type CommitInfo struct{
	Id int `json:"id"`
	Type string `json:"type"`
	Repoinfo Repo `json:"repo"`
	Payload Payload `json:"payload"`

}



func commitHandle(w http.ResponseWriter, r *http.Request){
	// fetch username from <url>/github/<username>
	path:=r.URL.Path
	path= strings.Replace(path, "/github/", "", 1)
	// trim away succeeding '/' if exists 
	path= strings.ReplaceAll(path, "/", "")



	url:="https://api.github.com/users/"+path+"/events/public"
	responseJSON,err:=http.Get(url)
	if err!=nil{
		log.Println("Error on fetching Github API: ",err)
		http.Error(w, "Github API request failed", 500)
		return 
	}



	// RESPONSE
	// Converting struct object to json
	dataJSON,err:=json.Marshal(reporesponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// setting header to let browser know that the response is json
	w.Header().Set("Content-Type", "application/json")
	// writing response
	w.Write(dataJSON)

}

func main(){
	
	http.HandleFunc("/commit/",commitHandle)

	fmt.Println("Starting server at 8080..")
	http.ListenAndServe(":8080",nil)
}