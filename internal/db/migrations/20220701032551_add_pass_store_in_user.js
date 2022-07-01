
exports.up = function(knex) {
    return knex.schema.alterTable("User", t => {
        t.string('password').notNullable();
    })
};

exports.down = function(knex) {
    return knex.schema.alterTable("User", t => {
        t.dropColumn('password');
    })
};
