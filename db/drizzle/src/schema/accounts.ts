import { pgTable, text, timestamp, uuid } from "drizzle-orm/pg-core";
import { users } from "./users";

export const account = pgTable("account", {
  id: uuid().primaryKey().notNull().defaultRandom(),
  accountId: text().notNull(),
  providerId: text().notNull(),
  userId: uuid()
    .notNull()
    .references(() => users.id),
  accessToken: text(),
  refreshToken: text(),
  idToken: text(),
  expiresAt: timestamp({
    mode: "date",
    withTimezone: true,
  }),
  password: text(),
});
