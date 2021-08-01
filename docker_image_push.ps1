docker build . -f ./server_Dockerfile -t pic-it/server:latest
docker tag pic-it/server:latest ericmarcantonio/pic-it-server:latest
docker push ericmarcantonio/pic-it-server:latest