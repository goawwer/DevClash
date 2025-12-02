// context/AuthContext.tsx
"use client";

import {
	createContext,
	useContext,
	useEffect,
	useState,
	ReactNode,
	useRef,
} from "react";
import { usePathname } from "next/navigation";
import { authCheck } from "@/features/api";

type AuthContextType = {
	isAuthenticated: boolean | null;
	isLoading: boolean;
	lastKnownAuthenticated: boolean; // ← новое поле
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
	const [isAuthenticated, setIsAuthenticated] = useState<boolean | null>(
		null
	);
	const [isLoading, setIsLoading] = useState(true);
	const lastKnownRef = useRef<boolean>(false); // ← храним последнее значение
	const pathname = usePathname();

	useEffect(() => {
		let cancelled = false;

		const verifyAuth = async () => {
			setIsLoading(true);
			try {
				const auth = await authCheck();
				if (!cancelled) {
					setIsAuthenticated(auth);
					lastKnownRef.current = auth; // ← обновляем последнее значение
				}
			} catch {
				if (!cancelled) {
					setIsAuthenticated(false);
					lastKnownRef.current = false;
				}
			} finally {
				if (!cancelled) {
					setIsLoading(false);
				}
			}
		};

		verifyAuth();

		return () => {
			cancelled = true;
		};
	}, [pathname]);

	return (
		<AuthContext.Provider
			value={{
				isAuthenticated,
				isLoading,
				lastKnownAuthenticated: lastKnownRef.current,
			}}
		>
			{children}
		</AuthContext.Provider>
	);
}

export function useAuth() {
	const context = useContext(AuthContext);
	if (!context) throw new Error("useAuth must be used within AuthProvider");
	return context;
}
