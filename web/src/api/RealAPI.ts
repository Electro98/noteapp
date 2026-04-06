import { Note, NoteDTO } from "../models/Note";
import { API, BaseCRUD } from "./API";

export class RealAPI implements API {
  async allNotes(): Promise<Note[]> {
    const data = fetchJSON(`${this.baseUrl}/note`);
    return data.then((obj) => {
      (obj as Array<Note>).forEach(toNote);
      return obj;
    }) as Promise<Note[]>;
  }
  note: BaseCRUD<Note, NoteDTO>;
  // Props
  baseUrl: string;

  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
    this.note = {
      create: async (obj: NoteDTO) => {
        const data = fetchJSON(`${this.baseUrl}/note`, {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(obj),
        });
        const n = await (data as Promise<Note>);
        return toNote(n);
      },
      read: async (id: number) => {
        const data = fetchJSON(`${this.baseUrl}/note/${id}`);
        const n = await (data as Promise<Note>);
        return toNote(n);
      },
      update: async (obj: NoteDTO) => {
        const data = fetchJSON(`${this.baseUrl}/note`, {
          method: "PUT",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(obj),
        });
        const n = await (data as Promise<Note>);
        return toNote(n);
      },
      delete: async (id: number) => {
        const resp = fetch(`${this.baseUrl}/note?id=${id}`, {
          method: "DELETE",
        });
        const _ = await resp;
      },
    };
  }
}

function toNote(n: Note): Note {
  n.createdAt = new Date(n.createdAt);
  n.updatedAt = new Date(n.updatedAt);
  return n;
}

async function fetchJSON(
  url: RequestInfo | URL,
  init?: RequestInit,
): Promise<object> {
  console.log("Request to ", url);
  const resp = await fetch(url, init);
  const data = await resp.json();
  if (typeof data === "object" && data !== null) {
    return data as object;
  } else {
    throw "Not an object!";
  }
}
