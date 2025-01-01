---
layout: post
title:  "How I made the worst programming language ever"
author: Elis Staaf
date:   2025-01-01 14:09:52 +0100
categories: projects
---
The year was 2024. I, as usual, had no idea what I was doing. And as a bonus,
I was incredibly sick. But then, an idea flew into my mind! Make C, but *WORSE*.
Everybody is going to love me! The worst adaptation of C, and I knew just the
programming language to write it in...

*Python*.

Python is quite possibly the worst programming language ever made, and if I plan to
beat that record, I'm going to have to use python to my advantage. No, really, I was 
just lazy and didn't want to write a C interpreter for my shitty language. So basically,
I was going to write an interpreted language on top of an interpreted language! PERFECT!

Step 1. Write python. This is pretty easy, so I decided to add an extra challenge: Make 
every name in the namespace awful. Take this for an example:

{% highlight python %}
def GetFileInfo(TheFileName: str):
    with open(TheFileName, "r") as TheFile:
        TheContentsOfTheFileInStrFormatInAListUsingTheReadLinesFunction = TheFile.readlines()
        TheFile.close()
    return TheContentsOfTheFileInStrFormatInAListUsingTheReadLinesFunction
{% endhighlight %}

Awful, right? Thank you. I decided I would start with reading a source file. I decided
to do that using the example above. Then came the fun part. Interpreting/Compiling.
I used [thesaurus](https://thesaurus.com) to get all of the words. This was my thought
process. I have this C function called `printf`. The "f" stands for "format". So, 
`printformat`. Then, see synonyms for print on thesaurus, get `lithograph`. So,
`lithographformat`. Then, make everything PascalCase so that the user has to shift
sometimes. So, `LithoGraphFormat`. So then I got a list of things to replace with
others. And this was how I did it, classic python... Awesomness...

{% highlight python %}
OneSelfsValue = OneSelfsLine\
.replace("Integer", "int")\
.replace("Character", "char")\
.replace("DoubleItAndGiveItToTheNextPerson", "double")\
.replace("MASSIVE", "long")\
.replace("Buoyant", "float")\
.replace("Deprived", "void")\
.replace("THESizeType", "size_t")\
.replace("NoteBook", "FILE")\
.replace("Architecture", "struct")\
.replace("Undisclosed", "unsigned")\
.replace("Persistent", "const")\
.replace("Worthless", "NULL")\
.replace("VarietyConstrue", "typedef")\
.replace("TroughoutTheTime", "while")\
.replace("Commence", "do")\
.replace("LoopThisWontYa?", "for")\
.replace("AssumingThat", "if")\
.replace("ButIfItsNotTrue", "else")\
.replace("Foreign", "extern")\
.replace("ReturnTheFollowing:", "return")\
.replace("LithoGraphFormat", "printf")\
.replace("EstablishString", "puts")\
.replace("Consultation", "scanf")\
.replace("ProportionOf", "sizeof")\
.replace("IncludeTheFollowingLibrary:", "#include")\
.replace("DefineTheFollowingMacro:", "#define")\
.replace("RaiseTheFollowingError:", "#error")\
.replace("Sober:", "#pragma")
{% endhighlight %}

Beauty. Then, add a main.py file and do some argv shit and write to C file
and Ta-Da! We've got a fully functioning terrible C adaptation, please *C*
the following. Please laugh.

{% highlight c %}
IncludeTheFollowingLibrary: <stdio.h>

Integer main(Integer argc, Character **argv) {
    LithoGraphFormat("Hello World!\n");
    ReturnTheFollowing: 0;
}
{% endhighlight %}

A few days pass. I continue work on my operating system and soon BadC becomes a distant memory.
Until I run out of ideas. Then, I got quite possibly the worst idea I have ever had in my life:
"Maybe I should rewrite BadC, *in* BadC". So that was what I was doing. I was still sick so I had
a lot of free time on my hands, and with that free time I... Ran into a bug. I ofcourse need a list
of the items to replace, like this:

{% highlight c %}
Character *keywords[] = {
    "Integer",                          "int",
    "Character",                        "char",
    "DoubleItAndGiveItToTheNextPerson", "double",
    "MASSIVE",                          "long",
    "Buoyant",                          "float",
    "Deprived",                         "void",
    "THESizeType",                      "size_t",
    "NoteBook",                         "FILE",
    "Architecture",                     "struct",
    "Undisclosed",                      "unsigned",
    "Persistent",                       "const",
    "Worthless",                        "NULL",
    "VarietyConstrue",                  "typedef",
    "TroughoutTheTime",                 "while",
    "Commence",                         "do",
    "LoopThisWontYa?",                  "for",
    "AssumingThat",                     "if",
    "ButIfItsNotTrue",                  "else",
    "Foreign",                          "extern",
    "ReturnTheFollowing:",              "return",
    "LithoGraphFormat",                 "printf",
    "EstablishString",                  "puts",
    "Consultation",                     "scanf",
    "ProportionOf",                     "sizeof",
    "IncludeTheFollowingLibrary:",      "#include",
    "DefineTheFollowingMacro:",         "#define",
    "RaiseTheFollowingError:",          "#error",
    "Sober:",                           "#pragma"
};
{% endhighlight %}

But uhh... This is were my python implementation falls short. It just substitutes a substring
with another string, which means that we're substituting with an identical string, and this
wont work. I realize something. I could possibly make an ignore thingy, that ignores! Like this:

{% highlight c %}
\do-ignore
/* ... */
\end-ignore
{% endhighlight %}

I managed to do that, using quite possibly the worst and most unreadable code I've ever written,
this was the implementation:

{% highlight python %}
for OneSelfsLine in self.OneSelfsCode:
    if self.OneSelfsIgnore and not OneSelfsLine.startswith("\\end-ignore"):
        OneSelfsCompiledCode += OneSelfsLine
        continue
    if OneSelfsLine.startswith("\\do-ignore"):
        self.OneSelfsIgnore = True
        continue
    self.OneSelfsIgnore = False
    if OneSelfsLine.startswith("\\end-ignore"):
        continue
{% endhighlight %}

I actually don't have any idea what that is doing. But anyways, then we use that, rewrite
using fairly standard (though shitty) C, and Ta-Da! Done with the BadC interpreter written
in BadC!

So, what did we learn? We learned that we should never let me create a programming language!
Anyway, see you later when I create my next programming language! No, but make sure to
check out my BadC implementation [here](https://github.com/ElisStaaf/BadC), and enjoy
the rest of your day!
