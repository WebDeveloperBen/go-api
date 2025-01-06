import { pgTable, serial, text, unique, uuid } from "drizzle-orm/pg-core";
import { timestamps } from "./columns/helpers";
import { roles } from "./roles";
import { users } from "./users";

export const user_roles = pgTable(
  "user_roles",
  {
    id: uuid().defaultRandom().primaryKey().notNull(),
    user_id: text()
      .notNull()
      .references(() => users.id, { onDelete: "cascade" }),
    role_id: serial()
      .notNull()
      .references(() => roles.role_id, { onDelete: "cascade" }),
    context_id: uuid().notNull(), // This can be a course, organization, etc.
    context_type: text().notNull(), // 'course', 'organization', etc.
    ...timestamps,
  },
  (table) => [
    unique("user_roles_user_id_role_key").on(table.user_id, table.role_id),
  ],
);
