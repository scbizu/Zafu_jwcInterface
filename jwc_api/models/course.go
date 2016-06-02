package models

import (
	"io/ioutil"
	"net/http"
)

//Get Course info.
func GetCourseData(c *http.Client) string {
	req, err := http.NewRequest("GET", courseURL, nil)
	//NICE
	req.Header.Set("Referer", courseURL)
	checkError(err)
	finalRes, err := c.Do(req)
	checkError(err)
	allData, err := ioutil.ReadAll(finalRes.Body)
	checkError(err)
	finalRes.Body.Close()
	return string(allData)
}
