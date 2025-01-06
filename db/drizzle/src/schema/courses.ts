import { integer, pgTable, text, uuid } from "drizzle-orm/pg-core";
import { timestamps } from "./columns/helpers";

export const courses = pgTable("courses", {
  courses_id: uuid().primaryKey().notNull().defaultRandom(),
  name: text().notNull(),
  description: text(),
  chapters: integer(),
  sections: integer(),
  duration: integer(),
  category: text(),
  image: text(),
  ...timestamps,
});
