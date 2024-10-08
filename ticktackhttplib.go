package ticktackhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GenericRequest makes an HTTP request with a generic body and decodes a generic response
func GenericRequest[T any, R any](method string, url string, headers map[string]string, body T) (R, error) {
    var result R

    // Convert request body to JSON
    reqBody, err := json.Marshal(body)
    if err != nil {
        return result, fmt.Errorf("failed to marshal request body: %v", err)
    }

    // Create new HTTP request
    req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
    if err != nil {
        return result, fmt.Errorf("failed to create request: %v", err)
    }

    // Set headers
    for key, value := range headers {
        req.Header.Set(key, value)
    }
    req.Header.Set("Content-Type", "application/json")

    // Send the request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return result, fmt.Errorf("request failed: %v", err)
    }
    defer resp.Body.Close()

    // Read and decode response body
    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return result, fmt.Errorf("failed to read response body: %v", err)
    }

    // Unmarshal the response into the generic response type
    err = json.Unmarshal(respBody, &result)
    if err != nil {
        return result, fmt.Errorf("failed to unmarshal response: %v", err)
    }

    return result, nil
}
