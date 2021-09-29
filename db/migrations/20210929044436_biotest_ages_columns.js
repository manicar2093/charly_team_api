
exports.up = function(knex) {
    return knex.schema.alterTable('Biotest', t => {
        t.integer('chronological_age').notNullable();
        t.integer('age_on_test').notNullable();
    });
};

exports.down = function(knex) {
    return knex.schema.alterTable('Biotest', t => {
        t.dropColumn('chronological_age');
        t.dropColumn('age_on_test');
    });
};
