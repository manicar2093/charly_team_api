module.exports = {
  client: 'postgresql',
  connection: process.env.DB_URL,
  migrations: {
    tableName: 'schema_migrations',
    directory: './internal/db/migrations',
  },
  seeds: {
    directory: './internal/db/seeds',
  },
};
