import { InfoCard } from "@/shared/templates";
import styles from "./InfoContainers.module.scss";

type Props = {
	started_at: string;
	ended_at: string;
};

export default function DataContainer({ started_at, ended_at }: Props) {
	return (
		<InfoCard
			title={<h2 className={styles.title}>Дата</h2>}
			content={
				<div className={styles.stack}>
					<p className={styles.content}>
						{started_at} - {ended_at}
					</p>
				</div>
			}
		/>
	);
}
