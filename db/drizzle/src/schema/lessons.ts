import {
  integer,
  jsonb,
  pgTable,
  text,
  unique,
  uuid,
} from "drizzle-orm/pg-core";
import { timestamps } from "./columns/helpers";
import { courses } from "./courses";
import { sections } from "./sections";

export const lessons = pgTable(
  "lessons",
  {
    lesson_id: uuid().defaultRandom().primaryKey().notNull(),
    section_id: uuid()
      .notNull()
      .references(() => sections.section_id, { onDelete: "cascade" }),
    course_id: uuid().references(() => courses.courses_id),
    title: text().notNull(),
    summary: text(),
    content: jsonb(),
    duration: integer(),
    order_in_subject: integer(),
    image: text(),
    ...timestamps,
  },
  (t) => [unique("unique_lesson_key").on(t.title, t.section_id, t.course_id)],
);
