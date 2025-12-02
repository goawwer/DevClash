"use client";

import GradientContainer from "@/shared/templates/pageGradient/PageGradient";
import styles from "./page.module.scss";
import { BaseButton } from "@/shared/ui";
import Image from "next/image";
import FileUploadInput from "@/shared/ui/inputs/InputFile";
import { EditUserForm } from "@/widgets/forms";
import { useEffect, useState } from "react";
import { profileUser, settingsUser } from "@/features/api";
import { UserProfileDto } from "@/entities/user.interface";

const data = {
	username: "Профиль",

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
	status: "Скиньте денег сбер +7-999-999-99-99",
	color: "blue",
	participant: 58,
	prizes: 12,
	stack: [
		{ id: 1, title: "React" },
		{ id: 2, title: "Java" },
		{ id: 3, title: "Spring Boot" },
	],
};

export default function Settings() {
	const [userData, setUser] = useState<UserProfileDto | undefined>();
	useEffect(() => {
		async function getProfile() {
			const data = await settingsUser();
			console.log("settings:", data);
			setUser(data);
		}

		getProfile();
	}, []);

	return (
		<>
			<div className={styles.page__titleContainer}>
				<h1 className={styles.page__title}>Настройки</h1>

				<h2 className={styles.page__subTitle}>Доступная информация</h2>
				<div className={styles.page__section}>
					<EditUserForm userData={userData} />
				</div>
			</div>

			<section className={styles.page__containersSection}></section>

			<section className={styles.page__section}></section>
			<GradientContainer color="blue" />
		</>
	);
}
