import { CreationOptional, InferAttributes, InferCreationAttributes, Model } from "sequelize";
import sequelize from "./sequelize";

class Room extends Model<InferAttributes<Room>, InferCreationAttributes<Room>> {
    declare id: CreationOptional<number>;
    declare roomCategoryId: number;
    declare hotelId: number;
    declare dateOfAvailability: Date;
    declare price: number;
    declare bookingId?: number | null;
    declare createdAt: CreationOptional<Date>;
    declare updatedAt: CreationOptional<Date>;
    declare deletedAt: CreationOptional<Date | null>;
}

Room.init({
    id: {
        type: "INTEGER",
        autoIncrement: true,
        primaryKey: true
    },
    roomCategoryId: {
        type: "INTEGER",
        allowNull: false,
        references: {
            model: 'room_categories',
            key: 'id'
        }
    },
    hotelId: {
        type: "INTEGER",
        allowNull: false,
        references: {
            model: 'hotels',
            key: 'id'
        }
    },
    dateOfAvailability: {
        type: "DATE",
        allowNull: false
    },
    price: {
        type: "FLOAT",
        allowNull: false
    },
    bookingId: {
        type: "INTEGER",
        allowNull: true
    },
    createdAt: {
        type: "TIMESTAMP",
        defaultValue: new Date()
    },
    updatedAt: {
        type: "TIMESTAMP",
        defaultValue: new Date()
    },
    deletedAt: {
        type: "TIMESTAMP",
        allowNull: true
    }
}, {
    tableName: "rooms",
    sequelize: sequelize,
    underscored: true, // createdAt -> created_at
    timestamps: true // createdAt, updatedAt
});

export default Room;