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

    async findByRoomCategoryIdAndDate(roomCategoryId: number, date: Date) {
        return this.model.findOne({
            where: {
                roomCategoryId,
                dateOfAvailability: date,
                deletedAt: null
            }
        });
    }

    async bulkCreate(rooms: CreationAttributes<Room>[]) {
        return this.model.bulkCreate(rooms);
    }

    async findLatestDateByRoomCategoryId(roomCategoryId: number) {
        const latestRoom = await this.model.findOne({
            where: {
                roomCategoryId,
                deletedAt: null
            },
            attributes: ['dateOfAvailability'],
            order: [['dateOfAvailability', 'DESC']]
        });
        return latestRoom ? latestRoom.dateOfAvailability : null;
    }

    async findLatestDatesForAllCategories() {
        const latestRooms = await this.model.findAll({
            where: {
                deletedAt: null
            },
            attributes: [
                'roomCategoryId',
                [this.model.sequelize!.fn('MAX', this.model.sequelize!.col('date_of_availability')), 'latestDate']
            ],
            group: ['roomCategoryId'],
            raw: true
        });
        return latestRooms.map((room: any) => ({
            roomCategoryId: room.roomCategoryId,
            latestDate: room.latestDate
        }));
    }
}

export default RoomRepository;