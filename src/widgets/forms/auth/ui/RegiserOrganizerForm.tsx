"use client";

import styles from "./Auth.module.scss";
import { useForm } from "react-hook-form";
import { User } from "@/entities/user.interface";
import {
	ColorInput,
	EmailInput,
	emailOptions,
	PasswordInput,
	passwordOptions,
	UsernameInput,
	usernameOptions,
} from "@/features/forms";
import { BaseButton, InputFile } from "@/shared/ui";
import Link from "next/link";
import { login, organizerSignUp } from "@/features/api";
import { Organizer } from "@/entities/organizer.interface";
import { useState } from "react";

type FormValues = User & {
	color: string;
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

	const [submitStatus, setStatus] = useState<
		"idle" | "pending" | "error" | "success"
	>("idle");

	const onSubmit = async (data: FormValues) => {
		const organizer: Organizer = {
			name: data.username,
			email: data.email,
			color: data.color,
			logo: data.logo,
			password: data.password,
		};

		try {
			setStatus("pending");
			await organizerSignUp(organizer);
			await login({
				email: organizer.email,
				password: organizer.password,
			});
			setStatus("success");
		} catch (error) {
			setStatus("error");
		}
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
					width={25}
					error={errors.username?.message}
					register={register("username", usernameOptions)}
				/>

				<EmailInput
					width={25}
					error={errors.email?.message}
					register={register("email", emailOptions)}
				/>

				<PasswordInput
					width={25}
					error={errors.password?.message}
					register={register("password", passwordOptions)}
				/>

				<PasswordInput
					width={25}
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

				<ColorInput
					setFormColor={(color: string) => setValue("color", color)}
				/>
			</div>
			<div className={styles.form__buttons}>
				<Link href={"/login"} className={styles.form__loginLink}>
					Уже есть аккаунт?
				</Link>
				<BaseButton type="submit" status={submitStatus}>
					Создать аккаунт
				</BaseButton>
			</div>
		</form>
	);
};

export default RegisterOrganizerForm;
