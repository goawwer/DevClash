import { IBM_Plex_Mono } from "next/font/google";
import { Genos } from "next/font/google";

export const ibm = IBM_Plex_Mono({
	subsets: ["latin"],
	weight: ["400", "500", "700"],
	display: "swap",
	variable: "--font-ibm",
});

export const genos = Genos({
	subsets: ["latin"],
	weight: ["400", "700"],
	display: "swap",
	variable: "--font-genos",
});
