---
layout: post
title:  "The stolen bytes"
author: Salmon
date:   2025-01-18 17:29:13 +0100
categories: ramblings
---

MSVC. Possibly the most insane tool out there. It is sold as a C & C++ compiler,
when in reality it is a torture compiler. I have never compiled a program with MSVC
that *hasn't* had any bug that wouldn't have come up in GCC or Clang. It's stupid!
Hell, the ``stdint.h`` library didn't exist for the longest time! And *still*, it
*barely* works! I use clang now, btw.

*"Well, what about the title of this post? 'The stolen bytes', what's that?"*.
I'll explain that to you now. Firstly, take a look at this code:

{% highlight cpp %}
int p;
int i1;
int i2;
i1 = 1 << 16;
i2 = 1 << 8;
p = int(&i1)+3;

cout << hex;
cout << "&i1: " << int(&i1) << endl;
cout << "&i2: " << int(&i2) << endl;
for(int i = 0; i < 16; i++)
  cout << p << ": " << uint(*((byte*)p--)) << endl;
{% endhighlight %}

Then, let's run it:

{% highlight cpp %}
&i1: 12fac8
&i2: 12fabc
12facb: 0
12faca: 1
12fac9: 0
12fac8: 0
12fac7: cc
12fac6: cc
12fac5: cc
12fac4: cc
12fac3: cc
12fac2: cc
12fac1: cc
12fac0: cc
12fabf: 0
12fabe: 0
12fabd: 1
12fabc: 0
{% endhighlight %}

Still not getting it? Okay, let me rephrase that text for you: See the ``cc``
values right there? They are useless addresses. ``0xcc`` is a commonly used
"nothing" address, as it is almost impossible for the kernel to invoke.
There are 8 useless bytes between the two aligned integers
on the stack. Why is this? These are padding bytes, they are used so that..?
Wait, what? MSVC is *stealing* 8 bytes for "padding" when they could *easily*
be compressed. MSVC is a *thief*. This makes a lot of programs *much* slower!
Think about this on a larger scale! I call the above code 1024 times!
8 * 1024 is around 8000 bytes! 8KB have now been used. Let's do that 1024 times.
8 * 1024 is still around 8000... MEGABYTES!!! 8MB. Let's repeat *that* 1024 times.
Now, 8 * 1024 continues to be around 8000... But the value increases to *GIGABYTES*!?
8GB of *nothing* is now on your system! You have wasted 8GB! This is of course
a deliberate misunderstanding of how memory works, but it's a useful hypothetical for
demonstrating how this *could* waste memory, temporarily. To take our hypothetical even
further, consider the following piece of code:

{% highlight cpp %}
void evil() {
    int p;
    int i1;
    int i2;
    i1 = 1 << 16;
    i2 = 1 << 8;
    p = int(&i1)+3;
}

int main() {
    int i = 0;
    while (true) {
        while (++i < 1024) {
            evil();
        }
    }
    return 0;
}
{% endhighlight %}

This would, in theory, fill your system with bytes upon bytes (even though
it basically wouldn't, it would just make some programs much slower) and
demonstrates how bad MSVC is, really.

In smaller projects, this probably wouldn't make too much of a difference
(premature optimization and yada-yada) but in larger projects this could
cause delays and if you are a really bad coder: **crashes**.

Anyway, that was just my rambling. I'm way to lazy to contribute
anything interesting to society. I was actually planning to make
a blog post called "I'm too lazy" but I was too lazy to even
do that. Anyway... Enjoy my incredible rambling about MSVC
wasting 8B of space. Wait... Was this blog post meaningless?
I'm going to contemplate my blog post now, bye!
