package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// 接收传来的JSON
type CatFactResponse struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

// 辅助函数，写JSON
func writeJSON(w http.ResponseWriter, status int, data any) {
	// 设置响应头
	w.Header().Set("Content-Type", "application/json")
	// 设置状态码
	w.WriteHeader(status)
	// data编码成JSON
	json.NewEncoder(w).Encode(data)
}

// 辅助函数，获取外部URL
func fetchCatFact() (CatFactResponse, error) {
	url := "https://catfact.ninja/fact"

	resp, err := http.Get(url)

	if err != nil {
		return CatFactResponse{}, err
	}

	// 状态码200为接收成功
	if resp.StatusCode != http.StatusOK {
		return CatFactResponse{}, fmt.Errorf("External API failed: %s", resp.Status)
	}

	// io接收字节流
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return CatFactResponse{}, err
	}

	var data CatFactResponse

	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return CatFactResponse{}, err
	}

	defer resp.Body.Close()

	return data, nil
}

func externalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "only GET method is allowed",
		})
		return
	}

	data, err := fetchCatFact()

	if err != nil {
		writeJSON(w, http.StatusBadGateway, map[string]any{
			"ok":    false,
			"error": "Fail to fetch data",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"ok":        true,
		"timeStamp": time.Now().UTC(),
		"external": map[string]any{
			"source": "CatFact-ninja",
			"fact":   data.Fact,
			"length": data.Length,
		},
	})
}

func main() {

	http.HandleFunc("/external", externalHandler)

	// 启动端口
	err := http.ListenAndServe(":8081", nil)
	fmt.Println(err)
}
