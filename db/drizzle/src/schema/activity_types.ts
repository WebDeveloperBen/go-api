import { boolean, pgTable, text, uuid } from "drizzle-orm/pg-core";

export const activity_types = pgTable("activity_types", {
  id: uuid().primaryKey().notNull().defaultRandom(),
  name: text().notNull(),
  published: boolean().default(false).notNull(),
  category: text(),
  card_color: text(),
});
