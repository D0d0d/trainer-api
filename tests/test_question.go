package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddQuestion(t *testing.T) {
	body := `{
        "question": "Какие языки подходят для фронтенда?",
        "answers": "[\"HTML\", \"CSS\", \"JavaScript\"]",
        "correct": "[0, 2]"
    }`

	req, _ := http.NewRequest("POST", "/api/questions", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.NotNil(t, response["data"])
	assert.Equal(t, response["data"].(map[string]interface{})["correct"], "[0,2]")
}
