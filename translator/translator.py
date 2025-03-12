from fastapi import FastAPI
from pydantic import BaseModel
from transformers import M2M100ForConditionalGeneration, M2M100Tokenizer
from langdetect import detect

app = FastAPI()

class TranslateRequest(BaseModel):
    text: str
    target_lang: str

class TranslateResponse(BaseModel):
    translated_text: str
    detected_lang: str

model_name = "facebook/m2m100_418M"
tokenizer = M2M100Tokenizer.from_pretrained(model_name)
model = M2M100ForConditionalGeneration.from_pretrained(model_name)

@app.post("/translate")
def translate(req: TranslateRequest):
    source_lang = detect(req.text)
    tokenizer.src_lang = source_lang

    encoded = tokenizer(req.text, return_tensors="pt")
    generated_tokens = model.generate(**encoded, forced_bos_token_id=tokenizer.get_lang_id(req.target_lang))
    translation = tokenizer.batch_decode(generated_tokens, skip_special_tokens=True)[0]

    return {
        "translated_text": translation,
        "detected_source_lang": source_lang
    }
