package models

import (
	"io/ioutil"
	"net/http"
)

func GetExaminfo(c *http.Client) string {
	req, err := http.NewRequest("GET", examURL, nil)
	//NICE
	req.Header.Set("Referer", examURL)
	checkError(err)
	finalRes, err := c.Do(req)
	checkError(err)
	allData, err := ioutil.ReadAll(finalRes.Body)
	checkError(err)
	finalRes.Body.Close()
	return string(allData)
}
