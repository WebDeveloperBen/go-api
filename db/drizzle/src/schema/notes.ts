import { pgTable, text, uuid } from "drizzle-orm/pg-core";
import { timestamps } from "./columns/helpers";
import { lessons } from "./lessons";
import { users } from "./users";

export const notes = pgTable("notes", {
  note_id: uuid().primaryKey().notNull().defaultRandom(),
  user_id: uuid().references(() => users.id),
  content: text(),
  lesson_id: uuid().references(() => lessons.lesson_id, {
    onDelete: "set null",
  }),
  ...timestamps,
});
