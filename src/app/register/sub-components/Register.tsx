import { Card, Chooser } from "@/shared/ui";
import { RegisterParticipantForm, RegisterOrganizerForm } from "@/widgets/forms";
import styles from "./Register.module.scss";

const registerVariants = [
	{
		name: "Участник",
		value: (
			<Card color="blue">
				<RegisterParticipantForm />
			</Card>
		),
	},
	{ name: "Организатор", value: (
			<Card color="blue">
				<RegisterOrganizerForm />
			</Card>
		) },
];

export default function Register() {
	return (
		<div className={styles.register}>
			<Chooser variants={registerVariants}></Chooser>
		</div>
	);
}
