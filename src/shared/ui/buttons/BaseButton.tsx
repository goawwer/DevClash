import React, { FC, ButtonHTMLAttributes, ReactNode } from "react";
import styles from "./BaseButton.module.scss";
import cn from "classnames";

type Props = ButtonHTMLAttributes<HTMLButtonElement> & {
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
	...props
}) => {
	const buttonClass = cn(styles.baseButton, styles[`baseButton--${size}`], {
		[styles["baseButton--pending"]]: status === "pending",
		[styles["baseButton--success"]]: status === "success",
		[styles["baseButton--error"]]: status === "error",
		[styles["baseButton--bordered"]]: variant === "bordered",
	});

	return (
		<button className={buttonClass} {...props}>
			{children}
		</button>
	);
};

export default BaseButton;
