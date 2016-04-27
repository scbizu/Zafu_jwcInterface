package main 

import (
	"net/http"
	"net/url"
	"log"
	"fmt"
	"os"
	"io"
	"regexp"
//	"strings"
	"io/ioutil"
	"net/http/cookiejar"
	 "github.com/scbizu/mahonia"
)
  
//全局Cookies
var cookies []*http.Cookie
//唯一标识
var VIEWSTATE string=""
const(
	//模拟登陆第一个入口地址
	login_url_gate0 string ="http://210.33.60.8/"
	//模拟登陆第一个入口验证码地址
	vrcode_url_gate0 string="http://210.33.60.8/CheckCode.aspx"
	//首页地址
	logged_url string="http://210.33.60.8/xs_main.aspx?xh=201305070123"
	//默认登录页
	default_url string ="http://210.33.60.8/default2.aspx"
	//用户名
	username string=""
	//密码
	password string=""


	)

func checkError(err error){
	if err!=nil{
		log.Fatal("Connection Fail")
		}
	}

func main() {
//第一次  是  特意拿cookie的~
        view,err:=http.Get(login_url_gate0)
        checkError(err)
        //去拿__VIEWSTATE
        body,err:=ioutil.ReadAll(view.Body)
        checkError(err)
        regular:=`<input.type="hidden".name="__VIEWSTATE".value="(.*)" />`
        pattern:=regexp.MustCompile(regular)        
        VIEWSTATE:=pattern.FindAllStringSubmatch(string(body),-1)
    //拿__VIEWSTATEGENERATOR
    	retor:=`<input.type="hidden".name="__VIEWSTATEGENERATOR".value="(.*)" />`
		patterntor:=regexp.MustCompile(retor)
		VIEWSTATEGENERATOR:=patterntor.FindAllStringSubmatch(string(body),-1)
		
		
    //获取登陆界面的cookie
    c := &http.Client{}
    req, _ := http.NewRequest("GET", login_url_gate0, nil)
    res, _ := c.Do(req)

    req.URL, _ = url.Parse(vrcode_url_gate0)
   var temp_cookies = res.Cookies()

    for _, v := range res.Cookies() {
        req.AddCookie(v)
    }
    // 获取验证码
    var verify_code string
    for {
   	//用刚才生成的cookie去爬 验证码   否则会504!!!!!
       res, _ = c.Do(req)      
       file, _ := os.Create("verify.gif")
       io.Copy(file, res.Body)
	   
      fmt.Println("请查看verify.gif， 然后输入验证码， 看不清输入0重新获取验证码")
       fmt.Scanf("%s", &verify_code)
       if verify_code != "0" {
           break
       }
       res.Body.Close()
		}
    //准备POST的数据
    postValue:=url.Values{}
  //  postValue.Add("Expires","-1")
  	cd:=mahonia.NewEncoder("gb2312")
  	rb:=cd.ConvertString("学生")
  	b1:=cd.ConvertString("登录")
    postValue.Add("txtUserName",username)
    postValue.Add("TextBox2",password)
    postValue.Add("txtSecretCode",verify_code)
    postValue.Add("__VIEWSTATE",VIEWSTATE[0][1])
    postValue.Add("__VIEWSTATEGENERATOR",VIEWSTATEGENERATOR[0][1])
 	postValue.Add("Button1",b1)
 	postValue.Add("lbLanguage","")
 	postValue.Add("hidPdrs","")
 	postValue.Add("hidsc","")
 	postValue.Add("RadioButtonList1",rb)
 	//开始POST   这次POST到登陆界面   带上第一次请求的cookie 和 验证码  和 一些必要的数据
 	postUrl,_:=url.Parse(login_url_gate0)
 	Jar,_:=cookiejar.New(nil)
 	Jar.SetCookies(postUrl,temp_cookies)
 	c.Jar=Jar
 	result,_:=c.PostForm(login_url_gate0,postValue)
 	cookies=result.Cookies()
 	
 	//拿到这个登录成功的cookie后  再带着这个cookie 再伪造一次请求去我们想要的URL
 	finalUrl,err:=url.Parse(default_url)
 	checkError(err)
 	finalJar,err:=cookiejar.New(nil)
 	checkError(err)
 	finalJar.SetCookies(finalUrl,cookies)
 	finalReq,err:=http.NewRequest("GET",default_url,nil)
 	checkError(err)
 	finalReq.Header.Set("Referer","http://210.33.60.8/default2.aspx")
 	//finalRes,_=c.Do(finalReq)
 	c.Jar=finalJar
 	
 	allData,err:=ioutil.ReadAll(result.Body)
 	checkError(err)
 	defer result.Body.Close()
 	fmt.Println(string(allData))
}
