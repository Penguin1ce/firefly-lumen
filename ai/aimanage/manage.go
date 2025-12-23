package aimanage

import (
	"context"
	"fireflybot/ai/aihelper"
	"fireflybot/model"
	"fmt"
	"log"
	"sync"
)

var ctx = context.Background()
var GlobalManage *AIManage
var once sync.Once

type AIManage struct {
	helpers map[string]*aihelper.AIHelper // map[会话ID]*AIHelper
	mu      sync.RWMutex
}

func NewAIManage() *AIManage {
	return &AIManage{
		helpers: make(map[string]*aihelper.AIHelper),
	}
}
func GetGlobalManager() *AIManage {
	once.Do(func() {
		GlobalManage = NewAIManage()
	})
	return GlobalManage
}

func (am *AIManage) GetOrCreateHelper(sid string) (*aihelper.AIHelper, error) {
	am.mu.Lock()
	helper, exists := am.helpers[sid]
	am.mu.Unlock()
	if exists {
		return helper, nil
	}
	// 创建一个 AIHelper
	am.mu.Lock()
	defer am.mu.Unlock()
	Client, err := model.NewOpenAIClient(ctx)
	if err != nil {
		log.Printf("model.NewOpenAIClient failed: %v\n", err)
		return nil, fmt.Errorf("新建 AIHelper 时，新建 model 失败")
	}
	var newHelper = aihelper.NewAIHelper(Client, sid)
	am.helpers[sid] = newHelper
	return newHelper, nil
}

func (am *AIManage) GetHelperBySid(sid string) (*aihelper.AIHelper, error) {
	am.mu.RLock()
	defer am.mu.RUnlock()
	helper, exists := am.helpers[sid]
	if !exists {
		return nil, fmt.Errorf("该 SID 对应的 Helper 不存在")
	}
	return helper, nil
}
