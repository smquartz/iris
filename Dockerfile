FROM alpine:3.2
ADD iris-srv /iris-srv
ENTRYPOINT [ "/iris-srv" ]
