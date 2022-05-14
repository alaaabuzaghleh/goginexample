package responses

type Group struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty" validate:"required"`
	ArName   string `json:"arName,omitempty" validate:"required"`
	Desc     string `json:"desc,omitempty" validate:"required"`
	ArDesc   string `json:"arDesc,omitempty" validate:"required"`
	SmallPic string `json:"smallPic,omitempty" validate:"required"`
	LargePic string `json:"laregePic,omitempty" validate:"required"`
}
