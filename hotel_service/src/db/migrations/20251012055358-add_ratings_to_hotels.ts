import { DataTypes, QueryInterface } from "sequelize";

module.exports = {
  async up (queryInterface: QueryInterface) {
    await queryInterface.addColumn('hotels', 'rating', {
      type: DataTypes.FLOAT,
      allowNull: true,
      defaultValue: null
    });

    await queryInterface.addColumn('hotels', 'rating_count', {
      type: DataTypes.INTEGER,
      allowNull: true,
      defaultValue: null
    });
  },
  async down (queryInterface: QueryInterface) {
    await queryInterface.removeColumn('hotels', 'rating');
    await queryInterface.removeColumn('hotels', 'rating_count');
  }
}