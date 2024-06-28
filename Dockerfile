FROM python:3.9.19-alpine3.20

WORKDIR /jin

COPY . /jin

RUN pip install --no-cache-dir -r requirements.txt

CMD [ "python", "main.py" ]