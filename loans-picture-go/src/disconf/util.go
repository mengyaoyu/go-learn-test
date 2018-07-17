package disconf

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getStringFromUrl(url string) string {

	var res string
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("%s\n", err)
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s\n", err)
			return ""
		}
		res = string(bodyBytes)
	}
	return res
}
