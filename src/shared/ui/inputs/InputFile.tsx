'use client'

import styles from "./InputFile.module.scss"
import { ChangeEvent, FC, useRef, useState } from 'react';

interface FileUploadInputProps {
	width?: number; 
  onFileSelect?: (file: File) => void;
  label?: string;
  formats?: string[]
}

const FileUploadInput: FC<FileUploadInputProps> = ({
  onFileSelect,
  label = "Изображение",
  width = 27,
  formats
}) => {
  const inputRef = useRef<HTMLInputElement>(null);
  const [fileURL, setFileURL] = useState<string>("");

  const handleFileChange = (e: ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];

    if (file && onFileSelect) {
	  const url = URL.createObjectURL(file)
	  setFileURL(url)
      onFileSelect(file);
    }
  };

  return (
    <div className={styles.inputFile} style={{width: `${width}rem`}}>
    	
		<div 
			className={styles.inputFile__preview} 
			style={fileURL ? 
				{width: `${width}rem`, backgroundImage: `url(${fileURL})`, backgroundColor: "var(--white-color)"}  :
				{width: `${width}rem`, backgroundColor: "var(--dark-blue-color)"}}>
			
			{fileURL ? 
				<></>: 
				(<p>
					{label} не выбран <br/> <br/> 
					Форматы: {formats && parseFormats(formats)}
				</p>)
			}
		</div>
	  
	  <label 
        htmlFor="file-upload"
		className={styles.inputFile__label}
      >
        {`Выберите ${label}`}
      </label>
      
      
	  <input
        id="file-upload"
        type="file"
        ref={inputRef}
        onChange={handleFileChange}
        accept=".jpg,.jpeg,.png,.svg"
        style={{ display: 'none' }}
      />

	  <p>ФАЙЛ: {fileURL}</p>
    </div>
  );
};

export default FileUploadInput;

function parseFormats(formats: string[]) {
	return (
	formats.map((format, id) => {
		if (id !== formats.length -1) {
			return <span key={id}>{`.${format}, `}</span>
		}
		
		return <span key={id}>{`.${format}`}</span>
	}))
}