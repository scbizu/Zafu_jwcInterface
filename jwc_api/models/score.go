package models

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/astaxie/beego"

	"gopkg.in/iconv.v1"
)

func GetScoreinfo(c *http.Client) string {
	ScoreURL := scoreURL + StuNo
	beego.Debug(ScoreURL)
	req, err := http.NewRequest("GET", ScoreURL, nil)
	//NICE
	req.Header.Set("Referer", ScoreURL)
	CheckError(err)
	finalRes, err := c.Do(req)
	CheckError(err)
	allData, err := ioutil.ReadAll(finalRes.Body)
	CheckError(err)
	finalRes.Body.Close()
	return string(allData)
}

func GetscoreVs(str string) string {
	//	beego.Debug(str)
	regular := `<input.type="hidden".name="__VIEWSTATE".value="(.*)" />`
	pattern := regexp.MustCompile(regular)
	res := pattern.FindAllStringSubmatch(str, -1)
	return res[0][1]
}

func GetscoreVg(str string) string {
	regular := `<input.type="hidden".name="__VIEWSTATEGENERATOR".value="(.*)" />`
	pattern := regexp.MustCompile(regular)
	res := pattern.FindAllStringSubmatch(str, -1)
	return res[0][1]
}

func FindOutScore(client *http.Client, Vs string, Vg string, xn string, xq string, btn_xq string, login_co []*http.Cookie) string {
	ScoreURL := scoreURL + StuNo
	beego.Debug(ScoreURL)
	getScore := url.Values{}
	cd, _ := iconv.Open("gbk", "utf-8")
	getScore.Add("__EVENTTARGET", "")
	getScore.Add("__EVENTARGUMENT", "")
	getScore.Add("hidLanguage", "")
	getScore.Add("__VIEWSTATE", Vs)
	getScore.Add("__VIEWSTATEGENERATOR", Vg)
	getScore.Add("ddlXN", "2015-2016")
	getScore.Add("ddlXQ", "2")
	getScore.Add("ddl_kcxz", "")
	getScore.Add("btn_xq", cd.ConvString("学期成绩"))

	req, err := http.NewRequest("POST", ScoreURL, strings.NewReader(getScore.Encode()))
	beego.Debug(getScore.Encode())
	if err != nil {
		beego.Debug(err)
	}
	req.Header.Set("Referer", ScoreURL)

	for _, v := range Login_co {
		req.AddCookie(v)
		beego.Debug(v)
	}

	if err != nil {
		beego.Debug(err)
	}
	Res, err := client.Do(req)
	if err != nil {
		beego.Debug(err)
	}
	beego.Debug(Res.Status)
	body, err := ioutil.ReadAll(Res.Body)
	defer Res.Body.Close()
	if err != nil {
		beego.Debug(err)
	}
	//beego.Debug(string(body))
	return string(body)
}
