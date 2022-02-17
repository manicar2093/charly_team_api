
exports.up = function(knex) {
    return knex.schema.alterTable('User', t => {
        t.string('avatar_url').notNullable().defaultTo("")
    })
};

exports.down = function(knex) {
    return knex.schema.alterTable('User', t => {
        t.dropColumn('avatar_url')
    })
};
