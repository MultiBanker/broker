package application

import (
	"net/http"

	_ "github.com/MultiBanker/broker/src/models"
	_ "github.com/MultiBanker/broker/src/servers/adminhttp/dto"
)

// @Summary      user application
// @Description  user application
// @Tags         User
// @Produce      json
// @Security     ApiTokenAuth
// @Param        auto       body      dto.UserApplication  true  "add car"
// @Success      204
// @Failure      400           {object}  httperrors.Response
// @Failure      401           {object}  httperrors.Response
// @Failure      404           {object}  httperrors.Response
// @Failure      500           {object}  httperrors.Response
// @Router       /api/v1/users/application/ [post]
func (res resource) createApplication(w http.ResponseWriter, r *http.Request) {}

// @Summary      user application
// @Description  user application
// @Tags         User
// @Produce      json
// @Security     ApiTokenAuth
// @Param        id  path      string  true  "application id"
// @Success      200           {object} models.UserApplication
// @Failure      400           {object}  httperrors.Response
// @Failure      401           {object}  httperrors.Response
// @Failure      404           {object}  httperrors.Response
// @Failure      500           {object}  httperrors.Response
// @Router       /api/v1/users/application/{id} [get]
func (res resource) getApplication(w http.ResponseWriter, r *http.Request) {}

// @Summary      user application
// @Description  user application
// @Tags         User
// @Produce      json
// @Security     ApiTokenAuth
// @Param limit query int false "pagination limit"
// @Param skip query int false "pagination skip"
// @Success      200           {object} dto.UserApplications
// @Failure      400           {object}  httperrors.Response
// @Failure      401           {object}  httperrors.Response
// @Failure      404           {object}  httperrors.Response
// @Failure      500           {object}  httperrors.Response
// @Router       /api/v1/users/application/list [get]
func (res resource) listApplications(w http.ResponseWriter, r *http.Request) {}

// @Summary      user application
// @Description  user application
// @Tags         User
// @Produce      json
// @Security     ApiTokenAuth
// @Param        id  path      string  true  "application id"
// @Success      204
// @Failure      400           {object}  httperrors.Response
// @Failure      401           {object}  httperrors.Response
// @Failure      404           {object}  httperrors.Response
// @Failure      500           {object}  httperrors.Response
// @Router       /api/v1/users/application/{id} [delete]
func (res resource) deleteApplication(w http.ResponseWriter, r *http.Request) {}
