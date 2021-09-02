
exports.seed = function(knex) {
  return knex('Gender').insert([
    {id: 1, description: 'MASCULINE'},
    {id: 2, description: 'FEMININE'},
  ]);
};
