
exports.up = function(knex) {
    return knex.schema.alterTable('Biotest', t => {
        t.integer('chronological_age').notNullable().defaultTo(0);
        t.integer('corporal_age').notNullable().defaultTo(0);
    });
};

exports.down = function(knex) {
    return knex.schema.alterTable('Biotest', t => {
        t.dropColumn('chronological_age');
        t.dropColumn('corporal_age');
    });
};
