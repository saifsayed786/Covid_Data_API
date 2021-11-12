package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"

	
)
type Data struct{
	CountryName string `json:"country_name"`
	Cases string `json:"cases"`
	Deaths string `json:"death"`
	ActiveCases string `json:"active_cases"`
}
func main() {

	url := "https://corona-virus-world-and-india-data.p.rapidapi.com/api"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "corona-virus-world-and-india-data.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "1a0f68620fmsh5ed906b3969d3b2p1bc410jsn7c582b2b0197")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	
	jsonString := string(body)
	//fmt.Println("API Response as String:\n" + jsonString)

	var result map[string]interface{}
	json.Unmarshal([]byte(jsonString), &result)

	var data Data
	for _,item:=range result["countries_stat"].([]interface{}) {
	
		data.CountryName = fmt.Sprintf("%s",item.(map[string]interface{})["country_name"])
		data.Cases = fmt.Sprintf("%s",item.(map[string]interface{})["cases"])  //assertion
		data.Deaths = fmt.Sprintf("%s",item.(map[string]interface{})["deaths"])
		data.ActiveCases = fmt.Sprintf("%s",item.(map[string]interface{})["active_cases"])
		
		jsonData, err := json.MarshalIndent(data,""," ")
        if err != nil {
                panic(err)
        }

        //fmt.Println(string(jsonData))
		f, err := os.OpenFile("covidData.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if _, err = f.Write(jsonData); err != nil {
			panic(err)
		}
		
	}
	fmt.Println("JSON data written")
	

}