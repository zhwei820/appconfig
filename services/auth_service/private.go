package auth_service

import (
	"github.com/zhwei820/appconfig/utils"
	"fmt"
	"strconv"
	"github.com/zhwei820/appconfig/utils/redisp"
	"github.com/rs/zerolog/log"

	"github.com/zhwei820/appconfig/models"
	"github.com/zhwei820/appconfig/utils/define"
	"encoding/json"
)

// 生成 jwt token, 保存到redis
func genTokenAndSave(user models.StaffUser) (string, error) {
	et := utils.EasyToken{
		Uid:     user.Id,
		Expires: utils.GetExpireTime(),
	}
	token, err := et.GetToken()

	userBytes, err := json.Marshal(user)

	c := redisp.CachePool.Get() // set redis session
	_, err = c.Do("SETEX", token, define.CacheDuration, userBytes)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("user_token_error: username: %s id: %s; %s", user.Username, strconv.Itoa(int(user.Id)), err.Error()))
	}
	return token, err
}
