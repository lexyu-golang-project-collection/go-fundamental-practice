package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	jsonFile, err := os.Open("conf.json")

	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	jsonStr, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	doc := make(map[string]interface{})

	if err := json.Unmarshal(jsonStr, &doc); err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc)

	fmt.Println(doc["KEY"])

	fmt.Println(doc["SECRET_KEY"])

	for k, v := range doc {
		fmt.Printf("%s=%s\n", k, v.(string))
	}

}
