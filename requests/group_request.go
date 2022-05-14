package requests

import (
	"mime/multipart"
)

type Group struct {
	Name     string                `form:"name" binding:"required"`
	ArName   string                `form:"arName" binding:"required"`
	Desc     string                `form:"desc" binding:"required"`
	ArDesc   string                `form:"arDesc" binding:"required"`
	SmallPic *multipart.FileHeader `form:"smallPic" binding:"required"`
	LargePic *multipart.FileHeader `form:"largePic" binding:"required"`
}
