import { fail, redirect } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";

import { auth } from "$lib/server/lucia";

export const load: PageServerLoad = async ({ locals }) => {
    const session = await locals.auth.validate();
    if (!session) throw redirect(302, "/login");

    return {
        userId: session?.user.userId,
        username: session?.user.username
    };
};

export const actions: Actions = {
    logout: async ({ locals }) => {
        const session = await locals.auth.validate();
        if (!session) return fail(401);
        await auth.invalidateSession(session.sessionId);
        locals.auth.setSession(null);
        throw redirect(302, "/login");
    }
};
