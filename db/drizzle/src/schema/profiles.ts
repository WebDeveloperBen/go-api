import { boolean, integer, pgTable, text, uuid } from 'drizzle-orm/pg-core'
import { timestamps } from './columns/helpers'

export const profiles = pgTable('profiles', {
  id: uuid().primaryKey().notNull(),
  username: text(),
  egl: integer(),
  enrolledAt: text(),
  termsAccepted: boolean().default(false).notNull(),
  ...timestamps,
})
