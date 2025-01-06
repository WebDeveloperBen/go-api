import { index, pgTable, text, timestamp, uuid } from "drizzle-orm/pg-core";
import { timestamps } from "./columns/helpers";
import { users } from "./users";
export const presence = pgTable(
  "presence",
  {
    id: uuid().primaryKey().notNull().defaultRandom(),
    user_id: uuid()
      .notNull()
      .references(() => users.id),
    last_status: text().notNull(),
    last_login: timestamp({
      mode: "date",
      withTimezone: true,
    }),
    last_logout: timestamp({
      mode: "date",
      withTimezone: true,
    }),
    ...timestamps,
  },
  (table) => [index("user_presence_index").on(table.user_id)],
);
