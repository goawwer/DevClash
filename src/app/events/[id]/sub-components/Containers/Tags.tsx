type Tag = {
	id: number;
	title: string;
};

type StackContainerProps = {
	tags: Tag[];
};

export default function Tags({ tags }: StackContainerProps) {
	return (
		<>
			{tags.map((tag) => (
				<p key={tag.id}>{tag.title}</p>
			))}
		</>
	);
}
