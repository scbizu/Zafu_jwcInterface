package models

import (
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/astaxie/beego"
)

//closure
// func InitClient() *http.Client {
// 	Client = &http.Client{}
// 	return Client
// }

func Getvrcode(client *http.Client) (*http.Client, []*http.Cookie) {
	req, _ := http.NewRequest("GET", login_url_gate0, nil)
	res, _ := client.Do(req)
	req.URL, _ = url.Parse(vrcode_url_gate0)

	for _, v := range res.Cookies() {
		req.AddCookie(v)
		beego.Debug(v)
	}
	Res, _ := client.Do(req)
	//test
	//beego.Debug(Res.Cookies())
	fileName := "code/" + StuNo + ".gif"
	file, _ := os.Create(fileName)
	io.Copy(file, Res.Body)

	return client, res.Cookies()
}
