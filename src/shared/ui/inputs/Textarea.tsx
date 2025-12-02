// shared/ui/inputs/Textarea.tsx
import { forwardRef, useCallback, useRef } from "react";
import styles from "./Input.module.scss";

type TextareaProps = {
	label?: string;
	error?: string;
	width?: number;
	placeholder?: string;
	className?: string;
} & React.TextareaHTMLAttributes<HTMLTextAreaElement>;

const Textarea = forwardRef<HTMLTextAreaElement, TextareaProps>(
	({ label, error, width = 24, className, ...props }, ref) => {
		const textareaRef = useRef<HTMLTextAreaElement>(null);

		const adjustHeight = useCallback(() => {
			const el = textareaRef.current;
			if (el) {
				el.style.height = "auto";
				el.style.height = `${el.scrollHeight}px`;
			}
		}, []);

		return (
			<div className={styles.input} style={{ width: `${width}rem` }}>
				{label && (
					<label className={styles.input__label} htmlFor={props.id}>
						{label}
					</label>
				)}
				<textarea
					{...props}
					ref={(node) => {
						if (typeof ref === "function") ref(node);
						else if (ref) ref.current = node;
						textareaRef.current = node;
						// Автоподстройка при монтировании (например, после reset)
						if (node) {
							// маленькая задержка, чтобы значение успело примениться
							setTimeout(() => adjustHeight(), 0);
						}
					}}
					onInput={adjustHeight}
					className={`${styles.input__input} ${
						styles.input__input_textarea
					} ${className || ""}`}
					rows={1}
					style={{
						resize: "none",
						minHeight: "4rem",
						height: "auto",
						width: `${width}rem`,
						...props.style,
					}}
				/>
				{error && (
					<span className={styles.input__error}>* {error}</span>
				)}
			</div>
		);
	}
);

Textarea.displayName = "Textarea";

export default Textarea;
