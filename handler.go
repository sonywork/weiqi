package weiqi

import (
	"github.com/dgf1988/weiqi/h"
	"net/http"
)

//默认处理器。处理首页访问
func handleDefault(w http.ResponseWriter, r *http.Request, args []string) {
	//从会话中获取用户信息，如果没登录，则为nil。
	u := getSessionUser(r)

	err := render_default(w, u)
	if err != nil {
		h.ServerError(w, err)
		return
	}
}

func defaultHtml() *Html {
	return defHtmlLayout().Append(
		defHtmlHead(),
		defHtmlHeader(),
		defHtmlContent(),
		defHtmlFooter(),
	)
}

func defaultData(u *U, posts []P, players []Player, sgfs []Sgf) *Data {
	data := defData()
	data.User = u
	data.Content["Posts"] = posts
	data.Content["Sgfs"] = sgfs
	data.Content["Players"] = players
	return data
}

func render_default(w http.ResponseWriter, u *U) error {
	html := defaultHtml()

	posts, err := Posts.List(40, 0)
	if err != nil {
		return err
	}

	//
	players, err := Posts.List(40, 0)
	if err != nil {
		return err
	}

	//
	sgfs, err := Sgfs.List(40, 0)
	if err != nil {
		return err
	}

	data := defaultData(u, posts, players, sgfs)
	return html.Execute(w, data, defFuncMap)
}
