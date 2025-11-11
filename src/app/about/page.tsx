import GradientContainer from "@/shared/pageGradient/PageGradient";
import styles from "./page.module.scss";
import { Logo } from "@/shared/ui";
import { Description } from "./sub-components";
import Authors from "./sub-components/authors/Authors";

export default function Home() {
	return (
		<>
			<h1 className={styles.title}>
				Кратко про <br /> <Logo fontSize={64} />
			</h1>

			<section className={styles.sectionDescription}>
				<Description />
			</section>

			<section className={styles.sectionAuthors}>
				<Authors />
			</section>

			<GradientContainer color="orange" />
		</>
	);
}
