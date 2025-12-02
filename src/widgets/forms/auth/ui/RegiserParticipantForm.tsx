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
import { BaseButton } from "@/shared/ui";
import Link from "next/link";
import { login, userSignUp } from "@/features/api";
import { useState } from "react";
import { useRouter } from "next/navigation";

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
	// внутри компонента:
	const router = useRouter();

	const [submitStatus, setStatus] = useState<
		"idle" | "pending" | "error" | "success"
	>("idle");

	const onSubmit = async (data: FormValues) => {
		const user = {
			email: data.email,
			name: data.username,
			password: data.password,
			username: data.username,
		};

		try {
			setStatus("pending");
			await userSignUp(user);
			await login({ email: user.email, password: user.password });
			setStatus("success");
			router.push("/in/profile");
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

export default RegisterForm;
