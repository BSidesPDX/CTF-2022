FROM python:3.10.6-alpine
WORKDIR /app
COPY index.html /app/
COPY assets/ /app/assets/
EXPOSE 8000
CMD ["python", "-m", "http.server"]
