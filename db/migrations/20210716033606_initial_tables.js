
exports.up = function(knex) {

    const biotypes = () => knex.schema.createTable('Biotype', t => {
        t.increments('id').primary();
        t.string('description').notNullable();
    });

    const userSchema = () => knex.schema.createTable("User", t => {
        t.increments('id').primary();
        t.integer('biotype_id').notNullable().references('Biotype.id');
        t.string('bone_density').notNullable();
        t.string('name').notNullable();
        t.string('last_name').notNullable();
        t.string('email').unique();
        t.string('username').unique().nullable();
        t.string('password').unique().notNullable();
        t.date('birthday').notNullable();
        t.timestamps(true);
    });

    return biotypes()
        .then(userSchema);
};

exports.down = function(knex) {
    const userSchema = () => knex.schema.dropTable('User');
    const biotypes = () => knex.schema.dropTable('Biotype');

    return userSchema()
        .then(biotypes);
};
