
exports.seed = function(knex) {
    return knex('Biotype').insert([
      {id: 1, description: 'ECTOMORPH'},
      {id: 2, description: 'MESOMORPH'},
      {id: 3, description: 'ENDOMORPH'},
    ]);
};
