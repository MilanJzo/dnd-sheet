import { drizzle } from "drizzle-orm/libsql";
import { createClient } from "@libsql/client";

export const client = createClient({ url: "file:./drizzle/db.sqlite" });
export const db = drizzle(client);
