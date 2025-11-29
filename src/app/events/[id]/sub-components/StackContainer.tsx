import { InfoCard } from "@/shared/templates";
import styles from "./InfoContainers.module.scss";

type Skill = {
	id: number;
	title: string;
};

type StackContainerProps = {
	stack: Skill[];
};

export default function StackContainer({ stack }: StackContainerProps) {
	return (
		<InfoCard
			title={<h2 className={styles.title}>Стэк</h2>}
			content={
				<div className={styles.stack}>
					{stack.map((skill) => (
						<p className={styles.content} key={skill.id}>
							{skill.title}
						</p>
					))}
				</div>
			}
		/>
	);
}
