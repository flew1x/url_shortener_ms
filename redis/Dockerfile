FROM redislabs/rebloom:2.6.12 as rebloom
FROM redis:7.2.4-alpine

ENV LD_LIBRARY_PATH /usr/lib/redis/modules

COPY --from=rebloom ${LD_LIBRARY_PATH}/*.so ${LD_LIBRARY_PATH}/

CMD ["--loadmodule", "/usr/lib/redis/modules/redisbloom.so"]