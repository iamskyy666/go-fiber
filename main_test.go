package main

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestCheckToken(t *testing.T){
	app:=InitApp()

	// Test req. without token 🟠
	req:=httptest.NewRequest("GET","check-token",nil)
	req.Header.Set("Authorization","")
	resp, _:=app.Test(req)
	if resp.StatusCode != fiber.StatusUnauthorized{
		t.Errorf("Expected Status: %d, Got: %d",fiber.StatusUnauthorized,resp.StatusCode)
	}


	// Test req. with token 🔵
	req=httptest.NewRequest("GET","check-token",nil)
	req.Header.Set("Authorization","Bearer mock-token")
	resp, _=app.Test(req)
	if resp.StatusCode != fiber.StatusOK{
		t.Errorf("Expected Status: %d, Got: %d",fiber.StatusUnauthorized,resp.StatusCode)
	}
}