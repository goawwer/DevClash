import { MDContainer } from "@/shared/templates";

export default function Description({
	description,
	color,
}: {
	description: string;
	color: string;
}) {
	return <MDContainer content={description} color={color} />;
}
