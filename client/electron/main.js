// electron/main.js
import { app, BrowserWindow } from "electron";
import path from "node:path";
import url from "node:url";
import "./ipc.js";

let mainWindow;

app.on("ready", () => {
  mainWindow = new BrowserWindow({
    width: 1000,
    height: 800,
    webPreferences: {
      contextIsolation: true,
      preload: path.join(process.cwd(), "electron", "preload.js"),
    },
  });

  const devUrl = process.env.ELECTRON_START_URL || "http://localhost:3000";
  const prodUrl = url.pathToFileURL(path.join(process.cwd(), "out/index.html"));

  mainWindow.loadURL(process.env.ELECTRON_START_URL ? devUrl : prodUrl.href);
});
