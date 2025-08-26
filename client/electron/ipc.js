
// electron/ipc.js
import { ipcMain } from "electron";

ipcMain.handle("ping", () => {
  return "pong from main";
});
