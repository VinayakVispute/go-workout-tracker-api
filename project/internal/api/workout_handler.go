package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/vinayakvispute/project/internal/store"
	"github.com/vinayakvispute/project/internal/utils"
)

type WorkoutHandler struct {
	workoutStore store.WorkoutStore
	logger       *log.Logger
}

func NewWorkoutHandler(workoutStore store.WorkoutStore, logger *log.Logger) *WorkoutHandler {
	return &WorkoutHandler{
		workoutStore: workoutStore,
		logger:       logger,
	}
}

func (wh *WorkoutHandler) HandleGetWorkoutByID(w http.ResponseWriter, r *http.Request) {

	workoutID, err := utils.ReadIDParam(r)

	if err != nil {
		wh.logger.Printf("Error: readIDParams: %v", err)
		utils.WriteJson(w, http.StatusBadRequest, utils.Envelop{"error": "invalid workout id"})
		return
	}

	workout, err := wh.workoutStore.GetWorkoutByID(workoutID)

	if err != nil {

		wh.logger.Printf("Error: getWorkoutByID: %v", err)
		utils.WriteJson(w, http.StatusInternalServerError, utils.Envelop{"error": "invalid server error"})
		return
	}
	utils.WriteJson(w, http.StatusOK, utils.Envelop{"workout": workout})
}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout store.Workout

	err := json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		wh.logger.Printf("Error: decodingCreateWorkout: %v", err)
		utils.WriteJson(w, http.StatusBadRequest, utils.Envelop{"error": "invalid server error"})
		return
	}
	createdWorkout, err := wh.workoutStore.CreateWorkout(&workout)
	if err != nil {
		wh.logger.Printf("Error: createWorkout: %v", err)
		utils.WriteJson(w, http.StatusInternalServerError, utils.Envelop{"error": "failed to create workout"})
		return
	}

	utils.WriteJson(w, http.StatusCreated, utils.Envelop{"workout": createdWorkout})

}

func (wh *WorkoutHandler) HandleUpdateWorkoutByID(w http.ResponseWriter, r *http.Request) {
	workoutID, err := utils.ReadIDParam(r)
	if err != nil {
		wh.logger.Printf("Error: readIDParams: %v", err)
		utils.WriteJson(w, http.StatusBadRequest, utils.Envelop{"error": "invalid workout id"})
		return
	}

	exisitingWorkout, err := wh.workoutStore.GetWorkoutByID(workoutID)
	if err != nil {
		wh.logger.Printf("Error: getWorkoutByID: %v", err)
		utils.WriteJson(w, http.StatusInternalServerError, utils.Envelop{"error": "invalid server error"})
		return

	}
	if exisitingWorkout == nil {
		wh.logger.Printf("Error: workout not found: %v", err)
		http.Error(w, "workout not found", http.StatusNotFound)
		return
	}

	// at this point we assume that we are able to find the existing workout
	// imp to discuss here that why string pointer
	var updateWorkoutRequest struct {
		Title           *string              `json:"title"`
		Description     *string              `json:"description"`
		DurationMinutes *int                 `json:"duration_minutes"`
		CaloriesBurned  *int                 `json:"calories_burned"`
		Entries         []store.WorkoutEntry `json:"entries"`
	}

	err = json.NewDecoder(r.Body).Decode(&updateWorkoutRequest)

	if err != nil {
		wh.logger.Printf("Error: decodingUpdateWorkout: %v", err)
		utils.WriteJson(w, http.StatusBadRequest, utils.Envelop{"error": "invalid server error"})
		return
	}

	if updateWorkoutRequest.Title != nil {
		exisitingWorkout.Title = *updateWorkoutRequest.Title
	}
	if updateWorkoutRequest.Description != nil {
		exisitingWorkout.Description = *updateWorkoutRequest.Description
	}
	if updateWorkoutRequest.DurationMinutes != nil {
		exisitingWorkout.DurationMinutes = *updateWorkoutRequest.DurationMinutes
	}
	if updateWorkoutRequest.CaloriesBurned != nil {
		exisitingWorkout.CaloriesBurned = *updateWorkoutRequest.CaloriesBurned
	}
	if updateWorkoutRequest.Entries != nil {
		exisitingWorkout.Entries = updateWorkoutRequest.Entries
	}

	err = wh.workoutStore.UpdateWorkout(exisitingWorkout)
	if err != nil {
		wh.logger.Printf("Error: updateWorkout: %v", err)
		utils.WriteJson(w, http.StatusInternalServerError, utils.Envelop{"error": "failed to update workout"})
		return
	}

	utils.WriteJson(w, http.StatusOK, utils.Envelop{"workout": exisitingWorkout})
}

func (wh *WorkoutHandler) HandleDeleteWorkoutByID(w http.ResponseWriter, r *http.Request) {

	workoutID, err := utils.ReadIDParam(r)
	if err != nil {
		wh.logger.Printf("Error: readIDParams: %v", err)
		utils.WriteJson(w, http.StatusBadRequest, utils.Envelop{"error": "invalid workout id"})
		return
	}

	err = wh.workoutStore.DeleteWorkoutByID(workoutID)

	if err == sql.ErrNoRows {
		wh.logger.Printf("Error: workout not found: %v", err)
		utils.WriteJson(w, http.StatusNotFound, utils.Envelop{"error": "workout not found"})
	} else if err != nil {
		wh.logger.Printf("Error: deleteWorkout: %v", err)
		utils.WriteJson(w, http.StatusInternalServerError, utils.Envelop{"error": "failed to delete workout"})
		return
	}

	utils.WriteJson(w, http.StatusNoContent, utils.Envelop{"message": "workout deleted successfully"})
	// No need to return a body for a 204 No Content response

}
