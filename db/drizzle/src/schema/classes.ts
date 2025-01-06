import { integer, pgTable, text, uuid } from "drizzle-orm/pg-core";
import { timestamps } from "./columns/helpers";

export const classes = pgTable("classes", {
  class_id: uuid().primaryKey().notNull().defaultRandom(),
  name: text().notNull(),
  description: text(),
  course_code: text().notNull(),
  duration: integer(),
  ...timestamps,
});
