"use client";

import styles from "./Auth.module.scss";
import { useForm } from "react-hook-form";
import { User } from "@/entities/user.interface";
import { EmailInput, PasswordInput } from "@/features/forms";
import { BaseButton } from "@/shared/ui";
import Link from "next/link";
import { login } from "@/features/api";
import { useState } from "react";
import { useRouter } from "next/navigation";

type FormValues = User & {
	confirmPassword: string;
};

const LoginForm = () => {
	const { register, handleSubmit } = useForm<FormValues>();

	const [submitStatus, setStatus] = useState<
		"idle" | "pending" | "error" | "success"
	>("idle");

	const router = useRouter();

	const onSubmit = async (data: FormValues) => {
		const user = {
			email: data.email,
			password: data.password,
		};

		try {
			setStatus("pending");
			await login(user);
			setStatus("success");
			router.push("/in/profile");
		} catch (error) {
			setStatus("error");
		}
	};

	return (
		<form
			className={styles.form}
			onSubmit={handleSubmit(onSubmit)}
			aria-label="Форма входа"
		>
			<div className={styles.form__inputs}>
				<EmailInput width={25} register={register("email")} />

				<PasswordInput width={25} register={register("password")} />
			</div>

			<div className={styles.form__buttons}>
				<Link href={"/register"} className={styles.form__loginLink}>
					ЕЩЕ НЕ ИМЕЕТЕ АККАУНТ?
				</Link>
				<BaseButton type="submit" status={submitStatus}>
					Войти в аккаунт
				</BaseButton>
			</div>
		</form>
	);
};

export default LoginForm;
