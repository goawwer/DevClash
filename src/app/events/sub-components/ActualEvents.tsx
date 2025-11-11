import { BaseButton } from "@/shared/ui";
import styles from "./ActualEvents.module.scss";
import Carusel from "@/widgets/event_carusel/Carusel";

const eventArray = [
	{
		title: "Сбербанк Хакатон",
		started_at: "12.11.2025",
		ended_at: "15.11.2025",
		stack: "React, Java Spring",
		photo_alt: "Сбербанк хакатон",
		photo_src: "/dj.jpg",
	},

	{
		title: "ВТБ Хакатон",
		started_at: "12.11.2025",
		ended_at: "15.11.2025",
		stack: "React, GO",
		photo_alt: "ВТБ Хакатон",
		photo_src: "/how.jpg",
	},

	{
		title: "Т-Банк Хакатон",
		started_at: "12.11.2025",
		ended_at: "15.11.2025",
		stack: "React, GO",
		photo_alt: "Т-Банк Хакатон",
		photo_src: "/ged.jpg",
	},
];

const ActualEvents = () => {
	return (
		<div className={styles.actualEvents}>
			<Carusel events={eventArray} />
			<BaseButton variant="bordered">Посмотреть все</BaseButton>
		</div>
	);
};

export default ActualEvents;
