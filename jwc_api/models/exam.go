package models

import (
	"io/ioutil"
	"net/http"
)

func GetExaminfo(c *http.Client) string {
	ExamURL := examURL + StuNo
	req, err := http.NewRequest("GET", ExamURL, nil)
	//NICE
	req.Header.Set("Referer", ExamURL)
	CheckError(err)
	finalRes, err := c.Do(req)
	CheckError(err)
	allData, err := ioutil.ReadAll(finalRes.Body)
	CheckError(err)
	finalRes.Body.Close()
	return string(allData)
}
