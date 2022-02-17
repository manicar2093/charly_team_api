
exports.seed = function(knex) {
      return knex('Role').insert([
        {id: 1, description: 'ADMIN'},
        {id: 2, description: 'COACH'},
        {id: 3, description: 'CUSTOMER'},
      ]);
};
