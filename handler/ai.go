package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type aiClient struct {
	url    string
	apiKey string
	http   *http.Client
}

type aiRequest struct {
	Prompt string `json:"prompt"`
}

type aiResponse struct {
	Reply string `json:"reply"`
}

func (c *aiClient) call(ctx context.Context, prompt string) (string, error) {
	payload, err := json.Marshal(aiRequest{Prompt: prompt})
	if err != nil {
		return "", fmt.Errorf("marshal ai request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url, bytes.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("build ai request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return "", fmt.Errorf("call ai api: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return "", fmt.Errorf("ai api status %d", resp.StatusCode)
	}

	var body aiResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return "", fmt.Errorf("decode ai response: %w", err)
	}

	if body.Reply == "" {
		return "", fmt.Errorf("ai 返回空内容")
	}

	return body.Reply, nil
}

// NewAIHandler 返回一个处理 Telegram 文本消息的 handler。
// 它会把用户输入转发给 AI 接口，并将回复再发回用户。
func NewAIHandler(aiURL, aiKey string) bot.HandlerFunc {
	cli := &aiClient{
		url:    aiURL,
		apiKey: aiKey,
		http:   &http.Client{Timeout: 10 * time.Second},
	}

	return func(ctx context.Context, b *bot.Bot, upd *models.Update) {
		msg := upd.Message
		if msg == nil || msg.Text == "" {
			return
		}

		reply, err := cli.call(ctx, msg.Text)
		if err != nil {
			log.Println("ai handler error:", err)
			reply = "抱歉，AI 暂时无法回复，请稍后重试。"
		}

		_, sendErr := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: msg.Chat.ID,
			Text:   reply,
		})
		if sendErr != nil {
			log.Println("send message error:", sendErr)
		}
	}
}
