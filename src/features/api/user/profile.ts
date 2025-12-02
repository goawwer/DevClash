import { User, UserProfileDto } from "@/entities/user.interface";
import { apiFetch } from "@/shared/api/fetchCongig";

export async function profile(): Promise<UserProfileDto> {
	const response = await apiFetch<UserProfileDto>("/api/users/me/profile", {
		method: "GET",
		cache: "no-cache",
	});

	return response;
}
