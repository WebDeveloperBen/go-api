import { pgTable, primaryKey, uuid } from "drizzle-orm/pg-core";
import { educational_standards } from "./educational_standards";
import { questions } from "./questions";

export const educations_standards_questions_mapping = pgTable(
  "educations_standards_questions_mapping",
  {
    question_id: uuid()
      .notNull()
      .references(() => questions.id),
    standard_id: uuid()
      .notNull()
      .references(() => educational_standards.id),
  },
  (table) => [
    primaryKey({
      columns: [table.question_id, table.standard_id],
      name: "educations_standards_questions_mapping_pkey",
    }),
  ],
);
