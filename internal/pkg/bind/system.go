package bind

type BodySystemDatabase struct {
	Driver string `form:"driver" json:"driver" binding:"required"`
	DSN    string `form:"dsn" json:"dsn" binding:"required"`
}
