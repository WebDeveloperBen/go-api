import { pgTable, text, timestamp, uuid } from "drizzle-orm/pg-core";

export const jwks = pgTable("jwks", {
  id: uuid().defaultRandom().primaryKey().notNull(),
  publicKey: text(),
  privateKey: text(),
  createdAt: timestamp().defaultNow().notNull(),
});
