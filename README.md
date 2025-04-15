# Byzantine Fault Tolerance

A basic simulation to show how a group of nodes can reach consensus in the presence of Byzantine nodes.

## What This Simulates

- **Honest nodes** propose the correct value.  
- **Byzantine nodes** propose garbage or inconsistent values.  
- The system reaches consensus **only if enough nodes agree on a value** — specifically, if **≥ n - f** nodes agree.  
- This demonstrates tolerance to **f = floor((n - 1) / 3)** Byzantine faults (as in [PBFT](https://en.wikipedia.org/wiki/Practical_Byzantine_Fault_Tolerance)).

## Example

Given `n = 7` nodes:  
- The system can tolerate up to `f = floor((7 - 1) / 3) = 2` Byzantine nodes.
- Consensus is achieved when at least `n - f = 5` nodes agree on a value.
