import GradientContainer from "@/shared/pageGradient/PageGradient";
import styles from "./page.module.scss";
import { AnimatedText, InfoContainer } from "./sub-components/";

export default function Home() {
	return (
		<>
			<h1 className={styles.title}>
				Присоединяйтесь <br /> к ПЕРЕДОВОМУ ИТ-сообществу
			</h1>

			<div>
				<AnimatedText />
			</div>

			<section className={styles.section}>
				<InfoContainer />
			</section>

			<GradientContainer color="blue" />
		</>
	);
}
