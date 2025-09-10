import os
from fastapi import FastAPI, Depends, HTTPException
from dotenv import load_dotenv
from app.models import ProgramRequest, ProgramResponse
from app.agent import run_agent

load_dotenv()
app = FastAPI(title="Hypertrophy Program Agent API (OpenAI Agents SDK)")

def ensure_openai():
    if not os.getenv("OPENAI_API_KEY"):
        raise HTTPException(status_code=500, detail="OPENAI_API_KEY missing")
    return True

@app.post("/programs/generate", response_model=ProgramResponse)
async def generate_program(req: ProgramRequest, _=Depends(ensure_openai)):
    try:
        plan = await run_agent(req)  # await
        return plan
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
