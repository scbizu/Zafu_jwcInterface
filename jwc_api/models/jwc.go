package models

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"

	"github.com/scbizu/mahonia"
)

var T_cookies []*http.Cookie

//init Client
var Client *http.Client

func InitClient() *http.Client {
	Client = &http.Client{}
	return Client
}

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
	//考试安排
	examURL string = "http://210.33.60.8/xskscx.aspx?xh=201305070123"
)

//Check Error
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Set T_cookies
func SetT_cookie(co []*http.Cookie) {
	T_cookies = co
	return
}

func GetT_cookie() []*http.Cookie {
	return T_cookies
}

func Get_Client() *http.Client {
	return Client
}

// func GetStu(default_url string, c *http.Client, username string, password string, verify_code string) {
// 	viewRes := jwcpkg.Getsp()
// 	VIEWSTATE := viewRes["VIEWSTATE"]
// 	VIEWSTATEGENERATOR := viewRes["VIEWSTATEGENERATOR"]
// 	_ = jwcpkg.Post(default_url, c, username, password, verify_code, VIEWSTATE, VIEWSTATEGENERATOR, T_cookies)
// }

// func Get_Code(c *http.Client) []*http.Cookie {
// 	T_cookies = jwcpkg.GetLoginCo(c)
// 	jwcpkg.GetVRcode(c, T_cookies)
// }

func Getsp() map[string]string {
	view, err := http.Get(login_url_gate0)
	checkError(err)
	//去拿__VIEWSTATE
	body, err := ioutil.ReadAll(view.Body)
	checkError(err)
	regular := `<input.type="hidden".name="__VIEWSTATE".value="(.*)" />`
	pattern := regexp.MustCompile(regular)
	VIEWSTATE := pattern.FindAllStringSubmatch(string(body), -1)
	//拿__VIEWSTATEGENERATOR
	retor := `<input.type="hidden".name="__VIEWSTATEGENERATOR".value="(.*)" />`
	patterntor := regexp.MustCompile(retor)
	VIEWSTATEGENERATOR := patterntor.FindAllStringSubmatch(string(body), -1)
	res := make(map[string]string)
	res["VIEWSTATE"] = VIEWSTATE[0][1]
	if VIEWSTATEGENERATOR != nil {
		res["VIEWSTATEGENERATOR"] = VIEWSTATEGENERATOR[0][1]
	}
	return res
}

/**
*模拟post表单
 */
func Post(c *http.Client, username string, password string, verify_code string, VIEWSTATE string, VIEWSTATEGENERATOR string, temp_cookies []*http.Cookie) ([]*http.Cookie, *http.Client) {
	postValue := url.Values{}
	cd := mahonia.NewEncoder("gb2312")
	rb := cd.ConvertString("学生")
	//准备POST的数据
	postValue.Add("txtUserName", username)
	postValue.Add("TextBox2", password)
	postValue.Add("txtSecretCode", verify_code)
	postValue.Add("__VIEWSTATE", VIEWSTATE)
	postValue.Add("__VIEWSTATEGENERATOR", VIEWSTATEGENERATOR)
	postValue.Add("Button1", "")
	postValue.Add("lbLanguage", "")
	postValue.Add("hidPdrs", "")
	postValue.Add("hidsc", "")
	postValue.Add("RadioButtonList1", rb)
	//开始POST   这次POST到登陆界面   带上第一次请求的cookie 和 验证码  和 一些必要的数据
	postUrl, _ := url.Parse(default_url)
	Jar, _ := cookiejar.New(nil)
	Jar.SetCookies(postUrl, temp_cookies)
	c.Jar = Jar
	resp, _ := c.PostForm(default_url, postValue)
	cookies := resp.Cookies()
	return cookies, c
}

//Get stu name
func GetStuName(c *http.Client) string {
	restuName := ""
	req, err := http.NewRequest("GET", logged_url, nil)
	checkError(err)
	finalRes, err := c.Do(req)
	checkError(err)
	allData, err := ioutil.ReadAll(finalRes.Body)
	checkError(err)
	defer finalRes.Body.Close()
	cd := mahonia.NewEncoder("gb2312")
	rb := cd.ConvertString("<span.id=\"xhxm\">(.*)同学</span>")
	//Regexp
	regular := rb
	pattern := regexp.MustCompile(regular)
	stuName := pattern.FindAllStringSubmatch(string(allData), -1)
	if len(stuName) > 0 {
		restuName = stuName[0][1]
	}
	return restuName
}

/**
*通过查看登录成功之后的名字来判断用户是否登录成功
 */
func Login_OK(c *http.Client) (bool, string) {

	stuName := GetStuName(c)
	if stuName == "" {
		return false, ""
	} else {
		return true, stuName
	}
}
