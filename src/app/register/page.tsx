import GradientContainer from "@/shared/pageGradient/PageGradient";
import styles from "./page.module.scss";
import Register from "./sub-components/Register";

export default function Home() {
	return (
		<>
			<h1 className={styles.title}>Регистрация</h1>

			<section className={styles.sectionRegister}>
				<Register />
			</section>

			<GradientContainer color="blue" />
		</>
	);
}
