require('dotenv').config()

module.exports = {
  client: 'postgresql',
  connection: {
    host: process.env.DB_HOST,
    port: process.env.DB_PORT,
    database: process.env.DB_NAME,
    user: process.env.DB_USER,
    password: process.env.DB_PASSWORD,
  },
  migrations: {
    tableName: 'schema_migrations',
    directory: './db/migrations',
  },
  seeds: {
    directory: './db/seeds',
  },
};