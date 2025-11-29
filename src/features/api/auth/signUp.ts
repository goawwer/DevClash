import { Organizer } from "@/entities/organizer.interface";
import { User } from "@/entities/user.interface";
import { apiFetch } from "@/shared/api/fetchCongig";

type ResponseType = {
	err: string;
	statusCode: number;
};

export async function userSignUp(user: User): Promise<ResponseType> {
	console.log(JSON.stringify(user));
	const response = await apiFetch<ResponseType>("/auth/signup/user", {
		headers: {},
		method: "POST",
		body: JSON.stringify(user),
	});

	console.log(response);

	return response;
}

export async function organizerSignUp(
	organizer: Organizer
): Promise<ResponseType> {
	const organizerFormData = new FormData();
	organizerFormData.append("email", organizer.email);
	organizerFormData.append("name", organizer.name);
	organizerFormData.append("password", organizer.password);
	if (organizer.color) {
		organizerFormData.append("color", organizer.color);
	}

	if (organizer.logo) {
		organizerFormData.append("logo", organizer.logo);
	}

	console.log(organizer);
	const response = await apiFetch<ResponseType>("/auth/signup/organizer", {
		method: "POST",
		body: organizerFormData,
	});

	console.log(response);

	return response;
}
