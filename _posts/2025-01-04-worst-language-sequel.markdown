---
layout: post
title:  "Worst language ever: The sequel"
author: Elis Staaf
date:   2025-01-04 17:21:56 +0100
categories: projects
---

Okay okay, if you're reading this, you probably know about my endeavour
in language making that I had in my
[last blog post](https://elisstaaf.github.io/projects/2025/01/01/worst-language.html),
and you probably hated me for that. But if you didn't know, in my last blog
post I wrote a terrible C adaptation in python. And for you to understand *this*
blog post, that blog post is required reading.

Okay, now that you've caught yourself up on the latest and greatest language,
it's time to explain. The year was *still* 2024, I wasn't sick now at least,
but I was *incredibly* bored, *frustrated* even. Also it was like in the middle
of the night so I couldn't like... Take a walk. Then I realized something. "Why
not take out my frustration trough a programming language!". So yeah... I was
all of a sudden thrown into another language-making endeavour. But were do I 
start?

Step 1. What should the language be about? I thought about this for quite
the while, actually. But then I realized: `screaming`. *That* is the peak of
frustration. Just imagine, something *terrible* happened and you need a way
to take out your frustration.

{% highlight c %}
UUUUUUUUuAUUUUuAUUAUUUAUUUAUEEEEOoAUAUAOAAUuEoEOoAAeAOOOeUUUUUUUeeUUUeAAeEOeEeUUUeOOOOOOeOOOOOOOOeAAUeAUUe
{% endhighlight %}

And now your frustration has suddenly stopped. You feel... Relieved. *That*, my
friend, is what I want to create.

Step 2. Okay, but what language should I write it *in*? C. C is the greatest
programming language of all time and I don't want to literally *torture* myself
by using a language like python or god forbid *javascript*. Eugh, I'm gonna puke...

Step 3. Okay, but... How? This is also a question I thought about for a while.
But then I realized. Brainfuck. Brainfuck is a programming language that
operates using like shifting and output of a data pointer. Any non-technical
person might not understand what I'm talking about. And the truth is... I
barely know either! Just know, this is a program that prints *Hello, World!* 
in the brainfuck programming language (split into 3 different lines
for your viewing pleasure):

{% highlight brainfuck %}
>++++++++[<+++++++++>-]<.>++++[<+++++++>-]<+.+++++++..+++.>>++++++[<+++++++>-]<+
+.------------.>++++++[<+++++++++>-]<+.<.+++.------.--------.>>>++++[<++++++++>-
]<+.
{% endhighlight %}

So what if *I* make a transalation of brainfuck but replace the little funny
symbols (ex. `>`) by existential crisis symbols (ex. `A`), meaning that e.x
`>>+<+..` could become `AAeEeUU`. So... I'm going to be writing a brainfuck
interpreter? Oh god no...

Let's write a brainfuck interpreter! Unlike my first language, I'm going to
be going more into the "Nit and gritty" with this one. Alright, the first thing
we need is a file for evaluation, let's call that one `src/scream.c` (we'll fix 
headers later). Okay, what now? We need to create a function to hold the
evaluation, let's call it `expr` for `evaluate EXPRession`. Then we need
to create a data pointer, that is just a fact, we just need it to be a 
``char*`` (e.g a string) to be able to print it later:

{% highlight c %}
void expr(char *command_pointer, char *input) {
    char *dp;
}
{% endhighlight %}

Okay, but what should `dp` be? Well it needs to hold data. But what
should the size of that data be. The truth is; no idea. I chose 1001
but I think whatever is okay. We need the data to be empty, e.g `{0}`,
then we need to take some of that data into `dp` (I chose 1/2):

{% highlight c %}
#define DATASIZE 1001

void expr(...) {
    char *data[DATASIZE] = {0};
    char *dp;
    dp = &data[DATASIZE/2];
}
{% endhighlight %}

Also, as to not clutter the code, let's create a header file called,
say `include/scream.h`, and then compile with the `-I./include`
option. Now:

{% highlight c %}
#include <scream.h>

void expr(...) {
    char *data[DATASIZE] = {0};
    char *dp;
    dp = &data[DATASIZE/2];
}
{% endhighlight %}

