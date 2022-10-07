package service

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"inwind-blog-server-v3/common/request"
	"inwind-blog-server-v3/common/serializer"
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/interner/model"
	"inwind-blog-server-v3/utils/dto"
)

type BlogService struct{}

func (b *BlogService) GetBlogList(params request.PageRequest) (blogs []model.Blog, total int64, err error) {

	var pageSize, pageIndex = params.PageSize, params.PageIndex
	// limit -1 offset -1 表示查询总的条数
	err = global.DB.Order("id desc").Scopes(dto.Paginate(pageIndex, pageSize)).Preload("Tags").Where("hidden = ?", 0).Find(&blogs).Limit(-1).Offset(-1).Count(&total).Error

	return blogs, total, err
}

func (b *BlogService) GetBlogDetail(params request.SelectBlogRequest) (blog model.Blog, err error) {
	var id = params.Id
	err = global.DB.Where("id = ?", id).Preload("Tags").First(&blog).Error
	return blog, err
}

func (b *BlogService) EditBlog(params request.EditBlog) error {
	//查询所有标签

	var allTag []model.Tag
	global.DB.Find(&allTag)

	// 关联模式 的更新
	newParams := serializer.BuildEditBlogParams(params, allTag)
	err := global.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&newParams).Error

	return err
}

func (b *BlogService) CreateBlog(params request.CreateBlog) error {
	newParams := serializer.BuildCreateBlogParams(params)
	err := global.DB.Create(&newParams).Error

	return err
}

func (b *BlogService) DeleteBlog(id uint) error {
	err := global.DB.Select(clause.Associations).Delete(&model.Blog{BasicModel: model.BasicModel{ID: id}}).Error
	return err
}
