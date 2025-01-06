import { DATABASE_URL } from "client";
import "dotenv/config";
import { Config, defineConfig } from "drizzle-kit";

export default defineConfig({
  schema: "./src/schema/*",
  out: "./src/migrations",
  dialect: "postgresql", // 'postgresql' | 'mysql' | 'sqlite'
  dbCredentials: {
    url: DATABASE_URL,
    ssl: false,
  },
  casing: "snake_case",
}) satisfies Config;
