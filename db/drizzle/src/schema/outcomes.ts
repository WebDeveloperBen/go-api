import { pgTable, serial, text, unique } from 'drizzle-orm/pg-core'
import { timestamps } from './columns/helpers'

export const outcomes = pgTable(
  'outcomes',
  {
    outcome_id: serial().primaryKey().notNull(),
    name: text().notNull(),
    code: text().notNull(),
    ...timestamps,
  },
  (table) => [unique('outcomes_code_key').on(table.code)]
)
