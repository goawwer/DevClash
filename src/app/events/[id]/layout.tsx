import { ReactNode } from "react";
import AuthCheck from "./sub-components/Auth/AuthCheck";

export const dynamic = "force-dynamic"; // чтобы fetch не кешировался

export default async function EventLayout({
	children,
}: {
	children: ReactNode;
}) {
	return <>{children}</>;
}
