
exports.up = function(knex) {
    const higherMuscleDensity = () => knex.schema.createTable('HigherMuscleDensity', t => {
        t.increments('id').primary()
        t.decimal('neck');
        t.decimal('shoulders');
        t.decimal('back');
        t.decimal('chest');
        t.decimal('back_chest');
        t.decimal('right_relaxed_bicep');
        t.decimal('right_contracted_bicep');
        t.decimal('left_relaxed_bicep');
        t.decimal('left_contracted_bicep');
        t.decimal('right_forearm');
        t.decimal('left_forearm');
        t.decimal('wrists');
        t.decimal('high_abdomen');
        t.decimal('lower_abdomen');
    });

    const lowerMuscleDensity = () => knex.schema.createTable('LowerMuscleDensity', t => {
        t.increments('id').primary()
        t.decimal('hips');
        t.decimal('right_leg');
        t.decimal('left_leg');
        t.decimal('right_calf');
        t.decimal('left_calf');
    });

    const skinFolds = () => knex.schema.createTable('SkinFolds', t => {
        t.increments('id').primary()
        t.integer('subscapular');
        t.integer('suprailiac');
        t.integer('bicipital');
        t.integer('tricipital');
    });

    return higherMuscleDensity()
        .then(lowerMuscleDensity)
        .then(skinFolds);

};

exports.down = function(knex) {
    const higherMuscleDensity = () => knex.schema.dropTable('HigherMuscleDensity');
    const lowerMuscleDensity = () => knex.schema.dropTable('LowerMuscleDensity');
    const skinFolds = () => knex.schema.dropTable('SkinFolds');
    return skinFolds()
        .then(lowerMuscleDensity)
        .then(higherMuscleDensity);
};
