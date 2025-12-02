import { ibm, genos } from "@/shared/styles/fonts";
import styles from "./layout.module.scss";
import "./globals.css";
import Header from "@/widgets/header/Header";
import Footer from "@/widgets/footer/Footer";
import { devClashMetadata } from "./metadata";
import { AuthProvider } from "@/app/context/AuthContext";

export const metadata = {
	...devClashMetadata,
};

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
						<AuthProvider>
							<Header />
						</AuthProvider>
						<main className={styles.layout__main}>{children}</main>
					</div>
					<Footer />
				</div>
			</body>
		</html>
	);
}
