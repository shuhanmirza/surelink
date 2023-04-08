package infrastructure

import "encoding/json"

type CaptchaModel struct {
	Val    string
	ImgB64 string
}

type HomePageStatModel struct {
	NumUrlMapCreatedLifetime    int64 `redis:"num_url_map_created_lifetime"`
	NumUrlMapRedirectedLifetime int64 `redis:"num_url_map_redirected_lifetime"`
}

func (model HomePageStatModel) MarshalBinary() ([]byte, error) {
	return json.Marshal(model)
}
