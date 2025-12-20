package service

import (
	"context"
	myconfig "firflybot/config"
	"io"
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

func Test_AI(t *testing.T) {
	t.Setenv("TELEGRAM_BOT_TOKEN", "dummy-token")
	cfg, err := myconfig.LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.AiUrl == "" || cfg.AiKey == "" {
		t.Skip("AI_URL or AI_KEY 未设置，跳过集成测试")
	}

	ctx := context.Background()
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  cfg.AiKey,
		Model:   "gpt-5-chat-latest",
		BaseURL: cfg.AiUrl, // 不填默认走官方
	})
	messages := createMessagesFromTemplate()
	result, err := chatModel.Generate(ctx, messages)
	if err != nil {
		t.Fatalf("chatModel generate error: %v", err)
	}
	turl := "https://api.telegram.org/bot" + cfg.TOKEN + "/sendMessage?chat_id=" + "SESSION" + "=" + url.QueryEscape(result.Content)
	resp, err := http.Post(turl, "", nil)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	t.Log(result)
}

func createTemplate() prompt.ChatTemplate {
	// 创建模板，使用 FString 格式
	return prompt.FromMessages(schema.FString,
		// 系统消息模板
		schema.SystemMessage("你是一个{role}。你需要用{style}的语气回答问题。你的目标是帮助程序员保持积极乐观的心态，提供技术建议的同时也要关注他们的心理健康。"),

		// 插入需要的对话历史（新对话的话这里不填）
		schema.MessagesPlaceholder("chat_history", true),

		// 用户消息模板
		schema.UserMessage("问题: {question}"),
	)
}

func createMessagesFromTemplate() []*schema.Message {
	template := createTemplate()

	// 使用模板生成消息
	messages, err := template.Format(context.Background(), map[string]any{
		"role":     "程序员鼓励师",
		"style":    "积极、温暖且专业",
		"question": "我的代码一直报错，感觉好沮丧，该怎么办？",
		// 对话历史（这个例子里模拟两轮对话历史）
		"chat_history": []*schema.Message{
			schema.UserMessage("你好"),
			schema.AssistantMessage("嘿！我是你的程序员鼓励师！记住，每个优秀的程序员都是从 Debug 中成长起来的。有什么我可以帮你的吗？", nil),
			schema.UserMessage("我觉得自己写的代码太烂了"),
			schema.AssistantMessage("每个程序员都经历过这个阶段！重要的是你在不断学习和进步。让我们一起看看代码，我相信通过重构和优化，它会变得更好。记住，Rome wasn't built in a day，代码质量是通过持续改进来提升的。", nil),
		},
	})
	if err != nil {
		log.Fatalf("format template failed: %v\n", err)
	}
	return messages
}
