package helper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// ExtractContent function extracts the problem title and content from the Project Euler website
func ExtractContent(problemString string) (string, string, error) {
	url := fmt.Sprintf("https://projecteuler.net/problem=%s", problemString)

	body, err := fetchURL(url)
	if err != nil {
		return "", "", err
	}

	title, err := extractTitle(body)
	if err != nil {
		return "", "", err
	}

	content, err := extractContent(body)
	if err != nil {
		return "", "", err
	}

	return strings.TrimSpace(title), strings.TrimSpace(content), nil
}

// fetchURL function sends an HTTP GET request and returns the response body
func fetchURL(url string) (string, error) {
	// Create HTTP client and request
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Set user agent header to avoid 403 error
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.9999.999 Safari/537.36")

	// Send request and read response body
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check if response status code is OK
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch URL: %s", resp.Status)
	}

	// Extract problem content from response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	body := string(bodyBytes)

	return body, nil
}

// extractTitle function extracts the problem title from the response body
func extractTitle(body string) (string, error) {
	reTitle := regexp.MustCompile(`(?s)<h2>(.*?)</h2>`)
	titleMatches := reTitle.FindStringSubmatch(body)
	if len(titleMatches) < 2 {
		return "", fmt.Errorf("failed to extract problem title")
	}
	title := titleMatches[1]

	return title, nil
}

// extractContent function extracts the problem content from the response body
func extractContent(body string) (string, error) {
	contentStart := strings.Index(body, `<div class="problem_content" role="problem">`)
	if contentStart == -1 {
		return "", fmt.Errorf("failed to find problem content")
	}
	contentStart += len(`<div class="problem_content" role="problem">`)

	openDivCount := 1
	contentEnd := contentStart

	for i := contentStart; i < len(body); i++ {
		if body[i] == '<' {
			if strings.HasPrefix(body[i:], "<div") {
				openDivCount++
			} else if strings.HasPrefix(body[i:], "</div") {
				openDivCount--
				if openDivCount == 0 {
					contentEnd = i
					break
				}
			}
		}
	}

	content := body[contentStart:contentEnd]

	return content, nil
}
