import { pgTable, primaryKey, uuid } from "drizzle-orm/pg-core";
import { lessons } from "./lessons";
import { resources } from "./resources";

export const lesson_resources = pgTable(
  "lesson_resources",
  {
    lesson_id: uuid()
      .notNull()
      .references(() => lessons.lesson_id, { onDelete: "cascade" }),
    resource_id: uuid()
      .notNull()
      .references(() => resources.resource_id, { onDelete: "cascade" }),
  },
  (table) => [
    primaryKey({
      columns: [table.lesson_id, table.resource_id],
      name: "lesson_resources_pkey",
    }),
  ],
);
