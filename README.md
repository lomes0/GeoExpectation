We have an experiment $T$ with probability of $p$ for success. We preform $T$ over and over again until the first success. If $X$ denote the number of experiments held, then $X$ is a random variable with geometric distribution. Let's show that $X$ has exception $\frac{1}{p}$.



<u>Proof</u>

Denote $X_i$ a random variable counting the number of success in the i-th experiment.

We have that $X = \sum\limits_{i\in \N} X_i = X_0 + \sum\limits_{i\in \N>0}X_i$. And therefore:
$$
E(X) = E(X_0) + E(\sum\limits_{i\in \N>0}X_i)
$$
Since the first experiment is always preformed, then $E(X_0) = 1$.

If $Y = \sum\limits_{i\in \N>0}X_i$, then we can say $Y$ counts the number of tries starting from the second experiment. And we can determine $E(Y) = (1-p)E(X)$.

Putting it all together, we have:
$$
1 = p\cdot E(X)
$$
Meaning $E(X) = \frac{1}{p}$.



<u>Example</u>

Say we have an unfair coin, with $\frac{1}{10}$ probability for tossing heads, and $\frac{9}{10}$ for tossing tails.

The next program simulate 100 rounds of this porcess.

```go
package main

import (
	"fmt"
	"math/rand"
)

const (
	HEAD  = iota
	TAILS = iota
)

type Flip struct {
	side uint8
}

func flip() Flip {

	if rand.Intn(10) == 0 {
		return Flip{0}
	}
	return Flip{1}
}

func toss_until() uint32 {

	var i uint32 = 0
	for true {
		i++
		var c = flip()
		if c.side == HEAD {
			return i
		}
	}
	return 0
}

func main() {

	const rounds = 100
	var results = make([]uint32, rounds)

	for i := 0; i < rounds; i++ {

		ntries := toss_until()
		results[i] = ntries
	}

	var toss_sum uint32 = 0
	for i := 0; i < rounds; i++ {
		fmt.Printf("Round %v took %v tosses\n", i, results[i])
		toss_sum += results[i]
	}
	fmt.Printf("Average number of tosses: %v\n", float64(toss_sum)/float64(rounds))
}
```



Yielding the output:

```
.
.
Round 94 took 33 tosses
Round 95 took 4 tosses
Round 96 took 24 tosses
Round 97 took 2 tosses
Round 98 took 18 tosses
Round 99 took 1 tosses
Average number of tosses: 9.18
```
