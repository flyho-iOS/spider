package spider

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Spider(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("fetch url error: ", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("fetch url status code: ", resp.StatusCode)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("read content fail: ", err)
		return
	}

	fmt.Println("body: ", string(body))
}
