"use client";

import { Event } from "@/entities/event.interface";
import React, { useEffect, useState } from "react";
import { getAll } from "@/features/api/event/getAll";

export default function AllEvents() {
	const [events, setEvents] = useState<Event[] | undefined>(undefined);

	useEffect(() => {
		async function getEvents() {
			const data = await getAll();
			console.log("events:", data);
			setEvents(data);
		}

		getEvents();
	}, []);

	return <div>{/* рендер событий */}</div>;
}
