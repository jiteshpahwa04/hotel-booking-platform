import { Worker } from "bullmq";
import { NotificationDto } from "../dtos/notification.dto";
import { MAILER_QUEUE } from "../queues/mailer.queue";
import { getRedisConnectionObject } from "../config/redis.config";
import { MAILER_PAYLOAD } from "../producers/email.producer";
import { renderMailTemplate } from "../templates/templates.handler";
import { sendEmail } from "../services/mailer.service";

export const setupMailerWorker = () => {

    const emailProcessor = new Worker<NotificationDto>(
        MAILER_QUEUE, // name of the queue
        async(job)=>{
            if(job.name !== MAILER_PAYLOAD) {
                throw new Error("Invalid job name");
            }
    
            // call the service layer here
            const payload = job.data;
            const emailContent = await renderMailTemplate(payload.templateId, payload.params);

            await sendEmail(payload.to, payload.subject, emailContent);
    
        }, // Process function
        {
            connection: getRedisConnectionObject()
        }
    )
    
    emailProcessor.on("failed", () => {
        console.error("Email processing failed");
    });
    
    emailProcessor.on("completed", () => {
        console.log("Email processing completed successfully!");
    });
}
