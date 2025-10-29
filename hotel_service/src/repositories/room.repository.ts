import { CreationAttributes, Op } from "sequelize";
import Room from "../db/models/room";
import BaseRepository from "./base.repository";

class RoomRepository extends BaseRepository<Room> {
    constructor() {
        super(Room);
    }

    async findByRoomCategoryIdAndDateRange(roomCategoryId: number, startDate: Date, endDate: Date) {
        return this.model.findAll({
            where: {
                roomCategoryId,
                dateOfAvailability: {
                    [Op.between]: [startDate, endDate]
                },
                deletedAt: null
            }
        });
    }

    async bulkCreate(rooms: CreationAttributes<Room>[]) {
        return this.model.bulkCreate(rooms);
    }
}

export default RoomRepository;