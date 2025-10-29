import logger from "../config/logger.config";
import Hotel from "../db/models/hotel";
import { NotFoundError } from "../utils/errors/app.error";
import BaseRepository from "./base.repository";

// export async function createHotel(hotelData : createHotelDto) {
//     const hotel = await Hotel.create({
//         name: hotelData.name,
//         address: hotelData.address,
//         location: hotelData.location,
//         rating: hotelData.rating,
//         ratingCount: hotelData.ratingCount
//     });

//     logger.info(`Hotel created: ${hotel.id}`);

//     return hotel;
// };

// export async function getHotelById(id : number) {
//     const hotel = await Hotel.findByPk(id);

//     if(!hotel){
//         logger.error(`Hotel not found: ${id}`);
//         throw new NotFoundError(`Hotel with id ${id} not found`);
//     }

//     logger.info(`Hotel found: ${id}`);

//     return hotel;
// }

// export async function getAllHotels() {
//     const hotels = await Hotel.findAll({
//         where: {
//             deletedAt: null
//         }
//     });
//     return hotels;
// }

// export async function softDeleteHotel(id:number) {
//     const hotel = await Hotel.findByPk(id);

//     if(!hotel){
//         logger.error(`Hotel not found: ${id}`);
//         throw new NotFoundError(`Hotel with id ${id} not found`);
//     }

//     hotel.deletedAt = new Date();
//     await hotel.save();
//     logger.info(`Hotel is soft deleted, id: ${hotel.id}`);
//     return hotel;
// }

export class HotelRepository extends BaseRepository<Hotel> {
    constructor() {
        super(Hotel);
    }

    // Override findAll to exclude soft-deleted records
    async findAll(): Promise<Hotel[]> {
        const hotels = await this.model.findAll({
            where: {
                deletedAt: null
            }
        });

        if(!hotels) {
            logger.error('No hotels found');
            throw new NotFoundError('No hotels found');
        }
        return hotels;
    }

    // Soft delete a hotel by ID
    async softDelete(id:number) {
        const hotel = await Hotel.findByPk(id);

        if(!hotel){
            logger.error(`Hotel not found: ${id}`);
            throw new NotFoundError(`Hotel with id ${id} not found`);
        }

        hotel.deletedAt = new Date();
        await hotel.save();
        logger.info(`Hotel is soft deleted, id: ${hotel.id}`);
        return hotel;
    }
}