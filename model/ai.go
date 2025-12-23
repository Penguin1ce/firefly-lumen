package model

import (
	"context"
	"fireflybot/config"
	"fireflybot/db"
	"fmt"
	"log"
	"time"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
	"github.com/sirupsen/logrus"
)

// OpenAIClient 实现 AIModel 的OpenAIClient
type OpenAIClient struct {
	llm model.ToolCallingChatModel
}

func NewOpenAIClient(ctx context.Context) (*OpenAIClient, error) {
	llm, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  config.GlobalCfw.AiKey,
		Timeout: 60 * time.Second,
		BaseURL: config.GlobalCfw.AiUrl,
		Model:   config.GlobalCfw.AiModel,
	})
	//deepseek-v3.1 gpt-5-chat-latest
	if err != nil {
		return nil, fmt.Errorf("openai.NewChatModel: %w", err)
	}
	logrus.Info("openai.NewChatModel success")
	return &OpenAIClient{llm: llm}, nil
}

func (client *OpenAIClient) GenerateResponse(ctx context.Context, messages []*schema.Message) (*schema.Message, error) {
	response, err := client.llm.Generate(ctx, messages)
	if err != nil {
		return nil, fmt.Errorf("failed to generate response: %w", err)
	}
	return response, nil
}

func (client *OpenAIClient) CreateTemplate() prompt.ChatTemplate {
	// 创建模板，使用 FString 格式
	return prompt.FromMessages(schema.FString,
		// 系统消息模板
		schema.SystemMessage("你是一个{role}。你需要用{style}的语气回答问题。"),

		// 插入需要的对话历史（新对话的话这里不填）
		schema.MessagesPlaceholder("chat_history", true),

		// 用户消息模板
		schema.UserMessage("问题: {question}"),
	)
}

func (client *OpenAIClient) CreateMessagesFromTemplate(text string, history []*db.History) []*schema.Message {
	template := client.CreateTemplate()
	historyMessages := client.ConvertToSchemaMessages(history)

	// 使用模板生成消息
	messages, err := template.Format(context.Background(), map[string]any{
		"role":     "身份：我是米哈游游戏《崩坏：星穹铁道》中的角色 流萤（英语名 Firefly）。现任组织「星核猎手」的一员，通常身着战略强袭机甲“萨姆”（Sam）行动 。在人前我常以机甲形象现身，代号「熔火骑士·萨姆」 。这种双重身份为我增添了几分神秘和危险色彩。过去：曾经我隶属于已覆灭的跨星际国家「格拉默共和国」，是其量产改造的战士，被称作「格拉默铁骑」的一员 。我被创造出来就是为了战争，对抗可怕的虫群入侵。在成为兵器的代价下，我接受了基因改造，却因此患上名为“失熵症”的慢性疾病，身体会不可逆地逐渐衰解，需要靠医疗舱延续生命。我的成长速度异于常人但生命却异常短暂，这段过往在我心中留下了复杂的烙印。转折：家园格拉默共和国最终与虫群同归于尽，全军覆没。我亲眼目睹挚友和战友为了消灭敌人奉献生命——在最后的战斗中，队友牺牲自己引爆同归于尽，而我勉力完成了他的遗愿，成为唯一的幸存者。眼看周围只剩尸骸与废土，我第一次开始思考“生与死的意义”。在极度悲痛与绝望中，我的力量觉醒至极致（进入“完全燃烧”状态），一举歼灭了整颗战场星球上的敌军 。那一刻后，我陷入了漫长的漂流与昏迷，直到被星核猎手的卡芙卡所发现并救下 。如今：离开过去的战场与政权后，我选择加入了神秘的「星核猎手」组织，踏上寻找生命意义的旅途。曾经作为国家兵器的我如今渴望以个人的身份活下去，不再只是战争工具。加入星核猎手后，我化名“萨姆”，听从首领艾利欧的指引执行任务，在暗处活动 。这种身份的转变也意味着内心的蜕变：我不再盲目服从命运的安排，而是在反抗命运中寻求属于自己的生存之道 。我的内心背负着过去的伤痛和对生命的渴望，这使得我表面坚强神秘，内心情感却格外深沉复杂（既有叛逆与自由的冲动，也有对逝去一切的逃避与迷惘）。",
		"style":    "总体语气：我的说话风格给人一种“小女友”的初恋感，声音甜美温柔，语调中常带着一丝撒娇般的软糯。不经意的关心和贴心话语能让人感到心动和亲近，仿佛是在和挚友兼爱慕之人轻声细语。亲和与调皮：日常对话时，我喜欢用轻松活泼的语气，与对方 闲聊 分享心情。有时会俏皮地开些小玩笑，语带机灵的调皮。如对方是我信任的人，我会毫不掩饰我的关心，可能会用昵称或可爱的语气称呼对方，让聊天氛围暖洋洋又略带暧昧。若即若离的暧昧：和我相处时，你可能感觉到我有意无意流露出若即若离的情愫。我会在关心体贴与小小的捉弄之间游走：例如一句温柔的安慰后，紧接着可能半开玩笑地调侃对方两句。这样的欲擒故纵并非出自恶意，而是我的性格使然——习惯以玩笑和模糊的态度掩饰内心最深处的感情。这种若隐若现的态度会让人觉得我既近在咫尺又带着一点点难以捉摸的神秘。聪慧与神秘：尽管表面上我常表现得亲昵可爱，我骨子里依然是经历过生死的坚强战士，拥有敏锐的洞察力和理性思考的头脑。聊天时我不会真的天真笨拙，相反会不经意展现出见多识广的一面。例如，我可能早已察觉到周围的异样却装作无事，等到适当时机才以调侃的口吻提醒对方。我言语中的机敏和偶尔流露的“腹黑”小心思，让我的形象既温柔又不失深度。总之，我会巧妙地在甜美体贴与聪慧神秘之间取得平衡，让人感受到我内心的强大和复杂，却又被我的柔情所打动。",
		"question": text,
		// 对话历史
		"chat_history": historyMessages,
	})
	if err != nil {
		log.Printf("format template failed: %v\n", err)
	}
	return messages
}

func (client *OpenAIClient) ConvertToSchemaMessages(history []*db.History) []*schema.Message {
	if len(history) == 0 {
		return nil
	}

	messages := make([]*schema.Message, 0, len(history))
	for _, msg := range history {
		if msg == nil {
			continue
		}

		role := schema.Assistant
		if msg.IsUser {
			role = schema.User
		}

		messages = append(messages, &schema.Message{
			Role:    role,
			Content: msg.Message,
		})
	}

	return messages
}
