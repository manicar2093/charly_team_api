
exports.up = function(knex) {

    const biotypes = () => knex.schema.createTable('Biotype', t => {
        t.increments('id').primary();
        t.string('description').notNullable();
        t.date('created_at').notNullable().defaultTo(knex.fn.now());
    });

    const bone_density = () => knex.schema.createTable('BoneDensity', t => {
        t.increments('id').primary();
        t.string('description').notNullable();
        t.date('created_at').notNullable().defaultTo(knex.fn.now());
    });

    const role = () => knex.schema.createTable('Role', t => {
        t.increments('id').primary();
        t.string('description').notNullable().unique();
        t.date('created_at').notNullable().defaultTo(knex.fn.now());
    });

    const customerSchema = () => knex.schema.createTable("Customer", t => {
        t.increments('id').primary();
        t.integer('biotype_id').notNullable().references('Biotype.id');
        t.integer('bone_density_id').notNullable().references('BoneDensity.id');
        t.string('bone_density').notNullable();
        t.string('name').notNullable();
        t.string('last_name').notNullable();
        t.string('email').unique().nullable();
        t.date('birthday').notNullable();
        t.timestamps(true);
    });

    const userSchema = () => knex.schema.createTable("User", t => {
        t.increments('id').primary();
        t.integer('rol_id').notNullable().references('Role.id');
        t.string('name').notNullable();
        t.string('last_name').notNullable();
        t.string('username').unique().nullable();
        t.string('password').unique().notNullable();
        t.timestamps(true);
    });

    return biotypes()
        .then(bone_density)
        .then(role)
        .then(customerSchema)
        .then(userSchema);
};

exports.down = function(knex) {
    const biotypes = () => knex.schema.dropTable('Biotype');
    const bone_density = () => knex.schema.dropTable('BoneDensity');
    const role = () => knex.schema.dropTable('Role');
    const customerSchema = () => knex.schema.dropTable('Customer');
    const userSchema = () => knex.schema.dropTable('User');

    return userSchema()
        .then(customerSchema)
        .then(biotypes)
        .then(bone_density)
        .then(role);
};
