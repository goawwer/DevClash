import { ibm, genos } from "@/shared/styles/fonts";
import styles from "./layout.module.scss";
import "./globals.css";
import HeaderUnauth from "@/widgets/header/HeaderUnauth";
import Footer from "@/widgets/footer/Footer";

export default function RootLayout({
	children,
}: Readonly<{
	children: React.ReactNode;
}>) {
	return (
		<html lang="ru" className={`${ibm.variable} ${genos.variable}`}>
			<body>
				<div className={styles.layout}>
					<div className={styles.layout__wrapper}>
						<HeaderUnauth />
						<main className={styles.layout__main}>{children}</main>
					</div>
					<Footer />
				</div>
			</body>
		</html>
	);
}
