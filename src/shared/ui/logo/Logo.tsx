export default function Logo({ fontSize = 32 }: { fontSize?: number }) {
	return (
		<span
			style={{
				fontFamily: `var(--font-genos)`,
				fontStyle: "italic",
				fontSize: `${fontSize}px`,
			}}
		>
			<span
				style={{
					fontWeight: "var(--bold-font-weight)",
				}}
			>
				Dev
			</span>
			<span
				style={{
					fontWeight: "var(--regular-font-weight)",
				}}
			>
				.clash
			</span>
		</span>
	);
}
