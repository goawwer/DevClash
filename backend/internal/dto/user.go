package dto

type UserProfile struct {
	Username            string   `json:"username"`
	ProfilePictureURL   *string  `json:"profile_picture_url"`
	Bio                 *string  `json:"bio"`
	ProfileStatus       *string  `json:"profile_status"`
	ParticipationsCount int      `json:"participations_count"`
	WinsCount           int      `json:"wins_count"`
	TechStack           []string `json:"tech_stack"`
}

type UserGetProfileSettings struct {
	Username          string   `json:"username"`
	Email             string   `json:"email"`
	Bio               *string  `json:"bio"`
	ProfileStatus     *string  `json:"profile_status"`
	TechStack         []string `json:"tech_stack"`
	ProfilePictureURL *string  `json:"profile_picture_url"`
}

type UserUpdateProfileSettings struct {
	Username          string   `json:"username"`
	Email             string   `json:"email"`
	Bio               *string  `json:"bio"`
	ProfileStatus     *string  `json:"profile_status"`
	TechStack         []string `json:"tech_stack"`
	OldPassword       string   `json:"old_password"`
	NewPassword       string   `json:"new_password"`
	ProfilePictureURL *string  `json:"profile_picture_url"`
}
