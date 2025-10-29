import { CreationOptional, InferAttributes, InferCreationAttributes, Model } from "sequelize";
import sequelize from "./sequelize";

export enum RoomType {
    SINGLE = 'Single',
    DOUBLE = 'Double',
    FAMILY = 'Family',
    DELUXE = 'Deluxe',
    SUITE = 'Suite'
}

class RoomCategory extends Model<InferAttributes<RoomCategory>, InferCreationAttributes<RoomCategory>> {
    declare id: CreationOptional<number>;
    declare roomType: RoomType;
    declare price: number;
    declare hotelId: number;
    declare roomCount: number;
    declare createdAt: CreationOptional<Date>;
    declare updatedAt: CreationOptional<Date>;
    declare deletedAt: CreationOptional<Date | null>;
}

RoomCategory.init({
    id: {
        type: "INTEGER",
        autoIncrement: true,
        primaryKey: true
    },
    roomType: {
        type: 'ENUM',
        values: [...Object.values(RoomType)]
    },
    price: {
        type: "FLOAT",
        allowNull: false
    },
    hotelId: {
        type: "INTEGER",
        allowNull: false,
        references: {
            model: 'hotels',
            key: 'id'
        }
    },
    roomCount: {
        type: "INTEGER",
        allowNull: false
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
    tableName: "room_categories",
    sequelize: sequelize,
    underscored: true, // createdAt -> created_at
    timestamps: true // createdAt, updatedAt
});

export default RoomCategory;