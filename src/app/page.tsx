import GradientContainer from "@/shared/pageGradient/PageGradient";
import styles from "./page.module.scss";
import { AnimatedText, InfoContainer } from "./sub-components/";

export default function Home() {
	return (
		<>
			<h1 className={styles.title}>
				станьте частью <br /> пЕРЕДОВОго ИТ-сообщества
			</h1>

			<div>
				<p className={styles.subText}>
					Здесь организуются ИТ-мероприятия любого уровня
				</p>
			</div>

			<section className={styles.section}>
				<InfoContainer />
			</section>

			<GradientContainer color="blue" />
		</>
	);
}
