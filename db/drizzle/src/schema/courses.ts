import { integer, pgTable, text, uuid } from 'drizzle-orm/pg-core'
import { timestamps } from './columns/helpers'
import { sql } from 'drizzle-orm'

export const courses = pgTable('courses', {
  courses_id: uuid()
    .primaryKey()
    .$defaultFn(() => sql`uuid_generate_v4()`),
  name: text().notNull(),
  description: text(),
  chapters: integer(),
  sections: integer(),
  duration: integer(),
  category: text(),
  image: text(),
  ...timestamps,
})
