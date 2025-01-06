import { integer, jsonb, pgTable, text, uuid } from 'drizzle-orm/pg-core'
import { timestamps } from './columns/helpers'
import { homework } from './homework'
import { users } from './users'

export const homework_submissions = pgTable('homework_submissions', {
  id: uuid().defaultRandom().primaryKey().notNull(),
  user_id: text().references(() => users.id),
  homework_id: uuid().references(() => homework.id),
  progress: integer().default(0),
  content: jsonb(),
  ...timestamps,
})
