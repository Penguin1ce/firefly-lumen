package service

import (
	"context"
	"fireflybot/model"
	"log"
)

func AiChatService(ctx context.Context, question string) string {
	Client, err := model.NewOpenAIClient(ctx)
	if err != nil {
		log.Printf("model.NewOpenAIClient failed: %v\n", err)
		return "错误"
	}
	log.Println("收到消息：" + question)
	message := Client.CreateMessagesFromTemplate(question)
	resp, err := Client.GenerateResponse(ctx, message)
	if err != nil {
		log.Printf("Client.GenerateResponse failed: %v\n", err)
		return "AI 生成错误"
	}
	log.Println("AI 输出成功")
	return resp.Content
}
