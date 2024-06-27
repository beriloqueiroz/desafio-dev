package web

import (
	"encoding/json"
	"github.com/beriloqueiroz/desafio-back/core/internal/usecase"
	"log/slog"
	"net/http"
	"time"
)

type SchedulerRoutes struct {
	CreateScheduleNotificationUseCase usecase.CreateScheduleNotificationUseCase
	DeleteScheduleNotificationUseCase usecase.DeleteScheduleNotificationUseCase
}

func NewSchedulerRoutes(
	createScheduleNotificationUseCase usecase.CreateScheduleNotificationUseCase,
	deleteScheduleNotificationUseCase usecase.DeleteScheduleNotificationUseCase) *SchedulerRoutes {
	return &SchedulerRoutes{
		CreateScheduleNotificationUseCase: createScheduleNotificationUseCase,
		DeleteScheduleNotificationUseCase: deleteScheduleNotificationUseCase,
	}
}

type createScheduleInputDto struct {
	StartTime time.Time `json:"start_time" default:"2024-06-24T23:58:50-03:00"`
}

// @Summary Create new Scheduler
// @Description Create new Scheduler
// @Tags scheduler
// @Accept  json
// @Produce  json
// @Param scheduler body createScheduleInputDto true "scheduler"
// @Success 200 {object} output
// @Router /schedule [post]
func (rs *SchedulerRoutes) CreateScheduleNotificationHandler(w http.ResponseWriter, r *http.Request) {
	var input createScheduleInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&output{
			Message: err.Error(),
		})
		return
	}
	err = rs.CreateScheduleNotificationUseCase.Execute(r.Context(), input.StartTime)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&output{
			Message: err.Error(),
		})
		return
	}
	output := &output{
		Message: "Insert Success",
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

// @Summary Delete scheduler
// @Description Delete scheduler
// @Tags scheduler
// @Accept  json
// @Produce  json
// @Param id path string true "scheduler id (uuid)"
// @Success 200 {object} output
// @Router /schedule/{id} [delete]
func (rs *SchedulerRoutes) DeleteScheduleNotificationHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := rs.DeleteScheduleNotificationUseCase.Execute(r.Context(), id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&output{
			Message: err.Error(),
		})
		return
	}
	output := &output{
		Message: "Delete Success",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
