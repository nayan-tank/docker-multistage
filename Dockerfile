
FROM ubuntu AS build

RUN apt-get update && apt-get install golang -y

ENV GO111MODULE=off

COPY . . 

RUN CGO_ENABLED=0 go build -o ./app .


#########################
#	Multistage 	#
#########################

FROM scratch

# copy the complied binary from build stage

COPY --from=build /app /app

# set entrypoint to the container to run the binary 
ENTRYPOINT ["/app"]



