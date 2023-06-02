package eclient

import (
	"encoding/json"
	"fmt"
	"main/common/table"
	"net/http"
	"time"
)

func GetTxInfoFromEtherScan(addr string, start, end int64, page, offset int, sort string, apiKey string) (*table.ResponseTxData, error) {
	url := fmt.Sprintf("https://api.etherscan.io/api?module=account&action=txlist&address=%s&startblock=%d&endblock=%d&page=%d&offset=%d&sort=%s&apikey=%s", addr, start, end, page, offset, sort, apiKey)
	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{Timeout: time.Second * 5}
	resp, err := client.Do(req)
	for err != nil {
		fmt.Println("get request failed:", err)
		time.Sleep(10 * time.Second)
		resp, err = client.Do(req)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error get:%d\n", resp.StatusCode)
		//fmt.Errorf("http code:%d", resp.StatusCode)
		return nil, fmt.Errorf("error http code:%d", resp.StatusCode)
	}

	blocks := table.ResponseTxData{}
	_ = json.NewDecoder(resp.Body).Decode(&blocks)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	return &blocks, nil
}
