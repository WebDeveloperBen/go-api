import { integer, pgTable, text, uuid } from 'drizzle-orm/pg-core'
import { timestamps } from './columns/helpers'
import { sql } from 'drizzle-orm'

export const classes = pgTable('classes', {
  class_id: uuid()
    .primaryKey()
    .$defaultFn(() => sql`uuid_generate_v4()`),
  name: text().notNull(),
  description: text(),
  course_code: text().notNull(),
  duration: integer(),
  ...timestamps,
})
