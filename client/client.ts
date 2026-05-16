/**
 * Client-side equivalents of the Go backend logic for weighted selection and range spinning.
 * These types and classes are used for client-side simulations and demos, functioning
 * similarly to the 'wr' and 'ranges' packages in Go.
 */

/**
 * Weight information for an item. Equivalent to 'wr.W' in Go.
 */
export interface W {
	id: number;
	weights: number;
	empty: boolean;
}

/**
 * Interface for weighted items. Equivalent to 'wr.Wer' in Go.
 */
export interface Wer {
	reward(): number;
	info(): W;
}

/**
 * Handles weighted selection logic. Equivalent to 'wr.Slots' in Go.
 */
export class Slots {
	lists: Wer[] = [];
	total: number = 0;

	/**
	 * Calculates the total weight from the provided lists.
	 */
	init() {
		this.total = 0;
		for (const v of this.lists) {
			this.total += v.info().weights;
		}
	}

	/**
	 * Performs a weighted random selection (spin).
	 */
	spin(): Wer {
		const rnd = Math.random() * this.total;
		let last = 0;
		let selected: Wer | undefined;

		for (const v of this.lists) {
			const info = v.info();
			const start = last;
			last += info.weights;
			if (rnd <= last && start <= rnd) {
				selected = v;
				break;
			}
		}
		if (!selected) {
			throw new Error("should have selected something");
		}
		return selected;
	}
}

/**
 * Range selection logic. Equivalent to 'ranges.R' in Go.
 */
export class R {
	range: number;

	constructor(range: number) {
		this.range = range;
	}

	/**
	 * Simulates a roll/spin within the defined range.
	 */
	spin(): number {
		return Math.random() * this.range;
	}
}

/**
 * Basic implementation of the Wer interface for simple mapping.
 */
export class Basic implements Wer {
	w: W;
	rwd: number;

	constructor(w: W, reward: number) {
		this.w = w;
		this.rwd = reward;
	}

	info(): W {
		return this.w;
	}

	reward(): number {
		return this.rwd;
	}
}
