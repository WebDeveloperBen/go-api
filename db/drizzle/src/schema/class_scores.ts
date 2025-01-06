import { integer, pgTable, primaryKey, uuid } from "drizzle-orm/pg-core";
import { classes } from "./classes";
import { users } from "./users";

export const class_scores = pgTable(
  "class_scores",
  {
    class_id: uuid()
      .notNull()
      .references(() => classes.class_id),
    user_id: uuid().references(() => users.id),
    score: integer().notNull(),
    week: integer().notNull(),
  },
  (table) => [
    primaryKey({
      columns: [table.class_id, table.user_id, table.week],
      name: "class_scores_pkey",
    }),
  ],
);
