import { BaseButton, Card, Chooser } from "@/shared/ui";
import styles from "./InfoContainer.module.scss";
import Link from "next/link";

const InfoContainer = () => {
	return (
		<Card color="blue">
			<h2 className={styles.info__title}>Кто вы?</h2>
			<Chooser
				variants={[
					{ name: "Участник", value: <ParticipantInfo /> },
					{ name: "Организатор", value: <OrganizerInfo /> },
				]}
			/>
		</Card>
	);
};

const ParticipantInfo = () => {
	return (
		<div className={styles.info__content}>
			<p className={styles.info__text}>
				ВЫ ИТ-специалист? Здесь вы можете участвовать в ИТ-мероприятиях
				любого уровня, находить единомышленников и взаимодействовать с
				представителями ИТ-компаний{" "}
			</p>

			<Link href={"/register"}>
				<BaseButton>Присоединиться</BaseButton>
			</Link>
		</div>
	);
};

const OrganizerInfo = () => {
	return (
		<div className={styles.info__content}>
			<p className={styles.info__text}>
				ВЫ представитель компании или энтуазиаст? Здесь вы можете
				организовать ИТ-мероприятие любого уровня, найти ценных кадров и
				держать руку на пульсе ИТ-сферы
			</p>

			<Link href={"/register"}>
				<BaseButton>Присоединиться</BaseButton>
			</Link>
		</div>
	);
};

export default InfoContainer;
