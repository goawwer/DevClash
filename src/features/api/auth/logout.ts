import { User } from "@/entities/user.interface";
import { apiFetch } from "@/shared/api/fetchCongig";

type ResponseType = {
	err: string;
	statusCode: number;
};

export async function logout(): Promise<ResponseType> {
	const response = await apiFetch<ResponseType>("/api/logout", {
		headers: {},
		method: "POST",
	});

	console.log(response);

	return response;
}
