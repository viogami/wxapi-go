package oaHandle

import (
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2/officialaccount"
	m "github.com/silenceper/wechat/v2/officialaccount/menu"
)

// CreateMenu 创建菜单
func CreateMenu(c *gin.Context, oa *officialaccount.OfficialAccount) {
	buttons := make([]*m.Button, 0)
	// 一级菜单
	btn1 := m.NewClickButton("bot信息", "vio--made by viogami")
	btn2 := m.NewViewButton("作者博客", "http://viogami.me/")
	buttons = append(buttons, btn1, btn2)

	err := oa.GetMenu().SetMenu(buttons)
	if err != nil {
		c.JSON(500, "菜单创建失败"+err.Error())
		return
	}
	c.JSON(200, "菜单创建成功")
}

// 菜单检查
func CheckMenu(c *gin.Context, oa *officialaccount.OfficialAccount) error {
	menu, err := oa.GetMenu().GetMenu()
	if menu.Menu.Button == nil {
		CreateMenu(c, oa)
	} else {
		c.JSON(500, "菜单已存在, 请勿重复创建")
	}
	return err
}

// DeleteMenu 删除菜单
func DeleteMenu(oa *officialaccount.OfficialAccount) error {
	oa.GetMenu().DeleteMenu()
	return nil
}

// AddConditionalMenu 添加个性化菜单
func AddConditionalMenu(oa *officialaccount.OfficialAccount, buttons []*m.Button, matchRule *m.MatchRule) error {
	oa.GetMenu().AddConditional(buttons, matchRule)
	return nil
}
