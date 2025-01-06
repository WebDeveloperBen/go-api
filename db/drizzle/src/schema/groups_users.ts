import { pgTable, text, uuid } from 'drizzle-orm/pg-core'
import { timestamps } from './columns/helpers'
import { groups } from './groups'
import { users } from './users'

export const groups_users = pgTable('groups_users', {
  group_id: uuid('group_id').references(() => groups.group_id),
  user_id: text('user_id').references(() => users.id),
  ...timestamps,
})
