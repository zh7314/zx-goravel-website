package admin

import (
	"errors"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type NewsService struct {
}

func NewNewsService() *NewsService {
	return &NewsService{}
}

func (r *NewsService) GetList(request requests.NewsRequest) (map[string]interface{}, error) {

	var list []models.News
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.NewsCateId) {
	orm = orm.Where("news_cate_id", request.NewsCateId)
}
if !gconv.IsEmpty(request.Title) {
	orm = orm.Where("title", request.Title)
}
if !gconv.IsEmpty(request.ShortTitle) {
	orm = orm.Where("short_title", request.ShortTitle)
}
if !gconv.IsEmpty(request.Content) {
	orm = orm.Where("content", request.Content)
}
if !gconv.IsEmpty(request.AdminId) {
	orm = orm.Where("admin_id", request.AdminId)
}
if !gconv.IsEmpty(request.Count) {
	orm = orm.Where("count", request.Count)
}
if !gconv.IsEmpty(request.IsShow) {
	orm = orm.Where("is_show", request.IsShow)
}
if !gconv.IsEmpty(request.Sort) {
	orm = orm.Where("sort", request.Sort)
}
if !gconv.IsEmpty(request.StartTime) {
	orm = orm.Where("start_time", request.StartTime)
}
if !gconv.IsEmpty(request.EndTime) {
	orm = orm.Where("end_time", request.EndTime)
}
if !gconv.IsEmpty(request.Pic) {
	orm = orm.Where("pic", request.Pic)
}
if !gconv.IsEmpty(request.Pics) {
	orm = orm.Where("pics", request.Pics)
}
if !gconv.IsEmpty(request.SeoTitle) {
	orm = orm.Where("seo_title", request.SeoTitle)
}
if !gconv.IsEmpty(request.SeoKeyword) {
	orm = orm.Where("seo_keyword", request.SeoKeyword)
}
if !gconv.IsEmpty(request.SeoDescription) {
	orm = orm.Where("seo_description", request.SeoDescription)
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

func (r *NewsService) GetAll(request requests.NewsRequest) ([]models.News, error) {

	var list []models.News

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.NewsCateId) {
	orm = orm.Where("news_cate_id", request.NewsCateId)
}
if !gconv.IsEmpty(request.Title) {
	orm = orm.Where("title", request.Title)
}
if !gconv.IsEmpty(request.ShortTitle) {
	orm = orm.Where("short_title", request.ShortTitle)
}
if !gconv.IsEmpty(request.Content) {
	orm = orm.Where("content", request.Content)
}
if !gconv.IsEmpty(request.AdminId) {
	orm = orm.Where("admin_id", request.AdminId)
}
if !gconv.IsEmpty(request.Count) {
	orm = orm.Where("count", request.Count)
}
if !gconv.IsEmpty(request.IsShow) {
	orm = orm.Where("is_show", request.IsShow)
}
if !gconv.IsEmpty(request.Sort) {
	orm = orm.Where("sort", request.Sort)
}
if !gconv.IsEmpty(request.StartTime) {
	orm = orm.Where("start_time", request.StartTime)
}
if !gconv.IsEmpty(request.EndTime) {
	orm = orm.Where("end_time", request.EndTime)
}
if !gconv.IsEmpty(request.Pic) {
	orm = orm.Where("pic", request.Pic)
}
if !gconv.IsEmpty(request.Pics) {
	orm = orm.Where("pics", request.Pics)
}
if !gconv.IsEmpty(request.SeoTitle) {
	orm = orm.Where("seo_title", request.SeoTitle)
}
if !gconv.IsEmpty(request.SeoKeyword) {
	orm = orm.Where("seo_keyword", request.SeoKeyword)
}
if !gconv.IsEmpty(request.SeoDescription) {
	orm = orm.Where("seo_description", request.SeoDescription)
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

func (r *NewsService) GetOne(id int64) (res models.News, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var news models.News
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&news)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return news, nil
}

func (r *NewsService) Add(request requests.NewsRequest) (bool, error) {

	var news models.News

	news.NewsCateId = request.NewsCateId
news.Title = html.EscapeString(request.Title)
news.ShortTitle = html.EscapeString(request.ShortTitle)
news.Content = html.EscapeString(request.Content)
news.AdminId = request.AdminId
news.Count = request.Count
news.IsShow = request.IsShow
news.Sort = request.Sort
news.StartTime = request.StartTime
news.EndTime = request.EndTime
news.Pic = html.EscapeString(request.Pic)
news.Pics = html.EscapeString(request.Pics)
news.SeoTitle = html.EscapeString(request.SeoTitle)
news.SeoKeyword = html.EscapeString(request.SeoKeyword)
news.SeoDescription = html.EscapeString(request.SeoDescription)
news.Platform = html.EscapeString(request.Platform)
news.Lang = html.EscapeString(request.Lang)


	err := facades.Orm().Query().Create(&news)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *NewsService) Save(request requests.NewsRequest) (bool, error) {

	var news models.News

	news.ID = request.ID
	news.NewsCateId = request.NewsCateId
news.Title = html.EscapeString(request.Title)
news.ShortTitle = html.EscapeString(request.ShortTitle)
news.Content = html.EscapeString(request.Content)
news.AdminId = request.AdminId
news.Count = request.Count
news.IsShow = request.IsShow
news.Sort = request.Sort
news.StartTime = request.StartTime
news.EndTime = request.EndTime
news.Pic = html.EscapeString(request.Pic)
news.Pics = html.EscapeString(request.Pics)
news.SeoTitle = html.EscapeString(request.SeoTitle)
news.SeoKeyword = html.EscapeString(request.SeoKeyword)
news.SeoDescription = html.EscapeString(request.SeoDescription)
news.Platform = html.EscapeString(request.Platform)
news.Lang = html.EscapeString(request.Lang)


	err := facades.Orm().Query().Save(&news)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *NewsService) Delete(id int64) (bool, error) {

	var admin models.News
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
