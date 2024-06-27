package http

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/vaberof/smartway-task/internal/app/entrypoint/http/views"
	"github.com/vaberof/smartway-task/internal/domain/employee"
	"github.com/vaberof/smartway-task/pkg/http/protocols/apiv1"
	"net/http"
	"strconv"
)

const (
	defaultLimit  = 10
	defaultOffset = 0
)

type listEmployeeResponseBody struct {
	Employees []*listEmployeeResponse `json:"employees"`
}

type listEmployeeResponse struct {
	Name       string         `json:"name"`
	Surname    string         `json:"surname"`
	Phone      string         `json:"phone"`
	CompanyId  int64          `json:"company_id"`
	Passport   ListPassport   `json:"passport"`
	Department ListDepartment `json:"department"`
}

type ListPassport struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type ListDepartment struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

//	@Summary		List employees by company id
//	@Tags			employees
//	@Description	List employees by company id
//	@ID				list-employees-by-company-id
//	@Accept			json
//	@Produce		json
//	@Param			limit		query		integer	false	"An optional query parameter 'limit' that limits total number of returned employees. By default 'limit' = 10"
//	@Param			offset		query		integer	false	"An optional query parameter 'offset' that indicates how many records should be skipped while listing employees. By default 'offset' = 0"
//	@Param			companyId	path		integer	true	"Path parameter 'companyId'"
//	@Success		200			{object}	listEmployeeResponse
//	@Failure		400			{object}	apiv1.Response
//	@Failure		404			{object}	apiv1.Response
//	@Failure		500			{object}	apiv1.Response
//	@Router			/employees/companies/{companyId} [get]
func (h *Handler) ListEmployeesByCompanyIdHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		var err error

		companyIdStr := chi.URLParam(request, "companyId")
		if companyIdStr == "" {
			views.RenderJSON(rw, request, http.StatusBadRequest, apiv1.Error(apiv1.CodeBadRequest, "Missing required url parameter 'companyId'"))

			return
		}

		companyId, err := strconv.ParseInt(companyIdStr, 10, 64)
		if err != nil {
			views.RenderJSON(rw, request, http.StatusBadRequest, apiv1.Error(apiv1.CodeBadRequest, "Employee`s id must be integer"))

			return
		}

		var limit, offset int

		limitStr := request.URL.Query().Get("limit")
		offsetStr := request.URL.Query().Get("offset")

		if limitStr == "" {
			limit = defaultLimit
		} else {
			limit, err = strconv.Atoi(limitStr)
			if err != nil {
				views.RenderJSON(rw, request, http.StatusBadRequest, apiv1.Error(apiv1.CodeBadRequest, "'limit' parameter must be non-negative integer"))

				return
			}
			if limit < 0 {
				views.RenderJSON(rw, request, http.StatusBadRequest, apiv1.Error(apiv1.CodeBadRequest, "'limit' parameter must be non-negative integer"))

				return
			}
		}

		if offsetStr == "" {
			offset = defaultOffset
		} else {
			offset, err = strconv.Atoi(offsetStr)
			if err != nil {
				views.RenderJSON(rw, request, http.StatusBadRequest, apiv1.Error(apiv1.CodeBadRequest, "'offset' parameter must be non-negative integer"))

				return
			}
			if offset < 0 {
				views.RenderJSON(rw, request, http.StatusBadRequest, apiv1.Error(apiv1.CodeBadRequest, "'offset' parameter must be non-negative integer"))

				return
			}
		}

		employees, err := h.employeeService.ListByCompanyId(companyId, limit, offset)
		if err != nil {
			if errors.Is(err, employee.ErrEmployeeNotFound) {
				views.RenderJSON(rw, request, http.StatusNotFound, apiv1.Error(apiv1.CodeNotFound, "Employee not found"))
			} else {
				views.RenderJSON(rw, request, http.StatusInternalServerError, apiv1.Error(apiv1.CodeInternalError, err.Error()))
			}

			return
		}

		payload, _ := json.Marshal(&listEmployeeResponseBody{
			Employees: h.buildListEmployeeResponseBody(employees),
		})

		views.RenderJSON(rw, request, http.StatusOK, apiv1.Success(payload))
	}
}

