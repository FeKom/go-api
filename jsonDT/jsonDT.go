package jsondt

type SearchApiResponse struct {
	Data []struct {
		ID           int    `json:"id"`
		SportID      int    `json:"sport_id"`
		CountryID    int    `json:"country_id"`
		VenueID      int    `json:"venue_id"`
		Gender       string `json:"gender"`
		Name         string `json:"name"`
		ShortCode    string `json:"short_code"`
		ImagePath    string `json:"image_path"`
		Founded      int    `json:"founded"`
		Type         string `json:"type"`
		Placeholder  bool   `json:"placeholder"`
		LastPlayedAt string `json:"last_played_at"`
	} `json:"data"`
	Pagination struct {
		Count       int  `json:"count"`
		PerPage     int  `json:"per_page"`
		CurrentPage int  `json:"current_page"`
		NextPage    any  `json:"next_page"`
		HasMore     bool `json:"has_more"`
	} `json:"pagination"`
	Subscription []struct {
		Meta  []any `json:"meta"`
		Plans []struct {
			Plan     string `json:"plan"`
			Sport    string `json:"sport"`
			Category string `json:"category"`
		} `json:"plans"`
		AddOns  []any `json:"add_ons"`
		Widgets []any `json:"widgets"`
	} `json:"subscription"`
	RateLimit struct {
		ResetsInSeconds int    `json:"resets_in_seconds"`
		Remaining       int    `json:"remaining"`
		RequestedEntity string `json:"requested_entity"`
	} `json:"rate_limit"`
	Timezone string `json:"timezone"`
}
