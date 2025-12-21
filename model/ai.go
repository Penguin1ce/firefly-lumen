package model

import (
	"firflybot/config"
	"net/http"
)

type AIClient struct {
	Config *config.Config
	Http   *http.Client
}

type aiRequest struct {
	Prompt string `json:"prompt"`
}

type aiResponse struct {
	Reply string `json:"reply"`
}
