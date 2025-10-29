import { RoomType } from "../db/models/roomCategory"

export type CreateRoomCategoryDto = {
    hotelId: number,
    roomType: RoomType,
    price: number,
    roomCount: number
}