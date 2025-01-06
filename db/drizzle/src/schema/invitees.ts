import { pgTable, text, uuid } from 'drizzle-orm/pg-core'
import { timestamps } from './columns/helpers'

export const invitees = pgTable('invitees', {
  id: uuid().primaryKey().notNull(),
  email: text().notNull(),
  name: text(),
  invited_by: text().notNull(),
  relation: text(),
  ...timestamps,
})
