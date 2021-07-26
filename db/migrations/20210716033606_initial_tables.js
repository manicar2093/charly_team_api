
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
        t.integer('role_id').notNullable().references('Role.id').defaultTo(2);
        t.string('name').notNullable();
        t.string('last_name').notNullable();
        t.string('email').unique().nullable();
        t.string('password').notNullable();
        t.date('birthday').notNullable();
        t.timestamps(true);
    });

    const userSchema = () => knex.schema.createTable("User", t => {
        t.increments('id').primary();
        t.integer('role_id').notNullable().references('Role.id');
        t.string('name').notNullable();
        t.string('last_name').notNullable();
        t.string('username').unique().nullable();
        t.string('password').notNullable();
        t.timestamps(true);
    });

    return biotypes()
        .then(bone_density)
        .then(role)
        .then(userSchema)
        .then(customerSchema);
};

exports.down = function(knex) {
    const biotypes = () => knex.schema.dropTable('Biotype');
    const bone_density = () => knex.schema.dropTable('BoneDensity');
    const role = () => knex.schema.dropTable('Role');
    const customerSchema = () => knex.schema.dropTable('Customer');
    const userSchema = () => knex.schema.dropTable('User');

    return customerSchema()
        .then(userSchema)
        .then(biotypes)
        .then(bone_density)
        .then(role);
};
