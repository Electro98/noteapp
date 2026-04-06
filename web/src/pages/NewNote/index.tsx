import { useContext, useId, useState } from "preact/hooks";
import { getElemById, routeTo } from "../../utils/utils";
import { NoteAPI } from "../../utils/constants";

export function NewNote() {
  const titleId = useId();
  const contentId = useId();
  const [uiActive, setUiActive] = useState(true);

  const api = useContext(NoteAPI);

  const saveNewNote = () => {
    if (!uiActive) {
      return;
    }
    const newNoteData = {
      title: getElemById<HTMLInputElement>(titleId).value,
      content: getElemById<HTMLTextAreaElement>(contentId).value,
    };
    api.note
      .create(newNoteData)
      .then((newNote) => routeTo(`/note/${newNote.id}`))
      .catch((reason) => console.log("Request failed reason: ", reason));
    setUiActive(false);
  };

  return (
    <section>
      <header>
        <h3>New Note</h3>
        <input
          placeholder="Your great title!!"
          id={titleId}
          disabled={!uiActive}
        />
      </header>
      <textarea
        id={contentId}
        placeholder="Dear diary, I ..."
        disabled={!uiActive}
      />
      <footer>
        <button onClick={saveNewNote} disabled={!uiActive}>
          Create new note
        </button>
      </footer>
    </section>
  );
}
