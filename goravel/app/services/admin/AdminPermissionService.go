package admin

import (
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type AdminPermissionService struct {
}

func NewAdminPermissionService() *AdminPermissionService {
	return &AdminPermissionService{}
}

func (r *AdminPermissionService) GetList(request requests.AdminPermissionRequest) (map[string]interface{}, error) {

	var list []models.AdminPermission
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Component) {
	orm.Where("component", request.Component)
}
if !gconv.IsEmpty(request.Hidden) {
	orm.Where("hidden", request.Hidden)
}
if !gconv.IsEmpty(request.Icon) {
	orm.Where("icon", request.Icon)
}
if !gconv.IsEmpty(request.IsMenu) {
	orm.Where("is_menu", request.IsMenu)
}
if !gconv.IsEmpty(request.Name) {
	orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.ParentId) {
	orm.Where("parent_id", request.ParentId)
}
if !gconv.IsEmpty(request.Path) {
	orm.Where("path", request.Path)
}
if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
}


	orm.Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *AdminPermissionService) GetAll(request requests.AdminPermissionRequest) ([]models.AdminPermission, error) {

	var list []models.AdminPermission

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.Component) {
	orm.Where("component", request.Component)
}
if !gconv.IsEmpty(request.Hidden) {
	orm.Where("hidden", request.Hidden)
}
if !gconv.IsEmpty(request.Icon) {
	orm.Where("icon", request.Icon)
}
if !gconv.IsEmpty(request.IsMenu) {
	orm.Where("is_menu", request.IsMenu)
}
if !gconv.IsEmpty(request.Name) {
	orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.ParentId) {
	orm.Where("parent_id", request.ParentId)
}
if !gconv.IsEmpty(request.Path) {
	orm.Where("path", request.Path)
}
if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
}


	orm.Order("id desc").Get(&list)

	return list, nil
}

func (r *AdminPermissionService) Add(request requests.AdminPermissionRequest) (bool, error) {

	var adminPermission models.AdminPermission

	adminPermission.Component = html.EscapeString(request.Component)
adminPermission.Hidden = request.Hidden
adminPermission.Icon = html.EscapeString(request.Icon)
adminPermission.IsMenu = request.IsMenu
adminPermission.Name = html.EscapeString(request.Name)
adminPermission.ParentId = request.ParentId
adminPermission.Path = html.EscapeString(request.Path)
adminPermission.Sort = request.Sort


	err := facades.Orm().Query().Create(&adminPermission)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *AdminPermissionService) Save(request requests.AdminPermissionRequest) (bool, error) {

	var adminPermission models.AdminPermission

	adminPermission.ID = request.ID
	adminPermission.Component = html.EscapeString(request.Component)
adminPermission.Hidden = request.Hidden
adminPermission.Icon = html.EscapeString(request.Icon)
adminPermission.IsMenu = request.IsMenu
adminPermission.Name = html.EscapeString(request.Name)
adminPermission.ParentId = request.ParentId
adminPermission.Path = html.EscapeString(request.Path)
adminPermission.Sort = request.Sort


	err := facades.Orm().Query().Save(&adminPermission)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *AdminPermissionService) Delete(id int64) (bool, error) {

	var admin models.AdminPermission
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
