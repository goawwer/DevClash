export async function apiFetch<T>(
	url: string,
	options: RequestInit = {}
): Promise<T> {
	const { headers, body, cache, ...rest } = options;

	const checkedHeaders =
		body instanceof FormData
			? headers || {}
			: { "Content-Type": "application/json", ...(headers || {}) };

	const response = await fetch(process.env.NEXT_PUBLIC_API_URL + url, {
		credentials: "include",
		headers: checkedHeaders,
		cache: cache ?? "no-cache",
		body: body,
		...rest,
	});

	if (!response.ok) {
		throw new Error(`API error: ${response.status}`);
	}

	const text = await response.text();
	return text ? JSON.parse(text) : (response.status as T);
}
