package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func putToStream(id string, value int, timestamp int64) {
	payload := map[string]interface{}{
		"random":    value,
		"timestamp": timestamp,
		"id":        id,
	}

	response, _ := json.Marshal(payload)

	file, err := os.OpenFile("logfile.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(string(response) + "\n")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	rand.NewSource(time.Now().UnixNano())

	for {
		value := rand.Intn(100) + 1
		timestamp := time.Now().Unix()
		id := "stream-1"

		putToStream(id, value, timestamp)

		// Wait for 5 seconds
		time.Sleep(5 * time.Second)
	}
}
