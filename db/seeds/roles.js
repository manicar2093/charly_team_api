
exports.seed = function(knex) {
      return knex('Role').insert([
        {id: 1, description: 'ADMIN'},
        {id: 2, description: 'CUSTOMER'},
      ]);
};
