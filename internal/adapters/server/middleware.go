package server

import (
	"fmt"
	"time"

	"github.com/flexer2006/orders-api/internal/logger"
	"github.com/gofiber/fiber/v3"
)

func LoggingMiddleware(log logger.Logger) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		duration := time.Since(time.Now())
		err := ctx.Next()
		if err != nil || ctx.Response().StatusCode() >= 400 {
			log.Info("HTTP request",
				"method", ctx.Method(),
				"path", ctx.Path(),
				"status", ctx.Response().StatusCode(),
				"duration_ms", duration.Milliseconds(),
				"ip", ctx.IP(),
			)
		}
		if err != nil {
			return fmt.Errorf("middleware next: %w", err)
		}
		return nil
	}
}
