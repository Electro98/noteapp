import { useLocation } from "preact-iso";

import searchIcon from "../assets/search.svg";
import addIcon from "../assets/add.svg";
import { IconButton } from "./IconButton";
import { routeTo } from "../utils/utils";

export function Header() {
  const { url } = useLocation();

  return (
    <header>
      <nav>
        <ul>
          <li>
            <a href="/">
              <h1>NoteApp</h1>
            </a>
          </li>
        </ul>
        <ul>
          <li>
            <IconButton icon={searchIcon} />
          </li>
          <li>
            <IconButton icon={addIcon} onClick={() => routeTo("/note/new")} />
          </li>
        </ul>
      </nav>
    </header>
  );
}
