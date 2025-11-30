import styles from "./MDContainer.module.scss";
import { FC, ReactNode } from "react";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";

type Props = {
	color?: string;
	title?: string;
	editButton?: ReactNode;
	content: string;
};

const MDContainer: FC<Props> = ({
	title = "Подробная информация",
	content,
	editButton,
	color = "white",
}) => {
	return (
		<div className={styles.mdContainer}>
			<div className={styles.mdContainer__header}>
				<p className={styles.mdContainer__title}>
					{title}
					<span style={{ color: `var(--${color}-color)` }}>.md</span>
				</p>

				{editButton && (
					<button style={{ color: `var(--${color}-color)` }}>
						Редактировать?
					</button>
				)}
			</div>
			<div className={styles.mdContainer__content}>
				<ReactMarkdown remarkPlugins={[remarkGfm]}>
					{content}
				</ReactMarkdown>
			</div>
		</div>
	);
};

export default MDContainer;
