
exports.seed = function(knex) {
    return knex('Biotype').insert([
      {id: 1, description: 'Ectomorph'},
      {id: 2, description: 'Mesomorph'},
      {id: 3, description: 'Endomorph'},
    ]);
};
