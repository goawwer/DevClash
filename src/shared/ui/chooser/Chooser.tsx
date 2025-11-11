"use client";

import { FC, ReactNode, useState } from "react";
import styles from "./Chooser.module.scss";
import cn from "classnames";

type Variant = {
	name: string;
	value: ReactNode;
};

type Props = {
	variants: Variant[];
};

const Chooser: FC<Props> = ({ variants }) => {
	const [currentVariant, setVariant] = useState(0);

	const changeVariant = (variantID: number) => setVariant(variantID);

	return (
		<div className={styles.chooser}>
			<div className={styles.chooser__variants}>
				{variants.map((variant, id) => {
					return (
						<button
							key={id}
							onClick={() => changeVariant(id)}
							className={cn(styles.chooser__button, {
								[styles[`chooser__button--active`]]:
									id === currentVariant,
							})}
						>
							{variant.name}
						</button>
					);
				})}
			</div>

			<>{variants[currentVariant].value}</>
		</div>
	);
};

export default Chooser;
