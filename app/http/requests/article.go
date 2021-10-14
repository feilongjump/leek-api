package requests

type ArticleForm struct {
	Title    string `json:"title" binding:"required"`
	Markdown string `json:"markdown"`
	Html     string `json:"html"`
}
