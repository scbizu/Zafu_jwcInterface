package jwcpkg

import (
	"regexp"
	"strconv"

	"gopkg.in/iconv.v1"

	"github.com/astaxie/beego"
)

var (
	examInfo  map[string]*Examinfo
	classInfo map[string]*Classinfo
)

//考试信息
type Examinfo struct {
	Class    string
	Deadline string
}

//课程信息
type Classinfo struct {
	Class    []string
	Teacher  []string
	timeline []string
	place    []string
}

func Fetchclass(str string) map[string]*Classinfo {
	class := new(Classinfo)
	// 先匹配全部的课程列表
	//初始化ClassInfo结构
	ClassInfo := make(map[string]*Classinfo)
	regular_ALL := `<td align="Center" rowspan="2" width="7%">(.*?)</td>`
	pattern_ALL := regexp.MustCompile(regular_ALL)
	data := pattern_ALL.FindAllStringSubmatch(str, -1)
	for i := 0; i < len(data); i++ {
		class = Split_Class(data[i][1])
		ClassInfo[strconv.Itoa(i)] = class
	}
	return ClassInfo
}

func Split_Class(classinfo string) *Classinfo {

	ClassNameArr := make([]string, 0)
	TeacherNameArr := make([]string, 0)
	TimelineArr := make([]string, 0)
	PlaceArr := make([]string, 0)
	regular_name := `<td align="Center" rowspan="2" width="7%">(.*?)<br>`
	pattern_name := regexp.MustCompile(regular_name)
	data_name := pattern_name.FindAllStringSubmatch(classinfo, -1)
	ClassNameArr = append(ClassNameArr, data_name[0][1])

	regular_normal := `<br>(.*?)<br>`
	pattern_normal := regexp.MustCompile(regular_normal)
	data_normal := pattern_normal.FindAllStringSubmatch(classinfo, -1)
	TeacherNameArr = append(TeacherNameArr, data_normal[1][1])
	TimelineArr = append(TimelineArr, data_normal[0][1])
	PlaceArr = append(PlaceArr, data_normal[2][1])
	//两节课在一个教室的case
	if len(data_normal) > 4 {
		ClassNameArr = append(ClassNameArr, data_normal[6][1])
		TeacherNameArr = append(TeacherNameArr, data_normal[8][1])
		TimelineArr = append(TimelineArr, data_normal[7][1])
		PlaceArr = append(PlaceArr, data_normal[9][1])
	}
	//返回class
	class := &Classinfo{ClassNameArr, TeacherNameArr, TimelineArr, PlaceArr}
	return class
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
		//	beego.Debug(i-1, class, deadline)
	}
	return examInfo
}
