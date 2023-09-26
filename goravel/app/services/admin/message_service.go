package admin

import (
	"errors"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type MessageService struct {
}

func NewMessageService() *MessageService {
	return &MessageService{}
}

func (r *MessageService) GetList(request requests.MessageRequest) (map[string]interface{}, error) {

	var list []models.Message
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Type) {
	orm = orm.Where("type", request.Type)
}
if !gconv.IsEmpty(request.Mobile) {
	orm = orm.Where("mobile", request.Mobile)
}
if !gconv.IsEmpty(request.RealName) {
	orm = orm.Where("real_name", request.RealName)
}
if !gconv.IsEmpty(request.Email) {
	orm = orm.Where("email", request.Email)
}
if !gconv.IsEmpty(request.Ip) {
	orm = orm.Where("ip", request.Ip)
}
if !gconv.IsEmpty(request.Status) {
	orm = orm.Where("status", request.Status)
}
if !gconv.IsEmpty(request.Title) {
	orm = orm.Where("title", request.Title)
}
if !gconv.IsEmpty(request.Content) {
	orm = orm.Where("content", request.Content)
}
if !gconv.IsEmpty(request.Pics) {
	orm = orm.Where("pics", request.Pics)
}
if !gconv.IsEmpty(request.IsSent) {
	orm = orm.Where("is_sent", request.IsSent)
}
if !gconv.IsEmpty(request.Remark) {
	orm = orm.Where("remark", request.Remark)
}
if !gconv.IsEmpty(request.Platform) {
	orm = orm.Where("platform", request.Platform)
}
if !gconv.IsEmpty(request.Lang) {
	orm = orm.Where("lang", request.Lang)
}


	orm.Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *MessageService) GetAll(request requests.MessageRequest) ([]models.Message, error) {

	var list []models.Message

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.Type) {
	orm = orm.Where("type", request.Type)
}
if !gconv.IsEmpty(request.Mobile) {
	orm = orm.Where("mobile", request.Mobile)
}
if !gconv.IsEmpty(request.RealName) {
	orm = orm.Where("real_name", request.RealName)
}
if !gconv.IsEmpty(request.Email) {
	orm = orm.Where("email", request.Email)
}
if !gconv.IsEmpty(request.Ip) {
	orm = orm.Where("ip", request.Ip)
}
if !gconv.IsEmpty(request.Status) {
	orm = orm.Where("status", request.Status)
}
if !gconv.IsEmpty(request.Title) {
	orm = orm.Where("title", request.Title)
}
if !gconv.IsEmpty(request.Content) {
	orm = orm.Where("content", request.Content)
}
if !gconv.IsEmpty(request.Pics) {
	orm = orm.Where("pics", request.Pics)
}
if !gconv.IsEmpty(request.IsSent) {
	orm = orm.Where("is_sent", request.IsSent)
}
if !gconv.IsEmpty(request.Remark) {
	orm = orm.Where("remark", request.Remark)
}
if !gconv.IsEmpty(request.Platform) {
	orm = orm.Where("platform", request.Platform)
}
if !gconv.IsEmpty(request.Lang) {
	orm = orm.Where("lang", request.Lang)
}


	orm.Order("id desc").Get(&list)

	return list, nil
}

func (r *MessageService) GetOne(id int64) (res models.Message, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var message models.Message
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&message)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return message, nil
}

func (r *MessageService) Add(request requests.MessageRequest) (bool, error) {

	var message models.Message

	message.Type = request.Type
message.Mobile = html.EscapeString(request.Mobile)
message.RealName = html.EscapeString(request.RealName)
message.Email = html.EscapeString(request.Email)
message.Ip = html.EscapeString(request.Ip)
message.Status = request.Status
message.Title = html.EscapeString(request.Title)
message.Content = html.EscapeString(request.Content)
message.Pics = html.EscapeString(request.Pics)
message.IsSent = request.IsSent
message.Remark = html.EscapeString(request.Remark)
message.Platform = html.EscapeString(request.Platform)
message.Lang = html.EscapeString(request.Lang)


	err := facades.Orm().Query().Create(&message)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *MessageService) Save(request requests.MessageRequest) (bool, error) {

	var message models.Message

	message.ID = request.ID
	message.Type = request.Type
message.Mobile = html.EscapeString(request.Mobile)
message.RealName = html.EscapeString(request.RealName)
message.Email = html.EscapeString(request.Email)
message.Ip = html.EscapeString(request.Ip)
message.Status = request.Status
message.Title = html.EscapeString(request.Title)
message.Content = html.EscapeString(request.Content)
message.Pics = html.EscapeString(request.Pics)
message.IsSent = request.IsSent
message.Remark = html.EscapeString(request.Remark)
message.Platform = html.EscapeString(request.Platform)
message.Lang = html.EscapeString(request.Lang)


	err := facades.Orm().Query().Save(&message)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *MessageService) Delete(id int64) (bool, error) {

	var admin models.Message
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
