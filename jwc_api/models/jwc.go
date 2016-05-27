package models

import (
	"net/http"

	"github.com/scbizu/Zafu_jwcInterface/jwc_api/jwcpkg"
)

var T_cookies []*http.Cookie

const (

	//模拟登陆第一个入口地址
	login_url_gate0 string = "http://210.33.60.8/"
	//模拟登陆第一个入口验证码地址
	vrcode_url_gate0 string = "http://210.33.60.8/CheckCode.aspx"
	//首页地址
	logged_url string = "http://210.33.60.8/xs_main.aspx?xh=201305070123"
	//默认登录页
	default_url string = "http://210.33.60.8/default2.aspx"
	//课程表
	courseURL string = "http://210.33.60.8/xskbcx.aspx?xh=201305070123"
)

func GetStu(default_url string, c *http.Client, username string, password string, verify_code string) {
	viewRes := jwcpkg.Getsp()
	VIEWSTATE := viewRes["VIEWSTATE"]
	VIEWSTATEGENERATOR := viewRes["VIEWSTATEGENERATOR"]
	_ = jwcpkg.Post(default_url, c, username, password, verify_code, VIEWSTATE, VIEWSTATEGENERATOR, T_cookies)
}

func Get_Code(c *http.Client) {
	T_cookies = jwcpkg.GetLoginCo(c)
	jwcpkg.GetVRcode(c, T_cookies)
}

//get first cookie
// func FirstCo(client *http.Client) []*http.Cookie {
// 	req, _ := http.NewRequest("GET", login_url_gate0, nil)
// 	res, _ := client.Do(req)
// 	return res.Cookies()
// }

/**
*通过查看登录成功之后的名字来判断用户是否登录成功
 */
// func Login_OK(c *http.Client,user,password) bool {
//
// 	stuName := jwcpkg.GetStuName(c,user,password)
// 	if stuName != "" {
// 		return false
// 	} else {
// 		return true
// 	}
// }
