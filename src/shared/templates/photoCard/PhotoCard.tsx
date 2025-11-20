import Image from "next/image";
import styles from "./PhotoCard.module.scss";
import { FC } from "react";

type Props = {
	width?: number;
	title: string;
	started_at: string;
	ended_at: string;
	photo_src: string;
	photo_alt: string;
	stack: string;
};

const PhotoCard: FC<Props> = ({
	width = 52,
	title,
	photo_alt,
	photo_src,
	started_at,
	ended_at,
	stack,
}) => {
	return (
		<div className={styles.photoCard} style={{ width: `${width}rem` }}>
			<div className={styles.photoCard__top}>
				<p className={styles.photoCard__title}>{title}</p>
				<p className={styles.photoCard__date}>
					{started_at} - {ended_at}
				</p>
			</div>

			<div className={styles.photoCard__bottom}>
				<p className={styles.photoCard__stack}>Стэк: {stack}</p>
			</div>
			<Image src={photo_src} alt={photo_alt} fill />
		</div>
	);
};

export default PhotoCard;
