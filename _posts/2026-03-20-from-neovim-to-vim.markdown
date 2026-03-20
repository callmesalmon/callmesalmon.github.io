---
layout: post
title:  "I switched from NeoVIM to VIM and it was one of the best decisions of my life"
author: Salmon
date:   2026-03-20 20:50:03 +0100
categories: dunno
---

In the ancient times, the old programming wizards used to debate whether EMACS or VI was the
best text editor. The EMACS camp was mostly centred around the GNU version of EMACS, and
likewise, the VI camp came to be centred around Bram Moolenar's cult classic: VIM.

EMACS lost the editor war.

In 2015, I believe one Thiago de Arruda submitted a patch to VIM adding multithreading, which
quickly gained attention from the community. However, the head maintainer/BDFL and creator of
VIM, Bram Moolenar, never merged the patch. This caused a fair chunk of the VIM community
to fork VIM, creating NeoVIM.

As time went by, NeoVIM added many new features alongside multithreading, among others:
- They added Lua as the new configuration language.
- They added a builtin LSP.
- They added MOUSE SUPPORT >:(((

Anyways, none of this really matters. The only reason this is of any
importance to me is that when I switched to using terminal-based text editors,
I settled on NeoVIM, not VIM and definitely not EMACS.

### The NeoVIM experience
For about 3 months, I used plain, unconfigured NeoVIM. Everytime it started up and I was going
to work on my Go project I would manually type ``:colorscheme koehler`` and I'd be happy.

This was my introduction to VIM, and I was considering just switching back to VSCode, which
is what I had been using before I switched to terminal based editors.

However, one day, I realized I could just use a community-maintained NeoVIM configuration.
I installed LunarVim and when I opened up NeoVIM I saw a big window saying ``lazy.nvim`` and
saying that like 5 trillion packages were installing. And boom, I was in. Life was good.
I could write my C project in peace. Until I wanted to change config.

I changed to NvChad.

One week later, I switched to LazyVim.

In the end, I decided to actually learn to configure NeoVIM and boy can I tell you it wasn't
fun. Call it a skill issue all you want but I'm sure I sat damn near 4 hours debugging my
stupid config until I got it to *resemble* a config. Then, for the next 5 weeks or so, I
fine-tuned my config. This is the config I lived with under the rest of my time in NeoVIM.

### Issues (what some dumbass sociologist would call a "push factor")
Over the time, issues were slowly creeping up, but the first big one was when I updated all my
lazy packages because TreeSitter refused to run otherwise, but *then* my LSPConfig errored out
BECAUSE IT HAD A FUCKING SYNTAX OVERHAUL IN ONE OF THE UPDATED I INSTALLED. I debugged my config
for 30 minutes and then it finally worked and I could finally work.

Then I updated once again, and my TreeSitter laid down and cried BECAUSE IT HAD ALSO GONE TROUGH
A SYNTAX OVERHAUL. Another 30 minutes wasted. Another impairment of what was supposed to be a
"relaxing" coding session.

Another little bit more hidden problem is that Lua is a slog and that it should never be used
for everything ever. There were **ZERO** reasons to switch to Lua in the first place, and so
I see all the reasons to switch back. You are not writing an app, you are writing a config.
There is no reason to use an actual language when there is already a perfectly fine
configuration language.

And then there is also the fact that the mouse is enabled by default. I found myself using
the mouse more than the faster VIM bindings, as I could never seem to remember to disable
the mouse, as I had more important things to do: like actually code.

The ONE thing that was making me stay was that I had a really good file manager installed
as a plugin which I couldn't live without. Then I remembered ``netrw``. And so there was
no reason to stay.

This is why I decided to switch to VIM.

### Switching to VIM
Installing VIM is easy-as. It is preinstalled on my system so I can just run ``vim`` and
enjoy! However, a VIM is not complete without a config, and unlike my NeoVIM config, this
wasn't even 100 lines. I started off with this:

{% highlight vim %}
set nu
set expandtab
set tabstop=4
set softtabstop=4
set shiftwidth=4
set autoindent
set smartindent

let mapleader = ','
{% endhighlight %}

Then came the big part: ``netrw``. I <3 ``netrw`` now, and there isn't too much to add, really:

{% highlight vim %}
" Unironically the best keymap ever
nnoremap <leader>e :E<CR>

" Configuring netrw
let g:netrw_localcopydircmd = 'cp -r'
let g:netrw_list_hide = '\(^\|\s\s\)\zs\.\S\+'
{% endhighlight %}

However, I realized I might need to see the files hidden by ``g:netrw_list_hide`` so I decided
to emulate a function of the file tree plugin I used in NeoVIM: press ``H`` to toggle showing
hidden files on/off. Pretty simple:

{% highlight vim %}
" Half of the config is just some code to toggle hiding
" files in netrw lol
function! NetrwToggleHidden()
  if exists('g:netrw_list_hide') && g:netrw_list_hide ==# ''
    let g:netrw_list_hide = '\(^\|\s\s\)\zs\.\S\+'
  else
    let g:netrw_list_hide = ''
  endif
  silent! e .
endfunction

augroup NetrwToggleHidden
  autocmd!
  autocmd FileType netrw nnoremap <buffer> H :call NetrwToggleHidden()<CR>
augroup EN
{% endhighlight %}

### Conclusion
I have been using VIM with this configuration for a few days now and safe to say: I like it *much*
more than NeoVIM, and I have found that I am also much faster and more efficient while writing
code, as I have been restricted to VIM bindings. ``netrw`` is just as good as my old file tree
and programming is genuinely much more fun when I can actually focus on... Well, programming.

In retrospect, I realize that I was a fool for thinking I needed IDE-like features like LSPs
and advanced code completion. It took a while to get over needing these features but now I am
more efficient than ever and enjoying programming much more.

Of course, there are a few small quality of life things that are included in NeoVIM and not
VIM, but those problems are minimal compared to the benefits of using plain old VIM. No more
big config to take care of and no more plugins breaking.

So my takeaway from this is: fuck EMACS.
