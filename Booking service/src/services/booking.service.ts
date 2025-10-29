import { CreateBookingDTO } from "../dto/booking.dto";
import {
  confirmBooking,
  createBooking,
  createIdempotencyKey,
  finalizeIdempotenceKey,
  getIdempotencyKeyWithLock,
} from "../repositories/booking.repository";
import { BadRequestError, InternalServerError, NotFoundError } from "../utils/errors/app.error";
import { generateIdempotencyKey } from "../utils/generateIdempotency";
import prismaClient from "../prisma/client";
import { redlock } from "../config/redis.config";
import { serverConfig } from "../config";

export async function createBookingService(createBookingDTO: CreateBookingDTO) {
  const bookingResource = `hotel:${createBookingDTO.hotelId}`;
  const ttl = serverConfig.LOCK_TTL;

  console.log(`Acquiring lock for resource: ${bookingResource} with TTL: ${ttl}`);

  try {
    await redlock.acquire([bookingResource], ttl);
    const booking = await createBooking({
      userId: createBookingDTO.userId,
      hotelId: createBookingDTO.hotelId,
      totalGuests: createBookingDTO.totalGuests,
      bookingAmount: createBookingDTO.bookingAmount,
    });

    const idempotencyKey = generateIdempotencyKey();

    await createIdempotencyKey(idempotencyKey, booking.id);

    return {
      bookingId: booking.id,
      idempotencyKey,
    };
  } catch (err) {
    throw new InternalServerError("Not able to acquire lock");
  }
}

export async function confirmBookingService(idempotencyKey: string) {
  return await prismaClient.$transaction(async (tx) => {
    const idempotencyKeyData = await getIdempotencyKeyWithLock(
      tx,
      idempotencyKey
    );

    if (!idempotencyKeyData || !idempotencyKeyData.bookingId) {
      throw new NotFoundError("Idempotency key not found");
    }

    if (idempotencyKeyData.finalized) {
      throw new BadRequestError("Idempotency key already finalized");
    }

    const booking = await confirmBooking(tx, idempotencyKeyData.bookingId);
    await finalizeIdempotenceKey(idempotencyKey);

    return booking;
  });
}
