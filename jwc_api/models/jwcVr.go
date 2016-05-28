package models

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

//closure
func InitClient() *http.Client {
	client := &http.Client{}
	return client
}

func Getvrcode(client *http.Client) func() (*http.Client, []*http.Cookie) {
	req, _ := http.NewRequest("GET", login_url_gate0, nil)
	res, _ := client.Do(req)
	req.URL, _ = url.Parse(vrcode_url_gate0)
	for _, v := range res.Cookies() {
		req.AddCookie(v)
	}
	res, _ = client.Do(req)
	file, _ := os.Create("code/verify.gif")
	io.Copy(file, res.Body)
	Vrcode := func() (*http.Client, []*http.Cookie) {
		res.Body.Close()
		return client, res.Cookies()
	}
	return Vrcode
}
