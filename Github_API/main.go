package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type RepoResponse struct{
	Language map[string]float32

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

	// fmt.Println(repolist[0].Name)
	// fmt.Println(repolist[0].Url)
	// fmt.Println(repolist[0].IsFork)
	// fmt.Println(repolist[0].Language)
	// fmt.Println(repolist[0].Id)

	// Making a new map to save the language and number of repos that uses it
	map_of_lang_with_number:=make(map[string]int)

	// Traverse through entire language and makes a map
	for i,_:=range repolist{
		if repolist[i].Language !=""{
			map_of_lang_with_number[repolist[i].Language]++
		}
	}

	for key,value:=range map_of_lang_with_number{
		fmt.Println(key," ",value)
	}



}

func main(){
	http.HandleFunc("/github/",githubHandle)

	fmt.Println("Starting server at 8080..")
	http.ListenAndServe(":8080",nil)
}