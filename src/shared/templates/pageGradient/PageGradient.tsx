import { FC } from "react";
import styles from "./PageGradient.module.scss";

type GradientContainerProps = {
	color?: string;
};

const GradientContainer: FC<GradientContainerProps> = ({
	color = "silver",
}) => {
	return (
		<div
			className={styles.fadeIn}
			style={{
				position: "absolute",
				zIndex: -1000,
				width: "100%",
				height: "fit-content",
				top: 0,
				left: 0,
			}}
		>
			<div
				style={{
					height: "100vh",
					background: `linear-gradient(to bottom, var(--black-color), var(--dark-${color}-color))`,
				}}
			></div>
			<div
				style={{
					height: "100vh",
					background: `linear-gradient(to bottom, var(--dark-${color}-color), var(--black-color))`,
				}}
			></div>
			<div
				style={{
					height: "100vh",
					background: `linear-gradient(to bottom, var(--black-color), var(--dark-${color}-color))`,
				}}
			></div>
			<div
				style={{
					height: "100vh",
					background: `linear-gradient(to bottom, var(--dark-${color}-color), var(--black-color))`,
				}}
			></div>
		</div>
	);
};

export default GradientContainer;
