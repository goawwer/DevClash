import { FC, ReactNode } from "react";
import styles from "./CardComb.module.scss";

type Props = {
	width: number;
	color: string;
	title?: ReactNode;
	topInfo?: ReactNode;
	subTitle?: ReactNode;
	content?: ReactNode;
};

const CardComb: FC<Props> = ({
	width,
	color,
	title,
	topInfo,
	subTitle,
	content,
}) => {
	return (
		<div className={styles.combCard}>
			<div className={styles.combCard__topInfo}>{topInfo && topInfo}</div>

			<div
				className={styles.combCard__main}
				style={{ width: `${width}rem` }}
			>
				{title && title}

				<div
					className={styles.combCard__subTitle}
					style={{ color: `var(--${color}-color)` }}
				>
					{subTitle && subTitle}
				</div>

				<div className={styles.combCard__content}>
					{content && content}
				</div>
			</div>
		</div>
	);
};

export default CardComb;
