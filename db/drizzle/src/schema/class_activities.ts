import { pgTable, text, timestamp, uuid } from "drizzle-orm/pg-core";
import { classes } from "./classes";
import { homework } from "./homework";
import { timestamps } from "./columns/helpers";

export const classes_activities = pgTable("classes_activities", {
  id: uuid().defaultRandom().primaryKey().notNull(),
  term: text(),
  startDate: timestamp({ withTimezone: true, mode: "string" }),
  endDate: timestamp({ withTimezone: true, mode: "string" }),
  homework_id: uuid()
    .notNull()
    .references(() => homework.id, {
      onDelete: "cascade",
      onUpdate: "cascade",
    }),
  day: text(),
  period: text(),
  class_id: uuid()
    .notNull()
    .references(() => classes.class_id),
  ...timestamps,
});
