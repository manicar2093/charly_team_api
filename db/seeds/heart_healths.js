
exports.seed = function(knex) {
    return knex('HeartHealths').insert([
        {id: 1, description: 'Good'},
        {id: 2, description: 'Bad'},
        {id: 3, description: 'Excellent'},
    ]);
};
