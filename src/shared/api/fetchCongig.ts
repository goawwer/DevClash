export async function apiFetch<T>(
	url: string,
	options: RequestInit = {}
): Promise<T> {
	const { headers, body, cache, ...rest } = options;

	const checkedHeaders =
		body instanceof FormData
			? headers || {}
			: { "Content-Type": "application/json", ...(headers || {}) };

	const response = await fetch("http://localhost:8080" + url, {
		credentials: "include",
		headers: checkedHeaders,
		cache: cache ?? "no-cache",
		body: body,
		...rest,
	});

	// Неуспешные статус-коды → ошибка
	if (!response.ok) {
		throw new Error(`API error: ${response.status}`);
	}

	// Пытаемся определить, есть ли JSON
	const contentType = response.headers.get("content-type") ?? "";
	const text = await response.text();

	// Если тело пустое → возвращаем пустой объект
	if (!text) {
		return {} as T;
	}

	// Если это JSON — парсим
	if (contentType.includes("application/json")) {
		try {
			return JSON.parse(text) as T;
		} catch {
			throw new Error("Invalid JSON from API");
		}
	}

	// Если это не JSON — возвращаем как строку
	return text as unknown as T;
}
