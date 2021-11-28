package logging

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
	"time"
)

// NewEndpointJsonLoggerMiddleware is a gin.HandlerFunc that logs some information of the incoming request and the consequent response.
func NewEndpointJsonLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now().UnixNano()

		// Request URL path
		path := c.Request.URL.Path
		if c.Request.URL.RawQuery != "" {
			path = path + "?" + c.Request.URL.RawQuery
		}

		// Process request
		c.Next()

		// Results
		end := time.Now().UnixNano()

		reqBody, _ := io.ReadAll(c.Request.Body)
		resStatus := c.Writer.Status()

		log, _ := json.MarshalIndent(map[string]interface{}{
			"time":       time.Now().Format("02/01/2006 - 15:04:05"),
			"start_time": start,
			"end_time":   end,
			"duration":   fmt.Sprintf("%s", time.Nanosecond*time.Duration(end-start)),
			"type":       "endpoint_logger",
			"request": map[string]interface{}{
				"path":    path,
				"method":  c.Request.Method,
				"headers": headersToMap(c.Request.Header),
				"body":    string(reqBody),
			},
			"response": map[string]interface{}{
				"status":  resStatus,
				"headers": headersToMap(c.Writer.Header()),
			},
		}, "", "  ")

		fmt.Println(string(log))
	}
}

func headersToMap(header http.Header) map[string]string {
	headers := map[string]string{}

	for key, val := range header {
		headers[key] = strings.Join(val, " ")
	}

	return headers
}
