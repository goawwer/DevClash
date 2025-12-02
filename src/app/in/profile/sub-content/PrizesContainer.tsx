import { InfoCard } from "@/shared/templates";
import styles from "./InfoContainers.module.scss";

export default function PrizesContainer({ number }: { number: number }) {
	return (
		<InfoCard
			title={<h2 className={styles.title}>Призерство</h2>}
			content={
				<div className={styles.container}>
					<p className={styles.content}>{number}</p>
				</div>
			}
		/>
	);
}
