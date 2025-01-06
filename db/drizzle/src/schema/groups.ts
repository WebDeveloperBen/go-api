import { pgTable, text, uuid } from 'drizzle-orm/pg-core'
import { timestamps } from './columns/helpers'

export const groups = pgTable('groups', {
  group_id: uuid('group_id').primaryKey(),
  name: text('name').notNull(),
  course_id: uuid('course_id'), //this would be the 'course', 'organisations', 'social' groups - hence no foreign key defined
  tags: text('tags').array(),
  ...timestamps,
})
