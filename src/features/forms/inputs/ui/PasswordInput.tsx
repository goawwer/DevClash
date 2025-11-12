import { Input, InputType } from "@/shared/ui";
import { forwardRef } from "react";
import type { UseFormRegisterReturn } from "react-hook-form";

interface PasswordInputProps extends Omit<InputType, "type"> {
	register?: UseFormRegisterReturn;
	placeholder?: string;
}

const PasswordInput = forwardRef<HTMLInputElement, PasswordInputProps>(
	(
		{ register, label = "пароль", placeholder = "Ваш пароль", ...props },
		ref
	) => {
		return (
			<Input
				label={label}
				type="password"
				placeholder={placeholder}
				ref={ref}
				{...register}
				{...props}
				aria-label={label}
			/>
		);
	}
);

PasswordInput.displayName = "Password_Input";

export default PasswordInput;
