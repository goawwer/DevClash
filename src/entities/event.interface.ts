export interface Event {
	title: string;
	type: string;
	is_online: boolean;
	is_free: boolean;
	number_of_teams: number;
	team_size: number;
	tech_stack: string[];
	description: string;
	prize: string;
	start_time: string; // ISO 8601 формат (например, "2025-11-30T12:48:39Z")
	end_time: string; // ISO 8601 формат
	picture: string;
}
