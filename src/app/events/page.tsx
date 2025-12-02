import GradientContainer from "@/shared/templates/pageGradient/PageGradient";
import styles from "./page.module.scss";
import ActualEvents from "./sub-components/ActualEvents";
import AllEvents from "./sub-components/AllEvents";

export default async function Home() {
	return (
		<>
			<h1 className={styles.title}>Актуальные мероприятия</h1>

			<section className={styles.sectionDescription}>
				<ActualEvents />
			</section>

			<AllEvents />

			<GradientContainer color="silver" />
		</>
	);
}
