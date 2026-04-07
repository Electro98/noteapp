import {
  LocationProvider,
  Router,
  Route,
  hydrate,
  prerender as ssr,
} from "preact-iso";

import { Header } from "./components/Header.jsx";
import { Home } from "./pages/Home/index.jsx";
import { NotFound } from "./pages/_404.jsx";
import { NoteDetails } from "./pages/NoteDetails/index.js";
import { MockAPI } from "./api/MockAPI.js";

import "./style.scss";
import { NoteAPI, SharedLinkID } from "./utils/constants.js";
import { NewNote } from "./pages/NewNote/index.js";
import { RealAPI } from "./api/RealAPI.js";

export function App() {
  return (
    <LocationProvider>
      <NoteAPI.Provider value={new RealAPI("http://0.0.0.0:8000/api")}>
        <Header />
        <main class="container">
          <Router>
            <Route path="/" component={Home} />
            <Route path="/note/new" component={NewNote} />
            <Route path="/note/:id" component={NoteDetails} />
            <Route default component={NotFound} />
          </Router>
        </main>
        <a id={SharedLinkID} href="/" hidden />
      </NoteAPI.Provider>
    </LocationProvider>
  );
}

if (typeof window !== "undefined") {
  hydrate(<App />, document.getElementById("app"));
}

export async function prerender(data) {
  return await ssr(<App {...data} />);
}
