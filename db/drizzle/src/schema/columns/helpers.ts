import { sql } from "drizzle-orm";
import { timestamp } from "drizzle-orm/pg-core";

export const timestamps = {
  updated_at: timestamp({ mode: "date", withTimezone: true })
    .defaultNow()
    .notNull()
    .$onUpdate(() => new Date()),
  created_at: timestamp({
    mode: "date",
    withTimezone: true,
  })
    .notNull()
    .defaultNow(),
};

export const uuid = sql`uuid_generate_v4()`;
