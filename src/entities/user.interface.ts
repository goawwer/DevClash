export interface User {
	username: string;
	email: string;
	password: string;
}

export interface UserProfileDto {
	username: string;
	email?: string;
	profile_picture_url: string | null;
	bio: string | null;
	profile_status?: string | null;
	participations_count?: number;
	wins_count?: number;
	tech_stack: string[] | null;
}
