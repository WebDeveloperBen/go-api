import { bigint, jsonb, pgTable, text, uuid } from "drizzle-orm/pg-core";
import { timestamps } from "./columns/helpers";

export const resources = pgTable("resources", {
  resource_id: uuid().primaryKey().notNull().defaultRandom(),
  resource_type: text(),
  size: bigint({ mode: "number" }).notNull(),
  metadata: jsonb(),
  etag: text(),
  provider: text().default("internal").notNull(),
  uri: text(),
  ...timestamps,
});
