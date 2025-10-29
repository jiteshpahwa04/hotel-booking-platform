import {z} from "zod";

export const createBookingSchema = z.object({
    userId: z.number({message: "User ID must be present"}),
    hotelId: z.number({message: "Hotel ID must be present"}),
    totalGuests: z.number({message: "Toal guests must be present"}).min(1, {message: "Total guests must be atlease 1"}),
    bookingAmount: z.number({message: "Booking amount must be present"}).min(0, {message: "Booking amount cannot be less than 0"})
})