import { fail, redirect } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
import { auth } from "$lib/server/lucia";

export const actions: Actions = {
    default: async ({ request, locals }) => {
        const formData = await request.formData();
        const username = formData.get("username");
        const password = formData.get("password");

        if (typeof username !== "string" || username.length < 4 || username.length > 31)
            return fail(400, { message: "username must be between 4 and 31 characters" });
        if (typeof password !== "string" || password.length < 6 || password.length > 255)
            return fail(400, { message: "password must be between 6 and 255 characters" });

        try {
            const user = await auth.createUser({
                key: {
                    providerId: "username",
                    providerUserId: username.toLocaleLowerCase(),
                    password
                },
                attributes: {
                    username
                }
            });
            const session = await auth.createSession({
                userId: user.userId,
                attributes: {}
            });
            locals.auth.setSession(session);
        } catch (e) {
            //TODO: check for unique constraint error in user table (duplicate username)
            if (e instanceof Error) return fail(500, { message: e.message });
        }

        throw redirect(302, "/");
    }
};

export const load: PageServerLoad = async ({ locals }) => {
    const session = await locals.auth.validate();
    if (session) throw redirect(302, "/");
    return {};
};
