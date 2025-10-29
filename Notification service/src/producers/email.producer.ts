import { NotificationDto } from "../dtos/notification.dto";
import { mailerQueue } from "../queues/mailer.queue";

export const MAILER_PAYLOAD = "payload:mail";

export const addEmailToQueue = async (payload: NotificationDto) => {
    await mailerQueue.add(
        MAILER_PAYLOAD, 
        payload, {
        attempts: 3,
        backoff: {
            type: 'exponential',
            delay: 1000
        }
    });
    console.log(`Email added to queue: ${JSON.stringify(payload)}`);
}