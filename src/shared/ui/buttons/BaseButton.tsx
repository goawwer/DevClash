import { FC, ButtonHTMLAttributes, ReactNode } from "react";
import styles from "./BaseButton.module.scss";
import cn from "classnames";

type Props = ButtonHTMLAttributes<HTMLButtonElement> & {
	noFocus?: boolean;
	status?: "idle" | "pending" | "success" | "error";
	variant?: "filled" | "bordered";
	size?: "low" | "default";
	children: ReactNode;
};

const BaseButton: FC<Props> = ({
	children,
	size = "default",
	variant = "filled",
	status,
	noFocus = false,
	tabIndex,
	...props
}) => {
	const buttonClass = cn(styles.baseButton, styles[`baseButton--${size}`], {
		[styles["baseButton--pending"]]: status === "pending",
		[styles["baseButton--success"]]: status === "success",
		[styles["baseButton--error"]]: status === "error",
		[styles["baseButton--bordered"]]: variant === "bordered",
		[styles["baseButton--noFocus"]]: noFocus === true,
	});

	return (
		<button
			className={buttonClass}
			tabIndex={noFocus ? -1 : tabIndex}
			{...props}
		>
			{children}
		</button>
	);
};

export default BaseButton;
