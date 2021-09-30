
exports.up = function(knex) {
    return knex.schema.alterTable('Biotest', t => {
        t.integer('chronological_age').notNullable();
        t.integer('corporal_age').notNullable();
    });
};

exports.down = function(knex) {
    return knex.schema.alterTable('Biotest', t => {
        t.dropColumn('chronological_age');
        t.dropColumn('corporal_age');
    });
};
