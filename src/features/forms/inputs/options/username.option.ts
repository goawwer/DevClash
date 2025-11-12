export const usernameOptions = {
	required: "Введите имя пользователя",
	pattern: {
		value: /^[a-zA-Z0-9]{3,31}$/,
		message: "Только латинские буквы и цифры, от 3 до 31 символа",
	},
};
