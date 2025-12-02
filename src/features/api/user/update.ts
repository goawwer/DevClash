// @/features/api/userApi.ts
import { UserProfileDto } from "@/entities/user.interface";
import { apiFetch } from "@/shared/api/fetchCongig";

export async function update(
	userData: Partial<UserProfileDto>,
	picture?: File
): Promise<UserProfileDto> {
	const formData = new FormData();

	// Собираем текстовые поля в объект payload
	const payload: Record<string, unknown> = {};

	if (userData.username != null) payload.username = userData.username.trim();
	if (userData.email != null) payload.email = userData.email.trim();
	if (userData.bio != null) payload.bio = userData.bio.trim();
	if (userData.profile_status != null)
		payload.profile_status = userData.profile_status.trim();

	if (userData.tech_stack) {
		payload.tech_stack = userData.tech_stack
			.map((tag) => tag.trim())
			.filter((tag) => tag !== "")
			.map(
				(tag) =>
					tag.charAt(0).toUpperCase() + tag.slice(1).toLowerCase()
			);

		console.log(payload.tech_stack);
	}
	// Позже, при отправке:

	if (userData.profile_picture_url != null)
		payload.profile_picture_url = userData.profile_picture_url;

	// Добавляем payload как строку JSON
	formData.append("payload", JSON.stringify(payload));

	// Добавляем файл, если он есть
	if (picture) {
		formData.append("picture", picture);
	}

	// 🔥 Не устанавливаем Content-Type вручную — браузер сам установит multipart/form-data с правильным boundary
	const response = await apiFetch<UserProfileDto>("/api/users/me", {
		method: "PUT",
		body: formData,
		cache: "no-cache",
	});

	return response;
}
