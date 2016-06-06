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
	checkError(err)
	finalRes, err := c.Do(req)
	checkError(err)
	allData, err := ioutil.ReadAll(finalRes.Body)
	checkError(err)
	finalRes.Body.Close()
	return string(allData)
}
