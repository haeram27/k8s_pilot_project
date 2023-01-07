package funcs

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func HttpRequestBody(url string, method string, headers url.Values, datas string) (int, []byte) {
	body := strings.NewReader(datas)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Println(err)
		return -1, nil
	}

	for key, val := range headers {
		req.Header.Set(key, strings.Join(val, ";"))
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return -1, nil
	}
	defer resp.Body.Close()

	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return -1, nil
	}

	return resp.StatusCode, output
}

func HttpRequestParams(url string, method string, headers url.Values, params url.Values) (int, []byte) {
	return HttpRequestBody(url, method, headers, params.Encode())
}

func HttpRequest(url string, method string, headers url.Values) (int, []byte) {
	return HttpRequestBody(url, method, headers, "")
}
