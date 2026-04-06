import { useContext } from "preact/hooks";
import "./style.css";
import editIcon from "../../assets/edit.svg";
import {
  LoadableResourceState,
  routeTo,
  useLoadableResource,
} from "../../utils/utils";
import { IconButton } from "../../components/IconButton";
import { LoadingScreen } from "../../components/LoadingScreen";
import { NoteAPI } from "../../utils/constants";

export function Home() {
  const api = useContext(NoteAPI);
  const [state, notes] = useLoadableResource(() => api.allNotes(), []);
  if (state === LoadableResourceState.Loading) {
    return LoadingScreen();
  }
  return (
    <section>
      {notes.current.map((note) => {
        return (
          <NoteCard
            title={note.title}
            content={note.content}
            updated={note.updatedAt.toLocaleDateString()}
            route={`/note/${note.id}`}
          />
        );
      })}
    </section>
  );
}

function NoteCard({ title, content, updated, route }) {
  return (
    <article class="NoteCard">
      <header>
        <h3>{title}</h3>
        <small>{updated}</small>
      </header>
      <p>{content}</p>
      <IconButton icon={editIcon} onClick={() => routeTo(route)}></IconButton>
    </article>
  );
}
