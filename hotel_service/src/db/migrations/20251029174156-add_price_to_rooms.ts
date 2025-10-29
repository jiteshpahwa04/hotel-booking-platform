import { QueryInterface } from "sequelize";

module.exports = {
  async up (queryInterface: QueryInterface) {
    await queryInterface.sequelize.query(
      `ALTER TABLE rooms ADD COLUMN price FLOAT NOT NULL DEFAULT 0;`
    );
  },
  async down (queryInterface: QueryInterface) {
    await queryInterface.sequelize.query(
      `ALTER TABLE rooms DROP COLUMN price;`
    );
  }
}