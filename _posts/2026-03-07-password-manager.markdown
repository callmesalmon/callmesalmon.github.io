---
layout: post
title:  "Making a somewhat functional password manager in C"
author: Salmon
date:   2026-03-07 20:03:00 +0100
categories: tutorials
---

C is the greatest programming language of all time. Nothing beats it. Nothing will ever
beat it. It's just so, so powerful. It can do *anything*

Just one small problem: How does one program C? I of course know this already, but as
a small excercise I think it can be useful to write a small "application" (glorified
CLI) in C. At the time of writing said application, I had also not programmed in a while.

So what project would we choose? Well, the little people on the internet gave me a few
suggestions:
- A TODO app
- A HTTP server
- A calculator
- A *password manager*

As I had not dabbled in encryption a lot before, I decided on writing a password manager,
and this semi-tutorial is a partial guide on how to replicate
[locksmith](https://github.com/callmesalmon/locksmith), my implementation of a password
manager for UNIX-like systems. You are expected to have basic knowledge of C, but you
don't need to be a 10x engineer by any means.

### 1. Encryption
So where do we begin? Well I would say we should start with a cryptography interface
for easy use in the more high-level parts of our project. Now I need you to listen to me
here: *do not write your own encryption library*. Use a functioning encryption library
instead. I have chosen [libsodium](https://libsodium.gitbook.io/doc/) for this tutorial,
so go ahead and install that.

Oh, and before I forget! All .c and .h files should probably be placed in a directory
called ``src/`` or something. That's what I did, at least.

Now, I gotta admit: I "borrowed" most of the cryptography code from
[pass](https://github.com/briandowns/pass). So there isn't much of an explanation here.
But you can go ahead and create your cryptography files: ``crypto.c``
and ``crypto.h``. Then we create a simple cryptography interface, starting with ``crypto.h``:

{% highlight c %}
#ifndef CRYPTO_H
#define CRYPTO_H

#include <sodium.h>

#define MAX_STRING_LEN 128 // Maximum string length: used all across the program.

int encrypt(const char *fname, const char *password, const unsigned char key[crypto_secretstream_xchacha20poly1305_KEYBYTES]);
unsigned char *decrypt(const char *source_file, const unsigned char key[crypto_secretstream_xchacha20poly1305_KEYBYTES]);
int create_key(const char *key_file);

#endif
{% endhighlight %}

We create a macro called ``MAX_STRING_LEN`` to hold a general maximum string length to be used across the program.
Set to ``128`` in my implementation. Then we define our core functions, which will be explained further when we
actually implement them. Lets start with ``encrypt()``

{% highlight c %}
#include <stdio.h>
#include <string.h>
#include <sodium.h>

#include "crypto.h"

int encrypt(const char *fname, const char *password, const unsigned char key[crypto_secretstream_xchacha20poly1305_KEYBYTES]) {
    unsigned char buf_out[MAX_STRING_LEN + crypto_secretstream_xchacha20poly1305_ABYTES];
    unsigned char header[crypto_secretstream_xchacha20poly1305_HEADERBYTES];

    crypto_secretstream_xchacha20poly1305_state st;

    unsigned long long out_len;
    unsigned char tag;
    
    FILE *fptr = fopen(fname, "wb");

    crypto_secretstream_xchacha20poly1305_init_push(&st, header, key);
    
    fwrite(header, 1, sizeof header, fptr);

    size_t pass_len = strlen(password);
    do {
        tag = EOF ? crypto_secretstream_xchacha20poly1305_TAG_FINAL : 0;
        crypto_secretstream_xchacha20poly1305_push(
            &st, buf_out, &out_len, (unsigned char *)password, pass_len, NULL, 0, tag);
        fwrite(buf_out, 1, (size_t)out_len, fptr);
    } while (!EOF);

    fclose(fptr);

    return 0;
}
{% endhighlight %}

The code looks pretty complicated, probably due to the long function names, but in reality it is quite simple.
- We define some needed variables (``buf_out``, ``header`` and ``st``) to be used for our cryptography magic
later.
- We open the file to be encrypted and we initialize our cryptography magic using that function with the
absurdly long name (``crypto_secretstream_xchacha20poly1305_init_push``) and then we write the header to the file.
- We get the length of the password, and then proceed to do our main cryptography magic inside a do/while loop.

Now, lets move on to decryption:

{% highlight c %}
unsigned char *decrypt(const char *source_file, const unsigned char key[crypto_secretstream_xchacha20poly1305_KEYBYTES]) {
    unsigned char buf_in[MAX_STRING_LEN + crypto_secretstream_xchacha20poly1305_ABYTES];
    static unsigned char buf_out[MAX_STRING_LEN];
    unsigned char header[crypto_secretstream_xchacha20poly1305_HEADERBYTES];

    crypto_secretstream_xchacha20poly1305_state st;

    unsigned long long out_len;
    int eof;
    unsigned char tag;

    FILE *fp_s = fopen(source_file, "rb");
    fread(header, 1, sizeof(header), fp_s);
    
    if (crypto_secretstream_xchacha20poly1305_init_pull(&st, header, key) != 0) {
        return NULL;
    }

    do {
        size_t rlen = fread(buf_in, 1, sizeof buf_in, fp_s);
        
        eof = feof(fp_s);
        
        if (crypto_secretstream_xchacha20poly1305_pull(&st, buf_out, &out_len, &tag, buf_in, rlen, NULL, 0) != 0) {
            return NULL;
        }

        if (tag == crypto_secretstream_xchacha20poly1305_TAG_FINAL && ! eof) {
            return NULL;
        }

    } while (!eof);

    fclose(fp_s);
    
    return buf_out;
}
{% endhighlight %}

Again, seems pretty complicated, but is quite simple:
- Again, we define the variables to be used for our cryptography magic.
- Then, we open the file to be decrypted and read it using ``fread``.
- Then, we loop trough the file, we attempt to pull decrypted info and if it succeeds, we return the info.
Otherwise, we return ``NULL``.

Now, we need to add a function to write our encryption key to a file (we wanna decrypt our passwords with
it, of course). It proves to be quite simple, naturally:

{% highlight c %}
int create_key(const char *key_file)  {
    unsigned char key[crypto_secretstream_xchacha20poly1305_KEYBYTES];
    crypto_secretstream_xchacha20poly1305_keygen(key);

    FILE *f = fopen(key_file, "w");
    if (!f) {
        return 1;
    }

    fwrite(key, 1, crypto_secretstream_xchacha20poly1305_KEYBYTES, f);
    fclose(f);

    return 0;
}
{% endhighlight %}

This code is pretty self explanatory.

### 2. The Actual Password-Handling
Now, lets use our low-level interface to create a higher-level interface for handling passwords.
The plan is to have a directory, ``~/.locksmith``, to save all of our information to. More
specifically: ``~/.locksmith/locksmith.key`` should hold the key and ``~/.locksmith/passwords/``
should hold the passwords.

Well, naturally, our first step here is to be able to access the home directory, or ``~``, since
C doesn't let us access it by default. So let's create a commons/utils library! get your
``commons.c`` and ``commons.h`` files ready! Go ahead and just create some (basically) boilerplate
for our header file:

{% highlight c %}
#ifndef COMMONS_H
#define COMMONS_H

char *get_home_dir();

#endif
{% endhighlight %}

Then, we are going to get our home dir, right? Yeah, but only for UNIX-likes, since I couldn't be
bothered to find the Windows implementation of this function. This TECHNICALLY makes locksmith
system-dependent. It's really simple though:

{% highlight c %}
#include <unistd.h>
#include <pwd.h>

#include "commons.h"

char *get_home_dir() {
    return getpwuid(getuid())->pw_dir;
}

{% endhighlight %}

And finally, we can get to implementing the interface, which will be implemented in
``password.c``/``password.h``. Lets begin with the basic "locksmith file system", which is
basically a set of macros and a few functions that will be explained later. But for now,
this'll do:

{% highlight c %}
#ifndef PASSWORD_H
#define PASSWORD_H

#include <string.h> // strcat

#include "commons.h"

#define locksmith_dir strcat(get_home_dir(), "/.locksmith/")
#define locksmith_passw_dir strcat(locksmith_dir, "passwords/")
#define locksmith_key_file strcat(locksmith_dir, "key.txt")

#define get_locksmith_passw_dir_filepath(name) strcat(locksmith_passw_dir, strcat(name, ".txt"))

#endif
{% endhighlight %}

It's the classic macro expands to another macro bullshittery and it looks AWFUL, but it works. The
``get_locksmith_passw-dir_filepath`` macro appends ``name`` to the directory name, e.g
``get_locksmith_passw_dir_filepath("hey.txt") `` becomes ``~/.locksmith/passwords/hey.txt
Now, we can move on to our password handling functions:

{% highlight c %}
#ifndef PASSWORD_H
#define PASSWORD_H

#include <string.h> // strcat
#include <sodium.h>

#include "commons.h"

// ...

int create_password(char name[], char password[], unsigned char key[crypto_secretstream_xchacha20poly1305_KEYBYTES]);
char *get_password(char name[], unsigned char key[crypto_secretstream_xchacha20poly1305_KEYBYTES]);
int delete_password(char name[]);
int list_passwords();

#endif
{% endhighlight %}

This is basically all we need for now, I suppose. Let's begin with implementing the ``create_password()`` function
then. Open ``password.c``:

{% highlight c %}
#include <stdio.h>

#include "crypto.h"

int create_password(char name[], char password[], unsigned char key[crypto_secretstream_xchacha20poly1305_KEYBYTES]) {
    char *fname = get_locksmith_passw_dir_filepath(name);

    FILE *fptr;

    fptr = fopen(fname, "w");
    if (fptr == NULL) {
        printf("Couldn't create password. Failed to write to file.");
        return 200;
    }
    fprintf(fptr, "%s", password);
    fclose(fptr); 

    encrypt(fname, password, key);

    return 0;
}
{% endhighlight %}

I find this code pretty self explanatory. It writes the encrypted password to ``~/.locksmith/passwords/{name}``
and exits. Simple as. In practice we got the most complicated stuff out of the way in the beginning.
``get_password()`` is even easier:

{% highlight c %}
char *get_password(char name[], unsigned char key[crypto_secretstream_xchacha20poly1305_KEYBYTES]) {
    char *fname = get_locksmith_passw_dir_filepath(name);

    unsigned char *buf = decrypt(fname, key);

    encrypt(fname, buf, key);

    return buf;
}
{% endhighlight %}

Additionally, we need to be able to delete passwords, hence ``delete_password()``:

{% highlight c %}
int delete_password(char name[]) {
    char *fname = get_locksmith_passw_dir_filepath(name);
    remove(fname);

    return 0;
}
{% endhighlight %}

And then one of the most important utilities, ``list_password()``:

{% highlight c %}
// ...

#include <string.h>
#include <dirent.h>

// ...

int list_passwords() {
    struct dirent *de;

    DIR *dir = opendir(locksmith_passw_dir);

    if (dir == NULL) {
        printf("Couldn't list passwords.");
        return 200;
    }

    while ((de = readdir(dir)) != NULL) {
        char *pname = strtok(de->d_name, ".");
        if (pname != NULL) printf("%s\n", pname);
    }

    closedir(dir);    

    return 0;
}
{% endhighlight %}

This function is actually quite complicated. Firstly, we need to include ``string.h`` and
``dirent.h``. Additionally, we need to remove the file extension from the password file (because
the password name is just the password filename minus the file extension). But really it is quite
simple, aswell.

Oh, and one last thing! In order to actually decrypt the passwords, we need to *get* the key, not
just create it. So go ahead and add one last function to ``password.h``:

{% highlight c %}
#include <sodium.h>

// ...

int get_key(unsigned char key[crypto_secretstream_xchacha20poly1305_KEYBYTES]);
{% endhighlight %}

And ``password.c``:

{% highlight c %}
#include <sodium.h>

// ...

int get_key(unsigned char key[crypto_secretstream_xchacha20poly1305_KEYBYTES]) {
    FILE *key_fd = fopen(locksmith_key_file, "r");
    if (key_fd == NULL) {
        exit(2);
    }

    fread(&key, crypto_secretstream_xchacha20poly1305_KEYBYTES, 1, key_fd);

    return 0;
}
{% endhighlight %}

### 3. The Locksmith Shell
Now, for the last part of our password manager, we're gonna create a small shell to run our
password-handling commands in. So go ahead and create ``cli.h`` and ``cli.c``. Firstly, we need
to create a "global key" (global var containing the contents of ``~/.locksmith/locksmith.key``)
to be passed to other functions:

{% highlight c %}
static unsigned char key[crypto_secretstream_xchacha20poly1305_KEYBYTES];

// high skill code amiright
int cli_init() {
    get_key(key);

    return 0;
}
{% endhighlight %}

Now, we need to basically create user-friendly wrappers around the ``password.h`` interface.
These are incredibly easy to implement, we just pass user input to the interface:

{% highlight c %}
#include <stdio.h>

#include "cli.h"
#include "crypto.h"
#include "password.h"

// ...

int cmd_get_password() {
    char pass_name[MAX_STRING_LEN];

    printf("What password do you want to get?\n"
            "List of passwords:\n");
    list_passwords();
    printf("> ");
    scanf("%s", pass_name);

    printf("Password: %s\n", get_password(pass_name, key));

    return 0;
}

int cmd_delete_password() {
    char pass_name[MAX_STRING_LEN];

    printf("What password do you want to delete?\n"
            "List of passwords:\n");
    list_passwords();
    printf("> ");
    scanf("%s", pass_name);

    delete_password(pass_name);
    
    return 0;
}
{% endhighlight %}

However, you might've noticed something: Where's ``cmd_create_password()``? Well, I want my password name
to have a certain syntax to it, more specifically ``{site name}:{user name}``, which means I have to take
3 inputs: sitename, username, and password. However, in the end it is pretty simple aswell, I just felt
like it needed a mention:

{% highlight c %}
#define format_password_filename(site, user) strcat(strcat(site, ":"), user)

int cmd_create_password() {
    char site_name[MAX_STRING_LEN], user_name[MAX_STRING_LEN], password[MAX_STRING_LEN];

    printf("Site:\n> ");
    scanf("%s", site_name);

    printf("Username:\n> ");
    scanf("%s", user_name);

    printf("Password:\n> ");
    scanf("%s", password);

    create_password(format_password_filename(site_name, user_name), password, key);

    return 0;
}
{% endhighlight %}

And then we of course need to do what any regular ol' shell does: *actually interpret commands.*
Since I'm... Well, y'know: *lazy*, I'm just gonna read an ``int`` as a command and interpret it 
using a plain switch statements. It's as simple as:

{% highlight c %}
#include <stdlib.h>

// ...

int cmd_interface(int *exit) {
    printf("\nEnter a command:\n"
           "[1]: Create password\n"
           "[2]: Get password\n"
           "[3]: Delete password\n"
           "[4]: Exit program\n");

    int command;
    printf("> ");
    scanf("%d", &command);

    char site_name[MAX_STRING_LEN];
    char user_name[MAX_STRING_LEN];
    char password[MAX_STRING_LEN];

    switch (command) {
        case 1:
            cmd_create_password();
            break;
        case 2:
            cmd_get_password();
            break;
        case 3:
            cmd_delete_password();
            break;
        case 4:
            printf("Exiting...\n");
            exit(0);
        default:
            printf("Invalid command! Exiting...\n");
            exit(0);
    }

    return 0;
}
{% endhighlight %}

Lastly, we expose the functions in ``cli.h``:

{% highlight c %}
#ifndef CLI_H
#define CLI_H

int cli_init();
int cmd_create_password();
int cmd_get_password();
int cmd_delete_password();
int cmd_interface();

#endif
{% endhighlight %}

### 4. Putting It All Together
The very, *very* last thing we need to do is to strap all of this together trough a ``main.c`` file
(and then a ``Makefile``) to make the program... Actually run. However, firstly, *how do we initalize
the locksmith filesystem?* Well, basically, we need to create all directories and keys, *if they
don't exist*. This can easily be solved with a new ``commons.c`` function:

{% highlight c %}
#include <sys/stat.h>

// ...

int directory_exists(char *dirname) {
    struct stat st;
    if (stat(dirname, &st) == 0) {
        return 1;
    }
    return 0;
}

int mkdirifnotexist(char *dirname) {
    if (directory_exists(dirname) == 0) {
        mkdir(dirname, 0755);
    }

    return 0;
}
{% endhighlight %}

Exposing the functions in ``commons.h``:
{% highlight c %}
// ...

int directory_exists(char *dirname);
int mkdirifnotexist(char *dirname);
{% endhighlight %}

And now we can create ``main.c`` and a basic ``main()`` function:

{% highlight c %}
#include <stdio.h>

#include "password.h"
#include "cli.h"
#include "crypto.h"

int main() {
    mkdirifnotexist(locksmith_dir);        // ~/.locksmith
    mkdirifnotexist(locksmith_passw_dir);  // ~/.locksmith/passwords

    create_key(locksmith_key_file);        // ~/.locksmith/locksmith.key

    cli_init();

    while (1) {
        cmd_interface();
    }

    return 0;
}
{% endhighlight %}

For extra personality, we can also add "Locksmith" in a random figlet font:

{% highlight c %}
#include <stdio.h>

#include "password.h"
#include "cli.h"
#include "crypto.h"

#define locksmith_title \
"88                                88                                     88         88         \n" \
"88                                88                                     \"\"   ,d    88         \n" \
"88                                88                                          88    88         \n" \
"88          ,adPPYba,   ,adPPYba, 88   ,d8  ,adPPYba, 88,dPYba,,adPYba,  88 MM88MMM 88,dPPYba, \n" \
"88         a8\"     \"8a a8\"     \"\" 88 ,a8\"   I8[    \"\" 88P'   \"88\"    \"8a 88   88    88P'    \"8a\n" \
"88         8b       d8 8b         8888[      `\"Y8ba,  88      88      88 88   88    88       88\n" \
"88         \"8a,   ,a8\" \"8a,   ,aa 88`\"Yba,  aa    ]8I 88      88      88 88   88,   88       88\n" \
"88888888888 `\"YbbdP\"'   \"Ybbd8\"' 88   `Y8a `\"YbbdP\"' 88      88      88 88   \"Y888 88       88\n\n" \

int main() {
    mkdirifnotexist(locksmith_dir);        // ~/.locksmith
    mkdirifnotexist(locksmith_passw_dir);  // ~/.locksmith/passwords

    create_key(locksmith_key_file);        // ~/.locksmith/locksmith.key

    cli_init();

    printf(locksmith_title);

    while (1) {
        cmd_interface();
    }

    return 0;
}
{% endhighlight %}

### Built to be flawed
And that brings us to the end of this semi-tutorial. I have one thing to say before you go: this
version of the software was **built to be flawed**.
[full-implementation locksmith](https://github.com/callmesalmon/locksmith) is much more complete
and much more memory-safe, but it isn't *flawed*. And being flawed is good. Being flawed means
there is so much more room for improvement. And there is some kind of beauty in that. Have fun.
