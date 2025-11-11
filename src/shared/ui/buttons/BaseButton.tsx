import React, { FC, ButtonHTMLAttributes } from "react";
import styles from "./BaseButton.module.scss";
import cn from "classnames";

type Props = ButtonHTMLAttributes<HTMLButtonElement> & {
	variant?: "filled" | "bordered";
	size?: "low" | "default";
};

const BaseButton: FC<Props> = ({
	children,
	size = "default",
	variant = "filled",
	...props
}) => {
	if (variant === "filled") {
		return (
			<button
				type="button"
				{...props}
				className={cn(styles.baseButton, styles[`baseButton--${size}`])}
			>
				{children}
			</button>
		);
	} else {
		return (
			<button
				type="button"
				{...props}
				className={cn(
					styles.baseButton,
					styles["baseButton--bordered"],
					styles[`baseButton--${size}`]
				)}
			>
				{children}
			</button>
		);
	}
};

export default BaseButton;
