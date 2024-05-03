package api

import (
	"log"
	"time"
	"userauthsystem/core"
	"userauthsystem/services"
)

func SignIn(username string, password string) (string, error) {
	isCorrect, id := services.CallIsUserCorrect(core.GetEnvrionment("HOST_PROFILE"), username, password)
	if isCorrect {
		var accessToken string = services.CreateAcctoken(id)
		core.RedisSetWithExpire(accessToken, id, 36000*time.Second)
		return accessToken, nil
	}
	log.Printf("isCorrect: %v", isCorrect)
	return "", nil
}
