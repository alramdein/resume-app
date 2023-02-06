package delivery

import (
	"net/http"

	"github.com/alramdein/karirlab-test/model"
	"github.com/labstack/echo/v4"
)

type resumeHandler struct {
	resumeUsecase model.ResumeUsecase
}

func NewResumeHandler(e *echo.Echo, ru model.ResumeUsecase) {
	handler := &resumeHandler{
		resumeUsecase: ru,
	}

	e.POST("/resume", handler.Create)
}

func (r *resumeHandler) Create(c echo.Context) error {
	// TODO: fetch data from body param and put it on create input model and send it to usecase
	return c.JSON(http.StatusOK, &Response{
		Message: "successfully add a resume",
	})

}
