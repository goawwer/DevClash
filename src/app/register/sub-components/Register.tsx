import { Card, Chooser } from "@/shared/ui";
import { RegisterForm } from "@/widgets/forms";
import styles from "./Register.module.scss";

const registerVariants = [
	{
		name: "Участник",
		value: (
			<Card color="blue">
				<RegisterForm />
			</Card>
		),
	},
	{ name: "Организатор", value: <></> },
];

export default function Register() {
	return (
		<div className={styles.register}>
			<Chooser variants={registerVariants}></Chooser>
		</div>
	);
}
