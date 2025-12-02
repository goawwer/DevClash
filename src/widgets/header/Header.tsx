// Header.tsx
"use client";

import { useAuth } from "@/app/context/AuthContext";
import HeaderUnauth from "./HeaderUnauth";
import HeaderAuthed from "./HeaderAuthed";

export default function Header() {
	const { isAuthenticated, isLoading, lastKnownAuthenticated } = useAuth();

	// Пока идёт проверка — показываем последний известный header
	if (isLoading) {
		return lastKnownAuthenticated ? <HeaderAuthed /> : <HeaderUnauth />;
	}

	// После загрузки — показываем актуальный
	return isAuthenticated ? <HeaderAuthed /> : <HeaderUnauth />;
}
