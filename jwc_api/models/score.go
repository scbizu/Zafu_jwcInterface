package models

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/scbizu/mahonia"
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
	// beego.Debug(ScoreURL)
	// beego.Debug(Vs)
	getScore := url.Values{}
	cd := mahonia.NewEncoder("gb2312")
	// getScore.Set("__EVENTTARGET", "")
	// getScore.Add("__EVENTARGUMENT", "")
	// getScore.Add("hidLanguage", "")
	getScore.Add("__VIEWSTATE", Vs)
	getScore.Add("__VIEWSTATEGENERATOR", Vg)
	// getScore.Add("ddlXN", "")
	// getScore.Add("ddlXQ", "")
	getScore.Add("ddl_kcxz", "")
	getScore.Add("btn_zcj", cd.ConvertString("历年成绩"))

	// postUrl, _ := url.Parse(ScoreURL)
	// Jar, _ := cookiejar.New(nil)
	// Jar.SetCookies(postUrl, login_co)
	// client.Jar = Jar

	req, err := http.NewRequest("POST", ScoreURL, bytes.NewBufferString(getScore.Encode()))
	//	beego.Debug(req)
	if err != nil {
		beego.Debug(err)
	}

	for _, v := range Login_co {
		req.AddCookie(v)
		beego.Debug(v)
	}
	req.Header.Add("Referer", ScoreURL)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(getScore.Encode())))
	Res, err := client.Do(req)
	if err != nil {
		beego.Debug(err)
	}

	// if Res.Status == "200 OK" {
	// 	return client, nil
	// } else {
	// 	return nil, errors.New("Post was denied")
	// }
	data, _ := ioutil.ReadAll(Res.Body)
	defer Res.Body.Close()
	return string(data)
}

func GetTestfunc(c *http.Client) string {
	ScoreURL := scoreURL + StuNo
	Req, err := http.NewRequest("GET", ScoreURL, nil)
	if err != nil {
		beego.Debug(err)
	}
	Req.Header.Set("Referer", ScoreURL)
	Res, err := c.Do(Req)
	if err != nil {
		beego.Debug(err)
	}
	data, _ := ioutil.ReadAll(Res.Body)
	defer Res.Body.Close()
	return string(data)
}
