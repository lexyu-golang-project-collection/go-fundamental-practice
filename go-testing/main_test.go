package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	got := strings.Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("No Equal!")
	} else {
		fmt.Println("Equal!")
	}
}
