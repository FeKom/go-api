package jsondto

type JsonToStruct struct {
	Result struct {
		Leagues []struct {
			ID     string
			Name   string
			Rounds []struct {
				ID       string
				Name     string
				Fixtures []struct {
					ID           int
					Name         string
					Participants []struct {
						ID        int
						SportID   int
						CountryID int
						VenueID   int
						Gender    string
						Name      string
						ShortCode string
						//ImagePath string `json:"image_path"`
					}
				}
			}
		}
	}
}
