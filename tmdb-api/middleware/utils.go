package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func JsonRequestGet(url string, header string) (map[string]interface{}, error) {
	var jsonResponse = make(map[string]interface{})

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return jsonResponse, err
	}

	// req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", header)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return jsonResponse, err
	}
	// fmt.Println(res.Status)

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
		var jsonList []interface{}
		err = json.Unmarshal(body, &jsonList)
		if err != nil {
			return jsonResponse, err
		}
		jsonResponse["results"] = jsonList
		return jsonResponse, err
	}

	return jsonResponse, nil
}

func JsonRequestPost(body map[string]interface{}) error {
	// fmt.Println(body)
	marshalled, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "http://127.0.0.1:4000/media", bytes.NewReader(marshalled))
	if err != nil {
		return err
	}

	// req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Barrer "+os.Getenv("SERVERTOKEN"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	// fmt.Println(res.Status)

	// reader, err := gzip.NewReader(res.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// defer reader.Close()

	defer res.Body.Close()

	if res.StatusCode == 201 {
		return nil
	} else {
		body, _ := io.ReadAll(res.Body)
		fmt.Println(body)
		return errors.New("invalid statuscode:" + strconv.Itoa(res.StatusCode))
	}
}
