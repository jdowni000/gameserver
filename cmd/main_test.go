package main

import (
	"net/http"
	"net/http/httptest"
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

func TestGetGame(t *testing.T) {
	req := httptest.NewRequest("GET", "/game?id=${2}", nil)
	rr := httptest.NewRecorder()

	// Call the handler
	GetGame(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "Game: Fallguys\nDescription: Fall Guys Creative is a level editor that allows you to create fiendish custom Rounds and share them with the wider community.\nID: 2\nCurrentPrice: 55\nSellerName: Epic Games store\nDeveloperName: Epic Games devel\nPublisherName: Epic Games publish\nThumbnailURL: https://www.fallguys.com/en-US/download\n"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, rr.Body.String())
	}
}

func TestListGames(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	rr := httptest.NewRecorder()

	// Call the handler
	ListGames(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "List of Games:\nFortnite\nFallguys\n"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, rr.Body.String())
	}
}
