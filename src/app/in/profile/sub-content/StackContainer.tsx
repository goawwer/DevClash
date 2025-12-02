import { InfoCard } from "@/shared/templates";
import styles from "./InfoContainers.module.scss";

type StackContainerProps = {
	stack: string[];
};

export default function StackContainer({ stack }: StackContainerProps) {
	return (
		<InfoCard
			title={<h2 className={styles.title}>Стэк</h2>}
			content={
				<div className={styles.stack}>
					{stack.map((skill, id) => (
						<p className={styles.content} key={id}>
							{skill}
						</p>
					))}
				</div>
			}
		/>
	);
}
