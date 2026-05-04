import React, { useState } from 'react';
import { calculate, Operation } from '../api';
import './Calculator.css';

const OPERATIONS: { label: string; value: Operation; unary?: boolean }[] = [
  { label: '+', value: 'add' },
  { label: '−', value: 'subtract' },
  { label: '×', value: 'multiply' },
  { label: '÷', value: 'divide' },
  { label: 'xʸ', value: 'exponentiate' },
  { label: '√', value: 'sqrt', unary: true },
  { label: '%', value: 'percentage' },
];

const Calculator: React.FC = () => {
  const [inputA, setInputA] = useState('');
  const [inputB, setInputB] = useState('');
  const [operation, setOperation] = useState<Operation>('add');
  const [result, setResult] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const isUnary = OPERATIONS.find((op) => op.value === operation)?.unary ?? false;

  const handleCalculate = async () => {
    setError(null);
    setResult(null);

    const a = parseFloat(inputA);
    if (isNaN(a)) {
      setError('Please enter a valid number for A');
      return;
    }

    let b: number | undefined;
    if (!isUnary) {
      b = parseFloat(inputB);
      if (isNaN(b)) {
        setError('Please enter a valid number for B');
        return;
      }
    }

    setLoading(true);
    try {
      const res = await calculate(operation, a, b);
      setResult(String(res));
    } catch (err: any) {
      setError(err.message || 'Calculation failed');
    } finally {
      setLoading(false);
    }
  };

  const handleClear = () => {
    setInputA('');
    setInputB('');
    setResult(null);
    setError(null);
  };

  return (
    <div className="calculator">
      <h1 className="calculator__title">Calculator</h1>

      <div className="calculator__operations">
        {OPERATIONS.map((op) => (
          <button
            key={op.value}
            className={`calculator__op-btn ${
              operation === op.value ? 'calculator__op-btn--active' : ''
            }`}
            onClick={() => {
              setOperation(op.value);
              setResult(null);
              setError(null);
            }}
            aria-label={op.value}
          >
            {op.label}
          </button>
        ))}
      </div>

      <div className="calculator__inputs">
        <input
          type="number"
          className="calculator__input"
          placeholder="A"
          value={inputA}
          onChange={(e) => setInputA(e.target.value)}
          aria-label="operand-a"
        />
        {!isUnary && (
          <input
            type="number"
            className="calculator__input"
            placeholder="B"
            value={inputB}
            onChange={(e) => setInputB(e.target.value)}
            aria-label="operand-b"
          />
        )}
      </div>

      <div className="calculator__actions">
        <button
          className="calculator__btn calculator__btn--primary"
          onClick={handleCalculate}
          disabled={loading}
        >
          {loading ? 'Calculating…' : '='}
        </button>
        <button className="calculator__btn calculator__btn--secondary" onClick={handleClear}>
          C
        </button>
      </div>

      {result !== null && (
        <div className="calculator__result" data-testid="result">
          {result}
        </div>
      )}

      {error && (
        <div className="calculator__error" role="alert">
          {error}
        </div>
      )}
    </div>
  );
};

export default Calculator;
