
exports.seed = function(knex) {
    return knex('HeartHealths').insert([
        {id: 1, description: 'BAD'},
        {id: 2, description: 'GOOD'},
        {id: 3, description: 'EXCELLENT'},
    ]);
};
