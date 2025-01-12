---
layout: post
title:  "I tried to make a game engine"
author: Elis Staaf
date:   2025-01-12 19:56:09 +0100
categories: projects
---

Alright. So... Game engines. Game engines are... Well, engines that
you use to make games. Some notable ones are Unity, Unreal engine
and my personal favourite; Godot. Wait! I'm getting a call! Oh, so
apparently Godot was involved in a massive controversy involving...
I don't remember, it was a month ago.

*"Wait! Why are you telling me all of this?"*. Well, my friend, I
wanted to make a game engine. And not just a game engine, the
worst engine of all time.

### Why game engines are inherently bad
Game engines aren't and will never be good. They're built on a principle that
you should put as many layers of abstraction on top of bare metal... That
doesn't work. You need control over the system, game engines don't give
you that. The only things game engines give you is a lack of challenge
and *way* too many useless tools. And with that out of the way,
let's make one!

### Making a game engine
Okay... How make game engine then? That's what I searched on Firefox.
I was looking for somewhat useful tutorials... That's not what I found.
The only thing I found was people screaming that you should either
use C++ to make your game engine, or give up.

Okay, so the question of what language to use is answered, but what
graphics library. And *no*, making my own graphics library is out
of the question. Well, I had used SFML in one of my older projects
before and it's still installed on my system, so I might as well.

And now... Let's code. Hmm... How do I say this? THIS WAS PURE FUCKING
TORTURE!!! I needed to create...

- An engine.
- A sprite object.
- A graphics object.
- Input.
- Frame updating.

Let's tackle them one at a time. Let's begin with the easier ones.
Keep in mind that right now, `engine.h` is an "abstract object"
(it doesn't exist bozo). This is also true with all of the 
"Bob" methods (where Bob is our sprite object). Input:

{% highlight cpp %}
#include "engine.h"
 
#include <SFML/Graphics.hpp>
using namespace sf;

void Engine::input()
{
    // Handle the player quitting
    if (Keyboard::isKeyPressed(Keyboard::Escape))
    {
        m_Window.close();
    }
 
    // Handle the player moving
    if (Keyboard::isKeyPressed(Keyboard::A))
    {
        m_Bob.moveLeft();
    }
    else
    {
        m_Bob.stopLeft();
    }
 
    if (Keyboard::isKeyPressed(Keyboard::D))
    {
        m_Bob.moveRight();
    }
    else
    {
        m_Bob.stopRight();
    }                       
 
}
{% endhighlight %}

Graphics object:

{% highlight cpp %}
#include "engine.h"

#include <SFML/Graphics.hpp>
using namespace sf;
 
void Engine::draw()
{
    // Rub out the last frame
    m_Window.clear(Color::White);
 
    // Draw the background
    m_Window.draw(m_BackgroundSprite);
    m_Window.draw(m_Bob.getSprite());
 
    // Show everything we have just drawn
    m_Window.display();
}
{% endhighlight %}

And frame updating:

{% highlight cpp %}
#include "engine.h"

#include <SFML/Graphics.hpp>
using namespace sf;
 
using namespace sf;
 
void Engine::update(float dtAsSeconds)
{
    m_Bob.update(dtAsSeconds);
}
{% endhighlight %}

Now to making a sprite object! This was not fun.
This was infact not fun at all. This was infact:

{% highlight cpp %}
#include "bob.h"

#include <SFML/Graphics.hpp>
using namespace sf;
 
Bob::Bob()
{
    // How fast does Bob move?
    m_Speed = 400;
 
    // Associate a texture with the sprite
    m_Texture.loadFromFile("img/bob.png");
    m_Sprite.setTexture(m_Texture);     
 
    // Set the Bob's starting position
    m_Position.x = 500;
    m_Position.y = 800;
 
}
 
// Make the private spite available to the draw() function
Sprite Bob::getSprite()
{
    return m_Sprite;
}
 
void Bob::moveLeft()
{
    m_LeftPressed = true;
}
 
void Bob::moveRight()
{
    m_RightPressed = true;
}
 
void Bob::stopLeft()
{
    m_LeftPressed = false;
}
 
void Bob::stopRight()
{
    m_RightPressed = false;
}
 
// Move Bob based on the input this frame,
// the time elapsed, and the speed
void Bob::update(float elapsedTime)
{
    if (m_RightPressed)
    {
        m_Position.x += m_Speed * elapsedTime;
    }
 
    if (m_LeftPressed)
    {
        m_Position.x -= m_Speed * elapsedTime;
    }
 
    // Now move the sprite to its new position
    m_Sprite.setPosition(m_Position);   
 
}
{% endhighlight %}

Okay, I have been trough worse. But DELTA TIME!? I don't
want to program delta time... Now to the actual engine:

{% highlight cpp %}
#include "engine.h"

#include <SFML/Graphics.hpp>
using namespace sf;
 
Engine::Engine()
{
    // Get the screen resolution and create an SFML window and View
    Vector2f resolution;
    resolution.x = VideoMode::getDesktopMode().width;
    resolution.y = VideoMode::getDesktopMode().height;
 
    m_Window.create(VideoMode(resolution.x, resolution.y),
        "Simple Game Engine",
        Style::Fullscreen);
 
    // Load the background into the texture
    // Be sure to scale this image to your screen size
    m_BackgroundTexture.loadFromFile("background.jpg");
 
    // Associate the sprite with the texture
    m_BackgroundSprite.setTexture(m_BackgroundTexture);
 
}

void Engine::start()
{
    // Timing
    Clock clock;
 
    while (m_Window.isOpen())
    {
        // Restart the clock and save the elapsed time into dt
        Time dt = clock.restart();
 
        // Make a fraction from the delta time
        float dtAsSeconds = dt.asSeconds();
 
        input();
        update(dtAsSeconds);
        draw();
    }
}
{% endhighlight %}

I couldn't be bothered to include any images, that's your job.
Anyway, after whipping up a quick ``main.cpp`` it was time
to build the project! Like so:

{% highlight shell %}
g++ src/*.cpp -o engine
{% endhighlight %}

And I got an error!? It said something about not being
able to create a DT_TEXTREL in a PIE, and I actually totally
agree. While I don't know what a DT_TEXTREL is, I don't want it
in my pie! That'll make it taste bad! No, I just needed to compile
with ``-lsfml-graphics`` and I was good to go, but it was real
scary. Okay, and now it wo-

WHAT THE!? ERROR!? Oh, okay! My ``$DISPLAY`` environment variable
wasn't set correctly!? It was ``:0.0``, and I still don't know
what it was talking about (plot twist: I needed to input
my IP address for some reason. Btw, my IP address is ********).
And then, my (non-)beautiful game engine was done! It was
(non-)beautiful! I could... Move left and right, and...
That was about it... Hmm...

Hehe, let's just ignore that and focus on that I DID IT!!!
Check out my project [here](https://github.com/ElisStaaf/ngin)
and have a... Day!
