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
	e.PUT("/resumes/:id", handler.Update)
	e.DELETE("/resumes/:id", handler.Delete)
	e.GET("/resumes/:id", handler.FindByID)
	e.GET("/resumes", handler.FindAllByFilter)
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

	resume, err := r.resumeUsecase.Create(c.Request().Context(), *request)
	if err != nil {
		logrus.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, &Response{
			Message: "something went wrong",
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Message: "successfully add a resume",
		Data:    resume,
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

	updatedResume, err := r.resumeUsecase.Update(c.Request().Context(), resumeID, *request)
	if err != nil {
		logrus.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, &Response{
			Message: "something went wrong",
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Message: "successfully update a resume",
		Data:    updatedResume,
	})
}

func (r *resumeHandler) Delete(c echo.Context) error {
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

	err = r.resumeUsecase.Delete(c.Request().Context(), resumeID)
	if err != nil {
		logrus.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, &Response{
			Message: "something went wrong",
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Message: "successfully deleted a resume",
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

func (r *resumeHandler) FindAllByFilter(c echo.Context) error {
	pageParam := c.QueryParam("page")
	sizeParam := c.QueryParam("size")

	if pageParam == "" || sizeParam == "" {
		return c.JSON(http.StatusBadRequest, &Response{
			Message: "page and size query required",
		})
	}

	page, err := strconv.ParseInt(pageParam, 10, 64)
	if err != nil {
		logrus.Error(err.Error())
		return c.JSON(http.StatusBadRequest, &Response{
			Message: "invalid page",
		})
	}

	size, err := strconv.ParseInt(sizeParam, 10, 64)
	if err != nil {
		logrus.Error(err.Error())
		return c.JSON(http.StatusBadRequest, &Response{
			Message: "invalid size",
		})
	}

	resumes, err := r.resumeUsecase.FindAllByFilter(c.Request().Context(), model.GetResumeFilter{
		Page: page,
		Size: size,
	})
	if err != nil {
		logrus.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, &Response{
			Message: "something went wrong",
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Data: resumes,
	})
}
