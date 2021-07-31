
exports.up = function(knex) {

    const weight_clasification = () => knex.schema.createTable('WeightClasifications', t => {
        t.increments('id').primary();
        t.string('description').notNullable();
        t.date('created_at').notNullable().defaultTo(knex.fn.now());
    });

    const heart_health = () => knex.schema.createTable('HeartHealths', t => {
        t.increments('id').primary()
        t.string('description').notNullable();
        t.date('created_at').notNullable().defaultTo(knex.fn.now());
    });

    const biotest = () => knex.schema.createTable('Biotest', t => {
        t.increments('id').primary();
        t.integer('higher_muscle_density_id').notNullable().references('HigherMuscleDensity.id');
        t.integer('lower_muscle_density_id').notNullable().references('LowerMuscleDensity.id');
        t.integer('skin_folds_id').notNullable().references('SkinFolds.id');
        t.integer('weight_clasification_id').notNullable().references('WeightClasifications.id');
        t.integer('heart_health_id').nullable().references('HeartHealths.id');
        t.integer('customer_id').notNullable().references('Customer.id');
        t.decimal('weight').notNullable();
        t.integer('height').notNullable();
        t.decimal('body_fat_percentage').notNullable();
        t.decimal('total_body_water').notNullable();
        t.decimal('body_mass_index').notNullable();
        t.decimal('oxygen_saturation_in_blood').notNullable();
        t.decimal('glucose').nullable();
        t.decimal('resting_heart_rate').nullable();
        t.decimal('maximum_heart_rate').nullable();
        t.string('observations').nullable();
        t.string('recommendations').nullable();
        t.string('front_picture').nullable();
        t.string('back_picture').nullable();
        t.string('right_side_picture').nullable();
        t.string('left_side_picture').nullable();
        t.date('next_evaluation').nullable();
        t.date('created_at').notNullable().defaultTo(knex.fn.now());
    });

    return weight_clasification()
        .then(heart_health)
        .then(biotest);
};

exports.down = function(knex) {
    const weight_clasification = () => knex.schema.dropTable('WeightClasifications');
    const biotest = () => knex.schema.dropTable('Biotest');
    const heart_health = () => knex.schema.dropTable('HeartHealths');


    return biotest()
        .then(weight_clasification)
        .then(heart_health);
};
