import { ReactNode } from "react";

export const dynamic = "force-dynamic"; // чтобы fetch не кешировался

export default async function EventLayout({
	children,
}: {
	children: ReactNode;
}) {
	return <>{children}</>;
}
