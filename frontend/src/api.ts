const API_BASE = process.env.REACT_APP_API_URL || '/api';

export type Operation =
  | 'add'
  | 'subtract'
  | 'multiply'
  | 'divide'
  | 'exponentiate'
  | 'sqrt'
  | 'percentage';

interface CalculatePayload {
  operation: Operation;
  a: number;
  b?: number;
}

interface CalculateResult {
  result: number;
}

interface ErrorResult {
  error: string;
}

export async function calculate(
  operation: Operation,
  a: number,
  b?: number
): Promise<number> {
  const payload: CalculatePayload = { operation, a };
  if (b !== undefined) {
    payload.b = b;
  }

  const response = await fetch(`${API_BASE}/calculate`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  });

  if (!response.ok) {
    const err: ErrorResult = await response.json();
    throw new Error(err.error || 'Unknown error');
  }

  const data: CalculateResult = await response.json();
  return data.result;
}
