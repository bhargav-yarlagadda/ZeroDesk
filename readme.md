Got it 👍 Since you’re building an **AnyDesk clone** for learning/system-level exploration, I’ll draft a **professional README.md** template you can use and adapt for your project.

Here’s a **starter README**:

```markdown
# ZeroDesk

ZeroDesk is a lightweight **remote desktop application** inspired by AnyDesk/TeamViewer.  
It allows users to **share their screen, transmit keystrokes, and control remote systems** in real time.  
This project is built as a **learning experiment** to understand networking, real-time communication, and system-level programming.

---

## 🚀 Features
- 🔑 **Authentication** – JWT-based login sessions  
- 👥 **Role Selection** – choose between `Viewer` (control) or `Sharer` (share screen)  
- 🖥 **Screen Sharing** – real-time transmission of screen pixels  
- ⌨️ **Keystroke Transmission** – send keyboard/mouse events from viewer to sharer  
- ⚡ **Low Latency** – powered by WebRTC / WebSockets (depending on config)  
- 🛡 **Secure Communication** – encrypted sessions  

---

## 🛠 Tech Stack
- **Frontend**: TypeScript + React / Electron (for desktop integration)  
- **Backend**: Go (Golang) + WebSockets (Socket.IO / custom)  
- **Auth**: JWT for session management  
- **Transport**: WebRTC (P2P) / fallback to WebSockets  

---

## 📂 Project Structure
```

zerodesk/
│── client/       # Frontend (Electron/React/TS)
│── server/       # Backend (Golang)
│── docs/         # Documentation & notes
│── README.md     # Project info

````

---

## 🔧 Setup & Installation

### Prerequisites
- Node.js (>= 18)  
- Go (>= 1.21)  
- npm / yarn  
- (Optional) Electron for desktop builds  

### Clone Repo
```bash
git clone https://github.com/bhargav-yarlagadda/zerodesk.git
cd zerodesk
````

### Backend (Go)

```bash
cd server
go mod tidy
go run main.go
```

### Frontend (React/Electron)

```bash
cd client
npm install
npm run dev
```
### docker
```bash
docker run -d --name coturn -p 3478:3478/udp instrumentisto/coturn turnserver -n --no-auth --no-dtls --no-tls --listening-port=3478 --fingerprint

```
---

## 🎮 Usage

1. **Login** using JWT session.
2. **Select Mode**:

   * `Sharer`: Share your screen with others.
   * `Viewer`: View and control a shared screen.
3. **Connect via Session ID**.
4. **Enjoy real-time remote access**.

---

## 📌 Roadmap

* [ ] Clipboard sharing
* [ ] File transfer support
* [ ] Multi-monitor support
* [ ] Audio streaming
* [ ] End-to-end encryption

---

## ⚠️ Disclaimer

ZeroDesk is a **learning project** built for educational purposes.
It is **not production-ready** and should not be used in critical environments.

---


