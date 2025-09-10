import json, uuid

from agents import Agent, Runner, ModelSettings
from app.models import ProgramRequest, ProgramResponse
from app.tools import tool_get_exercise_catalog, tool_validate_constraints
from app.prompts import SYSTEM_PROMPT

# By default the SDK uses OpenAI Responses API; set OPENAI_DEFAULT_MODEL or pass model="gpt-4.1"
# See: Models docs (choose model, temperature, etc.). :contentReference[oaicite:2]{index=2}
agent = Agent(
    name="HypertrophyPlanner",
    instructions=SYSTEM_PROMPT,
    tools=[tool_get_exercise_catalog, tool_validate_constraints],
    output_type=ProgramResponse,              # <- structured output (validated)
    model_settings=ModelSettings(temperature=0.2),
)

def build_input(req: ProgramRequest) -> str:
    """What we hand to the agent as the 'input' message."""
    return json.dumps({
        "task": "Generate hypertrophy program",
        "req": req.model_dump(),
        "notes": "Return only the final JSON; no prose."
    })

async def run_agent(req: ProgramRequest) -> ProgramResponse:
    result = await Runner.run(
        agent,
        input=build_input(req),
        max_turns=req.run_config.max_turns,
    )
    output: ProgramResponse = result.final_output
    # backfill id si absent
    if getattr(output, "program_id", None) in (None, ""):
        output.program_id = f"prg_{uuid.uuid4().hex[:10]}"
    try:
        output.trace_id = getattr(result, "trace_id", None) or None
    except Exception:
        pass
    return output