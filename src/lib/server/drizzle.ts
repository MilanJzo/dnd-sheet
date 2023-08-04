import { drizzle } from "drizzle-orm/libsql";
import { createClient } from "@libsql/client";
import { env } from "$env/dynamic/private";

export const client = createClient({ url: env.DATABASE_URL, authToken: env.DATABASE_AUTH_TOKEN });
// export const client = createClient({ url: "file:./drizzle/db.sqlite" });
export const db = drizzle(client);
