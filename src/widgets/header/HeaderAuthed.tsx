"use client";

import styles from "./Headers.module.scss";
import { BaseButton } from "@/shared/ui";
import { NavLink } from "@/shared/ui/";
import { Logo } from "@/shared/ui";
import { useState } from "react";
import { BiMenu, BiX } from "react-icons/bi";
import Link from "next/link";

export default function HeaderAuthed() {
	const [open, setOpen] = useState(false);
	return (
		<header className={styles.header}>
			<nav>
				<div className={styles.header__logo}>
					<Link href="/in/feed">
						<Logo fontSize={2} />
					</Link>
				</div>

				<div className={styles[`header__pageList--desktop`]}>
					<PageList />
				</div>

				<NavLink href={"/in/settings"}>
					<div className={styles.header__enteryButton}>
						<BaseButton variant="bordered" size="low" noFocus>
							⚙
						</BaseButton>
					</div>
				</NavLink>

				<button
					onClick={() => setOpen(!open)}
					aria-label="Toggle menu"
					className={styles.header__burgerButton}
				>
					{open ? (
						<BiX
							size="2rem"
							className={styles[`header__burgerButtonSVG`]}
						/>
					) : (
						<BiMenu
							size="2rem"
							className={styles[`header__burgerButtonSVG`]}
						/>
					)}
				</button>

				{open && (
					<>
						<div className={styles.header__mobilePagesList}>
							<div className={styles[`header__pageList--mobile`]}>
								<PageList />
								<NavLink href="/in/settings">настройки</NavLink>
							</div>
						</div>
						<button
							type="button"
							className={styles.header__mobileDarkBG}
							onClick={() => setOpen(!open)}
						></button>
					</>
				)}
			</nav>
		</header>
	);
}

const PageList = () => {
	return (
		<ol>
			<ul>
				<NavLink href="/in/feed" exact={true}>
					лента
				</NavLink>
			</ul>
			<ul>
				<NavLink href="/events">мероприятия</NavLink>
			</ul>
			<ul>
				<NavLink href="/in/profile">профиль</NavLink>
			</ul>
		</ol>
	);
};
