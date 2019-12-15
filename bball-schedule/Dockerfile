FROM python:3.8.0-alpine3.10

COPY ./cronjobs /etc/crontabs/root
COPY ./src/ /usr/share/bball_reminder/

RUN pip3 install requests boto3

# start crond with log level 8 in foreground, output to stderr
CMD ["crond", "-f", "-d", "8"]