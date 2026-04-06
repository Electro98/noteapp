import { Note, NoteDTO } from "../models/Note";

export interface API {
  allNotes(): Promise<Note[]>;
  note: BaseCRUD<Note, NoteDTO>;
}

export interface BaseCRUD<Model, TypeDTO> {
  create(obj: TypeDTO): Promise<Model>;
  read(id: number): Promise<Model>;
  update(obj: TypeDTO): Promise<Model>;
  delete(id: number): Promise<void>;
}

/** Can be thrown if client has made logic mistake.
 *      Such as trying to update inexistent ID.
 */
export class LogicError extends Error {
  constructor(message: string, options?: ErrorOptions) {
    super(message, options);
  }
}

/** Response from server failed to consistent
 *      with expected structure.
 *
 * Should not be expected in production code.
 */
export class ValidationError extends Error {
  constructor(message: string, options?: ErrorOptions) {
    super(message, options);
  }
}

export class NetworkError extends Error {
  constructor(message: string, options?: ErrorOptions) {
    super(message, options);
  }
}
