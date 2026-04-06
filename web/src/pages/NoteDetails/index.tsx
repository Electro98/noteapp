import {
  Dispatch,
  MutableRef,
  StateUpdater,
  useContext,
  useEffect,
  useId,
  useMemo,
  useRef,
  useState,
} from "preact/hooks";
import "./style.css";
import { useRoute } from "preact-iso";
import { getElemById, Getter, routeTo } from "../../utils/utils";
import { LoadingScreen } from "../../components/LoadingScreen";
import { NotFound } from "../_404";
import { Note, NoteDTO } from "../../models/Note";
import { NoteAPI } from "../../utils/constants";

enum NoteUIState {
  Loading,
  Failed,
  Loaded,
  Updating,
}

type UpdateNoteFunc = (n: Promise<Note>) => void;
type RawUiUpdateStateFunc = Dispatch<StateUpdater<NoteUIState>>;

function useNoteDetails(
  getInitialValue: Getter<Promise<Note>>,
): [NoteUIState, MutableRef<Note>, UpdateNoteFunc, RawUiUpdateStateFunc] {
  const [uiState, setUiState] = useState(NoteUIState.Loading);
  const resource: MutableRef<Note> = useRef(null);

  useMemo(() => {
    getInitialValue()
      .then((note) => {
        resource.current = note;
        setUiState(NoteUIState.Loaded);
      })
      .catch((_) => setUiState(NoteUIState.Failed));
  }, []);

  const updateNote = (newData: Promise<Note>) => {
    setUiState(NoteUIState.Updating);
    newData
      .then((note) => {
        resource.current = note;
        setUiState(NoteUIState.Loaded);
      })
      .catch((_) => setUiState(NoteUIState.Failed));
  };

  return [uiState, resource, updateNote, setUiState];
}

export function NoteDetails() {
  const route = useRoute();
  const id = parseInt(route.params["id"], 10);
  if (Number.isNaN(id)) {
    return NotFound();
  }

  const api = useContext(NoteAPI);
  const [uiState, refNote, updateNote, setUiState] = useNoteDetails(() =>
    api.note.read(id),
  );

  switch (uiState) {
    case NoteUIState.Loading: {
      return LoadingScreen();
    }
    case NoteUIState.Failed: {
      return LoadingFailed();
    }
    // case NoteUIState.Loaded: {
    // }
    // case NoteUIState.Updating: {
    // }
  }
  const sendNoteUpdate = (data: Note) => updateNote(api.note.update(data));
  const deleteNote = () => {
    setUiState(NoteUIState.Updating);
    api.note.delete(id).then(() => routeTo("/"));
  };
  const uiActive = uiState === NoteUIState.Loaded;
  return NoteDetailsElement(
    refNote.current,
    uiActive,
    sendNoteUpdate,
    deleteNote,
  );
}

function NoteDetailsElement(
  note: Note,
  uiActive: boolean,
  updateNote: (note: NoteDTO) => void,
  deleteNote: () => void,
) {
  const titleId = useId();
  const contentId = useId();
  const onSave = () => {
    if (!uiActive) {
      console.log("onSave: Get press when ui isn't active");
      return;
    }
    const updatedNote = {
      id: note.id,
      title: getElemById<HTMLInputElement>(titleId).value,
      content: getElemById<HTMLTextAreaElement>(contentId).value,
    };
    console.log("Sending an update for note:", updatedNote);
    updateNote(updatedNote);
  };
  return (
    <section>
      <header>
        <EditableTitle title={note.title} id={titleId} />
        <small>Last update: {note.updatedAt.toLocaleString()}</small>
      </header>
      <textarea id={contentId} disabled={!uiActive}>
        {note.content}
      </textarea>
      <footer>
        <button onClick={onSave} disabled={!uiActive}>
          Save
        </button>
        <button onClick={deleteNote} disabled={!uiActive}>
          Delete
        </button>
      </footer>
    </section>
  );
}

function LoadingFailed() {
  return <strong>This note doesn't exist or failed to load!</strong>;
}

function EditableTitle({
  title,
  id,
  disabled,
}: {
  title: string;
  id?: string;
  disabled?: boolean;
}) {
  const inputId = id || useId();
  const text = useRef(title);
  const [editMode, setEditMode] = useState(false);
  const startEditing = () => setEditMode(true);
  const onInput = (e) => (text.current = e.currentTarget.value);
  const onKeyDown = (e) => {
    if (e.key === "Enter") {
      setEditMode(false);
    }
  };
  useEffect(() => {
    if (editMode) {
      document.getElementById(inputId).focus();
    }
  }, [editMode]);
  return (
    <>
      <h2 class={editMode && "hidden"} onDblClick={startEditing}>
        {text.current}
      </h2>
      <input
        class={!editMode && "hidden"}
        type="text"
        name="title"
        placeholder="Your title"
        value={text.current}
        id={inputId}
        onInput={onInput}
        onKeyDown={onKeyDown}
        disabled={disabled}
      />
    </>
  );
}
