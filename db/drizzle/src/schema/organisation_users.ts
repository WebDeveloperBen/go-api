import { integer, pgTable, uuid } from "drizzle-orm/pg-core";
import { timestamps } from "./columns/helpers";
import { organisations } from "./organisations";
import { roles } from "./roles";
import { users } from "./users";

export const organisation_users = pgTable("organisation_users", {
  organisation_id: uuid().references(() => organisations.organisation_id),
  user_id: uuid().references(() => users.id),
  role_id: integer().references(() => roles.role_id),
  ...timestamps,
});
