package user

type UserProfile struct {
	Username            string  `json:"username"`
	ProfilePictureURL   *string `json:"profile_picture_url"`
	Bio                 *string `json:"bio"`
	ProfileStatus       *string `json:"profile_status"`
	ParticipationsCount int     `json:"participations_count"`
	WinsCount           int     `json:"wins_count"`
}
