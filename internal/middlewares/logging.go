package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"time"

	"github.com/labstack/echo/v4"
)

func (mw *Middleware) LoggingMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			// Track the start time of the request
			startTime := time.Now()

			if errCheck := func() error {
				// Check if the request has a body
				if req.Body != nil {
					// Read the request body
					reqBody, err := io.ReadAll(req.Body)
					if err != nil {
						return err
					}

					// Restore the request body so it can be used in subsequent handlers
					req.Body = io.NopCloser(bytes.NewBuffer(reqBody))

					// Include request body as JSON in the log
					var requestBodyJSON interface{}
					if err := json.Unmarshal(reqBody, &requestBodyJSON); err != nil {
						return err
					}

					// Include request body in the log
					mw.logger.Info("Received request",
						"method", req.Method,
						"path", req.URL.Path,
						"ip", req.RemoteAddr,
						"body", requestBodyJSON,
					)
				}

				return nil
			}(); errCheck != nil {
				// Log trace and span IDs using zerolog
				mw.logger.Info("Received request",
					"method", req.Method,
					"path", req.URL.Path,
					"ip", req.RemoteAddr,
				)
			}

			// Continue to the next middleware/handler
			err := next(c)

			// Log response details using zerolog
			mw.logger.Info("Sent response",
				"status", res.Status,
				"size", res.Size,
				"duration", time.Since(startTime),
			)

			return err
		}
	}
}
