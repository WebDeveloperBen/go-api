import { pgTable, uuid } from "drizzle-orm/pg-core";
import { homework } from "./homework";
import { questions } from "./questions";

export const homework_questions = pgTable("homework_questions", {
  id: uuid().defaultRandom().primaryKey().notNull(),
  question_id: uuid()
    .references(() => questions.id, { onDelete: "cascade" })
    .references(() => questions.id),
  homework_id: uuid()
    .references(() => homework.id, { onDelete: "cascade" })
    .references(() => homework.id),
});
