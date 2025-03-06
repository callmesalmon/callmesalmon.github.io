---
layout: base
permalink: /books/cyfraith-cath/cat-salesman-v
title: ""
---

# The cat salesman V
Now, with all of this "manual labour" and "hard work", you're starting to
look like a working class peasant! You're an entrepreneur, not a peasant! Yes,
why not just... Hire people? You probably wouldn't die if you were just sitting
in an office somewhere (in 900, at least, not *now*), and you would get super
rich! Yes! Why, not do that?

## The ethical way
Split the profits *evenly*. That's ethical, right? I couldn't think of anything
more ethical than a worker owned business, so that could work! Add that to our
ceca calculations and we're good to go (for x = # of workers)!

```
(800 / 50) / x
(80 / 5) / x
16 / x
```

If x -- for example -- is 4, our ceca is 16/4 or 8. But that isn't cash-money!
We need real cash-money, which we can do trough

## The unethical way
Underpay people! That's *another* really good idea, especially if you want to
earn the maximum number of cash-money in a ceca as possible! If we pay people
*unfairly* we can do something like (for x = # of workers, y = ceiniog/worker):

```
(800 / 50) - (x * y)
(80 / 5) - (x * y)
16 - (x * y)
```

So *now* if we have 4 workers and pay them 1 ceiniog we get (16 - (4 * 1)) * 2 ceiniog
or 28 ceiniog. See! Now we're *really* making cash-money! We've reduced our chance
of "muerto" (if you know what I mean) while severely underpaying workers to maximize
our money earned!

## Robbed and killed
28 ceiniog is of course a very generous estimate, as we would *probably* get murdered
and robbed along the way. What I mean is that it's probably a good idea to add our
probability calculations to our ceca calculations, something like (for x = # of workers,
y = ceiniog/worker):

```
((80 / 5) - (x * y)) * P(1)
(16 - (x * y)) * P(1)
(16 - (x * y)) * 0.6
```

Making it less ugly (Sh = Share {as in *your share*}, Pr = Probability remainder):

```
Sh(a, b) = { (80 / 5) - (a * b) },
Pr(a) = { a - P(1) },
Sh(x, y) * Pr(1)
```

And if you're lazy, you could put this into a function! I will be doing that since
I also want a better way of formulating cecas. Here it is:

```
Sh(a, b) = { (80 / 5) - (a * b) },
Pr(a) = { a - P(1) },
Ce(a, b) = { Sh(a, b) * Pr(1) }
```

So now with our four underpaid workers (``Ce(x, y)``), we'd get around 17 ceiniog on average.
Not good, but somehow better (for you) than splitting it evenly! The last part of this
actual calculation hell would be a last explanation of the ``Ce()`` function, as so:

```
Ce(x, y)
-> Sh(x, y) * Pr(1)
-> ((80 / 5) - (x * y)) * Pr(1)
-> (16 - (x * y)) * Pr(1)
-> (16 - (x * y)) * 0.6
```

## [Next chapter: The new code](/books/cyfraith-cath/new-code)

Sources:
- [Ancient laws and
institutes of Wales](https://archive.org/details/bub_gb_4_qi_6p1ZucC/page/27/mode/2up)
- [Youtube (Cambrian Chronicles) -
Medieval cat laws](https://www.youtube.com/watch?v=jD3b1s-s9bk&themeRefresh=1)
- [Hubpages - The cat legislation of the medieval king
"Hywel the Good"](https://discover.hubpages.com/animals/the-cat-legislation-of-the-medieval-king-hywel-the-good)
