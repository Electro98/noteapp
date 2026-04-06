export interface Note {
  id: number;
  title: string;
  content: string;
  createdAt: Date;
  updatedAt: Date;
}

export interface NoteDTO {
  id?: number;
  title: string;
  content: string;
}

export function isNote(obj: any): obj is Note {
  return false;
}
