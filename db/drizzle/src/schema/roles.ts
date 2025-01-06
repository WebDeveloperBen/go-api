import { pgTable, serial, text } from 'drizzle-orm/pg-core'
import { timestamps } from './columns/helpers'

export const roles = pgTable('roles', {
  role_id: serial().primaryKey().notNull(),
  name: text(),
  description: text(),
  ...timestamps,
})
