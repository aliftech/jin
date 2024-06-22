FROM python:3.11-slim

WORKDIR /jin

COPY . /jin

RUN pip install --no-cache-dir -r requirements.txt

CMD [ "python", "main.py" ]