"use client";

import GradientContainer from "@/shared/templates/pageGradient/PageGradient";
import styles from "./page.module.scss";
import { BaseButton } from "@/shared/ui";
import Image from "next/image";
import StackContainer from "./sub-content/StackContainer";
import ParticipantContainer from "./sub-content/ParticipantContainer";
import PrizesContainer from "./sub-content/PrizesContainer";
import Description from "./sub-content/Description";
import Link from "next/link";
import { profileUser } from "@/features/api";
import { useEffect, useState } from "react";
import { UserProfileDto } from "@/entities/user.interface";
import { logout } from "@/features/api/auth/logout";
import { redirect } from "next/navigation";

export default function Profile() {
	const [userData, setUser] = useState<UserProfileDto | undefined>(undefined);

	useEffect(() => {
		async function getProfile() {
			const data = await profileUser();
			console.log("PROFILE:", data);
			setUser(data);
		}

		getProfile();
	}, []);

	const logoutHandler = async () => {
		await logout();
		redirect("/");
	};

	return (
		<>
			<div className={styles.page__topWidget}>
				<div className={styles.page__photoContainer}>
					<Image
						src={userData?.profile_picture_url || "/pp.jpg"}
						alt="profile-pic"
						fill
					/>
				</div>

				<div className={styles.page__titleContainer}>
					<h1 className={styles.page__title}>
						@{userData?.username}
					</h1>

					<p className={styles.page__subText}>
						{userData?.profile_status || "нет статуса"}
					</p>

					<div className={styles.page__profileButton}>
						<Link href={"/in/settings"}>
							<BaseButton variant="bordered" noFocus>
								Редактировать
							</BaseButton>
						</Link>

						<BaseButton variant="bordered" onClick={logoutHandler}>
							{" "}
							←|
						</BaseButton>
					</div>
				</div>
			</div>

			<section className={styles.page__containersSection}>
				<StackContainer stack={userData?.tech_stack || ["не указан"]} />

				<ParticipantContainer
					number={userData?.participations_count || 0}
				/>
				<PrizesContainer number={userData?.wins_count || 0} />
			</section>

			<section className={styles.page__section}>
				<Description
					color="blue"
					description={
						userData?.bio ||
						"Пользователь не указал никакой информации о себе"
					}
				/>
			</section>
			<GradientContainer color="blue" />
		</>
	);
}
