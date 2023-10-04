// internal/app/rest/middleware/observability_middleware.go
package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
)

func TracingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		spanName := fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.Path)

		// Start a span for the incoming request
		tracer := otel.Tracer("http-server")
		ctx, span := tracer.Start(c.Request.Context(), spanName)
		defer span.End()

		// Add the span to the context
		c.Set("span", span)
		c.Request.WithContext(ctx) // add the new context into the request

		// Call the next middleware or handler
		c.Next()
	}
}
