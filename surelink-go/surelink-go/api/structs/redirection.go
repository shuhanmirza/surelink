package structs

import "time"

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

type SetMapRequestV2 struct {
	RecaptchaToken string `json:"recaptcha_token" binding:"required"`
	Url            string `json:"url" binding:"required"`
}

type SetMapResponse struct {
	ShortUrl string `json:"short_url"`
}

type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	ChallengeTs time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
}
