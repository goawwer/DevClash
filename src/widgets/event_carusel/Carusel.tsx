"use client";
import { FC, useState } from "react";
import cn from "classnames"; // 👈 импортируем classnames
import styles from "./Carusel.module.scss";
import animation from "./Animation.module.scss";
import { PhotoCard } from "@/shared/templates";

type Event = {
	title: string;
	started_at: string;
	ended_at: string;
	photo_src: string;
	photo_alt: string;
	stack: string;
};

type Props = {
	events: Event[];
};

const Carusel: FC<Props> = ({ events }) => {
	const [currentEvent, setEvent] = useState(0);

	const handleSwitch = (increseValue: number) => {
		if (currentEvent + increseValue >= events.length) {
			setEvent(0);
			return;
		}

		if (currentEvent + increseValue === -1) {
			setEvent(events.length - 1);
			return;
		}

		setEvent(currentEvent + increseValue);
	};

	const treeArray = () => {
		if (currentEvent - 1 === -1) {
			return {
				current: currentEvent,
				prev: events.length - 1,
				next: currentEvent + 1,
			};
		}

		if (currentEvent + 1 >= events.length) {
			return {
				current: currentEvent,
				prev: currentEvent - 1,
				next: 0,
			};
		}
		return {
			current: currentEvent,
			prev: currentEvent - 1,
			next: currentEvent + 1,
		};
	};

	return (
		<div className={styles.carusel}>
			<div className={styles.carusel__portableButtons}>
				<button
					className={styles.carusel__switchButton}
					type="button"
					onClick={() => handleSwitch(-1)}
				>
					{"<"}
				</button>

				<button
					className={styles.carusel__switchButton}
					type="button"
					onClick={() => handleSwitch(1)}
				>
					{">"}
				</button>
			</div>
			{/* current */}
			<div
				key={currentEvent}
				className={cn(styles.carusel__currentEvent, animation.current)}
			>
				<button
					className={styles.carusel__switchButton}
					type="button"
					onClick={() => handleSwitch(-1)}
				>
					{"<"}
				</button>

				<PhotoCard {...events[treeArray().current]} />

				<button
					className={styles.carusel__switchButton}
					type="button"
					onClick={() => handleSwitch(1)}
				>
					{">"}
				</button>
			</div>

			<div className={styles.carusel__otherEvents}>
				<div className={styles.carusel__otherEventsContainer}>
					{/* prev */}
					<div
						className={cn(
							styles.carusel__prevEvent,
							animation.secondary
						)}
						key={currentEvent - 1}
					>
						<PhotoCard
							photo_src={events[treeArray().prev].photo_src}
							width={25}
							photo_alt={events[treeArray().prev].photo_alt}
						/>
					</div>

					{/* next */}
					<div
						key={currentEvent + 1}
						className={cn(
							styles.carusel__nextEvent,
							animation.secondary
						)}
					>
						<PhotoCard
							photo_src={events[treeArray().next].photo_src}
							width={25}
							photo_alt={events[treeArray().next].photo_alt}
						/>
					</div>
				</div>
			</div>
		</div>
	);
};

export default Carusel;
