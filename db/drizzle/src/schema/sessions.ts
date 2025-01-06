import { pgTable, text, timestamp } from 'drizzle-orm/pg-core'
import { users } from './users'

export const sessions = pgTable('session', {
  id: text().primaryKey(),
  expiresAt: timestamp().notNull(),
  ipAddress: text(),
  userAgent: text(),
  userId: text()
    .notNull()
    .references(() => users.id),
})
