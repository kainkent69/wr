import { Slots, R, Basic, type Wer, type W } from './client.ts';

/**
 * Mirroring wr_test.go
 */
function testSlots() {
    console.log("--- Testing Slots (Weighted Selection) ---");
    
    const items: Wer[] = [
        new Basic({ id: 1, weights: 14000, empty: true }, 0),
        new Basic({ id: 2, weights: 2000, empty: false }, 300),
        new Basic({ id: 3, weights: 1400, empty: false }, 500),
        new Basic({ id: 4, weights: 600, empty: false }, 750),
    ];

    const slot = new Slots();
    slot.lists = items;
    slot.init();

    const iterations = 100000;
    const hits: Record<number, number> = {};

    for (let i = 0; i < iterations; i++) {
        const result = slot.spin();
        const id = result.info().id;
        hits[id] = (hits[id] || 0) + 1;
    }

    console.log(`Results after ${iterations} iterations:`);
    for (const item of items) {
        const id = item.info().id;
        const count = hits[id] || 0;
        const percentage = (count / iterations * 100).toFixed(2);
        const expected = (item.info().weights / slot.total * 100).toFixed(2);
        console.log(`ID ${id}: ${count} hits (${percentage}%) - Expected: ~${expected}%`);
    }
}

/**
 * Mirroring ranges/range_test.go
 */
class MockHilo {
    r: R;
    reward: number = 0;
    spins: number = 0;
    hits: number = 0;

    constructor(range: number) {
        this.r = new R(range);
    }

    run(prob: number, bet: number) {
        const acceptable = this.r.range / (100 / prob);
        const result = this.r.spin();
        this.reward = (bet * prob) / (100 + 3);
        this.spins++;
        if (result <= acceptable) {
            this.hits++;
        }
    }
}

function testRange() {
    console.log("\n--- Testing Range (Hilo Simulation) ---");
    
    const hilo = new MockHilo(10000);
    const prob = 50;
    const bet = 100;
    const iterations = 100000;

    for (let i = 0; i < iterations; i++) {
        hilo.run(prob, bet);
    }

    const hitRate = (hilo.hits / iterations * 100).toFixed(2);
    console.log(`Results after ${iterations} iterations:`);
    console.log(`Probability: ${prob}%`);
    console.log(`Hit Rate: ${hitRate}%`);
}

// Run tests
try {
    testSlots();
    testRange();
} catch (error) {
    console.error("Test failed:", error);
    process.exit(1);
}
