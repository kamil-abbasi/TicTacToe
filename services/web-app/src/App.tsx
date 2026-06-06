import { useEffect } from "react";
import Board from "./Board";

function App() {
	useEffect(() => {
		const socket = new WebSocket(`ws://localhost:8080/ws`);

		socket.addEventListener("open", () => {
			console.log("Connection with server established!");
		});

		socket.addEventListener("message", (event) => {
			console.log("Message from server: ", event.data);
		});

		return () => {
			socket.close();
			console.log("Disconnected from server");
		};
	}, []);

	return <Board />;
}

export default App;
