import { jsonb, pgTable, text, uuid } from "drizzle-orm/pg-core";
import { timestamps } from "./columns/helpers";
import { courses } from "./courses";
import { organisations } from "./organisations";
import { users } from "./users";

export const enrolments = pgTable("enrolments", {
  organisation_id: uuid().references(() => organisations.organisation_id),
  courseId: uuid().references(() => courses.courses_id),
  userId: uuid().references(() => users.id),
  name: text().notNull(),
  tags: text().array(),
  content: jsonb(),
  ...timestamps,
});
