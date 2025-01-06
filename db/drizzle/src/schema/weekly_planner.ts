import { integer, jsonb, pgTable, text, uuid } from 'drizzle-orm/pg-core'
import { classes } from './classes'
import { timestamps } from './columns/helpers'

export const weekly_planner = pgTable('weekly_planner', {
  id: uuid().defaultRandom().primaryKey().notNull(),
  class_id: uuid()
    .notNull()
    .references(() => classes.class_id),
  term: text().notNull(),
  week: integer().notNull(),
  content: jsonb(),
  ...timestamps,
})
