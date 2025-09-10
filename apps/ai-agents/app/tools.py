from typing import List
from agents import function_tool
from pydantic import BaseModel
from app.models import SessionPlan, Constraints

class ExerciseCatalogItem(BaseModel):
    name: str
    muscles: List[str]
    equipment: List[str]

# Small built-in demo DB (replace with your own source of truth later)
EXERCISE_DB = [
    {"name":"Barbell Bench Press","muscles":["chest","triceps","delts"],"equipment":["barbell"]},
    {"name":"Incline DB Press","muscles":["chest","delts"],"equipment":["dumbbell"]},
    {"name":"Weighted Pull-up","muscles":["back","biceps"],"equipment":["pullup_bar","dip_belt"]},
    {"name":"Barbell Row","muscles":["back","biceps"],"equipment":["barbell"]},
    {"name":"Seated Cable Row","muscles":["back"],"equipment":["cable"]},
    {"name":"Back Squat","muscles":["quads","glutes"],"equipment":["barbell"]},
    {"name":"Romanian Deadlift","muscles":["hamstrings","glutes"],"equipment":["barbell"]},
    {"name":"DB Lateral Raise","muscles":["delts"],"equipment":["dumbbell"]},
    {"name":"Cable Fly","muscles":["chest"],"equipment":["cable"]},
]

@function_tool
def tool_get_exercise_catalog(equipment: List[str] = [], muscles: List[str] = []) -> List[ExerciseCatalogItem]:
    """
    List available exercises filtered by equipment and target muscles.

    Args:
        equipment: A list of equipment names available to the user.
        muscles:   A list of primary muscles to prioritize.
    """
    eq, ms = set(equipment or []), set(muscles or [])
    out: List[ExerciseCatalogItem] = []
    for ex in EXERCISE_DB:
        if eq and not eq.intersection(ex["equipment"]):
            continue
        if ms and not ms.intersection(ex["muscles"]):
            continue
        out.append(ExerciseCatalogItem(**ex))
    return out

class ProgramDraft(BaseModel):
    sessions: List[SessionPlan] = []

class ValidationResult(BaseModel):
    ok_days: bool
    too_long_sessions: List[str]
    is_valid: bool

@function_tool
def tool_validate_constraints(program: ProgramDraft, constraints: Constraints) -> ValidationResult:
    """
    Validate that the built program fits user constraints.

    Args:
        program:     The candidate program JSON (sessions with exercises).
        constraints: User constraints (days/week, session duration, etc.)
    """
    max_days = int(constraints.available_days_per_week)
    session_limit = int(constraints.session_duration_min) + 15  # small margin
    ok_days = len(program.sessions or []) <= max_days

    def est_time(session: SessionPlan) -> int:
        sets = sum(ex.sets for ex in session.exercises)
        rests = sum(int(ex.rest_seconds) if ex.rest_seconds is not None else 90 for ex in session.exercises)
        return sets * 45 + rests  # ~45s per set + rests

    too_long = [s.title for s in (program.sessions or []) if est_time(s) > session_limit * 60]
    return ValidationResult(ok_days=ok_days, too_long_sessions=too_long, is_valid=(ok_days and not too_long))
