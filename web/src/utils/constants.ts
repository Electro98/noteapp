import { Context, createContext } from "preact";
import { API } from "../api/API";

export const NoteAPI: Context<API> = createContext(null);

export const SharedLinkID: string = "SHARED_LINK";
