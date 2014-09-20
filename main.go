package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	key      string
	token    string
	listId   string
	filepath string
)

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}

func registerFlag() {
	flag.StringVar(&key, "key", "", "application key")
	flag.StringVar(&key, "token", "", "application key")
	flag.StringVar(&listId, "l", "", "list id")
	flag.StringVar(&filepath, "csv", "", "csv filepath")
	flag.Parse()
}

func main() {

	registerFlag()

	file, err := os.Open(filepath)
	failOnError(err)
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read() // 1行読み出す
		if err == io.EOF {
			break
		} else {
			failOnError(err)
		}

		if len(record) == 0 {
			continue
		}

		name := record[0]
		var desc string
		if len(record) == 2 {
			desc = record[1]
		}

		client := &http.Client{}
		data := url.Values{
			"key":   {key},
			"token": {token},
			"name":  {name},
			"desc":  {desc},
		}

		url := "https://api.trello.com/1/lists/" + listId + "/cards"

		resp, _ := client.Post(
			url,
			"application/x-www-form-urlencoded",
			strings.NewReader(data.Encode()),
		)
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		fmt.Println(string(body))
	}
}
