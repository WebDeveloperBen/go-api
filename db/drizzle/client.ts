import { config } from "dotenv";
import { drizzle } from "drizzle-orm/node-postgres";
import pkg from "pg";
import { z } from "zod";

import * as schema from "./src/schema";

import path from "path";

// Explicitly load the .env file from the root
config({ path: path.resolve(__dirname, "../../.env") });

const { Pool } = pkg;

// Zod schema for validating the DATABASE_URL
const envSchema = z.object({
  DATABASE_URL: z.string().url(),
});

const env = envSchema.parse(process.env);

const DATABASE_URL = env.DATABASE_URL;

const pool = new Pool({
  connectionString: DATABASE_URL,
  ssl: false, // only used for migrations and schema management not used in production setting
});

const db = drizzle(pool, {
  logger: true,
  schema,
  casing: "snake_case",
});

const createTestMigrator = (conn: string) => {
  const pool = new Pool({
    connectionString: conn,
    ssl: false, // only used for test containers
  });
  return drizzle(pool, { logger: true, schema, casing: "snake_case" });
};

export { db, DATABASE_URL, createTestMigrator };
