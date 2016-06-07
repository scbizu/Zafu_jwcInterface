package models

import (
	"io/ioutil"
	"net/http"
)

//Get Course info.
func GetCourseData(c *http.Client) string {
	CourseURL := courseURL + StuNo
	req, err := http.NewRequest("GET", CourseURL, nil)
	//NICE
	req.Header.Set("Referer", CourseURL)
	CheckError(err)
	finalRes, err := c.Do(req)
	CheckError(err)
	allData, err := ioutil.ReadAll(finalRes.Body)
	CheckError(err)
	finalRes.Body.Close()
	return string(allData)
}
