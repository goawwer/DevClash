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
				maxWidth: "100%",
				height: "fit-content",
				backgroundColor: `var(--${color}-color)`,
				padding: "1.25rem",
				borderRadius: "1.25rem",
				
			}}
		>
			{children}
		</div>
	);
};

export default Card;
