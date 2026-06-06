import { useState } from "react";
import Tile from "./Tile";

export default function Board() {
	const [player, setPlayer] = useState<"x" | "o">("x");

	return (
		<div className="board">
			{Array.from({ length: 9 }).map((_, i) => (
				<Tile
					key={i}
					symbol={player}
					onTileClicked={() => {
						if (player === "x") setPlayer("o");
						else setPlayer("x");
					}}
				/>
			))}
		</div>
	);
}
