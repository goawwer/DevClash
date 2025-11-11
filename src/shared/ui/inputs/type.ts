import { InputHTMLAttributes } from "react";

export type InputType = InputHTMLAttributes<HTMLInputElement> & {
	label?: string;
	error?: string;
	width?: number;
};
