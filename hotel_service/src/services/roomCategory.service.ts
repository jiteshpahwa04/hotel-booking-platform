import { CreateRoomCategoryDto } from "../dtos/roomCategory.dto";
import { HotelRepository } from "../repositories/hotel.repository";
import RoomCategoryRepository from "../repositories/roomCategory.repository";
import { NotFoundError } from "../utils/errors/app.error";

const roomCategoryRepository = new RoomCategoryRepository();
const hotelRepository = new HotelRepository();
export const createRoomCategoryService = async (roomCategoryData: CreateRoomCategoryDto) => {
    const roomCategory = await roomCategoryRepository.create(roomCategoryData);
    return roomCategory;
}

export const getRoomCategoryByIdService = async (id: number) => {
    const roomCategory = await roomCategoryRepository.findById(id);
    return roomCategory;
}

export const getAllRoomCategoriesByHotelIdService = async (hotelId: number) => {
    const hotel = await hotelRepository.findById(hotelId);
    if (!hotel) {
        throw new NotFoundError(`Hotel with id ${hotelId} not found`);
    }

    const roomCategories = await roomCategoryRepository.findAllByHotelId(hotelId);
    return roomCategories;
}

export const deleteRoomCategoryByIdService = async (id: number) => {
    const roomCategory = await roomCategoryRepository.findById(id);
    if (!roomCategory) {
        throw new NotFoundError(`Room category with id ${id} not found`);
    }

    await roomCategoryRepository.deleteById({ id });
    return true;
}