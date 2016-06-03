package jwcpkg

import (
	"regexp"
	"strconv"

	"gopkg.in/iconv.v1"

	"github.com/astaxie/beego"
)

var (
	examInfo map[string]*Examinfo
)

type Examinfo struct {
	Class    string
	Deadline string
}

func Fetchclass(str string) string {
	// classInfo := ""
	regular := `<td align="Center" width="14%">(.*)</td><td align="Center" width="14%">`
	pattern := regexp.MustCompile(regular)
	data := pattern.FindAllStringSubmatch(str, -1)
	return data[0][1]
}

func FetchExam(str string) map[string]*Examinfo {
	cd, _ := iconv.Open("utf-8", "gbk") // convert gbk to utf-8
	defer cd.Close()
	regular := `<td>(.*?)</td>`
	pattern := regexp.MustCompile(regular)
	data := pattern.FindAllStringSubmatch(str, -1)
	examInfo = make(map[string]*Examinfo)
	beego.Debug(len(data))
	dataLen := len(data)
	for i := 1; i < dataLen/8; i++ {
		class := cd.ConvString(data[i*8+1][1])
		deadline := cd.ConvString(data[i*8+3][1])
		if deadline == "&nbsp;" {
			deadline = ""
		}
		examInfo[strconv.Itoa(i-1)] = &Examinfo{class, deadline}
		beego.Debug(i-1, class, deadline)
	}
	return examInfo
}
