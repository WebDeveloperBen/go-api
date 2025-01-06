import { integer, pgTable, text, uuid } from 'drizzle-orm/pg-core'
import { timestamps } from './columns/helpers'

export const quizzes = pgTable('quizzes', {
  title: text(),
  description: text(),
  id: uuid().defaultRandom().primaryKey().notNull(),
  total_questions: integer(),
  ...timestamps,
})
