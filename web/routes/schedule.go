package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"lab-argo-app/pkg/schedule"
	"lab-argo-app/web/handlers"
)

func ScheduleRouter(app fiber.Router, service schedule.Service) {
	app.Get("/", handlers.GetSchedule(service))
	app.Get("/lessons/:day", handlers.GetDay(service))

	// Етап 3: новий ендпоінт. Повертає версію образу, яку чарт передає
	// через змінну середовища APP_VERSION (= image.tag)
	app.Get("/version", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"version": os.Getenv("APP_VERSION")})
	})
}
