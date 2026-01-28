package internal

import (
	"cli-flix/internal/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand/v2"
	"net/http"
)

func GetRandomRecommendation(apiKey string) string {
	random_page := rand.IntN(9) + 1
	url := fmt.Sprintf("https://api.themoviedb.org/3/discover/movie?include_adult=false&include_video=false&language=en-US&page=%d&sort_by=popularity.desc&with_original_language=en", random_page)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Add("accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// fmt.Println(string(body))
	jsonDataBytes := []byte(body)
	var response models.TheMovieDBResponse
	err = json.Unmarshal(jsonDataBytes, &response)
	fmt.Println(response.Page)
	movieSlice := len(response.Results)

	fmt.Println(response.Results[rand.IntN(movieSlice-1)].Title)
	return "Random"
}
