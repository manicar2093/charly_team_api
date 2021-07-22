
exports.seed = function(knex) {
  // Deletes ALL existing entries
  return knex('WeightClasifications').del()
    .then(function () {
      // Inserts seed entries
      return knex('WeightClasifications').insert([
        {id: 1, description: 'Underweight'},
        {id: 2, description: 'Normal weight'},
        {id: 3, description: 'Overweight'},
        {id: 4, description: 'Obesity 1'},
        {id: 5, description: 'Obesity 2'},
        {id: 6, description: 'Obesity 3'},
      ]);
    });
};
