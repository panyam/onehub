Get and Peek messages return immediately and dont block for an event to be available
So we need the following construct:

1. Get some messages - from the log
2. Process it
3. Tell the log to commit the messages - otherwise if (2) had an error then we have loss
4. got to step 1

we only need the 1 + 2 + 3 if we care about atelast once deliver
Alternatively we could make (1) do the "commit" as well and focus on making (2) more robus
But (2) can always be buggy/failure prone making our log lossy

Another mode is we may not want to do (3) in 1 go but in multiple goes - this way
the log reader can work off smaller chunks
In eithe rcase - we want (1) to either be blocking - or there be a message that says
"WaitTillEvent" cal that blocks (interruptibly) until there are messages
Channels are good for this

```
checker := time.NewTimer(time.Millisecond * 50)
msgsize := 50
for {
    select {
    case <-checker.C:
        log.Println("Ping....")
    case evt := <-p.EventChannel():
        log.Println("Got New Event: ", evt)
    }
    peeked, err := p.GetMessages(msgsize, false)
    if err != nil {
        panic(err)
    }

    log.Println("Peeked: ", peeked)
    // Consume the messages
    p.GetMessages(msgsize, true)
}
```

Our data sources and sinks:

1. Event channel - gets us X events at a time
2. Event committer sync - lets us tell the event source when it can "commit" what it has consumed
3. Commands to pause/resume - This pauses/resumes log processing
4. Command to generate a "low" water mark + dumpid
5. Command to generate a "high" watermark + dumpid
6. Dump Interface For:
7.  GetAll(dumpid) -> returns all entries in a dump
8.  RemoveItem(dumpid, key) -> Removes items (by key) from a dump

When (1) happenes how should the event loader work?   When (1) is called - at the initial state
the loader can do a "peek" from the datasource (and may even buffer this or do literally 1 at a time).  Any subsequent peeks will return the same items as the client hasnt "comitted it yet" in (2).   If say 1000 items have been peeked and returned to the client, the client cannot do another peek until a commit has been issued.   But after some "max" storage call to commit (2)
will make more available - this lets us have at-least once semantics.   If we want atmost once
then we can make (1) a Get instead of a peek.