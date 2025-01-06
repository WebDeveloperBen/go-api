import { integer, pgTable, text, uuid } from 'drizzle-orm/pg-core'
import { timestamps } from './columns/helpers'

export const educational_standards = pgTable('educational_standards', {
  id: uuid().defaultRandom().primaryKey().notNull(),
  code: text().notNull(),
  description: text().notNull(),
  subject: text().notNull(),
  category: text().notNull(),
  sub_category: text().notNull(),
  level: integer().notNull(),
  ...timestamps,
})
