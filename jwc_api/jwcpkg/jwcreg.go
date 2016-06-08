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
	Class    map[string]string
	Teacher  map[string]string
	Timeline map[string]string
	Place    map[string]string
}

//成绩相关
type Scoreinfo struct {
	ClassName string
	Credit    string
	GPA       string
	Score     string
	Academy   string
	ReTest    string
	Rebuild   string
}

func HaveThreeClass(str string) (bool, int) {
	regular := `<td align="Center" rowspan="2">(.*?)</td>`
	pattern := regexp.MustCompile(regular)
	data := pattern.FindAllStringSubmatch(str, -1)
	if len(data) > 0 {
		return true, len(data)
	} else {
		return false, 0
	}
}

func Fetchclass(str string) map[string]*Classinfo {
	var i int
	cd, _ := iconv.Open("utf-8", "gbk") // convert gbk to utf-8
	defer cd.Close()
	class := new(Classinfo)
	// 先匹配全部的课程列表
	//初始化ClassInfo结构
	ClassInfo := make(map[string]*Classinfo)
	//上午12节课
	regular_M12 := `<td align="Center" rowspan="2" width="7%">(.*?)</td>`
	pattern_M12 := regexp.MustCompile(regular_M12)
	data_M12 := pattern_M12.FindAllStringSubmatch(str, -1)
	for i = 0; i < len(data_M12); i++ {
		beego.Debug(cd.ConvString(data_M12[i][1]))
		class = Split_Class(data_M12[i][1])
		ClassInfo[strconv.Itoa(i)] = class
	}

	//上午3节课

	regular_M3 := `<td align="Center">(.*?)</td>`
	pattern_M3 := regexp.MustCompile(regular_M3)
	data_M3 := pattern_M3.FindAllStringSubmatch(str, -1)
	beego.Debug(len(data_M3))

	for i2 := 5; i2 < 10; i2++ {
		beego.Debug(cd.ConvString(data_M3[i2][1]))
		class = Split_Class(data_M3[i2][1])
		ClassInfo[strconv.Itoa(i)] = class
		i++
	}
	//第五节课有无课

	return ClassInfo
}

func Split_Class(classinfo string) *Classinfo {

	ClassNameArr := make(map[string]string)
	TeacherNameArr := make(map[string]string)
	TimelineArr := make(map[string]string)
	PlaceArr := make(map[string]string)

	regular_normal := `(.*?)<br>`
	pattern_normal := regexp.MustCompile(regular_normal)
	data_normal := pattern_normal.FindAllStringSubmatch(classinfo, -1)
	beego.Debug(len(data_normal))
	ClassNameArr["first"] = data_normal[0][1]
	// TeacherNameArr = append(TeacherNameArr, data_normal[2][1])
	TeacherNameArr["first"] = data_normal[2][1]
	// TimelineArr = append(TimelineArr, data_normal[1][1])
	TimelineArr["first"] = data_normal[1][1]
	// PlaceArr = append(PlaceArr, data_normal[3][1])
	PlaceArr["first"] = data_normal[3][1]
	//两节课在一个教室的case
	if len(data_normal) > 5 {
		// ClassNameArr = append(ClassNameArr, data_normal[6][1])
		ClassNameArr["second"] = data_normal[6][1]
		// TeacherNameArr = append(TeacherNameArr, data_normal[8][1])
		TeacherNameArr["second"] = data_normal[8][1]
		// TimelineArr = append(TimelineArr, data_normal[7][1])
		TimelineArr["second"] = data_normal[7][1]
		// PlaceArr = append(PlaceArr, data_normal[9][1])
		PlaceArr["second"] = data_normal[9][1]
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

// func DirectTbody(str string) string {
// 	regular := `<table class="datelist" cellspacing="0" cellpadding="3" border="0" id="Datagrid1" style="DISPLAY:block">(.*?)</table>`
// 	pattern := regexp.MustCompile(regular)
// 	list := pattern.FindAllStringSubmatch(str, -1)
// 	return list[0][1]
// }
//
// func FetchScoreTR(str string) map[string]*Scoreinfo {
//
// 	scoreinfo := make(map[string]*Scoreinfo)
// 	cd, _ := iconv.Open("utf-8", "gbk")
// 	defer cd.Close()
// 	regular := `<tr>(.*?)</tr>`
// 	pattern := regexp.MustCompile(regular)
// 	list := pattern.FindAllStringSubmatch(str, -1)
// 	listLen := len(list)
//
// 	regularALT := `<tr class="alt">(.*?)</tr>`
// 	patternALT := regexp.MustCompile(regularALT)
// 	listALT := patternALT.FindAllStringSubmatch(str, -1)
// 	listLenALT := len(list)
// 	all := make([]string, listLen+listLenALT)
// 	for i := 0; i < listLen+listLenALT; i = i + 2 {
// 		all[i] = list[i/2][1]
// 		beego.Debug(all[i])
// 	}
// 	for i := 1; i < listLenALT+listLen; i = i + 2 {
// 		all[i] = listALT[(i-1)/2][1]
// 		beego.Debug(all[i])
// 	}
//
// 	for i := 0; i < len(all); i++ {
// 		Score := FetchScoreTD(all[i])
// 		beego.Debug(Score)
// 		scoreinfo[strconv.Itoa(i)] = Score
// 	}
// 	return scoreinfo
// }

func FetchScoreTD(str string) map[string]*Scoreinfo {
	scoreinfo := make(map[string]*Scoreinfo)
	cd, _ := iconv.Open("utf-8", "gbk")
	defer cd.Close()
	regular := `<td>(.*?)</td>`
	pattern := regexp.MustCompile(regular)
	info := pattern.FindAllStringSubmatch(str, -1)
	// infoLen:=	len(info)
	beego.Debug(len(info))
	for i := 1; i < (len(info)-18)/15; i++ {
		Classname := cd.ConvString(info[15*(i+1)+4][1])
		Credit := cd.ConvString(info[15*(i+1)+7][1])
		GPA := cd.ConvString(info[15*(i+1)+8][1])
		score := cd.ConvString(info[15*(i+1)+9][1])
		Academy := cd.ConvString(info[15*(i+1)+13][1])
		ReTest := cd.ConvString(info[15*(i+1)+11][1])
		Rebuild := cd.ConvString(info[15*(i+1)+12][1])
		Score := &Scoreinfo{Classname, Credit, GPA, score, Academy, ReTest, Rebuild}
		beego.Debug(Score)
		scoreinfo[strconv.Itoa(i-1)] = Score
	}
	return scoreinfo
}
