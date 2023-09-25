import type { Config } from "drizzle-kit";
import * as dotenv from "dotenv";
dotenv.config({ path: "./.env" });

export default {
    schema: "./src/lib/server/schema.ts",
    driver: "pg",
    out: "./drizzle/migrations",
    dbCredentials: {
        connectionString: process.env.DB_URL as string,
    },
} satisfies Config;
