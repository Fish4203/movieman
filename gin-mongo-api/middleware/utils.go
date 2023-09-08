package middleware 

import (
	"net/http"
	"io"
    "encoding/json"
    "os"
)


func JsonRequest(url string) (map[string]interface{}, error) {
    var jsonResponse map[string]interface{}

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return jsonResponse, err
    }

    // req.Header.Add("Accept-Encoding", "gzip, deflate, br")
    req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer " + os.Getenv("TMDB"))


    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return jsonResponse, err
    }


    // reader, err := gzip.NewReader(res.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// defer reader.Close()

    defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
    if err != nil {
        return jsonResponse, err
    }

    err = json.Unmarshal(body, &jsonResponse)
    if err != nil {
        return jsonResponse, err
    }

    return jsonResponse, nil
}