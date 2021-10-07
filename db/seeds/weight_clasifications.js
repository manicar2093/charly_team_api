
exports.seed = function(knex) {
      return knex('WeightClasifications').insert([
        {id: 1, description: 'UNDERWEIGHT'},
        {id: 2, description: 'NORMAL_WEIGHT'},
        {id: 3, description: 'OVERWEIGHT'},
        {id: 4, description: 'OBESITY_1'},
        {id: 5, description: 'OBESITY_2'},
        {id: 6, description: 'OBESITY_3'},
      ]);
};
