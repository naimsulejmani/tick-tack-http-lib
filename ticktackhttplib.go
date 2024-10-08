package ticktackhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

// GenericRequest makes an HTTP request and handles generic request bodies and responses
func GenericRequest[T any, R any](method string, url string, headers map[string]string, body T) (R, error) {
    var result R
    var reqBody *bytes.Buffer

    // Handle GET, DELETE requests without body
    if method == http.MethodGet || method == http.MethodDelete {
        reqBody = nil
    } else {
        // Convert request body to JSON
        jsonBody, err := json.Marshal(body)
        if err != nil {
            return result, fmt.Errorf("failed to marshal request body: %v", err)
        }
        reqBody = bytes.NewBuffer(jsonBody)
    }

    // Create new HTTP request
    req, err := http.NewRequest(method, url, reqBody)
    if err != nil {
        return result, fmt.Errorf("failed to create request: %v", err)
    }

    // Set headers
    for key, value := range headers {
        req.Header.Set(key, value)
    }

    if method != http.MethodGet && method != http.MethodDelete {
        req.Header.Set("Content-Type", "application/json")
    }

    // Send the request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return result, fmt.Errorf("request failed: %v", err)
    }
    defer resp.Body.Close()

    // Read response body
    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return result, fmt.Errorf("failed to read response body: %v", err)
    }

    // Determine the type of the result to properly handle it
    resultType := reflect.TypeOf(result)

    // If the expected result type is a string, return the raw body as a string
    if resultType.Kind() == reflect.String {
        result = any(string(respBody)).(R)
        return result, nil
    }

    // Unmarshal JSON if the expected result is a struct or slice
    err = json.Unmarshal(respBody, &result)
    if err != nil {
        return result, fmt.Errorf("failed to unmarshal response: %v", err)
    }

    return result, nil
}