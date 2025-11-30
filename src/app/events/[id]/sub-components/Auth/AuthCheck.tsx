"use client";
import { authCheck } from "@/features/api";
import styles from "./AuthCheck.module.scss";
import { useEffect, useState } from "react";
import { BaseButton } from "@/shared/ui";
import cn from "classnames";

export default function AuthCheck() {
	const [isAuthed, setAuthed] = useState<boolean | undefined>(undefined);
	useEffect(() => {
		async function check() {
			const response = await authCheck();
			setAuthed(response);
		}

		check();
	}, []);

	if (isAuthed === undefined) {
		return;
	}

	return (
		<>
			{isAuthed ? (
				<div className={styles.isAuthed}>
					<BaseButton>Участвовать</BaseButton>
				</div>
			) : (
				<div className={cn(styles.isAuthed, styles.unauthed)}>
					Для участия или связи с организаторами необходимо
					зарегистрироваться
				</div>
			)}
		</>
	);
}
