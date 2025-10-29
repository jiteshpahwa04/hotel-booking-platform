import { createHotelDto } from "../dtos/hotel.dto";
import { HotelRepository } from "../repositories/hotel.repository";

const hotelRepository = new HotelRepository();

export async function createHotelService(hotelData : createHotelDto) {
    const hotel = await hotelRepository.create(hotelData);
    return hotel;
}

export async function getHotelByIdService(id : number) {
    const hotel = await hotelRepository.findById(id);
    return hotel;
}

export async function getAllHotelsService() {
    return await hotelRepository.findAll();
}

export async function softDeleteHotelService(id:number) {
    return await hotelRepository.softDelete(id);
}