import { RoomGenerationJob } from "../dtos/roomGeneration.dto";
import { roomGenerationQueue } from "../queues/roomGeneration.queue";

export const ROOM_GENERATION_PAYLOAD = "payload:room-generation";

export const addRoomGenerationToQueue = async (payload: RoomGenerationJob) => {
    await roomGenerationQueue.add(
        ROOM_GENERATION_PAYLOAD,
        payload, {
        attempts: 3,
        backoff: {
            type: 'exponential',
            delay: 1000
        }
    });
    console.log(`Room added to queue: ${JSON.stringify(payload)}`);
}