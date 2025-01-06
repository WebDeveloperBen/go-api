import {
  boolean,
  integer,
  jsonb,
  pgTable,
  text,
  uuid,
} from "drizzle-orm/pg-core";
import { timestamps } from "./columns/helpers";

export const assets = pgTable("assets", {
  id: uuid().primaryKey().notNull().defaultRandom(),
  fileName: text().notNull().unique(),
  contentType: text().notNull(),
  eTag: text(), //for caching
  containerName: text().notNull(), //bucket name
  uri: text().notNull(),
  size: integer().notNull(),
  metadata: jsonb(),
  isPublic: boolean().notNull().default(true),
  published: boolean().default(true).notNull(),
  ...timestamps,
});
