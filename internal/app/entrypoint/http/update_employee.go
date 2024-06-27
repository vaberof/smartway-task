package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/vaberof/smartway-task/internal/app/entrypoint/http/views"
	"github.com/vaberof/smartway-task/internal/domain/employee"
	"github.com/vaberof/smartway-task/pkg/http/protocols/apiv1"
	"net/http"
	"strconv"
)

type updateEmployeeRequestBody struct {
	Name       *string           `json:"name"`
	Surname    *string           `json:"surname"`
	Phone      *string           `json:"phone"`
	CompanyId  *int64            `json:"company_id"`
	Passport   *UpdatePassport   `json:"passport"`
	Department *UpdateDepartment `json:"department"`
}

type UpdatePassport struct {
	Type   *string `json:"type"`
	Number *string `json:"number"`
}

type UpdateDepartment struct {
	Name  *string `json:"name"`
	Phone *string `json:"phone"`
}

type updateEmployeeResponseBody struct {
	Message string `json:"message"`
}

func (c *updateEmployeeRequestBody) Bind(req *http.Request) error {
	return nil
}

//	@Summary		Update a new employee
//	@Tags			employees
//	@Description	update a new employee
//	@ID				update-employee
//	@Accept			json
//	@Produce		json
//	@Param			input		body		updateEmployeeRequestBody	true	"Payload for updating an employee"
//	@Param			employeeId	path		integer						true	"Path parameter 'employeeId'"
//	@Success		200			{object}	updateEmployeeResponseBody
//	@Failure		400			{object}	apiv1.Response
//	@Failure		500			{object}	apiv1.Response
//	@Router			/employees/{employeeId} [patch]
func (h *Handler) UpdateEmployeeHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		employeeIdStr := chi.URLParam(request, "employeeId")
		if employeeIdStr == "" {
			views.RenderJSON(rw, request, http.StatusBadRequest, apiv1.Error(apiv1.CodeBadRequest, "Missing required url parameter 'id'"))

			return
		}

		employeeId, err := strconv.ParseInt(employeeIdStr, 10, 64)
		if err != nil {
			views.RenderJSON(rw, request, http.StatusBadRequest, apiv1.Error(apiv1.CodeBadRequest, "Employee`s id must be integer"))

			return
		}

		updateEmployeeReqBody := &updateEmployeeRequestBody{}
		if err = render.Bind(request, updateEmployeeReqBody); err != nil {
			views.RenderJSON(rw, request, http.StatusBadRequest, apiv1.Error(apiv1.CodeBadRequest, "Invalid request body"))

			return
		}

		var passportType, passportNumber *string
		var departmentName, departmentPhone *string

		if updateEmployeeReqBody.Passport != nil {
			passportType = updateEmployeeReqBody.Passport.Type
			passportNumber = updateEmployeeReqBody.Passport.Number
		}

		if updateEmployeeReqBody.Department != nil {
			departmentName = updateEmployeeReqBody.Department.Name
			departmentPhone = updateEmployeeReqBody.Department.Phone
		}

		err = h.employeeService.Update(
			employeeId,
			updateEmployeeReqBody.Name,
			updateEmployeeReqBody.Surname,
			updateEmployeeReqBody.Phone,
			updateEmployeeReqBody.CompanyId,
			passportType,
			passportNumber,
			departmentName,
			departmentPhone,
		)
		if err != nil {
			if errors.Is(err, employee.ErrEmployeeNotFound) {
				views.RenderJSON(rw, request, http.StatusNotFound, apiv1.Error(apiv1.CodeNotFound, "Employee not found"))
			} else {
				views.RenderJSON(rw, request, http.StatusInternalServerError, apiv1.Error(apiv1.CodeInternalError, err.Error()))
			}

			return
		}

		payload, _ := json.Marshal(&updateEmployeeResponseBody{
			Message: fmt.Sprintf("employee with id '%d' has updated", employeeId),
		})

		views.RenderJSON(rw, request, http.StatusOK, apiv1.Success(payload))
	}
}
