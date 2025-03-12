package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TranslateRequest struct {
	Text       string `json:"text"`
	TargetLang string `json:"target_lang"`
}

type TranslateResponse struct {
	TranslatedText     string `json:"translated_text"`
	DetectedSourceLang string `json:"detected_source_lang"`
}

func translateHandler(c *gin.Context) {
	var req TranslateRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	payload, _ := json.Marshal(req)
	resp, err := http.Post("http://translator-service:8000/translate", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("HTTP request error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Translation service unavailable"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading translation response"})
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Translation service error: %s", string(body))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Translation service returned error"})
		return
	}

	var translateResp map[string]interface{}
	json.Unmarshal(body, &translateResp)

	c.JSON(http.StatusOK, translateResp)
}

func main() {
	router := gin.Default()
	router.POST("/api/v1/translate", translateHandler)
	router.Run(":8080")
}
