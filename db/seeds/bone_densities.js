
exports.seed = function(knex) {
  return knex('BoneDensity').del()
    .then(function () {
      return knex('BoneDensity').insert([
        {id: 1, description: 'Big'},
        {id: 2, description: 'Medium'},
        {id: 3, description: 'Small'},
      ]);
    });
};
