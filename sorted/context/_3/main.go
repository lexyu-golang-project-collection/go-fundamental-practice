package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"time"
)

const (
	KEY = "trace_id"
)

func NonComplianceUid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	fmt.Println(uuid)
	return uuid
}
func NewRequestID() string {
	panic("")
	// return strings.Replace(uuid.New().String(), "-", "", -1)
}

func NewContextWithTraceID() context.Context {
	ctx := context.WithValue(context.Background(), KEY, NonComplianceUid())

	// ctx := context.WithValue(context.Background(), KEY, NewRequestID())

	return ctx
}

func PrintLog(ctx context.Context, message string) {
	fmt.Printf("%s | info | trace_id=%s | %s", time.Now().Format("2006-01-02 15:04:05"), GetContextValue(ctx, KEY), message)
}

func GetContextValue(ctx context.Context, k string) string {
	v, ok := ctx.Value(k).(string)
	if !ok {
		return ""
	}
	return v
}

func ProcessEnter(ctx context.Context) {
	PrintLog(ctx, "Go Context !")
}

func main() {

	// out, err := exec.Command("uuidgen").Output()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	ProcessEnter(NewContextWithTraceID())
}
