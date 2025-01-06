import { integer, jsonb, pgTable, uuid } from "drizzle-orm/pg-core";
import { timestamps } from "./columns/helpers";
import { homework } from "./homework";
import { users } from "./users";

export const homework_submissions = pgTable("homework_submissions", {
  id: uuid().defaultRandom().primaryKey().notNull(),
  user_id: uuid().references(() => users.id),
  homework_id: uuid().references(() => homework.id),
  progress: integer().default(0),
  content: jsonb(),
  ...timestamps,
});
