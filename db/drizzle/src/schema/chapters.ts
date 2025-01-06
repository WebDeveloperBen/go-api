import {
  boolean,
  integer,
  jsonb,
  pgTable,
  text,
  uuid,
} from 'drizzle-orm/pg-core'
import { timestamps } from './columns/helpers'
import { lessons } from './lessons'
import { sql } from 'drizzle-orm'

export const chapters = pgTable('chapters', {
  chapter_id: uuid()
    .primaryKey()
    .$defaultFn(() => sql`uuid_generate_v4()`),
  lesson_id: uuid()
    .notNull()
    .references(() => lessons.lesson_id),
  nav_item_name: text().notNull(),
  description: text(),
  order: integer(),
  content: jsonb(),
  published: boolean().default(false).notNull(),
  title: text().notNull(),
  chapter_number: integer(),
  ...timestamps,
})
