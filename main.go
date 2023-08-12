package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {

	url := "https://api.themoviedb.org/3/search/multi?query=star%20trek&include_adult=false&language=en-US&page=1"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJiODg2NGVhNGE4ZDYxMDMyNWQxMjgzODYzNWE5YTRmNyIsInN1YiI6IjY0ZDA4NmRmNGQ2NzkxMDBlMjQxMGQyZiIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.pFbR6640emAdm26k7wj7yQ46Q1-EiVjottAazjbkGNo")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

}
