package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const hitokotoUrl = "https://v1.hitokoto.cn/"

type Hitokoto struct {
	ID         int    `json:"id"`
	UUID       string `json:"uuid"`
	Hitokoto   string `json:"hitokoto"`
	Type       string `json:"type"`
	From       string `json:"from"`
	FromWho    string `json:"from_who"`
	Creator    string `json:"creator"`
	CreatorUID int    `json:"creator_uid"`
	Reviewer   int    `json:"reviewer"`
	CommitFrom string `json:"commit_from"`
	CreatedAt  string `json:"created_at"`
	Length     int    `json:"length"`
}

// GetHitokoto is a function that returns a random Hitokoto object and an error from v1.hitokoto.cn
func GetHitokoto(url string) (hitokoto Hitokoto, err error) {
	// Get the response from the URL
	response, err := http.Get(url)
	if err != nil {
		return hitokoto, err
	}
	// Close the response body when the function returns
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return hitokoto, err
	}

	// Unmarshal the response body into the Hitokoto object
	err = json.Unmarshal(body, &hitokoto)
	if err != nil {
		return hitokoto, err
	}

	return hitokoto, nil
}

// Get a random Hitokoto object by params c
func GetHitokotoByParams(c rune) (hitokoto Hitokoto, err error) {
	url := fmt.Sprintf("%s?c=%c", hitokotoUrl, c)
	return GetHitokoto(url)
}
