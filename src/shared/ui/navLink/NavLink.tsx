"use client";

import Link from "next/link";
import styles from "./NavLink.module.scss";
import { FC, ReactNode } from "react";
import { usePathname } from "next/navigation";
import cn from "classnames";

type Props = {
	href: string;
	children: ReactNode;
	exact?: boolean;
};

const NavLink: FC<Props> = ({ href, children, exact = false }) => {
	const pathname = usePathname();
	const isActive = exact
		? exact && pathname === href
		: pathname.startsWith(href);

	return (
		<Link
			href={href}
			className={cn(styles.navLink, {
				[styles[`navLink--active`]]: isActive,
			})}
		>
			{children}
		</Link>
	);
};

export default NavLink;
