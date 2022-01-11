package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wagnerww/go-gin-recipes-api.git/models"
)

func TestListRecipesHandler(t *testing.T) {

	url := "http://localhost:8080"

	resp, err := http.Get(fmt.Sprintf("%s/recipes", url))
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	var recipes []models.Recipe
	json.Unmarshal(data, &recipes)
	assert.Equal(t, len(recipes), 10)
}
