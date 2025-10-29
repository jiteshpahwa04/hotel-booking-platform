import { Worker } from "bullmq";
import { RoomGenerationJob } from "../dtos/roomGeneration.dto";
import { ROOM_GENERATION_QUEUE } from "../queues/roomGeneration.queue";
import { ROOM_GENERATION_PAYLOAD } from "../producers/roomGeneration.producer";
import { getRedisConnectionObject } from "../config/redis.config";
import { generateRooms } from "../services/roomGeneration.service";
import logger from "../config/logger.config";

export const setupRoomGenerationWorker = () => {

    const roomGenerationProcessor = new Worker<RoomGenerationJob>(
        ROOM_GENERATION_QUEUE, // name of the queue
        async(job)=>{
            if(job.name !== ROOM_GENERATION_PAYLOAD) {
                throw new Error("Invalid job name");
            }
    
            // call the service layer here
            const payload = job.data;
            logger.info(`Processing room generation for payload: ${JSON.stringify(payload)}`);

            await generateRooms(payload);

            logger.info(`Completed room generation for payload: ${JSON.stringify(payload)}`);
    
        }, // Process function
        {
            connection: getRedisConnectionObject()
        }
    )
    
    roomGenerationProcessor.on("failed", () => {
        logger.error("Room generation processing failed");
    });

    roomGenerationProcessor.on("completed", () => {
        logger.info("Room generation processing completed successfully!");
    });
}
