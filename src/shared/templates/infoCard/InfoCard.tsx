import { FC, ReactNode } from "react";
import { Card } from "@/shared/ui";

type Props = {
	title?: ReactNode;
	content?: ReactNode;
};

const InfoCard: FC<Props> = ({ title, content }) => {
	return (
		<Card color="dark-grey">
			{title}
			{content}
		</Card>
	);
};

export default InfoCard;
