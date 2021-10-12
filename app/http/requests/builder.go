package requests

type BuilderForm struct {
	Name string `json:"name" binding:"required"`
	Uri  string `json:"uri" binding:"required"`
}
