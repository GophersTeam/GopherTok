package test

import (
	"GopherTok/common/consts"
	"GopherTok/common/errorx"
	"GopherTok/common/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestJWT(t *testing.T) {
	AccessToken, RefreshToken := utils.GetToken(int64(9), uuid.New().String(), "xxx", "zzz")

	token := AccessToken + " " + RefreshToken
	fmt.Println(token)
}

func TestTokenValidationHandler(t *testing.T) {
	// Your token here for testing
	mockToken :=
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6OSwic3RhdGUiOiIwMGM5ZGVjMy1hNDcwLTQ3MzQtODU4NC1iYzIxNmE0MjEyYTEiLCJleHAiOjE2OTA3MzUyNDgsImlhdCI6MTY5MDczNTA2OCwiaXNzIjoiQVIifQ.jINaRT7uHejHONkj3EethqZqrmWlzI_fqAYOc69MX0o eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6OSwic3RhdGUiOiIwMGM5ZGVjMy1hNDcwLTQ3MzQtODU4NC1iYzIxNmE0MjEyYTEiLCJleHAiOjE2OTMzMjcwNjgsImlhdCI6MTY5MDczNTA2OCwiaXNzIjoiUlQifQ.FYHex1C9DVoa3QMPoPLfuAhDhlBpjt5SWDWFVqPQwE8"

	// Create a new HTTP request with the token
	req, err := http.NewRequest("GET", "http://example.com/some/path", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Form = make(map[string][]string)
	req.Form.Add("token", mockToken)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the function under test with the mocked request and response
	TokenValidationHandler(rr, req)

	// Check the response status code and body
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d; got %d", http.StatusOK, rr.Code)
	}

	// Add further assertions if needed based on the expected behavior of the token validation
}

func TokenValidationHandler(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	if token == "" {
		token = r.PostFormValue("token")
		if token == "" {
			http.Error(w, "Token is empty", http.StatusUnauthorized)
			return
		}
	}

	parts := strings.Split(token, " ")
	if len(parts) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"code": 30002, "message": "Invalid token format"}`))
		log.Fatal(errors.New("part !=2"))
		return
	}
	parseToken, isExpire, err := utils.ParseToken(parts[0], parts[1], "xx", "zzz")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err, _ := json.Marshal(errorx.NewCodeError(30003, errorx.ErrTokenProve))
		log.Fatal(err)
		w.Write(err)
		return
	}
	fmt.Println(parts)
	fmt.Println(isExpire)
	if isExpire {
		parts[0], parts[1] = utils.GetToken(parseToken.ID, parseToken.State, "xxx", "zzz")
		//w.Header().Set("Authorization", fmt.Sprintf("Bearer %s %s", parts[0], parts[1]))

	}
	token = parts[0] + " " + parts[1]
	r = r.WithContext(context.WithValue(r.Context(), consts.UserId, parseToken.ID))
	r = r.WithContext(context.WithValue(r.Context(), consts.Token, token))
	fmt.Println(parseToken.ID, "dd", token)
	// Your token validation logic goes here...
	// You can add further testing for different cases in this function
}
