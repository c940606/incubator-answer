package migrations

import (
	"encoding/json"
	"fmt"

	"github.com/answerdev/answer/internal/entity"
	"xorm.io/xorm"
)

func addThemeAndPrivateMode(x *xorm.Engine) error {
	loginConfig := map[string]bool{
		"allow_new_registrations": true,
		"login_required":          false,
	}
	loginConfigDataBytes, _ := json.Marshal(loginConfig)
	siteInfo := &entity.SiteInfo{
		Type:    "login",
		Content: string(loginConfigDataBytes),
		Status:  1,
	}
	exist, err := x.Get(&entity.SiteInfo{Type: siteInfo.Type})
	if err != nil {
		return fmt.Errorf("get config failed: %w", err)
	}
	if !exist {
		_, err = x.InsertOne(siteInfo)
	}
	return err
}