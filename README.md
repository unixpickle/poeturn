# poeturn

It's your turn to be a poet. With poeturn, you and a recurrent neural network can take turns writing lines of a poem.

# Usage

First, you must install and configure [Go](https://golang.org/doc/install).

Once you have Go, download this code as follows:

```
go get -d github.com/unixpickle/poeturn/...
```

Fetch poetry data as follows (saves to `/path/to/poems.json`):

```
$ cd $GOPATH/src/github.com/unixpickle/poeturn/fetch_data
$ go run *.go /path/to/poems.json 1000
```

Train a network on the data as follows (saves network to `/path/to/network`):

```
$ cd $GOPATH/src/github.com/unixpickle/poeturn/train
$ go run *.go -samples /path/to/poems.json -output /path/to/network
```

You may press Ctrl+C exactly once to stop/pause training. Pressing it multiple times will terminate without saving. It is suggested that you train the network for a while (perhaps 24 hours). Of course, you may pause and resume training several times during this process.

Once you have a trained network, you can compose a poem as follows:

```
$ cd $GOPATH/src/github.com/unixpickle/poeturn/compose
$ go run *.go -network /path/to/network
```

When you see a `>`, it is a prompt asking you to enter a line. You can have the network start off writing `n` lines by adding `-netstart n` to the arguments.

# Results

Here are some "poems" I wrote with the help of this program. They seem to make just as much sense as regular poetry, if not more. Lines beginning with `>` were my contributions:

```
Fare, from me, whose sheets only another;
> for I have to clean the sheets of my dead mother!
it is the start-window-wheel disday
> and I must ensure not to cause dismay.
Cross the tide. You falls adains.
> The tides are now growing stronger,
uneven and fierceal, laboured, calm belied:
> and then suddenly, I was on shore.
Uplass !ill you call the woods the leases
> which we cannot retract but only sublet.
```

```
> Roses are red
In broken drift of skin
> not attached to body.

Death's fresk predect,
> man's most perfect defect,
For Then.
>     
Of alas! yes, no, though brassies
> might never store a boolean.

Up, Knock, sorrow, all my well!
> Down, flock, morrow, all you tell!
Her two diviner's smatter!
> And his three sinners splatter!
Love, can fear her big your fate
> but not your mate.
```

```
In griming in the current on
> the seas off of London,
And the thing gives ye within home bereat,
> yet does not feel like a nice treat.

They starter brother,
> and ended mother,
We live of my naily year be:
> that he may one day see!

Do yoke, I lived hy oaked will not die
> because he's killed by a spy!   
Left and follets as the kingdom of cow,
> and the queendom of tao,
Just recolfing a hollyronness
> and making up other silly words.
```

```
When I might pray to be lovin' togethe
> with you and not any othe.
We pay you but a moment. How should have roam
> the lands doth done.

A wood takes you have you till tepart arrayed
> in a linear fashion along the shore.
Plead by the soul.
> Bleed by the sword.

Me wears that word where more excettion trudy.
> Me hears that turd which gives exception truly!
Die to the joy's eye Virtue?
> Or live to the boy's pride a statue.
```

Here's some poems I let the net write on its own:

```
Snow past the spirit blind befined,
I didn't in the old roof world!
the Walling of the valley,

The sun is servingly, and by the borth
Can find a living of the made hidverly
Had fled, of wealth of silver winds
Or a disery roff and opbisless?

Both out if the wild-pounding scit.

How maches back to kill that stone depress,
Red clay whats could revell for words.
```

```
So, my beats that I write she silent,
To thee unhappy as blue-'disal
The hills of the season may nirr
Just wait his blood with the both may read
a-weaffird human hast in the storm.
only I told he plashes your door--pates, alad!
Nor to his eyes when out of he is!
He pursue the valiant futter bad doth in death,
Because through marter-full contravious part,
E'en makes unifout that you've find to guez,
Leaving you alrost life . . . here was in
:
and he'll died -- the wail will be!"
And . . . and repited Unsayl's" (While be still,
Join' old words youth above the dere devantive.
Sea-foundest found it some sweet painted,
Sea made me growing guint escaped by;
I demicing sickness of the feel's abtorted :
I may not from the veil dread;
With cod clear as a shell
```

```
I'm just a school, my fear expine,
Where I will not leap, my love,
When one my keftless age,
And seement you rest again,

I must live my frazelly keen.'

Wade, I enter we must be,
Nor did reaches in hands
That selflic eyes, and Death.

'Tis mine own fire doth remember
The secon must enchant;
A cloudless name as they green'
Let me thy heart the hurranest orient,
Lest this are seen be bearned.
```
