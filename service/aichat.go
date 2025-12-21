package service

import (
	myconfig "firflybot/config"

	"github.com/sirupsen/logrus"
)

func AiChatService(config *myconfig.Config) string {
	logrus.Info(config.TOKEN)
	return "404"
}
