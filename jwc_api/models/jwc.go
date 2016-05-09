package models

import(
	"net/http"
	"jwc_api/jwcpkg"
	)

var T_cookies []*http.Cookie

func GetStu(default_url string,c *http.Client,username string,password string,verify_code string){
	viewRes:=jwcpkg.Getsp()	
	VIEWSTATE:=viewRes["VIEWSTATE"]
	VIEWSTATEGENERATOR:=viewRes["VIEWSTATEGENERATOR"]
	_=jwcpkg.Post(default_url,c,username,password,verify_code,VIEWSTATE,VIEWSTATEGENERATOR,T_cookies)	
	}

func Get_Code(c *http.Client){
	T_cookies=jwcpkg.GetLoginCo(c)
	jwcpkg.GetVRcode(c,T_cookies)	
	}