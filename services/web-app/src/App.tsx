function App() {
	const clientName = `player${Math.trunc(Math.random() * 100)}`;

	const socket = new WebSocket(`ws://localhost:8080/ws?name=${clientName}`);

	socket.addEventListener("open", (event) => {
		socket.send("Hello from client!");
	});

	socket.addEventListener("message", (event) => {
		console.log("Message from server: ", event.data);
	});

	return <p>Hello, World!</p>;
}

export default App;
