package request

type Blog struct {
	Id       string `form:"id" json:"id" binding:"required"`
	Title    string `form:"title" json:"title" binding:"required"`
	Content  string `form:"content" json:"content" binding:"required"`
	Src      string `form:"src" json:"src" binding:"required"`
	TagId    int    `form:"tagId" json:"tagId" binding:"required"`
	Overview string `form:"overview" json:"overview" binding:"required"`
}

type CreateBlog struct {
	Title    string `form:"title" json:"title" binding:"required"`
	Content  string `form:"content" json:"content" binding:"required"`
	Src      string `form:"src" json:"src" binding:"required"`
	TagId    int    `form:"tagId" json:"tagId" binding:"required"`
	Overview string `form:"overview" json:"overview" binding:"required"`
}