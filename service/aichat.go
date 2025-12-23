package service

import (
	"context"
	"fireflybot/ai/aimanage"
	"fireflybot/db"
	"log"
)

func AiChatService(ctx context.Context, sid string, question string) string {

	log.Println("收到消息：" + question)
	_, err := db.CreateSession(sid)
	if err != nil {
		return ""
	}
	manage := aimanage.GetGlobalManager()
	helper, err := manage.GetOrCreateHelper(sid)
	if err != nil {
		log.Println(err)
		return "获取 AIHelper 时出错"
	}
	message := helper.Client.CreateMessagesFromTemplate(question)
	resp, err := helper.Client.GenerateResponse(ctx, message)
	if err != nil {
		log.Printf("Client.GenerateResponse failed: %v\n", err)
		return "AI 生成错误"
	}
	log.Println("AI 输出成功")
	err = db.AppendMessage(sid, question, true)
	if err != nil {
		return "用户消息插入数据库出错"
	}
	err = db.AppendMessage(sid, resp.Content, false)
	if err != nil {
		return "AI 消息插入数据库出错"
	}
	return resp.Content
}