//	@Summary		List employees by department name
//	@Tags			employees
//	@Description	List employees by department name
//	@ID				list-employees-by-department-name
//	@Accept			json
//	@Produce		json
//	@Param			limit			query		integer	false	"An optional query parameter 'limit' that limits total number of returned employees. By default 'limit' = 10"
//	@Param			offset			query		integer	false	"An optional query parameter 'offset' that indicates how many records should be skipped while listing employees. By default 'offset' = 0"
//	@Param			companyId		path		integer	true	"Path parameter 'companyId'"
//	@Param			departmentName	path		integer	true	"Path parameter 'departmentName'"
//	@Success		200				{object}	listEmployeeResponse
//	@Failure		400				{object}	apiv1.Response
//	@Failure		404				{object}	apiv1.Response
//	@Failure		500				{object}	apiv1.Response
//	@Router			/employees/companies/{companyId}/departments/{departmentName} [get]
func (h *Handler) ListEmployeesByDepartmentNameHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		var err error

		companyIdStr := chi.URLParam(request, "companyId")
		if companyIdStr == "" {
			views.RenderJSON(rw, request, http.StatusBadRequest, apiv1.Error(apiv1.CodeBadRequest, "Missing required url parameter 'companyId'"))

			return
		}

		companyId, err := strconv.ParseInt(companyIdStr, 10, 64)
		if err != nil {
			views.RenderJSON(rw, request, http.StatusBadRequest, apiv1.Error(apiv1.CodeBadRequest, "Employee`s id must be integer"))

			return
		}

		departmentName := chi.URLParam(request, "departmentName")
		if departmentName == "" {
			views.RenderJSON(rw, request, http.StatusBadRequest, apiv1.Error(apiv1.CodeBadRequest, "Missing required url parameter 'departmentName'"))

			return
		}

		var limit, offset int

		limitStr := request.URL.Query().Get("limit")
		offsetStr := request.URL.Query().Get("offset")

		if limitStr == "" {
			limit = defaultLimit
		} else {
			limit, err = strconv.Atoi(limitStr)
			if err != nil {
				views.RenderJSON(rw, request, http.StatusBadRequest, apiv1.Error(apiv1.CodeBadRequest, "'limit' parameter must be non-negative integer"))

				return
			}
			if limit < 0 {
				views.RenderJSON(rw, request, http.StatusBadRequest, apiv1.Error(apiv1.CodeBadRequest, "'limit' parameter must be non-negative integer"))

				return
			}
		}

		if offsetStr == "" {
			offset = defaultOffset
		} else {
			offset, err = strconv.Atoi(offsetStr)
			if err != nil {
				views.RenderJSON(rw, request, http.StatusBadRequest, apiv1.Error(apiv1.CodeBadRequest, "'offset' parameter must be non-negative integer"))

				return
			}
			if offset < 0 {
				views.RenderJSON(rw, request, http.StatusBadRequest, apiv1.Error(apiv1.CodeBadRequest, "'offset' parameter must be non-negative integer"))

				return
			}
		}

		employees, err := h.employeeService.ListByDepartmentName(companyId, departmentName, limit, offset)
		if err != nil {
			if errors.Is(err, employee.ErrEmployeeNotFound) {
				views.RenderJSON(rw, request, http.StatusNotFound, apiv1.Error(apiv1.CodeNotFound, "Employee not found"))
			} else {
				views.RenderJSON(rw, request, http.StatusInternalServerError, apiv1.Error(apiv1.CodeInternalError, err.Error()))
			}

			return
		}

		payload, _ := json.Marshal(&listEmployeeResponseBody{
			Employees: h.buildListEmployeeResponseBody(employees),
		})

		views.RenderJSON(rw, request, http.StatusOK, apiv1.Success(payload))
	}
}

func (h *Handler) buildListEmployeeResponseBody(domainEmployees []*employee.Employee) []*listEmployeeResponse {
	employees := make([]*listEmployeeResponse, len(domainEmployees))
	for i := range domainEmployees {
		employees[i] = h.buildListEmployeeResponse(domainEmployees[i])
	}
	return employees
}

func (h *Handler) buildListEmployeeResponse(domainEmployee *employee.Employee) *listEmployeeResponse {
	return &listEmployeeResponse{
		Name:      domainEmployee.Name,
		Surname:   domainEmployee.Surname,
		Phone:     domainEmployee.Phone,
		CompanyId: domainEmployee.CompanyId,
		Passport: ListPassport{
			Type:   domainEmployee.Passport.Type,
			Number: domainEmployee.Passport.Number,
		},
		Department: ListDepartment{
			Name:  domainEmployee.Department.Name,
			Phone: domainEmployee.Department.Phone,
		},
	}
}
