export const usernameOptions = {
	required: "Введите имя пользователя",
	pattern: {
		value: /^[a-zA-Z0-9]{1,31}$/,
		message: "Только латинские буквы и цифры, до 31 символа",
	},
};
