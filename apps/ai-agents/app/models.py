from pydantic import BaseModel, Field
from typing import List, Literal, Optional

Experience = Literal["beginner", "intermediate", "advanced"]

class Constraints(BaseModel):
    available_days_per_week: int = Field(ge=1, le=7)
    session_duration_min: int = Field(ge=30, le=120)
    equipment: List[str] = []

class Preferences(BaseModel):
    experience_level: Experience
    target_muscle_priority: List[str] = []
    injuries: List[str] = []
    rest_days: List[Literal["mon","tue","wed","thu","fri","sat","sun"]] = []

class RunConfig(BaseModel):
    max_turns: int = 6
    trace: bool = False

class ProgramRequest(BaseModel):
    user_id: str
    goal: Literal["hypertrophy"]
    constraints: Constraints
    preferences: Preferences
    run_config: RunConfig = RunConfig()

class ProgressionEntry(BaseModel):
    metric: str
    values: List[float]

class ExerciseItem(BaseModel):
    name: str
    muscles: List[str]
    sets: int
    reps: str
    rir: Optional[int] = None
    rpe: Optional[float] = None
    rest_seconds: int
    tempo: Optional[str] = None
    notes: Optional[str] = None
    # Avoid dict/object types in schema to prevent additionalProperties
    # Use a structured list instead of Dict[str, List[float]]
    # e.g., metric="load", values=[60.0, 62.5, 65.0]
    progression: Optional[List[ProgressionEntry]] = None

class SessionPlan(BaseModel):
    day_index: int  # 1..7
    title: str
    focus: List[str]
    exercises: List[ExerciseItem]

class ProgramResponse(BaseModel):
    program_id: str
    user_id: str
    goal: Literal["hypertrophy"]
    macrocycle_weeks: int
    split: List[str]
    sessions: List[SessionPlan]
    notes: Optional[str] = None
    trace_id: Optional[str] = None
