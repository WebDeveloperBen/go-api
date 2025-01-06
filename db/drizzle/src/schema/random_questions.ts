import { integer, jsonb, pgTable, text, uuid } from "drizzle-orm/pg-core";

export const random_questions = pgTable("random_questions", {
  id: uuid().primaryKey().notNull().defaultRandom(),
  egl: integer(),
  curriculum_reference: text().array(),
  cognitive_skill: text().array(),
  reading_ability: integer(),
  writing_ability: integer(),
  listening_ability: integer(),
  title: text(),
  question: text(),
  answer: text(),
  options: jsonb(),
  type: text(),
});
