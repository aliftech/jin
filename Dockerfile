FROM python:3.11.2-buster
COPY . /jin
WORKDIR /jin
RUN pip install -r requirements.txt
CMD [ "python main.py" ]