import { pgTable, text, timestamp } from 'drizzle-orm/pg-core'
import { users } from './users'

export const account = pgTable('account', {
  id: text().primaryKey(),
  accountId: text().notNull(),
  providerId: text().notNull(),
  userId: text()
    .notNull()
    .references(() => users.id),
  accessToken: text(),
  refreshToken: text(),
  idToken: text(),
  expiresAt: timestamp(),
  password: text(),
})
