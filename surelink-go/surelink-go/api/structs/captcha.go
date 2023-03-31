package structs

type GetCaptchaResponse struct {
	Uuid string `json:"uuid"`
	Img  string `json:"img"`
}
