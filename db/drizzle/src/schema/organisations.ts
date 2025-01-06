import { pgTable, text, uuid } from 'drizzle-orm/pg-core'
import { timestamps } from './columns/helpers'

export const organisations = pgTable('organisations', {
  organisations_id: uuid().primaryKey().notNull(),
  name: text().notNull(),
  address: text(),
  contact_details: text(),
  ...timestamps,
})
