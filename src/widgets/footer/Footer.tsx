import styles from "./Footer.module.scss";
import { Logo } from "@/shared/ui";

export default function Footer() {
	return (
		<footer className={styles.footer}>
			<p>
				Сайт <Logo fontSize={20} /> является собственностью творческой
				организации «Jonklers®»
			</p>
		</footer>
	);
}
