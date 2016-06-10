# Zafu_jwcInterface
## About
这是关于浙江农林大学教务处的一个爬虫,用于爬取各个入口下指定学号学生的信息
##Install
* 首先这是一个Go项目 必须需要**Golang**的环境
* 其次这个项目引入了beego框架 需要安装下**beego**的框架 beego的安装看[这里](http://beego.me/) 
* 接下来就是
		 go get  github.com/scbizu/Zafu_jwcInterface
在项目根目录下
		 bee run

##Usage

* /vrcode： 获取**验证码**的路由

操作:Get
需要key:stuno

描述:在login界面,用户输入学号后,获取到这个学号,通过**Get**拼接学号到地址栏
(e.g. /vrcode?stuno=xxxxxx),然后 验证码图片的名称即为 **此学号.gif**

*/jwc:	提交信息路由

操作:POST
Key List:
(学号)stuno
(密码)password
(验证码)vrcode

描述:这里的vrcode 即为请求/vrcode获取的验证码

返回:

true(string)表示登入成功
Failed(string)表示登入失败

* /exam:获取该学生考试信息的路由

操作:Get

**No Key**

描述:
在/jwc 返回true之后 再请求此接口即可

返回:
JSON数据
（“0”：{
	"Class":"xxxx",
	"Deadline":"xxxx"
}...）

* /course:获取学生课程表

**give up**

理由:poor regexp regular

* /score:获取学生成绩

操作:Get

**No key**

描述:
在/jwc 返回true之后 再请求此接口即可

返回:
JSON数据
（“0”：{
	"ClassName":"xxx"
	"Credit":"xxx"
	"GPA":"xxx"
	"Score":"xxx"
	(开课学院)"Academy":"xxx"
	(补考分数)"ReTest":"xxx"
	(重修分数)"Rebuild": "xxx"
}...）

* Other interface will be continued....

Written in Golang 
Author：scnace
