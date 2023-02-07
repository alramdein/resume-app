package delivery

import (
	"net/http"
	"strconv"

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
	e.GET("/resumes/:id", handler.FindByID)
	e.PUT("/resumes/:id", handler.Update)
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

func (r *resumeHandler) Update(c echo.Context) error {
	id := c.Param("id")
	resumeID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logrus.Error(err.Error())
		return c.JSON(http.StatusBadRequest, &Response{
			Message: "invalid resumeID",
		})
	}

	resume, err := r.resumeUsecase.FindByID(c.Request().Context(), resumeID)
	if err != nil {
		logrus.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, &Response{
			Message: "something went wrong",
		})
	}
	if resume == nil {
		return c.JSON(http.StatusNotFound, &Response{
			Message: "resume not found",
		})
	}

	request := new(model.CreateResumeInput)
	err = c.Bind(request)
	if err != nil {
		logrus.Error(err.Error())
		return c.JSON(http.StatusBadRequest, &Response{
			Message: "invalid json input",
		})
	}

	err = r.resumeUsecase.Update(c.Request().Context(), resumeID, *request)
	if err != nil {
		logrus.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, &Response{
			Message: "something went wrong",
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Message: "successfully update a resume",
	})
}

func (r *resumeHandler) FindByID(c echo.Context) error {
	id := c.Param("id")
	resumeID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logrus.Error(err.Error())
		return c.JSON(http.StatusBadRequest, &Response{
			Message: "invalid resumeID",
		})
	}

	resume, err := r.resumeUsecase.FindByID(c.Request().Context(), resumeID)
	if err != nil {
		logrus.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, &Response{
			Message: "something went wrong",
		})
	}
	if resume == nil {
		return c.JSON(http.StatusNotFound, &Response{
			Message: "resume not found",
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Data: resume,
	})
}
