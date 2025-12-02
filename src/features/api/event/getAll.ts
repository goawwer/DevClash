import { Event } from "@/entities/event.interface";
import { apiFetch } from "@/shared/api/fetchCongig";

export async function getAll(): Promise<Event[]> {
	const response = await apiFetch<Event[]>("/api/events/all", {
		method: "GET",
		cache: "no-cache",
	});

	return response;
}
