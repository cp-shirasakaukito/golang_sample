package http

import (
	"net/http"
	"net/http/httputil"
	"fmt"
)

func Request(requestUrl string) {
	//リクエストをセット
	req, _ := http.NewRequest("GET", requestUrl, nil)

	//リクエストヘッダーを出力
	dump, _ := httputil.DumpRequestOut(req, true)
	fmt.Printf("%s", dump)

	//リクエストを実行
	client := new(http.Client)
	resp, _ := client.Do(req)

	//レスポンスを出力
	dumpResp, _ := httputil.DumpResponse(resp, true)
	fmt.Printf("%s", dumpResp)
}
