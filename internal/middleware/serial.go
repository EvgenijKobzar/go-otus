package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const KeyResponse = "response"
const KeyError = "error"

type Serial struct {
	start    time.Time
	end      time.Time
	duration time.Duration
}

func Process(c *gin.Context) {

	err, ok := Internalize(c)
	if ok {
		start := time.Now()
		c.Next()
		responseError, _ := c.Get(KeyError)

		if responseError == nil {
			end := time.Now()
			duration := time.Since(start)

			response, _ := c.Get(KeyResponse)

			if !c.Writer.Written() {
				c.JSON(c.Writer.Status(), gin.H{
					"result": response,
					"time": gin.H{
						"start":             start.UnixNano(),
						"end":               end.UnixNano(),
						"duration":          duration,
						"duration_datetime": duration.String(),
						"start_datetime":    start.String(),
						"end_datetime":      end.String(),
					},
				})
			}
		} else {
			err = responseError.(error)
		}
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
}
