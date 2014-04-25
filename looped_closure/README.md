# Looped Closures

Pretty simple example. Set up X number of Go channels via an anonymous closure. Perform some work (in this case, format a string), post it to a results channel.

Only exit the program once we have received the results from all channels, or have received an interrupt or SIGTERM.

The loop prints results from the results channel to stdout.
