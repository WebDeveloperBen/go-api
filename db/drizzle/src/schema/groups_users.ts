import { pgTable, uuid } from "drizzle-orm/pg-core";
import { timestamps } from "./columns/helpers";
import { groups } from "./groups";
import { users } from "./users";

export const groups_users = pgTable("groups_users", {
  groupId: uuid().references(() => groups.groupId),
  userId: uuid().references(() => users.id),
  ...timestamps,
});
