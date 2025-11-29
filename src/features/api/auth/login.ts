import { User } from "@/entities/user.interface";
import { apiFetch } from "@/shared/api/fetchCongig";

type ResponseType = {
	err: string;
	statusCode: number;
};

export async function login(
	user: Pick<User, "email" | "password">
): Promise<ResponseType> {
	console.log(JSON.stringify(user));
	const response = await apiFetch<ResponseType>("/auth/login", {
		headers: {},
		method: "POST",
		body: JSON.stringify(user),
	});

	console.log(response);

	return response;
}
