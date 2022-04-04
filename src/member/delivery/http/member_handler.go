package http

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sendaljpt/go-clean/src/domain"
	"github.com/sendaljpt/go-clean/src/helper"
	"github.com/sirupsen/logrus"
)

type ResponseError struct {
	Message string `json:"message"`
}

type MemberHandler struct {
	MemberUsecase domain.MemberUsecase
}

func NewMemberHandler(e *echo.Echo, ms domain.MemberUsecase) {
	handler := &MemberHandler{
		MemberUsecase: ms,
	}

	e.GET("/member", handler.FetchMember)
}

func (m *MemberHandler) FetchMember(c echo.Context) error {
	ctx := c.Request().Context()

	listMember, err := m.MemberUsecase.Fetch(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err), helper.Response(getStatusCode(err), "Failed get member", nil, err.Error()))
	}

	data := helper.Response(200, "Success get member", listMember, nil)

	return c.JSON(http.StatusOK, data)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
