package admin

import (
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type AdminGroupService struct {
}

func NewAdminGroupService() *AdminGroupService {
	return &AdminGroupService{}
}

func (r *AdminGroupService) GetList(request requests.AdminGroupRequest) (map[string]interface{}, error) {

	var list []models.AdminGroup
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.ParentId) {
	orm.Where("parent_id", request.ParentId)
}
if !gconv.IsEmpty(request.Name) {
	orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.PermissionIds) {
	orm.Where("permission_ids", request.PermissionIds)
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

func (r *AdminGroupService) GetAll(request requests.AdminGroupRequest) ([]models.AdminGroup, error) {

	var list []models.AdminGroup

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.ParentId) {
	orm.Where("parent_id", request.ParentId)
}
if !gconv.IsEmpty(request.Name) {
	orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.PermissionIds) {
	orm.Where("permission_ids", request.PermissionIds)
}
if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
}


	orm.Order("id desc").Get(&list)

	return list, nil
}

func (r *AdminGroupService) Add(request requests.AdminGroupRequest) (bool, error) {

	var adminGroup models.AdminGroup

	adminGroup.ParentId = request.ParentId
adminGroup.Name = html.EscapeString(request.Name)
adminGroup.PermissionIds = html.EscapeString(request.PermissionIds)
adminGroup.Sort = request.Sort


	err := facades.Orm().Query().Create(&adminGroup)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *AdminGroupService) Save(request requests.AdminGroupRequest) (bool, error) {

	var adminGroup models.AdminGroup

	adminGroup.ID = request.ID
	adminGroup.ParentId = request.ParentId
adminGroup.Name = html.EscapeString(request.Name)
adminGroup.PermissionIds = html.EscapeString(request.PermissionIds)
adminGroup.Sort = request.Sort


	err := facades.Orm().Query().Save(&adminGroup)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *AdminGroupService) Delete(id int64) (bool, error) {

	var admin models.AdminGroup
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}