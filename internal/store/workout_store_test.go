package store

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=postgres port=5433 sslmode=disable")
	if err != nil {
		t.Fatalf("opening test db: %v", err)
	}
	t.Log("connected to test db")
	// runs the migrations for test db
	err = Migrate(db, "../../migrations")
	if err != nil {
		t.Fatalf("migrating test db error: %v", err)
	}

	_, err = db.Exec("TRUNCATE workouts, workout_entries CASCADE")
	if err != nil {
		t.Fatalf("truncating tables: %v", err)
	}
	fmt.Printf("Test database setup complete\n")
	return db
}

func TestCreateWorkout(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	store := NewPostgresWorkoutStore(db)

	tests := []struct {
		name    string
		workout *Workout
		wantErr bool
	}{
		{
			name: "valid workout",
			workout: &Workout{
				Title:           "Push Day",
				Description:     "Upper body day",
				DurationMinutes: 60,
				CaloriesBurned:  200,
				Entries: []WorkoutEntry{
					{
						ExerciseName: "Brench Press",
						Sets:         3,
						Reps:         IntPtr(10),
						Weight:       FloatPrt(135.5),
						Notes:        "warm up properly",
						OrderIndex:   1,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "workout with invalid entries",
			workout: &Workout{
				Title:           "full body",
				Description:     "complete workout done",
				DurationMinutes: 90,
				CaloriesBurned:  500,
				Entries: []WorkoutEntry{
					{
						ExerciseName: "Plank",
						Sets:         3,
						Reps:         IntPtr(60),
						Notes:        "Keep the form",
						OrderIndex:   1,
					},
					{
						ExerciseName:    "Squats",
						Sets:            4,
						Reps:            IntPtr(12),
						DurationSeconds: IntPtr(30),
						Notes:           "Full depts",
						OrderIndex:      2,
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createdWorkout, err := store.CreateWorkout(tt.workout)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.workout.Title, createdWorkout.Title)
			assert.Equal(t, tt.workout.Description, createdWorkout.Description)
			assert.Equal(t, tt.workout.DurationMinutes, createdWorkout.DurationMinutes)

			retrieved, err := store.GetWorkoutByID(int64(createdWorkout.ID))
			require.NoError(t, err)

			assert.Equal(t, createdWorkout.ID, retrieved.ID)
			assert.Equal(t, len(tt.workout.Entries), len(retrieved.Entries))

			for i := range retrieved.Entries {
				assert.Equal(t, tt.workout.Entries[i].ExerciseName, retrieved.Entries[i].ExerciseName)
				assert.Equal(t, tt.workout.Entries[i].Reps, retrieved.Entries[i].Reps)
				assert.Equal(t, tt.workout.Entries[i].Sets, retrieved.Entries[i].Sets)
				assert.Equal(t, tt.workout.Entries[i].Weight, retrieved.Entries[i].Weight)
				assert.Equal(t, tt.workout.Entries[i].Notes, retrieved.Entries[i].Notes)
				assert.Equal(t, tt.workout.Entries[i].OrderIndex, retrieved.Entries[i].OrderIndex)
				assert.Equal(t, tt.workout.Entries[i].DurationSeconds, retrieved.Entries[i].DurationSeconds)

			}
		})
	}
}

func IntPtr(i int) *int {
	return &i
}

func FloatPrt(i float64) *float64 {
	return &i
}
