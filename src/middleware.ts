import { authCheck } from "@/shared/api/middleware/authCheck";
import { NextRequest, NextResponse } from "next/server";

export async function middleware(req: NextRequest) {
	const authResult = await authCheck(req);
	if (authResult) return authResult;

	return NextResponse.next();
}

export const config = {
	matcher: ["/in:path*"],
};
