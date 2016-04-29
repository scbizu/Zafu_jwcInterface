package main 

import (
	"net/http"
	"net/url"
	"log"
	"fmt"
	"os"
	"io"
	"regexp"
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
	username string="YOUR USERNAME"
	//密码
	password string="PASSWORD"


	)

func checkError(err error){
	if err!=nil{
		log.Fatal("Connection Fail")
		}
	}
/**
* 获取这两个不知道干什么的值
*/
func getsp()map[string]string{
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
		res:=make(map[string]string)
		res["VIEWSTATE"]=VIEWSTATE[0][1]
		res["VIEWSTATEGENERATOR"]=VIEWSTATEGENERATOR[0][1]
		return res	
	}
/**
*模拟post表单
*/
func post(c *http.Client,username string,password string,verify_code string,VIEWSTATE string,VIEWSTATEGENERATOR string,temp_cookies []*http.Cookie) []*http.Cookie{
	
    postValue:=url.Values{}
  	cd:=mahonia.NewEncoder("gb2312")
  	rb:=cd.ConvertString("学生")
  	//准备POST的数据
    postValue.Add("txtUserName",username)
    postValue.Add("TextBox2",password)
    postValue.Add("txtSecretCode",verify_code)
    postValue.Add("__VIEWSTATE",VIEWSTATE)
    postValue.Add("__VIEWSTATEGENERATOR",VIEWSTATEGENERATOR)
 	postValue.Add("Button1","")
 	postValue.Add("lbLanguage","")
 	postValue.Add("hidPdrs","")
 	postValue.Add("hidsc","")
 	postValue.Add("RadioButtonList1",rb)
 	//开始POST   这次POST到登陆界面   带上第一次请求的cookie 和 验证码  和 一些必要的数据
 	postUrl,_:=url.Parse(default_url)
 	Jar,_:=cookiejar.New(nil)
 	Jar.SetCookies(postUrl,temp_cookies)
 	c.Jar=Jar
 	result,_:=c.PostForm(default_url,postValue)
 	cookies=result.Cookies()	
 	return cookies
	}
/*
*测试结果
*/
func Testpage(cookies []*http.Cookie,c *http.Client) string{
 	//拿到这个登录成功的cookie后  再带着这个cookie 再伪造一次请求去我们想要的URL
 	
 	req,err:=http.NewRequest("GET",logged_url,nil)
 	checkError(err)
 	//traversal last cookies
 	for _,fv:=range cookies{
 		req.AddCookie(fv)
 		fmt.Println(fv)
 		}
    finalRes,err:=c.Do(req)
 	checkError(err)
 	allData,err:=ioutil.ReadAll(finalRes.Body)
 	checkError(err)
 	defer finalRes.Body.Close()
 	return string(allData)	
	}


//MAIN

func main() {
	viewRes:=getsp()	
	VIEWSTATE:=viewRes["VIEWSTATE"]
	VIEWSTATEGENERATOR:=viewRes["VIEWSTATEGENERATOR"]
	
    //获取登陆界面的cookie
    c := &http.Client{}
    req, _ := http.NewRequest("GET", login_url_gate0, nil)
    res, _ := c.Do(req)
   var temp_cookies = res.Cookies()
	//第二次 带着登陆界面的cookie去验证码页面拿验证码
   	req.URL, _ = url.Parse(vrcode_url_gate0)
    for _, v := range res.Cookies() {
        req.AddCookie(v)
        fmt.Println(v)
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
    
	cookies:=post(c,username,password,verify_code,VIEWSTATE,VIEWSTATEGENERATOR,temp_cookies)
 	
 	data:=Testpage(cookies,c)
 	
	fmt.Println(data)
}
