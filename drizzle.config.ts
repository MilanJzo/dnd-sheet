import type { Config } from "drizzle-kit";

export default {
    schema: "./src/lib/server/schema.ts",
    driver: "turso",
    out: "./drizzle/migrations",
    dbCredentials: {
        url: "file:./drizzle/db.sqlite"
    }
} satisfies Config;
