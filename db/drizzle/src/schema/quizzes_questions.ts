import { integer, jsonb, pgTable, text, uuid } from 'drizzle-orm/pg-core'
import { quizzes } from './quizzes'
import { timestamps } from './columns/helpers'

export const quizzes_questions = pgTable('quizzes_questions', {
  id: uuid().primaryKey().notNull(),
  quizzes_id: uuid().references(() => quizzes.id),
  question: text(),
  answer: text(),
  options: jsonb(),
  type: text(),
  question_order: integer(),
  resources: jsonb().array(),
  ...timestamps,
})
