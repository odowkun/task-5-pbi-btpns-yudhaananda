package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/odowkun/task-5-pbi-btpns-yudhaananda/database"
	"github.com/odowkun/task-5-pbi-btpns-yudhaananda/models"
	"github.com/odowkun/task-5-pbi-btpns-yudhaananda/router"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int()

	database.InitDB()

	api := "/api/v1"
	r := router.RouteInit()

	randUsername := fmt.Sprintf("Kadek Yudha Ananda Putra %d", randNum)
	randEmail := fmt.Sprintf("dummy%d@gmail.com", randNum)
	newUser := models.User{
		Username: randUsername,
		Email:    randEmail,
		Password: "rootuser",
	}
	jsonValue, _ := json.Marshal(newUser)

	req, _ := http.NewRequest("POST", api+"/users/register", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
