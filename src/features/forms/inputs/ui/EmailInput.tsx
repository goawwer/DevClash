import { Input, InputType } from "@/shared/ui";
import { forwardRef } from "react";
import type { UseFormRegisterReturn } from "react-hook-form";

interface EmailInputProps extends Omit<InputType, "type"> {
	register?: UseFormRegisterReturn;
}

const EmailInput = forwardRef<HTMLInputElement, EmailInputProps>(
	({ register, ...props }, ref) => {
		return (
			<Input
				label="электронная почта"
				type="email"
				placeholder="Ваш email"
				ref={ref}
				{...register}
				{...props}
				aria-label="Email"
			/>
		);
	}
);

EmailInput.displayName = "Email_Input";

export default EmailInput;
