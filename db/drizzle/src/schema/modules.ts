import { bigint, pgTable, text } from 'drizzle-orm/pg-core'
import { timestamps } from './columns/helpers'

export const modules = pgTable('modules', {
  id: bigint({ mode: 'number' }).primaryKey().notNull(),
  title: text(),
  description: text(),
  asset: text(),
  ...timestamps,
})
