import { pgTable, text, timestamp } from 'drizzle-orm/pg-core'

export const jwks = pgTable('jwks', {
  id: text().primaryKey(),
  publicKey: text(),
  privateKey: text(),
  createdAt: timestamp().defaultNow().notNull(),
})
