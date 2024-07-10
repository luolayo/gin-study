package Model

type Test struct {
	Msg string `form:"msg" json:"msg" binding:"required"`
}
