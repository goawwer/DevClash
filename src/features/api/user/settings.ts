import { UserProfileDto } from "@/entities/user.interface";
import { apiFetch } from "@/shared/api/fetchCongig";

export async function settings(): Promise<UserProfileDto> {
	const response = await apiFetch<UserProfileDto>("/api/users/me/settings", {
		method: "GET",
		cache: "no-cache",
	});

	return response;
}
