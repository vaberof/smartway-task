package http

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/vaberof/smartway-task/internal/domain/employee"
)

type Handler struct {
	employeeService employee.EmployeeService
}

func NewHandler(employeeService employee.EmployeeService) *Handler {
	return &Handler{employeeService: employeeService}
}

func (h *Handler) InitRoutes(router chi.Router) chi.Router {
	router.Route("/api/v1", func(apiV1 chi.Router) {

		apiV1.Route("/employees", func(employees chi.Router) {
			employees.Post("/", h.CreateEmployeeHandler())
			employees.Patch("/{employeeId}", h.UpdateEmployeeHandler())
			employees.Delete("/{employeeId}", h.DeleteEmployeeHandler())

			employees.Get("/companies/{companyId}", h.ListEmployeesByCompanyIdHandler())
			employees.Get("/companies/{companyId}/departments/{departmentName}", h.ListEmployeesByDepartmentNameHandler())
		})
	})

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json")))

	return router
}
