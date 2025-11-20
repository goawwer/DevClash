"use client";

import styles from "./Headers.module.scss";
import { BaseButton } from "@/shared/ui";
import { NavLink } from "@/shared/ui/";
import { Logo } from "@/shared/ui";
import { useState } from "react";
import { BiMenu, BiX } from "react-icons/bi";

export default function HeaderUnauth() {
	const [open, setOpen] = useState(false);
	return (
		<header className={styles.header}>
			<nav>
				<div className={styles.header__logo}>
					<Logo />
				</div>

				<div className={styles[`header__pageList--desktop`]}>
					<PageList />
				</div>

				<div className={styles.header__enteryButton}>
					<BaseButton size="low">войти</BaseButton>
				</div>

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
								<NavLink href="/auth/entery">войти</NavLink>
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
				<NavLink href="/" exact={true}>
					главная
				</NavLink>
			</ul>
			<ul>
				<NavLink href="/events">мероприятия</NavLink>
			</ul>
			<ul>
				<NavLink href="/about">о проекте</NavLink>
			</ul>
		</ol>
	);
};
