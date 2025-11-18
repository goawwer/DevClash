"use client";
import { FC, useState } from "react";
import styles from "./ColorInput.module.scss";
import cn from "classnames";

type MainProps = {
	setFormColor: (color: string) => void;
};

type ButtonProps = {
	setFormColor: (color: string) => void;
	ownColor: string;
	currentColor: string;
};

const colorArray = [
	"blue",
	"green",
	"yellow",
	"orange",
	"red",
	"purple",
	"white",
];

const ColorInput: FC<MainProps> = ({ setFormColor }) => {
	const [currentColor, setColor] = useState<string>("white");

	const setNewColor = (color: string) => {
		setColor(color);
		setFormColor(color);
	};

	return (
		<div className={styles.colorInput}>
			<p className={styles.colorInput__label}>Брендинг</p>
			<div className={styles.colorInput__body}>
				<p className={styles.colorInput__title}>Выберите цвет</p>
				<div className={styles.colorInput__container}>
					{colorArray.map((color, id) => (
						<ColorButton
							key={id}
							ownColor={color}
							currentColor={currentColor}
							setFormColor={setNewColor}
						/>
					))}
				</div>
			</div>
		</div>
	);
};

const ColorButton: FC<ButtonProps> = ({
	ownColor,
	currentColor,
	setFormColor,
}) => {
	const chooseColor = () => {
		setFormColor(ownColor);
	};
	return (
		<button
			type="button"
			className={cn(styles.button, {
				[styles["button--current"]]: ownColor === currentColor,
			})}
			onClick={chooseColor}
			style={{ border: `0.2rem solid var(--${ownColor}-color)` }}
		></button>
	);
};

export default ColorInput;
