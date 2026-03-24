---
layout: post
title:  "Why do all C alternatives suck?"
author: Salmon
date:   2026-03-24 22:27:03 +0100
categories: dunno
---

I write pretty much all of my software in C nowadays. And the reason for that is clear: you can do *anything* in
C, and you can do it **BLAZINGLY FAST**. However it comes with some drawbacks: C can be a bit weird sometimes.
``segmentation fault: core dumped`` amiright? But then the question becomes: Why hasn't C been replaced? Why
hasn't a fast, memory safe and syntactically sound language replaced it?

I do not know, but we can take a look at the biggest alternatives:

### C++
Just bloated C DLC, next.

### Rust
Rust is an interesting thing: it aims to be as powerful and low-level as C, yet aims to be memory safe. It does
this using the *borrow checker*. The borrow checker is *awful*. Simply put:

1. Variables are in charge of freeing their own resources.
2. Therefore, resources can only have *one* "owner".
3. When doing assigning variables or passing arguments by value, the ownership
of the resources is transferred.

It is impossible to write things with the borrow checker. It is simply impossible. To illustrate how strange
borrow checking is, observe:

{% highlight rust %}
fn main() {
    let a = Box::new(5i32);
    let b = a; // b fucken steals the value!!!

    // this gives an error >:(
    println!("val of a: {}", a);
}
{% endhighlight %}

And what's more:

{% highlight rust %}
// Does nothing. At all.
fn abstain(c: Box<i32>) {  }

fn main() {
    let a = Box::new(5i32);

    abstain(a);

    // will _also_ give an error...
    println!("val of a: {}", a);
}
{% endhighlight %}

Rust is hopeless. And what's more: the compile time is so slow you could take a short nap before it has
finished compiling. **Never use rust**.

### Odin
Now I gotta admit, I haven't programmed in Odin that much, but look at this syntax:

{% highlight odin %}
package main

import "core:fmt"
import "core:net"
import "core:thread"

is_ctrl_d :: proc(bytes: []u8) -> bool {
	return len(bytes) == 1 && bytes[0] == 4
}

is_empty :: proc(bytes: []u8) -> bool {
	return(
		(len(bytes) == 2 && bytes[0] == '\r' && bytes[1] == '\n') ||
		(len(bytes) == 1 && bytes[0] == '\n') \
	)
}

is_telnet_ctrl_c :: proc(bytes: []u8) -> bool {
	return(
		(len(bytes) == 3 && bytes[0] == 255 && bytes[1] == 251 && bytes[2] == 6) ||
		(len(bytes) == 5 &&
				bytes[0] == 255 &&
				bytes[1] == 244 &&
				bytes[2] == 255 &&
				bytes[3] == 253 &&
				bytes[4] == 6) \
	)
}
{% endhighlight %}

Yeah. Good luck programming in this Go++-ass language.

### Zig
Zig is actually pretty cool. It has some pretty readable syntax, and it has some pretty good
properties. *However*, ZIG DOES NOT HAVE A STRING TYPE. ``[]u8`` is the string type, here, which I do
*not* enjoy. However, that would be okay, if ZIG DIDN'T HAVE LIKE 3 OTHER DESIGN FLAWS. Also: Zig
typing overall is absolutely killing me.

However, Zig *does* have the capability for *custom dynamic memory allocators*, which I find absolutely
awesome! Out of all languages on this list, I think Zig is the best language of them all, despite the
flaws. If I could just suggest... Some improvements. The biggest ones of all: *1)* Mandatory typing,
*2)* Actually good string typing:

{% highlight c %}
const str: [32]char = "Zig is pretty cool, tho.";
{% endhighlight %}

### C3
C3 *also* has some pretty useful features (I *need* good defers in my C code!!!), however, it falls
for the famed "bloated syntax" trap:

{% highlight c3 %}
module hello_world;
import std::io;

fn void main() {
    io::printn("Hello, world!");
}
{% endhighlight %}

``module {x}`` should not have to be written on line 1 of every project and if I may add, even if
I have neglected to mention this earlier; I do not find any reason to have a ``fn`` keyword.
Furthermore, C3 finds it necessary to have a ``foreach`` keyword.

{% highlight c3 %}
fn void example_foreach(float[] values) {
    foreach (index, value : values) {
        io::printfn("%d: %f", index, value);
    }
}
{% endhighlight %}

For what reason? Why not just use a for-loop? It's not even much harder. *This is the problem
with C3*. It is a bloated language that pretends to be sleek.

### The design of C
Furthermore, C alternatives aren't just bad, C is also a very sleek language, designwise. It
skips out on the ``fn`` keywords and the ``foreach`` loops. It is very, very simple. But that is
actually a good things, moreoften.

There doesn't need to be a "C alternative".
