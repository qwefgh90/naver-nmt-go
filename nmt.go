package nmt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	apiUrl = "http://labspace.naver.com/api/n2mt/translate"

	client      = "labspace"
	userAgent   = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.59 Safari/537.36" // XXX - fake user agent
	contentType = "application/x-www-form-urlencoded;charset=UTF-8"
	origin      = "http://labspace.naver.com"
	referrer    = "http://labspace.naver.com/nmt/"
)

type Language string

const (
	Korean  Language = "ko"
	English Language = "en"
)

type ApiResponse struct {
	Message Message `json:"message,omitempty"`

	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type Message struct {
	Type    string `json:"@type"`
	Service string `json:"@service"`
	Version string `json:"@version"`
	Result  Result `json:"result"`
}

type Result struct {
	SourceLangType string `json:"srcLangType"`
	TargetLangType string `json:"tarLangType"`
	TranslatedText string `json:"translatedText"`
}

// translate given string
func Translate(str string, from, to Language) (translated string, err error) {
	// form
	form := url.Values{}
	form.Add("source", string(from))
	form.Add("target", string(to))
	form.Add("text", str)

	var req *http.Request
	if req, err = http.NewRequest("POST", apiUrl, strings.NewReader(form.Encode())); err == nil {
		// headers
		headers := http.Header{
			"User-Agent":        []string{userAgent},
			"Origin":            []string{origin},
			"Referer":           []string{referrer},
			"content-type":      []string{contentType},
			"x-naver-client-id": []string{client},
		}
		req.Header = headers

		var resp *http.Response
		client := &http.Client{}
		if resp, err = client.Do(req); err == nil {
			defer resp.Body.Close()

			var bytes []byte
			if bytes, err = ioutil.ReadAll(resp.Body); err == nil {
				var apiResponse ApiResponse

				if err = json.Unmarshal(bytes, &apiResponse); err == nil {
					if resp.StatusCode == 200 {
						translated = apiResponse.Message.Result.TranslatedText
					} else {
						err = fmt.Errorf("HTTP %d: [%s] %s", resp.StatusCode, apiResponse.ErrorCode, apiResponse.ErrorMessage)
					}
				}
			}
		}
	}

	return translated, err
}
