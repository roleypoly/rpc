docker pull roleypoly/rpc-builder
docker run -it --rm --mount type=bind,src="$(pwd)",dst=/src --mount type=volume,dst=/src/node_modules roleypoly/rpc-builder

