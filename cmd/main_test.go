package main

import (
	"reflect"
	"testing"
)

func TestJsonRootGameInfo(t *testing.T) {
	// Get json data
	resp, err := JsonRootGameInfo("games.json")
	if err != nil {
		t.Error(err)
	}
	// Verify struct type
	if reflect.TypeOf(resp) != reflect.TypeOf(RootGameInfo{}) {
		t.Fatalf("Not of struct type RootGameInfo")
	}
}

func TestJsonGameInfo(t *testing.T) {
	// Get json data
	resp, err := JsonGameInfo("games.json")
	if err != nil {
		t.Error(err)
	}
	// Verify struct type
	if reflect.TypeOf(resp) != reflect.TypeOf(GameInfo{}) {
		t.Fatalf("Not of struct type GameInfo")
	}
}
