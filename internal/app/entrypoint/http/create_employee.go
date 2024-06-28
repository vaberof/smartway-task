package http

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/vaberof/smartway-task/internal/app/entrypoint/http/views"
	"github.com/vaberof/smartway-task/internal/domain/employee"
	"github.com/vaberof/smartway-task/pkg/http/protocols/apiv1"
	"net/http"
)

type createEmployeeRequestBody struct {
	Name       string           `json:"name" validate:"required"`
	Surname    string           `json:"surname" validate:"required"`
	Phone      string           `json:"phone" validate:"required"`
	CompanyId  int64            `json:"company_id" validate:"required"`
	Passport   CreatePassport   `json:"passport" validate:"required"`
	Department CreateDepartment `json:"department" validate:"required"`
}

type CreatePassport struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type CreateDepartment struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type createEmployeeResponseBody struct {
	Id int64 `json:"id"`
}

func (c *createEmployeeRequestBody) Bind(req *http.Request) error {
	return nil
}

// @Summary		Create a new employee
// @Tags			employees
// @Description	Create a new employee
// @ID				create-employee
// @Accept			json
// @Produce		json
// @Param			input	body		createEmployeeRequestBody	true	"Payload for creating an employee"
// @Success		200		{object}	createEmployeeResponseBody
// @Failure		400		{object}	apiv1.Response
// @Failure		500		{object}	apiv1.Response
// @Router			/employees [post]
func (h *Handler) CreateEmployeeHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		createEmployeeReqBody := &createEmployeeRequestBody{}
		if err := render.Bind(request, createEmployeeReqBody); err != nil {
			views.RenderJSON(rw, request, http.StatusBadRequest, apiv1.Error(apiv1.CodeBadRequest, "Invalid request body"))

			return
		}

		validate := validator.New()
		if err := validate.Struct(createEmployeeReqBody); err != nil {
			views.RenderJSON(rw, request, http.StatusBadRequest, apiv1.Error(apiv1.CodeBadRequest, "Invalid request body"))

			return
		}

		id, err := h.employeeService.Create(
			createEmployeeReqBody.Name,
			createEmployeeReqBody.Surname,
			createEmployeeReqBody.Phone,
			createEmployeeReqBody.CompanyId,
			createEmployeeReqBody.Passport.Type,
			createEmployeeReqBody.Passport.Number,
			createEmployeeReqBody.Department.Name,
			createEmployeeReqBody.Department.Phone)
		if err != nil {
			if errors.Is(err, employee.ErrCompanyNotFound) {
				views.RenderJSON(rw, request, http.StatusNotFound, apiv1.Error(apiv1.CodeNotFound, "Company not found"))
			} else {
				views.RenderJSON(rw, request, http.StatusInternalServerError, apiv1.Error(apiv1.CodeInternalError, err.Error()))
			}

			return
		}

		payload, _ := json.Marshal(&createEmployeeResponseBody{
			Id: id,
		})

		views.RenderJSON(rw, request, http.StatusOK, apiv1.Success(payload))
	}
}
