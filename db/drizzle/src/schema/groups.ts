import { pgTable, text, uuid } from "drizzle-orm/pg-core";
import { timestamps } from "./columns/helpers";

export const groups = pgTable("groups", {
  groupId: uuid().primaryKey().notNull().defaultRandom(),
  name: text().notNull(),
  courseId: uuid(), //this would be the 'course', 'organisations', 'social' groups - hence no foreign key defined
  tags: text().array(),
  ...timestamps,
});
