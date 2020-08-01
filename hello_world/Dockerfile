FROM python:3.6-alpine

WORKDIR /home/hello_world
COPY app app
COPY migrations migrations
COPY main.py config.py startup.sh ./

# Set up python environment
COPY requirements.txt requirements.txt
RUN python -m venv venv
RUN venv/bin/pip install -r requirements.txt
RUN venv/bin/pip install gunicorn

EXPOSE 5000
ENTRYPOINT [ "/bin/sh", "startup.sh" ]
