FROM python:3.11-slim AS builder

WORKDIR /app
COPY requirements.txt .

RUN pip install --upgrade pip && \
    pip install --no-cache-dir -r requirements.txt

RUN python -c "from transformers import M2M100ForConditionalGeneration, M2M100Tokenizer; model_name='facebook/m2m100_418M'; M2M100ForConditionalGeneration.from_pretrained(model_name); M2M100Tokenizer.from_pretrained(model_name)"

FROM python:3.11-slim

COPY --from=0 /root/.cache /root/.cache

WORKDIR /app
COPY . .

RUN pip install --upgrade pip && \
    pip install --no-cache-dir -r requirements.txt

EXPOSE 8000

CMD ["uvicorn", "translator:app", "--host", "0.0.0.0", "--port", "8000"]
