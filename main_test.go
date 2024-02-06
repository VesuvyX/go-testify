package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenStatusOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	require.Equal(t, responseRecorder.Code, http.StatusOK)
	require.NotEmpty(t, responseRecorder.Body)

}
func TestMainHandlerWhenCityNotExpected(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=novosibirsk", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)

	require.Equal(t, "wrong city value", responseRecorder.Body.String())
}
func TestMainHandlerWhenCountMoreThenTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	require.Equal(t, responseRecorder.Code, http.StatusOK)
	require.NotEmpty(t, responseRecorder.Body)

	cafes := strings.Split(responseRecorder.Body.String(), ", ")
	assert.Len(t, len(cafes), totalCount)
}
