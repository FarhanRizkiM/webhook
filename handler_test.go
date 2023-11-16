package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aiteung/module/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRandomString(t *testing.T) {
	stringssss := GetRandomString(stringreply)
	fmt.Println(stringssss)
}

func TestPostBalasan(t *testing.T) {
	// Mock the request data
	mockMsg := model.IteungMessage{
		Message:      "Hello",
		Phone_number: "123456789",
		Latitude:     37.7749,
		Longitude:    -122.4194,
		Alias_name:   "John",
		LiveLoc:      false,
	}

	// Encode the mock message to JSON
	jsonData, err := json.Marshal(mockMsg)
	if err != nil {
		t.Fatal(err)
	}

	// Create a request with the mock data
	req, err := http.NewRequest("POST", "/post-balasan", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	// Set the Secret header
	req.Header.Set("Secret", "bebas")

	// Create a response recorder to capture the response
	res := httptest.NewRecorder()

	// Call the handler function
	PostBalasan(res, req)

	// Check the status code
	if res.Code != http.StatusOK {
		t.Errorf("Expected status 200; got %d", res.Code)
	}

	// You can add more assertions based on your specific use case
	// For example, you can check the response body or headers.
}

// You can create similar test functions for other functions like Liveloc, GetRandomString, etc.
