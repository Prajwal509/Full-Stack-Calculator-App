import React from 'react';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import Calculator from '../components/Calculator';
import * as api from '../api';

// Mock the API module
jest.mock('../api');
const mockCalculate = api.calculate as jest.MockedFunction<typeof api.calculate>;

describe('Calculator', () => {
  beforeEach(() => {
    mockCalculate.mockReset();
  });

  it('renders the title', () => {
    render(<Calculator />);
    expect(screen.getByText('Calculator')).toBeInTheDocument();
  });

  it('renders all operation buttons', () => {
    render(<Calculator />);
    expect(screen.getByLabelText('add')).toBeInTheDocument();
    expect(screen.getByLabelText('subtract')).toBeInTheDocument();
    expect(screen.getByLabelText('multiply')).toBeInTheDocument();
    expect(screen.getByLabelText('divide')).toBeInTheDocument();
    expect(screen.getByLabelText('exponentiate')).toBeInTheDocument();
    expect(screen.getByLabelText('sqrt')).toBeInTheDocument();
    expect(screen.getByLabelText('percentage')).toBeInTheDocument();
  });

  it('shows two inputs for binary operations', () => {
    render(<Calculator />);
    expect(screen.getByLabelText('operand-a')).toBeInTheDocument();
    expect(screen.getByLabelText('operand-b')).toBeInTheDocument();
  });

  it('shows only one input for sqrt (unary)', () => {
    render(<Calculator />);
    fireEvent.click(screen.getByLabelText('sqrt'));
    expect(screen.getByLabelText('operand-a')).toBeInTheDocument();
    expect(screen.queryByLabelText('operand-b')).not.toBeInTheDocument();
  });

  it('validates empty input A', async () => {
    render(<Calculator />);
    fireEvent.click(screen.getByText('='));
    expect(screen.getByRole('alert')).toHaveTextContent('Please enter a valid number for A');
  });

  it('validates empty input B for binary ops', async () => {
    render(<Calculator />);
    await userEvent.type(screen.getByLabelText('operand-a'), '5');
    fireEvent.click(screen.getByText('='));
    expect(screen.getByRole('alert')).toHaveTextContent('Please enter a valid number for B');
  });

  it('calls API and displays result on success', async () => {
    mockCalculate.mockResolvedValue(8);
    render(<Calculator />);
    await userEvent.type(screen.getByLabelText('operand-a'), '5');
    await userEvent.type(screen.getByLabelText('operand-b'), '3');
    fireEvent.click(screen.getByText('='));

    await waitFor(() => {
      expect(screen.getByTestId('result')).toHaveTextContent('8');
    });
    expect(mockCalculate).toHaveBeenCalledWith('add', 5, 3);
  });

  it('displays error from API', async () => {
    mockCalculate.mockRejectedValue(new Error('division by zero'));
    render(<Calculator />);
    fireEvent.click(screen.getByLabelText('divide'));
    await userEvent.type(screen.getByLabelText('operand-a'), '10');
    await userEvent.type(screen.getByLabelText('operand-b'), '0');
    fireEvent.click(screen.getByText('='));

    await waitFor(() => {
      expect(screen.getByRole('alert')).toHaveTextContent('division by zero');
    });
  });

  it('clears inputs and result on C click', async () => {
    mockCalculate.mockResolvedValue(42);
    render(<Calculator />);
    await userEvent.type(screen.getByLabelText('operand-a'), '6');
    await userEvent.type(screen.getByLabelText('operand-b'), '7');
    fireEvent.click(screen.getByText('='));

    await waitFor(() => {
      expect(screen.getByTestId('result')).toBeInTheDocument();
    });

    fireEvent.click(screen.getByText('C'));
    expect(screen.queryByTestId('result')).not.toBeInTheDocument();
    expect(screen.getByLabelText('operand-a')).toHaveValue(null);
    expect(screen.getByLabelText('operand-b')).toHaveValue(null);
  });

  it('calls sqrt with only one operand', async () => {
    mockCalculate.mockResolvedValue(4);
    render(<Calculator />);
    fireEvent.click(screen.getByLabelText('sqrt'));
    await userEvent.type(screen.getByLabelText('operand-a'), '16');
    fireEvent.click(screen.getByText('='));

    await waitFor(() => {
      expect(screen.getByTestId('result')).toHaveTextContent('4');
    });
    expect(mockCalculate).toHaveBeenCalledWith('sqrt', 16, undefined);
  });
});
