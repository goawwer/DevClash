import GradientContainer from "@/shared/templates/pageGradient/PageGradient";
import styles from "./page.module.scss";
import Login from "./sub-components/Login";

export default function Home() {
	return (
		<>
			<h1 className={styles.title}>Вход</h1>

			<section className={styles.sectionRegister}>
				<Login />
			</section>

			<GradientContainer color="blue" />
		</>
	);
}
