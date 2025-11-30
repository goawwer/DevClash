import GradientContainer from "@/shared/templates/pageGradient/PageGradient";
import styles from "./page.module.scss";
import { FC } from "react";
import Image from "next/image";
import StackContainer from "./sub-components/Containers/StackContainer";
import DateContainer from "./sub-components/Containers/DateContainer";
import Tags from "./sub-components/Containers/Tags";
import Description from "./sub-components/Containers/Description";
import AuthCheck from "./sub-components/Auth/AuthCheck";

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
	description: `# Frontend Meetup: Современный React
Обсуждаем реальные кейсы и практики в разработке интерфейсов.

**Дата:** 2025-12-12  
**Время:** 18:30 — 21:30 (Europe/Moscow)  
**Место:** Москва, ул. Примерная, 10, LoftHub / онлайн (https://meet.example.com/frontend)

---

## О мероприятии
Встреча разработчиков, где мы разберём новые подходы в React, оптимизацию рендеринга и реальные migration-кейсы с больших проектов. Полезно для middle+ frontend-инженеров и тимлидов.

## Для кого
- Frontend-разработчики  
- Архитекторы SPA  
- Техлиды и менеджеры продуктов

## Программа
**18:30 — 19:00** — Регистрация и welcome-кофе  
**19:00 — 19:10** — Открытие — Антон Иванов (организатор)  
**19:10 — 19:40** — Доклад: *Оптимизация рендеринга в React* — Мария Петрова (XYZ)  
**19:40 — 20:10** — Доклад: *Миграция на TypeScript: сложности и решения* — Олег Смирнов (ACME)  
**20:10 — 20:25** — Перерыв  
**20:25 — 21:00** — Панель: *Архитектура крупных фронтенд-приложений* — спикеры + Q&A  
**21:00 — 21:30** — Нетворкинг, вопросы спикерам

## Спикеры
- **Мария Петрова** — Senior Frontend Engineer, XYZ — работает над производительностью интерфейсов более 8 лет.  
- **Олег Смирнов** — Tech Lead, ACME — ведущий инженер миграционных проектов.

## Регистрация
- Стоимость: бесплатно (ограничено 100 мест)  
- Регистрация: https://example.com/register

## Что взять с собой
- Ноутбук (если планируешь кодить в перерывах)  
- Визитки для нетворкинга

## Контакты
Организатор: Антон Иванов — anton@example.com — +7 900 123-45-67  
Сайт: https://example.com

## Как добраться
Метро «Примерная» — 5 минут пешком; паркинг рядом (платный).

## Партнёры и спонсоры
- DevTools — спонсор кофе-паузы  
- CloudCorp — технологический партнёр

## Часто задаваемые вопросы
**Q:** Нужен ли билет?  
**A:** Да, регистрация обязательна, но бесплатна.

---

_Дата публикации описания: 2025-11-30_

`,
	tags: [
		{ id: 1, title: "Хакатон" },
		{ id: 2, title: "Вживую" },
		{ id: 3, title: "Оплата за призерство" },
	],
	color: "green",
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

				<div className={styles.page__check}>
					<AuthCheck />
				</div>
			</div>

			<Description description={data.description} color={data.color} />

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
