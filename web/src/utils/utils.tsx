import { MutableRef, useMemo, useRef, useState } from "preact/hooks";
import { SharedLinkID } from "./constants";

export type Getter<T> = () => T;

export enum LoadableResourceState {
  Loading,
  Ready,
  Failed,
}

export function useLoadableResource<T>(
  getRes: Getter<Promise<T>>,
  initialState?: T,
): [LoadableResourceState, MutableRef<T>] {
  const [resourceState, setResourceState] = useState(
    LoadableResourceState.Loading,
  );
  const resource: MutableRef<T> = useRef(initialState);

  useMemo(() => {
    getRes()
      .then((res) => {
        resource.current = res;
        setResourceState(LoadableResourceState.Ready);
      })
      .catch((_) => setResourceState(LoadableResourceState.Failed));
  }, []);

  return [resourceState, resource];
}

export function getElemById<T>(id: string): T {
  return document.getElementById(id) as T;
}

export function routeTo(route: string) {
  const link = getElemById<HTMLAnchorElement>(SharedLinkID);
  link.href = route;
  link.click();
}
