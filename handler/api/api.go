package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)




func AmbilAPI(c *gin.Context) {
	// Membuat permintaan GET ke URL API
	resp, err := http.Get("https://pokeapi.co/api/v2/ability")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch API"})
		return
	}
	defer resp.Body.Close()

	// Memparsing respon sebagai JSON
	var apiResponse model.APIResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse API response"})
		return
	}

	c.JSON(http.StatusOK, apiResponse)
}