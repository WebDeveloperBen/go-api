import { integer, pgTable, text, uuid } from 'drizzle-orm/pg-core'
import { courses } from './courses'
import { timestamps } from './columns/helpers'
import { sql } from 'drizzle-orm'

export const sections = pgTable('sections', {
  section_id: uuid()
    .primaryKey()
    .$defaultFn(() => sql`uuid_generate_v4()`),
  courses_id: uuid()
    .notNull()
    .references(() => courses.courses_id, { onDelete: 'cascade' }),
  name: text().notNull(),
  description: text(),
  order: integer(),
  ...timestamps,
})
