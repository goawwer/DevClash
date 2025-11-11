import { forwardRef } from "react";
import styles from "./Input.module.scss";
import type { InputType } from "./type";

const Input = forwardRef<HTMLInputElement, InputType>(
	({ label, error, width = 25, ...props }, ref) => {
		return (
			<div className={styles.input}>
				{label && (
					<label className={styles.input__label} htmlFor={props.id}>
						{label}
					</label>
				)}
				<input
					{...props}
					className={styles.input__input}
					ref={ref}
					style={{ width: `${width}rem` }}
				/>
				{error && <span className={styles.input__error}>{error}</span>}
			</div>
		);
	}
);

Input.displayName = "Custom_Input";

export default Input;
