"use client";

import styles from "./RegiserForm.module.scss";
import { useForm } from "react-hook-form";
import { User } from "@/entities/user.interface";
import {
	EmailInput,
	emailOptions,
	PasswordInput,
	passwordOptions,
	UsernameInput,
	usernameOptions,
} from "@/features/forms";
import { BaseButton } from "@/shared/ui";
import { useState } from "react";
import Link from "next/link";

type FormValues = User & {
	confirmPassword: string;
};

const RegisterForm = () => {
	const {
		register,
		handleSubmit,
		formState: { errors },
		watch,
	} = useForm<FormValues>();

	const [a, setA] = useState(0);

	const onSubmit = (data: FormValues) => {
		setA(a + 1);
		return console.log(data);
	};

	const password = watch("password");

	return (
		<form className={styles.form} onSubmit={handleSubmit(onSubmit)} aria-label="Форма регистрации">
			<div className={styles.form__inputs}>
				<UsernameInput
					width={27}
					error={errors.username?.message}
					register={register("username", usernameOptions)}
				/>

				<EmailInput
					width={27}
					error={errors.email?.message}
					register={register("email", emailOptions)}
				/>

				<PasswordInput
					width={27}
					error={errors.password?.message}
					register={register("password", passwordOptions)}
				/>

				<PasswordInput
					width={27}
					label="подтвердите пароль"
					error={errors.confirmPassword?.message}
					register={register("confirmPassword", {
						validate: (value) =>
							value === password || "Пароли не совпадают",
					})}
				/>
			</div>
			<div className={styles.form__buttons}>
				<Link href={"/login"} className={styles.form__loginLink}>Уже есть аккаунт?</Link>
				<BaseButton type="submit">Создать аккаунт</BaseButton>
			</div>
			<h2>{a}</h2>
		</form>
	);
};

export default RegisterForm;
