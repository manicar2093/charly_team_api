
exports.up = function(knex) {
    return knex.schema.alterTable('Biotest', t => {
        t.decimal('height').notNullable().alter();
    })
};

exports.down = function(knex) {
    return knex.schema.alterTable('Biotest', t => {
        t.integer('height').notNullable().alter();
    })
};
