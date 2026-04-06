import { Note, NoteDTO } from "../models/Note";
import { API, BaseCRUD, LogicError } from "./API";

const wait = (time: number) =>
  new Promise((resolve) => setTimeout(resolve, time));

export class MockAPI implements API {
  _storage: Note[] = [
    {
      id: 1,
      title: "Title 1",
      content: "Hello world!",
      createdAt: new Date(2024, 12, 23),
      updatedAt: new Date(2025, 3, 25),
    },
    {
      id: 2,
      title: "Title 2",
      content: "Hello worlds!",
      createdAt: new Date(2024, 12, 23),
      updatedAt: new Date(2025, 3, 25),
    },
  ];
  allNotes(): Promise<Note[]> {
    console.log(`API: allNotes`);
    return wait(2000).then(() => this._storage);
  }
  note: BaseCRUD<Note, NoteDTO>;
  constructor() {
    let storage = this._storage;
    this.note = {
      create: (obj: NoteDTO): Promise<Note> => {
        console.log(`API: create ${obj}`);
        const now = new Date(Date.now());
        const note = {
          id: storage[storage.length - 1].id + 1,
          title: obj.title,
          content: obj.content,
          createdAt: now,
          updatedAt: now,
        };
        storage[storage.length] = note;
        console.log(storage);
        return Promise.resolve(note);
      },
      read: (id: number): Promise<Note> => {
        console.log(`API: read ${id} id`);
        const note = storage.find((n) => n.id == id);
        if (note === undefined) {
          return Promise.reject(new LogicError("No such elem"));
        }
        return Promise.resolve(note);
      },
      update: async (obj: NoteDTO): Promise<Note> => {
        console.log(`API: update ${obj.id} id`);
        const index = storage.findIndex((n) => n.id == obj.id);
        if (index == -1) {
          return Promise.reject(new LogicError("No such elem"));
        } else {
          const now = new Date(Date.now());
          if (obj.title) {
            storage[index].title = obj.title;
          }
          if (obj.content) {
            storage[index].content = obj.content;
          }
          storage[index].updatedAt = now;
          await wait(2500);
            return storage[index];
        }
      },
      delete: (id: number): Promise<void> => {
        console.log(`API: delete ${id} id`);
        const index = storage.findIndex((n) => n.id == id);
        if (index != -1) {
          storage.splice(index, 1);
          return Promise.resolve();
        }
        return Promise.reject(new LogicError("No such elem"));
      },
    };
  }
}
