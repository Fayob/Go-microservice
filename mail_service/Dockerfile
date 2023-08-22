FROM alpine:latest

RUN mkdir /app

COPY mailerApp /app

COPY api/templates ./api/templates

CMD [ "/app/mailerApp" ]