package structs

type GetMapRequest struct {
	Uid string `form:"uid" json:"uid" binding:"required"`
}

type GetMapResponse struct {
	Url string `json:"url"`
}

type SetMapRequest struct {
	CaptchaUuid  string `json:"captcha_uuid" binding:"required"`
	CaptchaValue string `json:"captcha_value" binding:"required"`
	Url          string `json:"url" binding:"required"`
}

type SetMapResponse struct {
	ShortUrl string `json:"short_url"`
}
