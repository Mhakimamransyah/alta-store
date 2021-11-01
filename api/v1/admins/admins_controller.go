package admins

import (
	"altaStore/api/common"
	"altaStore/api/middleware"
	"altaStore/api/v1/admins/response"
	"altaStore/business/admins"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	adminService admins.Service
}

func InitAdminController(service admins.Service) *Controller {
	return &Controller{
		adminService: service,
	}
}

func (admin_service *Controller) LoginController(c echo.Context) error {
	admin, err := admin_service.adminService.LoginAdmin(c.FormValue("username"), c.FormValue("password"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessResponse(response.ConvertAdminToAdminLogin(admin)))
}

func (admin_service *Controller) GetAdminController(c echo.Context) error {

	page, errorPage := strconv.Atoi(c.QueryParam("page"))
	per_page, errorPerPage := strconv.Atoi(c.QueryParam("per_page"))
	if errorPage != nil {
		page = 0
	}
	if errorPerPage != nil {
		per_page = 100
	}
	list_admin, err := admin_service.adminService.FindAllAdmin(page, per_page)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessResponse(response.ConvertListAdminToDataAdmin(list_admin)))
}

func (admin_service *Controller) GetAdminByUsername(c echo.Context) error {
	admin, err := admin_service.adminService.FindAdminByUsername(c.Param("username"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessResponse(response.ConvertAdminToDataAdmin(admin)))
}

func (admin_service *Controller) CreateAdminController(c echo.Context) error {
	admin_spec := admins.AdminSpec{}
	c.Bind(&admin_spec)

	id_admin := int(middleware.ExtractTokenKey(c, "id").(float64))
	err := admin_service.adminService.InsertAdmin(admin_spec, id_admin)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessCreated())
}

func (admin_service *Controller) ModifyAdminController(c echo.Context) error {
	admin_updatable := admins.AdminUpdatable{}
	c.Bind(&admin_updatable)
	id_admin := int(middleware.ExtractTokenKey(c, "id").(float64))
	err := admin_service.adminService.ModifyAdmin(c.Param("username"), admin_updatable, id_admin)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessResponseWithoutData())
}
