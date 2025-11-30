import { apiFetch } from "@/shared/api/fetchCongig";

type ResponseType = {
	err: string;
	statusCode: number;
};

export async function authCheck(): Promise<boolean> {
	try {
		const response = await apiFetch<ResponseType>("/api/users/me/profile", {
			method: "GET",
			cache: "no-cache",
		});

		return !response.err;
	} catch {
		return false;
	}
}
