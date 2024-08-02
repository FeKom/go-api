package service

import (
	"encoding/json"
	"io"
	"net/http"

	au "github.com/fekom/go-api/apiUrlTrataments"
	j "github.com/fekom/go-api/jsonDT"
)

type SearchResponse struct {
	errorMessage string
	Response     j.SearchApiResponse
}

func SearchByName(name string) SearchResponse {
	ApiURL := au.ApiSearchBaseUrl + name + "?api_token=" + au.ApiToken
	Resp, err := http.Get(ApiURL)
	if err != nil {
		return SearchResponse{errorMessage: "failed to fetch data from external API", Response: j.SearchApiResponse{}}
	}
	defer Resp.Body.Close()
	if Resp.StatusCode != http.StatusOK {
		return SearchResponse{errorMessage: "Received non-200 response from external API", Response: j.SearchApiResponse{}}
	}
	Data, err := io.ReadAll(Resp.Body)
	if err != nil {
		return SearchResponse{errorMessage: "Failed to read response body", Response: j.SearchApiResponse{}}
	}
	var Target j.SearchApiResponse
	err = json.Unmarshal(Data, &Target)
	if err != nil {
		return SearchResponse{errorMessage: "Failed to parse response", Response: j.SearchApiResponse{}}
	}
	return SearchResponse{errorMessage: "", Response: Target}
}
