import { sql } from "drizzle-orm";
import { boolean, pgTable, text } from "drizzle-orm/pg-core";
import { timestamps } from "./columns/helpers";

export const users = pgTable("users", {
  id: text()
    .primaryKey()
    .default(sql`uuid_generate_v4()`),
  fullname: text().notNull(),
  email: text().notNull().unique(),
  emailVerified: boolean().default(false),
  image: text(),
  ...timestamps,
});
