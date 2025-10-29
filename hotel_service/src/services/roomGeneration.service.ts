import { CreationAttributes } from "sequelize";
import Room from "../db/models/room";
import RoomCategory from "../db/models/roomCategory";
import { RoomGenerationJob } from "../dtos/roomGeneration.dto";
import RoomRepository from "../repositories/room.repository";
import RoomCategoryRepository from "../repositories/roomCategory.repository";
import { BadRequestError, NotFoundError } from "../utils/errors/app.error";
import logger from "../config/logger.config";

const roomCategoryRepository = new RoomCategoryRepository();
const roomRepository = new RoomRepository();

export async function generateRooms(jobData: RoomGenerationJob) {
    logger.info("Generating rooms with the following data:", jobData);
    
    // Check if the category exists
    const roomCategory = await roomCategoryRepository.findById(jobData.roomCategoryId);
    if (!roomCategory) {
        throw new NotFoundError(`Room category with ID ${jobData.roomCategoryId} does not exist.`);
    }

    logger.info(`Room category found: ${JSON.stringify(roomCategory)}`);

    // Validate date range
    const startDate = new Date(jobData.startDate);
    const endDate = new Date(jobData.endDate);
    if (startDate >= endDate) {
        throw new BadRequestError("Start date must be before end date.");
    }
    if(startDate < new Date()) {
        throw new BadRequestError("Start date must be in the future.");
    }

    logger.info(`Date range is valid: ${startDate.toISOString()} to ${endDate.toISOString()}`);

    const totalDays = Math.ceil((endDate.getTime() - startDate.getTime()) / (1000 * 60 * 60 * 24)) + 1;
    logger.info(`Total days to process: ${totalDays}`);

    const batchSize = jobData.batchSize || 100;
    const currentDate = new Date(startDate);

    let totalRoomsCreated = 0;
    let totalDatesProcessed = 0;
    while (currentDate <= endDate) {
        const batchEndDate = new Date(currentDate);
        batchEndDate.setDate(batchEndDate.getDate() + batchSize - 1); // 20 to 23 if batchSize is 4
        if (batchEndDate > endDate) {
            batchEndDate.setTime(endDate.getTime());
        }

        logger.info(`Processing batch from ${currentDate.toISOString().split('T')[0]} to ${batchEndDate.toISOString().split('T')[0]}`);

        const { roomsCreated, daysProcessed } = await processDateBatch(roomCategory, currentDate, batchEndDate, jobData.priceOverride);
        logger.info(`Batch processed. Rooms created: ${roomsCreated}, Days processed: ${daysProcessed}`);

        currentDate.setTime(batchEndDate.getTime());
        currentDate.setDate(currentDate.getDate() + 1);
        totalRoomsCreated += roomsCreated;
        totalDatesProcessed += daysProcessed;
    }

    return {
        totalRoomsCreated,
        totalDatesProcessed
    };
}

export async function processDateBatch(roomCategory: RoomCategory, startDate: Date, endDate: Date, priceOverride?: number) {
    let roomsCreated = 0;
    let daysProcessed = 0;
    const roomsToCreate: CreationAttributes<Room>[] = []

    let currentDate = new Date(startDate);
    const existingRooms = await roomRepository.findByRoomCategoryIdAndDateRange(roomCategory.id, currentDate, endDate);
    const existingDateMap = new Map<string, boolean>();
    existingRooms.forEach(room => {
        existingDateMap.set(room.dateOfAvailability.toISOString().split('T')[0], true);
    });

    while (currentDate <= endDate) {
        if (!existingDateMap.has(currentDate.toISOString().split('T')[0])) {
            roomsToCreate.push({
                hotelId: roomCategory.hotelId,
                roomCategoryId: roomCategory.id,
                dateOfAvailability: new Date(currentDate), // Since the currentDate is a date object, it will be same if not cloned
                price: priceOverride || roomCategory.price
            })
        }
        currentDate.setDate(currentDate.getDate() + 1);
        daysProcessed++;
    }

    if (roomsToCreate.length > 0) {
        const createdRooms = await roomRepository.bulkCreate(roomsToCreate);
        roomsCreated += createdRooms.length;
    }

    return { roomsCreated, daysProcessed };
}