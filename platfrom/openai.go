package platfrom

import (
	"io"
	"net/http"
	"strings"
)

func CheckOpenai(httpClient *http.Client) (bool, error) {
	resp, err := httpClient.Get("https://android.chat.openai.com")
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 403 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return false, err
		}
		if strings.Contains(string(body), "Request is not allowed. Please try again later.") {
			return true, nil
		}
	}

	return false, nil

}
