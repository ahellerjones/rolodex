FROM python:3.11-slim
WORKDIR /code
COPY ./requirements.txt /code/requirements.txt
RUN pip install -r /code/requirements.txt
COPY . /code
CMD ["uvicorn", "slick_backend.main:app", "--host", "0.0.0.0", "--port", "80"]


