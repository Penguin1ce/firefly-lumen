package aimanage

import (
	"fireflybot/ai/aihelper"
	"sync"
)

type AIManage struct {
	helpers map[string]*aihelper.AIHelper // map[会话ID]*AIHelper
	mu      sync.RWMutex
}
