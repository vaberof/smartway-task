package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/vaberof/smartway-task/internal/app/entrypoint/http/views"
	"github.com/vaberof/smartway-task/internal/domain/employee"
	"github.com/vaberof/smartway-task/pkg/http/protocols/apiv1"
	"net/http"
	"strconv"
)

type deleteEmployeeResponseBody struct {
	Message string `json:"message"`
}

// @Summary		Delete an employee
// @Tags			employees
// @Description	Delete an employee
// @ID				delete-employee
// @Accept			json
// @Produce		json
// @Param			id	path		integer	true	"Employees`s id that needs to be deleted"
// @Success		200	{object}	deleteEmployeeResponseBody
// @Failure		400	{object}	apiv1.Response
// @Failure		404	{object}	apiv1.Response
// @Failure		500	{object}	apiv1.Response
// @Router			/employees/{id} [delete]
func (h *Handler) DeleteEmployeeHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		employeeIdStr := chi.URLParam(request, "employeeId")
		if employeeIdStr == "" {
			views.RenderJSON(rw, request, http.StatusBadRequest, apiv1.Error(apiv1.CodeBadRequest, "Missing required url parameter 'employeeId'"))

			return
		}

		employeeId, err := strconv.ParseInt(employeeIdStr, 10, 64)
		if err != nil {
			views.RenderJSON(rw, request, http.StatusBadRequest, apiv1.Error(apiv1.CodeBadRequest, "Employee`s id must be integer"))

			return
		}

		err = h.employeeService.Delete(employeeId)
		if err != nil {
			if errors.Is(err, employee.ErrEmployeeNotFound) {
				views.RenderJSON(rw, request, http.StatusNotFound, apiv1.Error(apiv1.CodeNotFound, "Employee not found"))
			} else {
				views.RenderJSON(rw, request, http.StatusInternalServerError, apiv1.Error(apiv1.CodeInternalError, err.Error()))
			}

			return
		}

		payload, _ := json.Marshal(&deleteEmployeeResponseBody{
			Message: fmt.Sprintf("employee with id '%d' has deleted", employeeId),
		})

		views.RenderJSON(rw, request, http.StatusOK, apiv1.Success(payload))
	}
}
