"use client";

import styles from "./InputFile.module.scss";
import { ChangeEvent, forwardRef, useState } from "react";

interface FileUploadInputProps {
	onFileSelect: (file: File) => void;
	width?: number;
	label?: string;
	name?: string;
	formats?: string[];
}

const FileUploadInput = forwardRef<HTMLInputElement, FileUploadInputProps>(
	(
		{ label = "Изображение", width = 27, formats, name, onFileSelect },
		ref
	) => {
		const [fileURL, setFileURL] = useState<string>("");

		const handleFileChange = (e: ChangeEvent<HTMLInputElement>) => {
			const file = e.target.files?.[0];

			if (file && onFileSelect) {
				const url = URL.createObjectURL(file);
				setFileURL(url);
				onFileSelect?.(file);
			}
		};

		return (
			<div className={styles.inputFile} style={{ width: `${width}rem` }}>
				<p className={styles.inputFile__label}>{label}</p>
				<div
					className={styles.inputFile__preview}
					style={
						fileURL
							? {
									width: `${width}rem`,
									backgroundImage: `url(${fileURL})`,
									backgroundColor: "var(--white-color)",
							  }
							: {
									width: `${width}rem`,
									backgroundColor: "var(--dark-blue-color)",
							  }
					}
				>
					{fileURL ? (
						<></>
					) : (
						<p>
							Файл не выбран <br /> <br />
							Форматы: {formats && parseFormats(formats)}
						</p>
					)}
				</div>

				<label
					htmlFor={`file-upload-${name}`}
					className={styles.inputFile__button}
				>
					{`Выберите ${label}`}
				</label>

				<input
					id={`file-upload-${name}`}
					type="file"
					onChange={(e) => {
						handleFileChange(e);
					}}
					accept=".jpg,.jpeg,.png,.svg"
					style={{ display: "none" }}
					ref={ref}
				/>
			</div>
		);
	}
);

FileUploadInput.displayName = "Custom_File_Input";

export default FileUploadInput;

function parseFormats(formats: string[]) {
	return formats.map((format, id) => {
		if (id !== formats.length - 1) {
			return <span key={id}>{`.${format}, `}</span>;
		}

		return <span key={id}>{`.${format}`}</span>;
	});
}
