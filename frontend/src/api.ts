import { API_URL } from './config';

export interface CreateNoteResponse {
  id: string;
  url: string;
  expires_at: string;
}

export interface GetNoteResponse {
  id: string;
  content: string;
  created_at: string;
  expires_at: string;
}

export interface ApiError {
  error: string;
  message: string;
}

export async function createNote(content: string): Promise<CreateNoteResponse> {
  const response = await fetch(`${API_URL}/api/notes`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ content }),
  });

  if (!response.ok) {
    const error: ApiError = await response.json();
    throw new Error(error.message);
  }

  return response.json();
}

export async function getNote(id: string): Promise<GetNoteResponse> {
  const response = await fetch(`${API_URL}/api/notes/${id}`);

  if (!response.ok) {
    if (response.status === 404) {
      throw new Error('Заметка не найдена или срок её хранения истёк');
    }
    const error: ApiError = await response.json();
    throw new Error(error.message);
  }

  return response.json();
}

