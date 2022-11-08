package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestNotUser : ユーザ情報が不正な場合のテスト(関数テスト)
func TestNotUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	// エラーが発生し、ステータスコードが401であることを確認
	if assert.Error(t, Login()(c)) {
		assert.Equal(t, echo.ErrUnauthorized, Login()(c))
	}
}

// TestUserCorrect : ユーザ情報が正しい場合のテスト(関数テスト)
func TestUserCorrect(t *testing.T) {
	const USERNAME = "admin"
	const PASSWORD = "admin"

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	req.Form = map[string][]string{
		"username": {USERNAME},
		"password": {PASSWORD},
	}
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	// エラーが発生しないことを確認
	if assert.NoError(t, Login()(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	// レスポンスの取得
	var loginResponse LoginResponse
	if assert.NoError(t, json.NewDecoder(rec.Body).Decode(&loginResponse)) {
		assert.NotEmpty(t, loginResponse.Token)
	}
}

// LoginResponse : ログイン時のレスポンス
type LoginResponse struct {
	Token string `json:"token"`
}

type RestrictedResponse struct {
	Message string `json:"message"`
}

// TestRestricted : トークンのテスト(HTTP テスト)
func TestRestricted(t *testing.T) {
	const HOST = "http://localhost:1323"
	const USERNAME = "admin"
	const PASSWORD = "admin"

	// ログイン
	resp, err := http.PostForm(HOST+"/login", map[string][]string{
		"username": {USERNAME},
		"password": {PASSWORD},
	})
	if err != nil {
		t.Log("サーバが起動していない可能性があります。")
		t.Log("make run_server でサーバを起動してからテストを実行してください。")
		t.Fatal(err)
	}

	// レスポンスの取得
	var loginResponse LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResponse); err != nil {
		t.Fatal(err)
	}

	// トークンの取得
	token := loginResponse.Token

	// トークンを使用してリクエスト
	req, err := http.NewRequest(http.MethodGet, HOST+"/restricted/hello", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := new(http.Client)
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// レスポンスを取得
	var restrictedResponse RestrictedResponse
	if err := json.NewDecoder(resp.Body).Decode(&restrictedResponse); err != nil {
		t.Fatal(err)
	}

	// レスポンスの確認
	assert.Equal(t, "Hello "+USERNAME+"!", restrictedResponse.Message)
}
