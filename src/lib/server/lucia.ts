import { lucia } from "lucia";
import { sveltekit } from "lucia/middleware";
import { dev } from "$app/environment";
import { pg } from "@lucia-auth/adapter-postgresql";
import { pool } from "$lib/server/drizzle";

export const auth = lucia({
    env: dev ? "DEV" : "PROD",
    middleware: sveltekit(),
    adapter: pg(pool, {
        user: "auth_user",
        key: "user_key",
        session: "user_session"
    }),
    getUserAttributes: (data) => {
        return {
            username: data.username
        };
    }
});

export type Auth = typeof auth;
