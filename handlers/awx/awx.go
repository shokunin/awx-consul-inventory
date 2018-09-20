package awx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AwxSearch struct {
	Results []struct {
		ID                int    `json:"id"`
		Name              string `json:"name"`
		HasActiveFailures bool   `json:"has_active_failures"`
	} `json:"results"`
}

func deleteHosts(hostList []int) {
	var username = os.Getenv("AWX_USER")
	var passwd = os.Getenv("AWX_PW")
	url := os.Getenv("AWX_URL") + "/api/v2/hosts/"
	for _, id := range hostList {
		fullPath := url + strconv.Itoa(id) + "/"
		client := &http.Client{}
		req, _ := http.NewRequest("DELETE", fullPath, nil)
		req.SetBasicAuth(username, passwd)
		res, err := client.Do(req)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(res)
	}

}

/*** GenInventory does a query on consul ***/
func GetFailedHosts(c *gin.Context) {

	var username = os.Getenv("AWX_USER")
	var passwd = os.Getenv("AWX_PW")
	url := os.Getenv("AWX_URL") + "/api/v2/hosts/?page_size=100000"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, passwd)
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("fetchfailed")
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("bodyFailed")
		panic(err.Error())
	}

	var data AwxSearch
	json.Unmarshal(body, &data)
	vsf := make([]int, 0)
	for _, v := range data.Results {
		if v.HasActiveFailures {
			vsf = append(vsf, v.ID)
		}
	}
	deleteHosts(vsf)
	c.JSON(200, vsf)
}
