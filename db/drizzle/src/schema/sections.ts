import { integer, pgTable, text, uuid } from "drizzle-orm/pg-core";
import { courses } from "./courses";
import { timestamps } from "./columns/helpers";

export const sections = pgTable("sections", {
  section_id: uuid().primaryKey().defaultRandom().notNull(),
  courses_id: uuid()
    .notNull()
    .references(() => courses.courses_id, { onDelete: "cascade" }),
  name: text().notNull(),
  description: text(),
  order: integer(),
  ...timestamps,
});
