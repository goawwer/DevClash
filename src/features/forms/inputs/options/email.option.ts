export const emailOptions = {
	required: "Введите email",
	pattern: {
		value: /^[^\s@]+@[^\s@]+\.[^\s@]+$/,
		message: "Некорректный email",
	},
};
