import { useState } from "react";

export type TileProps = {
	symbol?: string;
	onTileClicked: () => void;
};

export default function Tile(props: TileProps) {
	const [enabled, setEnabled] = useState(true);
	const [symbol, setSymbol] = useState<string | undefined>(undefined);

	function handleClick() {
		setEnabled(false);
		setSymbol(props.symbol);
		props.onTileClicked();
	}

	return (
		<button
			className="tile"
			type="button"
			onClick={handleClick}
			disabled={!enabled}
		>
			{symbol}
		</button>
	);
}
