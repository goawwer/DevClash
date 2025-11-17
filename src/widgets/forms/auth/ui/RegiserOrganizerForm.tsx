"use client";

import styles from "./Auth.module.scss";
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
import { BaseButton, InputFile } from "@/shared/ui";
import { useState } from "react";
import Link from "next/link";

type FormValues = User & {
	logo: File | undefined;
	confirmPassword: string;
};

const RegisterOrganizerForm = () => {
	const {
		register,
		handleSubmit,
		formState: { errors },
		watch,
		setValue,
	} = useForm<FormValues>();

	const [a, setA] = useState(0);

	const onSubmit = (data: FormValues) => {
		setA(a + 1);
		return console.log(data);
	};

	const password = watch("password");

	return (
		<form
			className={styles.form}
			onSubmit={handleSubmit(onSubmit)}
			aria-label="Форма регистрации"
		>
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

				<InputFile
					label="Логотип"
					formats={["png", "jpg", "jpeg", "svg"]}
					name="logo"
					onFileSelect={(file: File) => setValue("logo", file)}
				/>
			</div>
			<div className={styles.form__buttons}>
				<Link href={"/login"} className={styles.form__loginLink}>
					Уже есть аккаунт?
				</Link>
				<BaseButton type="submit">Создать аккаунт</BaseButton>
			</div>
			<h2>Количество: {a}</h2>
		</form>
	);
};

export default RegisterOrganizerForm;
