FROM node:20

WORKDIR /app

COPY . .
RUN npm i

ENV SCHEDULING_SYSTEM=${SCHEDULING_SYSTEM}
ENV NOTIFICATION_SYSTEM=${NOTIFICATION_SYSTEM}

RUN echo "env: ${NOTIFICATION_SYSTEM}"

EXPOSE 3000

CMD npm run dev
