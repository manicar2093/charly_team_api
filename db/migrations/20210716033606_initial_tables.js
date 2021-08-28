
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

    const userSchema = () => knex.schema.createTable("User", t => {
        t.increments('id').primary();
        t.integer('biotype_id').nullable().references('Biotype.id');
        t.integer('bone_density_id').nullable().references('BoneDensity.id');
        t.integer('role_id').notNullable().references('Role.id');
        t.string('name').notNullable();
        t.string('last_name').notNullable();
        t.string('email').unique().notNullable();
        t.date('birthday').nullable();
        t.bool('is_created').defaultTo(false);
        t.timestamps(true);
    });

    return biotypes()
        .then(bone_density)
        .then(role)
        .then(userSchema);
};

exports.down = function(knex) {
    const biotypes = () => knex.schema.dropTable('Biotype');
    const bone_density = () => knex.schema.dropTable('BoneDensity');
    const role = () => knex.schema.dropTable('Role');
    const userSchema = () => knex.schema.dropTable('User');

    return userSchema()
        .then(biotypes)
        .then(bone_density)
        .then(role);
};
