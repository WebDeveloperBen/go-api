import { integer, jsonb, pgTable, text, uuid } from 'drizzle-orm/pg-core'

export const questions = pgTable('questions', {
  id: uuid().defaultRandom().primaryKey().notNull(),
  egl: integer(),
  curriculum_reference: text().array(),
  cognitive_skill: text().array(),
  reading_ability: integer(),
  writing_ability: integer(),
  listening_ability: integer(),
  title: text(),
  question: text(),
  answer: text(),
  options: jsonb(),
  type: text(),
})
