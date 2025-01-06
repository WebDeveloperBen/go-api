import { jsonb, pgTable, text, uuid } from 'drizzle-orm/pg-core'
import { timestamps } from './columns/helpers'
import { users } from './users'

export const homework = pgTable('homework', {
  id: uuid().defaultRandom().primaryKey().notNull(),
  name: text().notNull(),
  created_by: text().references(() => users.id), //user
  tags: text().array(),
  content: jsonb(),
  ...timestamps,
})
