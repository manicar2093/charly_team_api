
exports.seed = function(knex) {
      return knex('BoneDensity').insert([
        {id: 1, description: 'BIG'},
        {id: 2, description: 'MEDIUM'},
        {id: 3, description: 'SMALL'},
      ]);
};
