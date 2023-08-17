FROM golang:latest

ARG USER_ID
ARG USERNAME
ARG GROUP_ID

RUN apt update && apt upgrade -y && \
    apt install -y git \
    make openssh-client

WORKDIR /client

RUN addgroup --gid ${USER_ID} ${USERNAME}
RUN adduser -disabled-password --gecos '' -uid ${USER_ID} --gid ${GROUP_ID} ${USERNAME}
USER ${USERNAME}

CMD ["go", "run", "cmd/client/main.go"]