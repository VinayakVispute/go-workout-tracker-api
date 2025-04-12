package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/vinayakvispute/project/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", app.HealthCheck)
	r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutByID)

	r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)
	r.Post("/users", app.UserHandler.HandleRegisterUser)

	r.Put("/workouts/{id}", app.WorkoutHandler.HandleUpdateWorkoutByID)

	r.Delete("/workouts/{id}", app.WorkoutHandler.DeleteWorkoutByID)

	return r
}
