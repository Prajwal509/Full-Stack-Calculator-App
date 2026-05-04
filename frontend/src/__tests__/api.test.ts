import { calculate } from '../api';

// Mock global fetch
const mockFetch = jest.fn();
global.fetch = mockFetch;

describe('api.calculate', () => {
  beforeEach(() => {
    mockFetch.mockReset();
  });

  it('sends correct payload for binary operation', async () => {
    mockFetch.mockResolvedValue({
      ok: true,
      json: async () => ({ result: 5 }),
    });

    const result = await calculate('add', 2, 3);
    expect(result).toBe(5);
    expect(mockFetch).toHaveBeenCalledWith('/api/calculate', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ operation: 'add', a: 2, b: 3 }),
    });
  });

  it('sends correct payload for unary operation', async () => {
    mockFetch.mockResolvedValue({
      ok: true,
      json: async () => ({ result: 4 }),
    });

    const result = await calculate('sqrt', 16);
    expect(result).toBe(4);
    expect(mockFetch).toHaveBeenCalledWith('/api/calculate', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ operation: 'sqrt', a: 16 }),
    });
  });

  it('throws error on non-ok response', async () => {
    mockFetch.mockResolvedValue({
      ok: false,
      json: async () => ({ error: 'division by zero' }),
    });

    await expect(calculate('divide', 10, 0)).rejects.toThrow('division by zero');
  });
});
