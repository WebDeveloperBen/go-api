import pkg from "pg";

import { drizzle } from "drizzle-orm/node-postgres";
import { migrate } from "drizzle-orm/node-postgres/migrator";
import path from "path";

(async () => {
  const { Pool } = pkg;
  try {
    // Retrieve arguments from command line
    const [connStr, migrationsDir] = process.argv.slice(2);

    if (!connStr || !migrationsDir) {
      console.error(
        "Usage: node migrate.js <connectionString> <migrationsDir>",
      );
      process.exit(1);
    }

    // Create a connection pool
    const pool = new Pool({
      connectionString: connStr,
      ssl: false, // Only for local/test use
    });

    // Set up Drizzle ORM
    const db = drizzle(pool, {
      logger: false,
      casing: "snake_case",
    });

    // Run migrations
    const migrationsPath = path.resolve(migrationsDir);
    console.log(`Running migrations from: ${migrationsPath}`);
    await migrate(db, { migrationsFolder: migrationsPath });

    console.log("Migrations applied successfully!");
    process.exit(0);
  } catch (error) {
    console.error("Failed to run migrations:", error);
    process.exit(1);
  }
})();
