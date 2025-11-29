import { Card } from "@/shared/ui";
import { LoginForm } from "@/widgets/forms";
import styles from "./Login.module.scss";

export default function Login() {
	return (
		<div className={styles.register}>
			<Card color="blue">
				<LoginForm />
			</Card>
		</div>
	);
}
