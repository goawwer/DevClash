import styles from "./Headers.module.scss";
import { BaseButton } from "@/shared/ui";
import { NavLink } from "@/shared/ui/";
import { Logo } from "@/shared/ui";

export default function HeaderUnauth() {
	return (
		<header className={styles.header}>
			<nav>
				<Logo />

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

				<BaseButton size="low">войти</BaseButton>
			</nav>
		</header>
	);
}
