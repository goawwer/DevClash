import { NextRequest, NextResponse } from "next/server";
import { apiFetch } from "@/shared/api/fetchCongig";

type ResponseType = {
	err: string;
	statusCode: number;
};

export async function authCheck(req: NextRequest) {
	const cookies = req.headers.get("cookie");

	try {
		console.log("tryng");
		const response = await apiFetch<ResponseType>("/api/users/me/profile", {
			method: "GET",
			headers: { ...(cookies ? { Cookie: cookies } : {}) },
			cache: "no-cache",
		});

		if (!response.err) {
			return;
		}
	} catch (error: any) {
		// apiFetch выбросил ошибку → статус НЕ 2xx
		const status = error.message.includes("API error: ")
			? parseInt(error.message.split("API error: ")[1], 10)
			: 500;
		console.log("error", status);
	}

	try {
		console.log("refresh");
		const refreshRes = await fetch("http://localhost:8080/auth/refresh", {
			method: "POST",
			headers: { ...(cookies ? { Cookie: cookies } : {}) },
			cache: "no-cache",
		});

		if (refreshRes.status === 200) {
			// Получаем ВСЕ куки из заголовка Set-Cookie
			const setCookies = refreshRes.headers.getSetCookie();

			// Клонируем текущий URL, чтобы редиректнуть на него же
			const url = req.nextUrl.clone();

			// Создаём редирект-ответ
			const response = NextResponse.redirect(url);

			// Устанавливаем КАЖДУЮ куку в заголовки ответа
			for (const cookie of setCookies) {
				response.headers.append("set-cookie", cookie);
			}

			return response; // 👈 ВАЖНО: возвращаем ответ с куками и редиректом
		} else {
			console.log("Refresh failed with status:", refreshRes.status);
			return redirectToRoot(req);
		}
	} catch {
		console.log("Refresh request error:");
		return redirectToRoot(req);
	}
}

function redirectToRoot(req: NextRequest) {
	const url = req.nextUrl.clone();
	url.pathname = "/";
	return NextResponse.redirect(url);
}