Now, brackets exist, so let's add a bracket_flag in. It should
be an `int` since we're going to iterate with it later. Also,
add a command variable to control command. It should be a
`char` since every command is a `char` (e.x `'>'`). Let's
update the `expr` function to include all these new additions:

{% highlight c %}
void expr(...) {
    int bracket_flag;
    char command;
    char *data[DATASIZE] = {0};
    char *dp;
    dp = &data[DATASIZE/2];
}
{% endhighlight %}

Now, let's iterate! We need to infinitely iterate using `while (command = *command_pointer)`,
then add all of the commands! I'm not going to go trough all of the commands, I'm just
going to show you the while loop:

{% highlight c %}
  while (command = *command_pointer++)
    switch (command) {
    case '>':
      dp++;
      break;
    case '<':
      dp--;
      break;
    case '+':
      (*dp)++;
      break;
    case '-':
      (*dp)--;
      break;
    case '.':
      printf("%c", *dp);
      break;
    case ',': 
      *dp = *input++;
      break;
    case '[':
      if (!*dp) { 
        for (bracket_flag=1; bracket_flag; command_pointer++)
          if (*command_pointer == '[')
            bracket_flag++;
          else if (*command_pointer == ']')
            bracket_flag--;
      } 
      break;
    case ']':
      if (*dp) {
        command_pointer -= 2;  
        for (bracket_flag=1; bracket_flag; command_pointer--)
          if (*command_pointer == ']')
            bracket_flag++;
          else if (*command_pointer == '[')
            bracket_flag--;
        command_pointer++;     
      }
      break;  
    }
{% endhighlight %}

Oh, and let's add a `printf("\n")` at the end
of the function for good measure:

{% highlight c %}
void expr(...) {
    int bracket_flag;
    char command;
    char *data[DATASIZE] = {0};
    char *dp;
    dp = &data[DATASIZE/2];
    while(...) {...}
    printf("\n");
}
{% endhighlight %}

Then, let's just add a main file to strap all of this together!
But before we do that, we need to add a function prototype to
`include/scream.h`:

{% highlight c %}
#include <stdio.h>

#define DATASIZE 1001

void expr(char *command_pointer, char *input);
{% endhighlight %}

Now, let's add a main file in `src/main.c`:

{% highlight c %}
#include <stdio.h>
#include <scream.h>

#define BUF_LEN 65536

int main(int argc, char **argv) {
    if (argc <= 1) {
        printf("USAGE: scream <file>\n");
        return 1;
    }
    FILE *file_ptr;
    char buf[BUF_LEN];

    file_ptr = fopen(argv[1], "r");

    if (NULL == file_ptr) {
        printf("File can't be opened \n");
    }

    while (fgets(buf, BUF_LEN, file_ptr) != NULL) {}
	
	char *input = "";
	expr(buf, input);
	return 0;
}
{% endhighlight %}

Just some bootstrapping, nothing special.

Step 4. ... Profit? No, not quite yet. In order to completely sell the
language to people, we need some examples. And before you ask, *NO*. I
will *never* write examples for this language. What I *can* do is take
brainfuck examples from the internet. I use the "Neovim" text editor,
which uses "Vim Script" as it's scripting language, so I simply wrote
a quick vim script to find and replace brainfuck instructions with
screamlang instructions. It looks like this:

{% highlight vim %}
%s/>/A/g
%s/</E/g
%s/+/U/g
%s/-/O/g
%s/\./e/g
%s/,/a/g
%s/\[/u/g
%s/\]/o/g
%s/\n//g
%s/ //g
{% endhighlight %}

And magically we have a *Hello World* program in screamlang:

{% highlight c %}
UUUUUUUUuAUUUUuAUUAUUUAUUUAUEEEEOoAUAUAOAAUuEoEOoAAeAOOOeUUUUUUUeeUUUeAAeEOeEeUUUeOOOOOOeOOOOOOOOeAAUeAUUe
{% endhighlight %}

After a quick `Makefile` my interpreter was good to go! So I released it onto github,
check it out [here](https://github.com/ElisStaaf/scream). And with that, I'd like to
present to you, the last step of this language endeavour:

*Step 5. Profit.*
