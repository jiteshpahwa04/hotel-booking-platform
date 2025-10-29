import IORedis, { Redis } from 'ioredis';
import Redlock from 'redlock';
import { serverConfig } from '.';

export function connectToRedis() {
    try {
        let connection: Redis;

        // singleton object
        return ()=>{
            if(!connection) {
                connection = new IORedis(serverConfig.REDIS_SERVER_URL);
            } 
            
            return connection;
        }

    } catch(err) {
        console.error("Could not connect to redis: ", err);
        throw err;
    }
}

export const getRedisConnectionObject = connectToRedis();

export const redlock = new Redlock([getRedisConnectionObject()], {
    driftFactor: 0.01,  // time in ms
    retryCount: 10,
    retryDelay: 200,  // time in ms
    retryJitter: 200  // time in ms
})