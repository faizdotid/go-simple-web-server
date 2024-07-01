package test

import (
	"go-go/app/service"
	"testing"
)

func TestEncodeJwt(t *testing.T) {
	lol := map[string]interface{}{
		"username": "admin",
		"role":     "admin",
	}
	token, err := service.EncrytData(lol)
	if err != nil {
		t.Error(err)
	}
	t.Log(token)

}
