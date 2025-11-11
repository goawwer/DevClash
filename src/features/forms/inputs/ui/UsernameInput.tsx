import { Input, InputType } from "@/shared/ui";
import { forwardRef } from "react";
import type { UseFormRegisterReturn } from "react-hook-form";

interface UsernameInputProps extends Omit<InputType, "type"> {
	register?: UseFormRegisterReturn;
}

const UsernameInput = forwardRef<HTMLInputElement, UsernameInputProps>(
	({ register, ...props }, ref) => {
		return (
			<Input
				label="имя пользователя"
				type="text"
				placeholder="Ваше имя пользователя"
				ref={ref}
				{...register}
				{...props}
			/>
		);
	}
);

UsernameInput.displayName = "Username_Input";

export default UsernameInput;
