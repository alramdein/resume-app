package delivery

import (
	"net/http"

	"github.com/alramdein/karirlab-test/model"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type resumeHandler struct {
	resumeUsecase model.ResumeUsecase
}

func NewResumeHandler(e *echo.Echo, ru model.ResumeUsecase) {
	handler := &resumeHandler{
		resumeUsecase: ru,
	}

	e.POST("/resumes", handler.Create)
}

func (r *resumeHandler) Create(c echo.Context) error {
	request := new(model.CreateResumeInput)
	err := c.Bind(request)
	if err != nil {
		logrus.Error(err.Error())
		return c.JSON(http.StatusBadRequest, &Response{
			Message: "invalid json input",
		})
	}

	err = r.resumeUsecase.Create(c.Request().Context(), *request)
	if err != nil {
		logrus.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, &Response{
			Message: "something went wrong",
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Message: "successfully add a resume",
	})

}
