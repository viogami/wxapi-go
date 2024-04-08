package menu

var URL_MENU_CREATE = "https://api.weixin.qq.com/cgi-bin/menu/addconditional?access_token=ACCESS_TOKEN"
var URL_MENU_DELETE = "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS"
var URL_MENU_GET = "https://api.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS"

// Menu 菜单
type Menu struct {
	Button []Button `json:"button"`
}

// CreateMenu 创建菜单
func CreateMenu(menu Menu) error {
	return nil
}

// DeleteMenu 删除菜单
func DeleteMenu() error {
	return nil
}

// GetMenu 获取菜单
func GetMenu() (Menu, error) {
	return Menu{}, nil
}
