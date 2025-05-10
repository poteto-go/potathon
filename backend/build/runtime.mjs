import { connect } from "cloudflare:sockets";
import mod from "./app.wasm";

export async function loadModule() {
  return mod;
}

export function createRuntimeContext({ env, ctx, binding }) {
  return {
    env,
    ctx,
    connect,
    binding,
  };
}
