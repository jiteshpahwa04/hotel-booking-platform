import Redis from 'ioredis';
import { serverConfig } from '.';
import logger from './logger.config';

export function connectToRedis() {
    try {
        let connection: Redis;
        const redisConfig = {
            port: serverConfig.REDIS_PORT,
            host: serverConfig.REDIS_HOST,
            maxRetriesPerRequest: null
        };

        // singleton object
        return ()=>{
            if(!connection) {
                connection = new Redis(redisConfig);
            } 
            
            return connection;
        }

    } catch(err) {
        logger.error("Could not connect to redis", err);
        console.error("Could not connect to redis: ", err);
        throw err;
    }
}

export const getRedisConnectionObject = connectToRedis(); // given the value of function => internal function => it will get the already having value of connection object, and not initialize a new object