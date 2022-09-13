FROM python:3.8-slim-buster

WORKDIR /app
COPY requirements.txt requirements.txt
RUN pip3 install -r requirements.txt && /usr/local/bin/python -m pip install --upgrade pip
COPY wsgi.py .
ENV FLASK_APP wsgi 
CMD [ "python3", "-m" , "flask", "run", "--host=0.0.0.0"]
