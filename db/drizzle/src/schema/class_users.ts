import { pgTable, text, uuid } from 'drizzle-orm/pg-core'
import { classes } from './classes'
import { users } from './users'

export const class_users = pgTable('class_users', {
  class_id: uuid()
    .notNull()
    .references(() => classes.class_id),
  user_id: text()
    .notNull()
    .references(() => users.id),
})
