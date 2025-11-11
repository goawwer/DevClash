import { FC, ReactNode } from "react";

type Props = {
	color: string;
	children: ReactNode;
};

const Card: FC<Props> = ({ color = `blue`, children }) => {
	return (
		<div
			style={{
				width: "fit-content",
				height: "fit-content",
				backgroundColor: `var(--${color}-color)`,
				padding: "24px",
				borderRadius: "24px",
			}}
		>
			{children}
		</div>
	);
};

export default Card;
