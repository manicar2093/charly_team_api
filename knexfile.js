require('dotenv').config()

module.exports = {
  client: 'postgresql',
  connection: process.env.DB_URL,
  migrations: {
    tableName: 'schema_migrations',
    directory: './db/migrations',
  },
  seeds: {
    directory: './db/seeds',
  },
};
