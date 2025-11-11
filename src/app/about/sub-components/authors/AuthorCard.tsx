import { CardComb } from "@/shared/templates";
import { FC } from "react";

type Props = {
	first_name: string;
	last_name: string;
	role: string;
};

const AuthorCard: FC<Props> = ({ first_name, last_name, role }) => {
	return (
		<CardComb
			width={25}
			color="orange"
			title={
				<h3>
					{first_name}
					<br />
					{last_name}
				</h3>
			}
			subTitle={<p>АИС-22-1</p>}
			topInfo={<p>{role}</p>}
		></CardComb>
	);
};

export default AuthorCard;
