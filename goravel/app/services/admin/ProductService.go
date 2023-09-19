package admin

import (
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type ProductService struct {
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (r *ProductService) GetList(request requests.ProductRequest) (map[string]interface{}, error) {

	var list []models.Product
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.AdminId) {
	orm.Where("admin_id", request.AdminId)
}
if !gconv.IsEmpty(request.Content) {
	orm.Where("content", request.Content)
}
if !gconv.IsEmpty(request.EndTime) {
	orm.Where("end_time", request.EndTime)
}
if !gconv.IsEmpty(request.IsShow) {
	orm.Where("is_show", request.IsShow)
}
if !gconv.IsEmpty(request.Lang) {
	orm.Where("lang", request.Lang)
}
if !gconv.IsEmpty(request.Pic) {
	orm.Where("pic", request.Pic)
}
if !gconv.IsEmpty(request.Pics) {
	orm.Where("pics", request.Pics)
}
if !gconv.IsEmpty(request.Platform) {
	orm.Where("platform", request.Platform)
}
if !gconv.IsEmpty(request.ProductCateId) {
	orm.Where("product_cate_id", request.ProductCateId)
}
if !gconv.IsEmpty(request.ShortTitle) {
	orm.Where("short_title", request.ShortTitle)
}
if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
}
if !gconv.IsEmpty(request.StartTime) {
	orm.Where("start_time", request.StartTime)
}
if !gconv.IsEmpty(request.Title) {
	orm.Where("title", request.Title)
}
if !gconv.IsEmpty(request.Url) {
	orm.Where("url", request.Url)
}
if !gconv.IsEmpty(request.VideoUrl) {
	orm.Where("video_url", request.VideoUrl)
}
if !gconv.IsEmpty(request.ViewCount) {
	orm.Where("view_count", request.ViewCount)
}


	orm.Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *ProductService) GetAll(request requests.ProductRequest) ([]models.Product, error) {

	var list []models.Product

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.AdminId) {
	orm.Where("admin_id", request.AdminId)
}
if !gconv.IsEmpty(request.Content) {
	orm.Where("content", request.Content)
}
if !gconv.IsEmpty(request.EndTime) {
	orm.Where("end_time", request.EndTime)
}
if !gconv.IsEmpty(request.IsShow) {
	orm.Where("is_show", request.IsShow)
}
if !gconv.IsEmpty(request.Lang) {
	orm.Where("lang", request.Lang)
}
if !gconv.IsEmpty(request.Pic) {
	orm.Where("pic", request.Pic)
}
if !gconv.IsEmpty(request.Pics) {
	orm.Where("pics", request.Pics)
}
if !gconv.IsEmpty(request.Platform) {
	orm.Where("platform", request.Platform)
}
if !gconv.IsEmpty(request.ProductCateId) {
	orm.Where("product_cate_id", request.ProductCateId)
}
if !gconv.IsEmpty(request.ShortTitle) {
	orm.Where("short_title", request.ShortTitle)
}
if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
}
if !gconv.IsEmpty(request.StartTime) {
	orm.Where("start_time", request.StartTime)
}
if !gconv.IsEmpty(request.Title) {
	orm.Where("title", request.Title)
}
if !gconv.IsEmpty(request.Url) {
	orm.Where("url", request.Url)
}
if !gconv.IsEmpty(request.VideoUrl) {
	orm.Where("video_url", request.VideoUrl)
}
if !gconv.IsEmpty(request.ViewCount) {
	orm.Where("view_count", request.ViewCount)
}


	orm.Order("id desc").Get(&list)

	return list, nil
}

func (r *ProductService) Add(request requests.ProductRequest) (bool, error) {

	var product models.Product

	product.AdminId = request.AdminId
product.Content = html.EscapeString(request.Content)
product.EndTime = request.EndTime
product.IsShow = request.IsShow
product.Lang = html.EscapeString(request.Lang)
product.Pic = html.EscapeString(request.Pic)
product.Pics = html.EscapeString(request.Pics)
product.Platform = html.EscapeString(request.Platform)
product.ProductCateId = request.ProductCateId
product.ShortTitle = html.EscapeString(request.ShortTitle)
product.Sort = request.Sort
product.StartTime = request.StartTime
product.Title = html.EscapeString(request.Title)
product.Url = html.EscapeString(request.Url)
product.VideoUrl = html.EscapeString(request.VideoUrl)
product.ViewCount = request.ViewCount


	err := facades.Orm().Query().Create(&product)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *ProductService) Save(request requests.ProductRequest) (bool, error) {

	var product models.Product

	product.ID = request.ID
	product.AdminId = request.AdminId
product.Content = html.EscapeString(request.Content)
product.EndTime = request.EndTime
product.IsShow = request.IsShow
product.Lang = html.EscapeString(request.Lang)
product.Pic = html.EscapeString(request.Pic)
product.Pics = html.EscapeString(request.Pics)
product.Platform = html.EscapeString(request.Platform)
product.ProductCateId = request.ProductCateId
product.ShortTitle = html.EscapeString(request.ShortTitle)
product.Sort = request.Sort
product.StartTime = request.StartTime
product.Title = html.EscapeString(request.Title)
product.Url = html.EscapeString(request.Url)
product.VideoUrl = html.EscapeString(request.VideoUrl)
product.ViewCount = request.ViewCount


	err := facades.Orm().Query().Save(&product)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *ProductService) Delete(id int64) (bool, error) {

	var admin models.Product
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
