"use client";

import styles from "./EditUserForm.module.scss";
import { useForm } from "react-hook-form";
import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { BaseButton, InputFile } from "@/shared/ui";
import { Input } from "@/shared/ui";
import type { UserProfileDto } from "@/entities/user.interface";
import {
	EmailInput,
	emailOptions,
	UsernameInput,
	usernameOptions,
} from "@/features/forms";
import Textarea from "@/shared/ui/inputs/Textarea";
import { update } from "@/features/api/user/update";

type FormValues = {
	username: string;
	email: string;
	bio: string;
	profile_status: string;
	tech_stack: string; // отображаем как строку: "Go, React"
	picture?: File;
};

interface EditUserFormProps {
	userData: UserProfileDto | undefined;
}

const EditUserForm = ({ userData }: EditUserFormProps) => {
	const {
		register,
		handleSubmit,
		reset,
		formState: { errors },
		setValue,
	} = useForm<FormValues>({
		defaultValues: {
			username: userData?.username || "",
			email: userData?.email || "",
			bio: userData?.bio || "",
			profile_status: userData?.profile_status || "",
			tech_stack: userData?.tech_stack?.join(", ") || "",
		},
	});

	const router = useRouter();
	const [submitStatus, setStatus] = useState<
		"idle" | "pending" | "error" | "success"
	>("idle");

	const [selectedFile, setSelectedFile] = useState<File | undefined>(
		undefined
	);

	useEffect(() => {
		if (userData) {
			reset({
				username: userData.username || "",
				email: userData.email || "",
				bio: userData.bio || "",
				profile_status: userData.profile_status || "",
				tech_stack: userData.tech_stack?.join(", ") || "",
			});
		}
	}, [userData, reset]);

	useEffect(() => {
		if (selectedFile) {
			setValue("picture", selectedFile, { shouldValidate: true });
		}
	}, [selectedFile, setValue]);

	const onSubmit = async (data: FormValues) => {
		// Преобразуем строку tech_stack в массив
		const techStackArray = data.tech_stack
			.split(",")
			.map((tag) => tag.trim())
			.filter((tag) => tag.length > 0);

		const preparedData: Partial<UserProfileDto> = {
			username: data.username,
			email: data.email,
			bio: data.bio,
			profile_status: data.profile_status,
			tech_stack: techStackArray.length > 0 ? techStackArray : undefined,
		};

		try {
			setStatus("pending");
			await update(preparedData, selectedFile);
			setStatus("success");
			router.refresh();
		} catch (error) {
			console.error("Ошибка обновления профиля:", error);
			setStatus("error");
		}
	};

	const handleFileSelect = (file: File) => {
		setSelectedFile(file);
	};

	return (
		<form
			className={styles.form}
			onSubmit={handleSubmit(onSubmit)}
			aria-label="Форма редактирования профиля"
			encType="multipart/form-data"
		>
			<div className={styles.form__inputs}>
				<div className={styles.form__picture}>
					<InputFile
						label="аватар"
						formats={["png", "jpg", "jpeg", "svg"]}
						name="picture"
						onFileSelect={handleFileSelect}
					/>

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
				</div>

				<Input
					label="Статус профиля"
					width={1000}
					error={errors.profile_status?.message}
					{...register("profile_status")}
				/>

				<Input
					label="Технологический стек (через запятую)"
					width={1000}
					placeholder="Например: Go, React, PostgreSQL"
					error={errors.tech_stack?.message}
					{...register("tech_stack")}
				/>

				<Textarea
					label="О себе (md)"
					width={1000}
					error={errors.bio?.message}
					{...register("bio")}
					placeholder="Расскажите о себе..."
				/>
			</div>

			<div className={styles.form__buttons}>
				<BaseButton type="submit" status={submitStatus}>
					Сохранить профиль
				</BaseButton>
			</div>
		</form>
	);
};

export default EditUserForm;
