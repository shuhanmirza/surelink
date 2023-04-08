package structs

type GetHomePageStatsResponse struct {
	NumUrlMapCreatedLifetime    int64 `json:"num_url_map_created_lifetime"`
	NumUrlMapRedirectedLifetime int64 `json:"num_url_map_redirected_lifetime"`
}
