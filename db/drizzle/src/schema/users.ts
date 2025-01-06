import { boolean, pgTable, text, uuid } from "drizzle-orm/pg-core";
import { timestamps } from "./columns/helpers";

export const users = pgTable("users", {
  id: uuid().primaryKey().notNull().defaultRandom(),
  fullname: text().notNull(),
  email: text().notNull().unique(),
  emailVerified: boolean().notNull().default(false),
  image: text(),
  ...timestamps,
});
