FROM archlinux:base-devel

RUN pacman -Syu --noconfirm --needed \
    bash \
    git \
    make \
    go \
    ca-certificates

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV HOME=/tmp
ENV CGO_ENABLED=0
ENV TERM=xterm-256color

CMD ["bash"]