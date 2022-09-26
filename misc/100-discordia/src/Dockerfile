FROM python:3.10.7-alpine3.15
WORKDIR /app/
COPY requirements.txt bot.py ./
RUN pip install -r requirements.txt
CMD python3 bot.py
