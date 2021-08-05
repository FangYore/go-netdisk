# gofi会在工作目录生成gofi.db文件用于存储gofi相关配置,可以将/app目录映射到host中以便持久化
# 由于gofi的默认文件仓库为工作目录下的storage文件夹,所以可以将/app/storage映射到host中,以便初始化Gofi时无需再指定storage path
FROM alpine:3.14
COPY ./output/gofi-linux-amd64 /usr/local/bin/gofi
WORKDIR /app
CMD [ "gofi","-p=8080","-ip=127.0.0.1" ]
EXPOSE 8080