import { integer, pgTable, text, uuid } from 'drizzle-orm/pg-core'
import { timestamps } from './columns/helpers'
import { organisations } from './organisations'
import { roles } from './roles'
import { users } from './users'

export const organisation_users = pgTable('organisation_users', {
  organisation_id: uuid().references(() => organisations.organisations_id),
  user_id: text().references(() => users.id),
  role_id: integer().references(() => roles.role_id),
  ...timestamps,
})
