import GradientContainer from "@/shared/templates/pageGradient/PageGradient";
import styles from "./page.module.scss";
import { FC } from "react";
import Image from "next/image";
import StackContainer from "./sub-components/StackContainer";
import DateContainer from "./sub-components/DateContainer";
import Tags from "./sub-components/Tags";

interface Props {
	params: { slug: string };
}

const data = {
	title: "Хакатон классный крутой и т.д.",
	brand: "Сбербанк",
	logo_url: "/dc-logo.png",
	baner_url: "/sber.png",
	stack: [
		{ id: 1, title: "React" },
		{ id: 2, title: "Java" },
		{ id: 3, title: "Spring Boot" },
	],
	started_at: "12.12.2025",
	ended_at: "15.12.2025",
	description: "Ало, да-да, описание",
	tags: [
		{ id: 1, title: "Хакатон" },
		{ id: 2, title: "Вживую" },
		{ id: 3, title: "Оплата за призерство" },
	],
};

const EventDetails: FC<Props> = ({ params }) => {
	const { slug } = params;

	return (
		<>
			<div className={styles.page__titleContainer}>
				<Image
					className={styles.page__logo}
					src={data.logo_url}
					alt={data.title + "-logo"}
					width={160}
					height={160}
				/>

				<h1 className={styles.page__title}>
					{data.brand}.
					<br /> {data.title}
				</h1>

				<div className={styles.page__tags}>
					<Tags tags={data.tags} />
				</div>

				<div className={styles.page__infoContainers}>
					<StackContainer stack={data.stack} />
					<DateContainer
						started_at={data.started_at}
						ended_at={data.ended_at}
					/>
				</div>
			</div>
			<section className={styles.page__sectionDescription}></section>

			<GradientContainer color="green" />
			<Image
				src={data.baner_url}
				alt={data.brand + "-baner"}
				className={styles.page__banner}
				width={1920}
				height={1170}
			/>
		</>
	);
};

export default EventDetails;
