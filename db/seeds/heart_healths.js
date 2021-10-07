
exports.seed = function(knex) {
    return knex('HeartHealths').insert([
        {id: 1, description: 'GOOD'},
        {id: 2, description: 'BAD'},
        {id: 3, description: 'EXCELLENT'},
    ]);
};
