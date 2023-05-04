package structs

type GetLinkPreviewRequest struct {
	Uid string `form:"uid" json:"uid" binding:"required"`
}

type GetLinkPreviewResponse struct {
	Url         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
