import { integer, sqliteTable, text } from "drizzle-orm/sqlite-core";
// https://orm.drizzle.team/docs/sql-schema-declaration

export const user = sqliteTable("user", {
    id: text("id", { length: 31 }).primaryKey().notNull(),
    username: text("username", { length: 255 }).notNull().unique()
});

export const user_key = sqliteTable("user_key", {
    id: text("id", { length: 255 }).primaryKey().notNull(),
    user_id: text("user_id", { length: 15 })
        .notNull()
        .references(() => user.id),
    hashed_password: text("hashed_password", { length: 255 })
});

export const user_session = sqliteTable("user_session", {
    id: text("id", { length: 127 }).primaryKey().notNull(),
    user_id: text("user_id", { length: 15 })
        .notNull()
        .references(() => user.id),
    active_expires: integer("active_expires").notNull(),
    idle_expires: integer("idle_expires").notNull()
});
